package quando

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		// ISO format (YYYY-MM-DD)
		{"ISO: basic date", "2026-02-09", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"ISO: year end", "2024-12-31", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"ISO: year start", "2020-01-01", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"ISO: leap year", "2024-02-29", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
		{"ISO: month boundary", "2026-06-30", time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},

		// ISO with slash (YYYY/MM/DD)
		{"ISO slash: basic date", "2026/02/09", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"ISO slash: year end", "2024/12/31", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"ISO slash: year start", "2020/01/01", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"ISO slash: leap year", "2024/02/29", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},

		// EU format (DD.MM.YYYY)
		{"EU: basic date", "09.02.2026", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"EU: year end", "31.12.2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{"EU: year start", "01.01.2020", time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"EU: leap year", "29.02.2024", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
		{"EU: month boundary", "30.06.2026", time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) unexpected error: %v", tt.input, err)
			}

			// Compare the time values
			if !result.Time().Equal(tt.expected) {
				t.Errorf("Parse(%q) = %v, want %v", tt.input, result.Time(), tt.expected)
			}

			// Verify default language is EN
			if result.lang != EN {
				t.Errorf("Parse(%q) lang = %v, want %v", tt.input, result.lang, EN)
			}
		})
	}
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantError bool
		checkMsg  string // substring to check in error message
	}{
		// Ambiguous formats
		{"ambiguous: 01/02/2026", "01/02/2026", true, "ambiguous"},
		{"ambiguous: 31/12/2024", "31/12/2024", true, "ambiguous"},
		{"ambiguous: 15/06/2025", "15/06/2025", true, "ambiguous"},
		{"ambiguous: 01/01/2020", "01/01/2020", true, "ambiguous"},

		// Invalid formats
		{"invalid: not a date", "not-a-date", true, ""},
		{"invalid: empty string", "", true, "empty"},
		{"invalid: only whitespace", "   ", true, ""},
		{"invalid: incomplete date", "2026-02", true, ""},
		{"invalid: wrong separator", "2026_02_09", true, ""},
		{"invalid: extra characters", "2026-02-09 extra", true, ""},

		// Invalid date components
		{"invalid: month 13", "2026-13-01", true, ""},
		{"invalid: month 00", "2026-00-01", true, ""},
		{"invalid: day 00", "2026-02-00", true, ""},
		{"invalid: day 32", "2026-01-32", true, ""},
		{"invalid: Feb 30", "2026-02-30", true, ""},
		{"invalid: non-leap year Feb 29", "2023-02-29", true, ""},
		{"invalid: April 31", "2026-04-31", true, ""},

		// EU format invalid dates
		{"invalid EU: Feb 30", "30.02.2026", true, ""},
		{"invalid EU: month 13", "01.13.2026", true, ""},

		// Edge cases
		{"invalid: just numbers", "20260209", true, ""},
		{"invalid: wrong length", "26-02-09", true, ""},
		{"invalid: mixed separators", "2026-02/09", true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)

			if tt.wantError {
				if err == nil {
					t.Fatalf("Parse(%q) expected error, got nil (result: %v)", tt.input, result)
				}

				// Verify it's the right error type
				if !errors.Is(err, ErrInvalidFormat) {
					t.Errorf("Parse(%q) error = %v, want error wrapping ErrInvalidFormat", tt.input, err)
				}

				// Check for specific message substring if provided
				if tt.checkMsg != "" && !containsSubstring(err.Error(), tt.checkMsg) {
					t.Errorf("Parse(%q) error message %q does not contain %q", tt.input, err.Error(), tt.checkMsg)
				}
			} else {
				if err != nil {
					t.Fatalf("Parse(%q) unexpected error: %v", tt.input, err)
				}
			}
		})
	}
}

func TestParseRFC2822(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{
			"RFC2822: basic",
			"Mon, 09 Feb 2026 00:00:00 +0000",
			time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			"RFC2822: with time",
			"Mon, 09 Feb 2026 15:30:45 +0000",
			time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC),
		},
		{
			"RFC1123: different month",
			"Fri, 31 Dec 2024 23:59:59 GMT",
			time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) unexpected error: %v", tt.input, err)
			}

			if !result.Time().Equal(tt.expected) {
				t.Errorf("Parse(%q) = %v, want %v", tt.input, result.Time(), tt.expected)
			}
		})
	}
}

