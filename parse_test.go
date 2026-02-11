package quando

import (
	"errors"
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

// BenchmarkParse benchmarks the Parse function with different formats
func BenchmarkParse(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"ISO format", "2026-02-09"},
		{"ISO slash", "2026/02/09"},
		{"EU format", "09.02.2026"},
		{"RFC2822", "Mon, 09 Feb 2026 00:00:00 +0000"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := Parse(bm.input)
				if err != nil {
					b.Fatalf("Parse failed: %v", err)
				}
			}
		})
	}
}

// BenchmarkParseError benchmarks error case (ambiguous format)
func BenchmarkParseError(b *testing.B) {
	input := "01/02/2026" // Ambiguous format

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := Parse(input)
		if err == nil {
			b.Fatal("Expected error for ambiguous format")
		}
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
