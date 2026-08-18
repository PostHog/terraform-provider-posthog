package core

import (
	"fmt"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

// Quiet hours are the times of day PostHog does not evaluate an alert. Two resources
// expose them under different attributes, but the windows mean the same thing, so the
// rules live here once.
//
// Only one kind of rule belongs here: the shapes PostHog accepts and then rewrites.
// PostHog merges the windows onto one 24-hour timeline and derives new ones, so a
// rewritten config reads back different from the plan and the apply fails. There is no
// server error to pass on, so the provider has to catch these itself.
//
// PostHog's own limits on window length and count are not repeated here. Copying a server
// constant means a provider release whenever PostHog changes it.
const MinutesPerDay = 24 * 60

// QuietHoursTimePattern is the only definition of the HH:MM grammar. Schema validators and
// the parser below both use it, so they cannot disagree about what a valid time is.
// time.Parse on its own would accept a single-digit hour.
var QuietHoursTimePattern = regexp.MustCompile(`^([01]\d|2[0-3]):[0-5]\d$`)

// QuietHoursWindow is one window as written in configuration. Times stay as strings so
// diagnostics can name the window the way the user wrote it.
type QuietHoursWindow struct {
	Start string
	End   string
}

func (w QuietHoursWindow) label() string { return w.Start + "-" + w.End }

// QuietHoursMinutes converts HH:MM to minutes past midnight. The second value reports
// whether the string was a time at all.
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

// span is one blocked range inside a single day. A window crossing midnight makes two.
// Each carries the index of its window, so the halves of one window are never read as a
// conflict between two.
type span struct {
	window     int
	start, end int
}

// ValidateQuietHoursWindows reports every way PostHog would rewrite the given windows.
// attrPath is where the diagnostics attach, which differs between resources.
//
// Windows whose times do not parse are skipped. The schema's format validator already
// reports those.
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

		// Equal bounds block no time. Skip them. Treating one as a span would read
		// 00:00-00:00 as a whole day. PostHog rejects them on apply.
		if start == end {
			continue
		}

		parsed++
		switch {
		case end > start:
			spans = append(spans, span{i, start, end})
		case end == 0:
			// Ends at the close of the day and does not continue into the next morning, so
			// it needs no second half. An empty [0, 0) half would collide with every
			// window starting at midnight.
			spans = append(spans, span{i, start, MinutesPerDay})
		default:
			crossings++
			spans = append(spans, span{i, start, MinutesPerDay}, span{i, 0, end})
		}
	}

	diags.Append(quietHoursOverlapErrors(spans, windows, attrPath)...)

	// A window written as crossing midnight becomes two spans. PostHog rejoins them only
	// when nothing else is on the timeline.
	if crossings > 0 && parsed > 1 {
		diags.AddAttributeError(
			attrPath,
			"Quiet-hours window crossing midnight must be the only window",
			"A window whose end is before its start crosses midnight, and PostHog stores it as one window "+
				"per side of midnight unless it is the only window configured. Either make it the only "+
				"window, or split it yourself into a window ending at 00:00 and one starting at 00:00. "+
				"Splitting costs a window slot, so PostHog's cap on the number of windows applies.",
		)
		return diags
	}

	// The rule below is the only one a skipped window can make wrong rather than
	// incomplete, because it keys off there being exactly two spans. The rules above can
	// only miss a problem, never invent one.
	if parsed != len(windows) {
		return diags
	}

	// Two windows blocking both sides of midnight are rejoined into one crossing window,
	// but only while they are the whole timeline. A third window leaves all of them stored
	// as written.
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

// quietHoursOverlapErrors reports windows that overlap or touch. PostHog merges on
// `next.start <= prev.end`, so 00:00-06:00 and 06:00-09:00 become one window.
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