func TestParseImmutability(t *testing.T) {
	input := "2026-02-09"

	// Parse twice
	date1, err1 := Parse(input)
	date2, err2 := Parse(input)

	if err1 != nil || err2 != nil {
		t.Fatalf("Parse(%q) unexpected errors: %v, %v", input, err1, err2)
	}

	// Verify they're equal
	if !date1.Time().Equal(date2.Time()) {
		t.Errorf("Parse(%q) produced different results: %v vs %v", input, date1, date2)
	}

	// Modify date1 by adding time
	modified := date1.Add(1, Days)

	// Verify original is unchanged
	if !date1.Time().Equal(date2.Time()) {
		t.Errorf("Modifying result of Parse affected original date")
	}

	// Verify modification worked
	expected := date1.Time().AddDate(0, 0, 1)
	if !modified.Time().Equal(expected) {
		t.Errorf("Add operation failed: got %v, want %v", modified.Time(), expected)
	}
}

func TestParseWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{"leading whitespace", "  2026-02-09", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"trailing whitespace", "2026-02-09  ", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"both whitespace", "  2026-02-09  ", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"tab whitespace", "\t2026-02-09\t", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Parse(%q) unexpected error: %v", tt.input, err)
			}

			if !result.Time().Equal(tt.expected) {
				t.Errorf("Parse(%q) = %v, want %v", tt.input, result.Time(), tt.expected)
			}
		})
	}
}


// containsSubstring is a helper function to check if a string contains a substring
func containsSubstring(s, substr string) bool {
	return len(substr) == 0 || len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsSubstringHelper(s, substr))
}

func containsSubstringHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestParseWithLayout(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		layout      string
		expectYear  int
		expectMonth time.Month
		expectDay   int
	}{
		// Disambiguating US vs EU slash formats
		{
			name:        "US format 01/02/2026 -> Jan 2",
			input:       "01/02/2026",
			layout:      "01/02/2006",
			expectYear:  2026,
			expectMonth: time.January,
			expectDay:   2,
		},
		{
			name:        "EU format 01/02/2026 -> Feb 1",
			input:       "01/02/2026",
			layout:      "02/01/2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   1,
		},
		{
			name:        "EU format 31/12/2025",
			input:       "31/12/2025",
			layout:      "02/01/2006",
			expectYear:  2025,
			expectMonth: time.December,
			expectDay:   31,
		},

		// Custom formats
		{
			name:        "Custom format with English month name",
			input:       "9. February 2026",
			layout:      "2. January 2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},
		{
			name:        "Custom format with short month",
			input:       "15-Mar-2026",
			layout:      "02-Jan-2006",
			expectYear:  2026,
			expectMonth: time.March,
			expectDay:   15,
		},

		// ISO 8601 with time
		{
			name:        "ISO 8601 with time",
			input:       "2026-02-09T14:30:00",
			layout:      "2006-01-02T15:04:05",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},

		// Different separators
		{
			name:        "Dash format MM-DD-YYYY",
			input:       "02-09-2026",
			layout:      "01-02-2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},
		{
			name:        "Space separator",
			input:       "09 02 2026",
			layout:      "02 01 2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},

		// Whitespace handling
		{
			name:        "Leading whitespace",
			input:       "  09.02.2026",
			layout:      "02.01.2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},
		{
			name:        "Trailing whitespace",
			input:       "09.02.2026  ",
			layout:      "02.01.2006",
			expectYear:  2026,
			expectMonth: time.February,
			expectDay:   9,
		},

		// Edge cases
		{
			name:        "Leap year Feb 29",
			input:       "29/02/2024",
			layout:      "02/01/2006",
			expectYear:  2024,
			expectMonth: time.February,
			expectDay:   29,
		},
		{
			name:        "Year boundary",
			input:       "01/01/2026",
			layout:      "02/01/2006",
			expectYear:  2026,
			expectMonth: time.January,
			expectDay:   1,
		},
		{
			name:        "Year end",
			input:       "31/12/2026",
			layout:      "02/01/2006",
			expectYear:  2026,
			expectMonth: time.December,
			expectDay:   31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, err := ParseWithLayout(tt.input, tt.layout)
			if err != nil {
				t.Fatalf("ParseWithLayout(%q, %q) unexpected error: %v", tt.input, tt.layout, err)
			}

			tm := date.Time()
			if tm.Year() != tt.expectYear {
				t.Errorf("Year = %d, want %d", tm.Year(), tt.expectYear)
			}
			if tm.Month() != tt.expectMonth {
				t.Errorf("Month = %v, want %v", tm.Month(), tt.expectMonth)
			}
			if tm.Day() != tt.expectDay {
				t.Errorf("Day = %d, want %d", tm.Day(), tt.expectDay)
			}

			// Verify default language is set
			if date.lang != EN {
				t.Errorf("lang = %v, want EN", date.lang)
			}
		})
	}
}

func TestParseWithLayoutErrors(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		layout string
	}{
		{
			name:   "Empty input",
			input:  "",
			layout: "02/01/2006",
		},
		{
			name:   "Whitespace only",
			input:  "   ",
			layout: "02/01/2006",
		},
		{
			name:   "Invalid date for layout",
			input:  "99/99/2026",
			layout: "02/01/2006",
		},
		{
			name:   "Wrong layout for input",
			input:  "2026-02-09",
			layout: "02/01/2006",
		},
		{
			name:   "Invalid month",
			input:  "15/13/2026",
			layout: "02/01/2006",
		},
		{
			name:   "Invalid day",
			input:  "32/01/2026",
			layout: "02/01/2006",
		},
		{
			name:   "Feb 30 (invalid)",
			input:  "30/02/2026",
			layout: "02/01/2006",
		},
		{
			name:   "Feb 29 on non-leap year",
			input:  "29/02/2026",
			layout: "02/01/2006",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseWithLayout(tt.input, tt.layout)
			if err == nil {
				t.Errorf("ParseWithLayout(%q, %q) expected error, got nil", tt.input, tt.layout)
			}

			// Verify error wraps ErrInvalidFormat
			if !errors.Is(err, ErrInvalidFormat) {
				t.Errorf("error should wrap ErrInvalidFormat, got: %v", err)
			}
		})
	}
}

func TestParseWithLayoutImmutability(t *testing.T) {
	// Parse the same date twice with the same layout
	date1, err1 := ParseWithLayout("01/02/2026", "01/02/2006")
	if err1 != nil {
		t.Fatalf("ParseWithLayout failed: %v", err1)
	}

	date2, err2 := ParseWithLayout("01/02/2026", "01/02/2006")
	if err2 != nil {
		t.Fatalf("ParseWithLayout failed: %v", err2)
	}

	// Modify date1
	modified := date1.Add(5, Days)

	// Verify date2 is unchanged
	if date2.Unix() != date1.Unix() {
		t.Error("date2 should not be affected by operations on date1")
	}
	if modified.Unix() == date1.Unix() {
		t.Error("Add should return a new Date instance")
	}
}

