package core

import (
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

// Quiet hours are local time windows during which PostHog does not evaluate an alert.
// Two resources expose them with different Terraform shapes - posthog_alert nests them
// under schedule_restriction, posthog_logs_alert has them at the top level - but the
// windows themselves mean the same thing and hit the same server-side handling, so the
// rules live here once.
//
// PostHog does not store the windows it is given. It lays them on a single 24-hour
// timeline, merges anything that touches or overlaps, and re-derives windows from the
// merged result. A configuration it would reshape reads back different from the plan,
// which Terraform reports as "Provider produced inconsistent result after apply" rather
// than as recoverable drift. Every rule below exists to reject one such shape at plan
// time, and each was confirmed against a live PostHog instance.
const (
	MinutesPerDay = 24 * 60
	// MinQuietHoursWindowMinutes is the shortest window PostHog accepts.
	MinQuietHoursWindowMinutes = 30
	// MaxQuietHoursWindows is the most windows PostHog stores for one alert.
	MaxQuietHoursWindows = 5
)

// QuietHoursTimePattern is the single definition of the HH:MM grammar. Schema validators
// and the parser below both check against it, so the layers cannot disagree about what a
// valid time is. time.Parse alone is not enough: it accepts a single-digit hour.
var QuietHoursTimePattern = regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)

// QuietHoursWindow is one window as written in configuration. Times are HH:MM strings so
// callers can pass what the practitioner typed and get diagnostics naming it back.
type QuietHoursWindow struct {
	Start string
	End   string
}

func (w QuietHoursWindow) label() string { return w.Start + "-" + w.End }

// QuietHoursMinutes converts an HH:MM string to minutes past midnight, reporting whether
// it is a time this package recognises.
func QuietHoursMinutes(value string) (int, bool) {
	if !QuietHoursTimePattern.MatchString(value) {
		return 0, false
	}
	parsed, err := time.Parse("15:04", value)
	if err != nil {
		return 0, false
	}
	return parsed.Hour()*60 + parsed.Minute(), true
}

// span is one contiguous blocked range within a single day. A window crossing midnight
// produces two, both tagged with the index of the window they came from so the two halves
// of one window are never mistaken for a conflict between two windows.
type span struct {
	window     int
	start, end int
}

