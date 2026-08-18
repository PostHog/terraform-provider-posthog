package core

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQuietHoursMinutes(t *testing.T) {
	tests := map[string]struct {
		in   string
		want int
		ok   bool
	}{
		"midnight":            {in: "00:00", want: 0, ok: true},
		"morning":             {in: "06:30", want: 390, ok: true},
		"last minute of day":  {in: "23:59", want: 1439, ok: true},
		"not a time":          {in: "nope"},
		"negative hour":       {in: "-1:30"},
		"hour out of range":   {in: "24:00"},
		"minute out of range": {in: "12:60"},
		"trailing junk":       {in: "12:30extra"},
		// time.Parse would read this as 09:30, which the pattern does not allow. Both
		// layers have to agree on what a valid time is.
		"single-digit hour": {in: "9:30"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, ok := QuietHoursMinutes(test.in)
			assert.Equal(t, test.ok, ok)
			if test.ok {
				assert.Equal(t, test.want, got)
			}
		})
	}
}

// The rule matrix lives here now that both resources share it. Every expectation was
// checked against PostHog's own normalizer, and the rejections against the live API.
func TestValidateQuietHoursWindows(t *testing.T) {
	const (
		overlaps  = "Quiet-hours windows overlap"
		crossesMN = "Quiet-hours window crossing midnight must be the only window"
		meetsAtMN = "Quiet-hours windows meeting at midnight are stored as one"
	)

	w := func(pairs ...[2]string) []QuietHoursWindow {
		out := make([]QuietHoursWindow, len(pairs))
		for i, p := range pairs {
			out[i] = QuietHoursWindow{Start: p[0], End: p[1]}
		}
		return out
	}

	tests := map[string]struct {
		windows       []QuietHoursWindow
		wantSummaries []string
	}{
		"single window":                             {windows: w([2]string{"01:00", "05:00"})},
		"lone crossing window":                      {windows: w([2]string{"22:00", "07:00"})},
		"ends at midnight plus daytime":             {windows: w([2]string{"19:00", "00:00"}, [2]string{"12:00", "13:00"})},
		"gap between windows":                       {windows: w([2]string{"01:00", "05:00"}, [2]string{"12:00", "13:00"})},
		"starts at 0 and ends at 23:59":             {windows: w([2]string{"00:00", "06:00"}, [2]string{"22:00", "23:59"})},
		"midnight pair with a third":                {windows: w([2]string{"22:00", "00:00"}, [2]string{"00:00", "07:00"}, [2]string{"12:00", "13:00"})},
		"exactly thirty minutes":                    {windows: w([2]string{"02:00", "02:30"})},
		"crossing window measured across midnight":  {windows: w([2]string{"23:50", "00:30"})},
		"unrecognised times are left to the schema": {windows: w([2]string{"nonsense", "06:00"}, [2]string{"05:00", "09:00"})},

		// Window length and count are PostHog's to enforce, so this is not a finding here.
		// The API rejects it on apply.
		"short window is left to the API": {windows: w([2]string{"02:00", "02:15"})},
		// Equal bounds block no time. Skipped, because 00:00-00:00 would otherwise read as
		// a whole day.
		"zero length is skipped": {windows: w([2]string{"02:00", "02:00"})},
		"touching":               {windows: w([2]string{"00:00", "06:00"}, [2]string{"06:00", "09:00"}), wantSummaries: []string{overlaps}},
		"overlapping":            {windows: w([2]string{"00:00", "06:00"}, [2]string{"05:00", "09:00"}), wantSummaries: []string{overlaps}},
		"contained":              {windows: w([2]string{"00:00", "09:00"}, [2]string{"02:00", "03:00"}), wantSummaries: []string{overlaps}},
		"crossing plus daytime":  {windows: w([2]string{"22:00", "07:00"}, [2]string{"12:00", "13:00"}), wantSummaries: []string{crossesMN}},
		"midnight pair alone":    {windows: w([2]string{"00:00", "06:00"}, [2]string{"22:00", "00:00"}), wantSummaries: []string{meetsAtMN}},
		"midnight pair reversed": {windows: w([2]string{"22:00", "00:00"}, [2]string{"00:00", "06:00"}), wantSummaries: []string{meetsAtMN}},
		// A whole-day block always touches, so the overlap rule already stops it.
		"whole day still caught as a reshape": {
			windows:       w([2]string{"00:00", "12:00"}, [2]string{"12:00", "00:00"}),
			wantSummaries: []string{overlaps, meetsAtMN},
		},

		// Both fire: the crossing window overlaps the morning one, and it is not alone.
		"crossing window overlaps morning": {
			windows:       w([2]string{"22:00", "07:00"}, [2]string{"06:00", "08:00"}),
			wantSummaries: []string{overlaps, crossesMN},
		},
		// Short windows join the timeline like any other, so they overlap normally.
		"short window overlaps its neighbour": {
			windows:       w([2]string{"02:00", "02:10"}, [2]string{"00:00", "06:00"}, [2]string{"22:00", "00:00"}),
			wantSummaries: []string{overlaps},
		},
		"short window does not hide an overlap": {
			windows:       w([2]string{"02:00", "02:10"}, [2]string{"08:00", "12:00"}, [2]string{"11:00", "14:00"}),
			wantSummaries: []string{overlaps},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			diags := ValidateQuietHoursWindows(test.windows, path.Root("blocked_windows"))

			var got []string
			for _, d := range diags.Errors() {
				got = append(got, d.Summary())
			}
			assert.Equal(t, test.wantSummaries, got)
		})
	}
}

// An empty set means quiet hours are off, which is always valid.
func TestValidateQuietHoursWindows_EmptyIsValid(t *testing.T) {
	diags := ValidateQuietHoursWindows(nil, path.Root("blocked_windows"))
	require.False(t, diags.HasError(), "unexpected diagnostics: %v", diags)
}