func TestParseRelativeWithClock(t *testing.T) {
	// Fixed clock for deterministic testing
	fixedTime := time.Date(2026, 2, 15, 14, 30, 45, 0, time.UTC) // Saturday, Feb 15, 2:30 PM
	clock := NewFixedClock(fixedTime)

	tests := []struct {
		name          string
		input         string
		expectedYear  int
		expectedMonth time.Month
		expectedDay   int
	}{
		// Simple keywords
		{
			name:          "today",
			input:         "today",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   15,
		},
		{
			name:          "tomorrow",
			input:         "tomorrow",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   16,
		},
		{
			name:          "yesterday",
			input:         "yesterday",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   14,
		},

		// Case-insensitive keywords
		{
			name:          "TODAY (uppercase)",
			input:         "TODAY",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   15,
		},
		{
			name:          "Tomorrow (mixed case)",
			input:         "Tomorrow",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   16,
		},

		// Relative offsets - days
		{
			name:          "+1 day",
			input:         "+1 day",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   16,
		},
		{
			name:          "+2 days",
			input:         "+2 days",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   17,
		},
		{
			name:          "-1 day",
			input:         "-1 day",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   14,
		},
		{
			name:          "+7 days",
			input:         "+7 days",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   22,
		},

		// Relative offsets - weeks
		{
			name:          "+1 week",
			input:         "+1 week",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   22,
		},
		{
			name:          "+2 weeks",
			input:         "+2 weeks",
			expectedYear:  2026,
			expectedMonth: time.March,
			expectedDay:   1,
		},
		{
			name:          "-1 week",
			input:         "-1 week",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   8,
		},

		// Relative offsets - months
		{
			name:          "+1 month",
			input:         "+1 month",
			expectedYear:  2026,
			expectedMonth: time.March,
			expectedDay:   15,
		},
		{
			name:          "+3 months",
			input:         "+3 months",
			expectedYear:  2026,
			expectedMonth: time.May,
			expectedDay:   15,
		},
		{
			name:          "-1 month",
			input:         "-1 month",
			expectedYear:  2026,
			expectedMonth: time.January,
			expectedDay:   15,
		},

		// Relative offsets - quarters
		{
			name:          "+1 quarter",
			input:         "+1 quarter",
			expectedYear:  2026,
			expectedMonth: time.May,
			expectedDay:   15,
		},
		{
			name:          "-1 quarter",
			input:         "-1 quarter",
			expectedYear:  2025,
			expectedMonth: time.November,
			expectedDay:   15,
		},

		// Relative offsets - years
		{
			name:          "+1 year",
			input:         "+1 year",
			expectedYear:  2027,
			expectedMonth: time.February,
			expectedDay:   15,
		},
		{
			name:          "-1 year",
			input:         "-1 year",
			expectedYear:  2025,
			expectedMonth: time.February,
			expectedDay:   15,
		},

		// Whitespace variations
		{
			name:          "leading whitespace",
			input:         "  today",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   15,
		},
		{
			name:          "trailing whitespace",
			input:         "+2 days  ",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   17,
		},
		{
			name:          "extra spaces between parts",
			input:         "+2   days",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   17,
		},

		// Case-insensitive units
		{
			name:          "+2 DAYS (uppercase)",
			input:         "+2 DAYS",
			expectedYear:  2026,
			expectedMonth: time.February,
			expectedDay:   17,
		},
		{
			name:          "+1 Month (mixed case)",
			input:         "+1 Month",
			expectedYear:  2026,
			expectedMonth: time.March,
			expectedDay:   15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date, err := ParseRelativeWithClock(tt.input, clock)
			if err != nil {
				t.Fatalf("ParseRelativeWithClock(%q) unexpected error: %v", tt.input, err)
			}

			tm := date.Time()
			if tm.Year() != tt.expectedYear {
				t.Errorf("Year = %d, want %d", tm.Year(), tt.expectedYear)
			}
			if tm.Month() != tt.expectedMonth {
				t.Errorf("Month = %v, want %v", tm.Month(), tt.expectedMonth)
			}
			if tm.Day() != tt.expectedDay {
				t.Errorf("Day = %d, want %d", tm.Day(), tt.expectedDay)
			}

			// Verify time is 00:00:00 (StartOf(Days) behavior)
			if tm.Hour() != 0 || tm.Minute() != 0 || tm.Second() != 0 {
				t.Errorf("Time should be 00:00:00, got %02d:%02d:%02d", tm.Hour(), tm.Minute(), tm.Second())
			}

			// Verify default language is set
			if date.lang != EN {
				t.Errorf("lang = %v, want EN", date.lang)
			}
		})
	}
}