// ValidateQuietHoursWindows reports every way the given windows would be reshaped by
// PostHog. attrPath is the attribute the diagnostics are attached to, which differs
// between resources.
//
// Windows whose times this package does not recognise are skipped: the schema's own
// format validator reports those, and re-reporting them here would duplicate the error.
func ValidateQuietHoursWindows(windows []QuietHoursWindow, attrPath path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	var spans []span
	parsed := 0
	crossings := 0

	for i, w := range windows {
		start, startOK := QuietHoursMinutes(w.Start)
		end, endOK := QuietHoursMinutes(w.End)
		if !startOK || !endOK {
			continue
		}

		length := end - start
		if length < 0 {
			length += MinutesPerDay
		}
		if length < MinQuietHoursWindowMinutes {
			diags.AddAttributeError(
				attrPath,
				"Quiet-hours window is too short",
				fmt.Sprintf(
					"Window %s spans %d minutes. PostHog requires each window to span at least %d minutes. "+
						"PostHog can produce such a window itself when it splits a window crossing midnight, so "+
						"an imported alert may carry one it will not accept back. Widen or drop it: the alert "+
						"cannot be applied as stored.",
					w.label(), length, MinQuietHoursWindowMinutes,
				),
			)
			continue
		}

		parsed++
		switch {
		case end > start:
			spans = append(spans, span{i, start, end})
		case end == 0:
			// Runs to the end of the day without continuing into the next morning, so it
			// needs no second half. An empty [0, 0) half would collide with every window
			// starting at midnight in the overlap check.
			spans = append(spans, span{i, start, MinutesPerDay})
		default:
			crossings++
			spans = append(spans, span{i, start, MinutesPerDay}, span{i, 0, end})
		}
	}

	// Checked before overlap and reported alone: a whole-day block always overlaps too,
	// and "combine them into one window" is the wrong advice here, since doing so yields a
	// zero-length window that fails again for an unrelated reason.
	if covered := quietHoursCoverageError(spans, attrPath); covered.HasError() {
		diags.Append(covered...)
		return diags
	}

	diags.Append(quietHoursOverlapErrors(spans, windows, attrPath)...)

	// A window written as crossing midnight always expands to two spans, and PostHog only
	// rejoins them when nothing else is on the timeline.
	if crossings > 0 && parsed > 1 {
		diags.AddAttributeError(
			attrPath,
			"Quiet-hours window crossing midnight must be the only window",
			fmt.Sprintf(
				"A window whose end is before its start crosses midnight, and PostHog stores it as one window "+
					"per side of midnight unless it is the only window configured. Either make it the only "+
					"window, or split it yourself into a window ending at 00:00 and one starting at 00:00. "+
					"Splitting costs a window slot, so it needs you to be under the limit of %d.",
				MaxQuietHoursWindows,
			),
		)
		return diags
	}

	// The rule below is the one a skipped window can make wrong rather than merely
	// incomplete: it keys off there being exactly two spans, which a skipped window can
	// fabricate. The rules above only ever miss a problem on a partial timeline.
	if parsed != len(windows) {
		return diags
	}

	// Two windows that between them block both sides of midnight are rejoined into a
	// single crossing window, but only while they are the whole timeline. A third window
	// anywhere in the day leaves all of them stored as written.
	if len(spans) == 2 && spans[0].window != spans[1].window &&
		((spans[0].start == 0 && spans[1].end == MinutesPerDay) ||
			(spans[1].start == 0 && spans[0].end == MinutesPerDay)) {
		diags.AddAttributeError(
			attrPath,
			"Quiet-hours windows meeting at midnight are stored as one",
			fmt.Sprintf(
				"Windows %s and %s together block midnight from both sides, and PostHog rejoins them into a "+
					"single crossing window when they are the only two, so the alert would not match this "+
					"configuration. Write them as one window crossing midnight, add a third window elsewhere "+
					"in the day, or end the evening window at 23:59.",
				windows[spans[0].window].label(), windows[spans[1].window].label(),
			),
		)
	}

	return diags
}

// quietHoursCoverageError reports windows that between them block the whole day. PostHog
// refuses that outright, and the overlap message would otherwise advise combining the
// windows, which produces a zero-length window that fails again for another reason.
func quietHoursCoverageError(spans []span, attrPath path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	blocked := 0
	var covered [MinutesPerDay]bool
	for _, s := range spans {
		for m := s.start; m < s.end && m < MinutesPerDay; m++ {
			if !covered[m] {
				covered[m] = true
				blocked++
			}
		}
	}
	if blocked >= MinutesPerDay {
		diags.AddAttributeError(
			attrPath,
			"Quiet-hours windows cover the whole day",
			"These windows block every minute of the day, so the alert could never run. PostHog rejects "+
				"this. Leave at least one gap when the alert is allowed to be evaluated.",
		)
	}
	return diags
}

// quietHoursOverlapErrors reports windows that overlap or merely touch. PostHog merges on
// `next.start <= prev.end`, so 00:00-06:00 and 06:00-09:00 are saved as one window.
func quietHoursOverlapErrors(spans []span, windows []QuietHoursWindow, attrPath path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	for i := range spans {
		for j := i + 1; j < len(spans); j++ {
			// The two halves of one crossing window cannot conflict: the first ends at
			// midnight and the second starts there. Comparing window indexes rather than
			// the rendered label keeps identity out of presentation.
			if spans[i].window == spans[j].window {
				continue
			}
			if spans[i].start <= spans[j].end && spans[j].start <= spans[i].end {
				diags.AddAttributeError(
					attrPath,
					"Quiet-hours windows overlap",
					fmt.Sprintf(
						"Windows %s and %s overlap or run straight into each other. PostHog merges them into "+
							"a single window when saving, so the alert would not match this configuration. "+
							"Combine them into one window.",
						windows[spans[i].window].label(), windows[spans[j].window].label(),
					),
				)
			}
		}
	}
	return diags
}