func TestParseRelativeErrors(t *testing.T) {
	clock := NewFixedClock(time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"whitespace only", "   "},
		{"unknown keyword", "now"},
		{"unknown keyword 2", "currently"},
		{"invalid format - no sign", "2 days"},
		{"invalid format - wrong parts", "+2"},
		{"invalid format - too many parts", "+2 days ago"},
		{"invalid offset - not a number", "+abc days"},
		{"invalid offset - float", "+1.5 days"},
		{"invalid unit", "+2 fortnights"},
		{"invalid unit 2", "+1 decade"},
		{"invalid unit 3", "+3 hours"}, // hours not supported in Phase 1
		{"sign without number", "+ days"},
		{"number without unit", "+2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseRelativeWithClock(tt.input, clock)
			if err == nil {
				t.Errorf("ParseRelativeWithClock(%q) expected error, got nil", tt.input)
			}

			// Verify error wraps ErrInvalidFormat
			if !errors.Is(err, ErrInvalidFormat) {
				t.Errorf("error should wrap ErrInvalidFormat, got: %v", err)
			}
		})
	}
}

func TestParseRelative(t *testing.T) {
	// Test the production function (uses system clock)
	// Just verify it doesn't error on valid inputs
	validInputs := []string{
		"today",
		"tomorrow",
		"yesterday",
		"+1 day",
		"-2 weeks",
		"+3 months",
	}

	for _, input := range validInputs {
		t.Run(input, func(t *testing.T) {
			_, err := ParseRelative(input)
			if err != nil {
				t.Errorf("ParseRelative(%q) unexpected error: %v", input, err)
			}
		})
	}
}

func TestParseRelativeImmutability(t *testing.T) {
	clock := NewFixedClock(time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC))

	// Parse the same expression twice
	date1, err1 := ParseRelativeWithClock("tomorrow", clock)
	date2, err2 := ParseRelativeWithClock("tomorrow", clock)

	if err1 != nil || err2 != nil {
		t.Fatalf("ParseRelativeWithClock failed: %v, %v", err1, err2)
	}

	// Modify date1
	modified := date1.Add(5, Days)

	// Verify date2 is unchanged
	if date2.Unix() != date1.Unix() {
		t.Error("date2 should not be affected by operations on date1")
	}
	if modified.Unix() == date1.Unix() {
		t.Error("Add should return a new Date instance")
	}
}


func TestMustParse_Success(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Time
	}{
		{"ISO format", "2026-02-09", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"ISO with slash", "2026/02/09", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"EU format", "09.02.2026", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
		{"RFC2822", "Sun, 09 Feb 2026 00:00:00 +0000", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MustParse(tt.input)

			if !result.Time().Equal(tt.expected) {
				t.Errorf("MustParse(%q) = %v, want %v", tt.input, result.Time(), tt.expected)
			}
		})
	}
}

func TestMustParse_Panic(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty string", ""},
		{"invalid format", "not-a-date"},
		{"ambiguous slash format", "01/02/2026"},
		{"invalid date", "2026-13-01"},
		{"malformed", "2026-02-"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("MustParse(%q) expected panic, but didn't panic", tt.input)
				}
			}()

			// This should panic
			_ = MustParse(tt.input)
		})
	}
}

func TestMustParse_PanicMessage(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprint(r)
			// Verify panic message includes the input string
			if !containsSubstring(msg, "invalid-input") {
				t.Errorf("panic message %q should include input string %q", msg, "invalid-input")
			}
			if !containsSubstring(msg, "MustParse") {
				t.Errorf("panic message %q should include function name %q", msg, "MustParse")
			}
		} else {
			t.Error("expected panic but didn't panic")
		}
	}()

	MustParse("invalid-input")
}

// TestIsYearPrefix_EdgeCases tests edge cases for the isYearPrefix function
func TestIsYearPrefix_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid year 2026", "2026", true},
		{"valid year 1000", "1000", true},
		{"invalid year 0999", "0999", false}, // Starts with 0
		{"invalid year 999", "999", false},   // Too short
		{"invalid non-digits", "abcd", false},
		{"invalid mixed", "20a6", false},
		{"empty string", "", false},
		{"year 9999", "9999", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isYearPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("isYearPrefix(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
