package quando

import (
	"math"
	"strings"
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	start := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 18, 0, 0, 0, time.UTC)

	dur := Diff(start, end)

	if dur.start != start {
		t.Errorf("Diff() start = %v, want %v", dur.start, start)
	}
	if dur.end != end {
		t.Errorf("Diff() end = %v, want %v", dur.end, end)
	}
}

func TestDurationSeconds(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int64
	}{
		{
			name:     "1 minute",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 1, 0, 0, time.UTC),
			expected: 60,
		},
		{
			name:     "1 hour",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC),
			expected: 3600,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -3600,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Seconds()
			if result != tt.expected {
				t.Errorf("Seconds() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationMinutes(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int64
	}{
		{
			name:     "1 hour",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC),
			expected: 60,
		},
		{
			name:     "1 day",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 1440,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -1440,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Minutes()
			if result != tt.expected {
				t.Errorf("Minutes() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationHours(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int64
	}{
		{
			name:     "1 day",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 24,
		},
		{
			name:     "1 week",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: 168,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Hours()
			if result != tt.expected {
				t.Errorf("Hours() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationDays(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			name:     "1 day",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "7 days",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: 7,
		},
		{
			name:     "365 days",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 365,
		},
		{
			name:     "leap year (366 days)",
			start:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 366,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 1, 8, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Days()
			if result != tt.expected {
				t.Errorf("Days() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationWeeks(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			name:     "1 week",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 8, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "4 weeks",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 29, 0, 0, 0, 0, time.UTC),
			expected: 4,
		},
		{
			name:     "52 weeks",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 52,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 1, 29, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Weeks()
			if result != tt.expected {
				t.Errorf("Weeks() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationMonths(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			name:     "1 month",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "11 months",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
			expected: 11,
		},
		{
			name:     "12 months",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 12,
		},
		{
			name:     "13 months",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 13,
		},
		{
			name:     "month-end to month-end",
			start:    time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 2, 28, 0, 0, 0, 0, time.UTC),
			expected: 0, // Less than a full month
		},
		{
			name:     "month-end across full month",
			start:    time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "leap year February",
			start:    time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "negative duration",
			start:    time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Months()
			if result != tt.expected {
				t.Errorf("Months() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationYears(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			name:     "1 year",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "2 years",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2028, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: 2,
		},
		{
			name:     "11 months (less than 1 year)",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 12, 1, 0, 0, 0, 0, time.UTC),
			expected: 0,
		},
		{
			name:     "13 months (1 year)",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 2, 1, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			name:     "negative duration",
			start:    time.Date(2028, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Years()
			if result != tt.expected {
				t.Errorf("Years() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDurationMonthsFloat(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		minExpect float64
		maxExpect float64
	}{
		{
			name:      "1 month exactly",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC),
			minExpect: 1.0,
			maxExpect: 1.0,
		},
		{
			name:      "1.5 months approximately",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2026, 2, 16, 0, 0, 0, 0, time.UTC),
			minExpect: 1.4,
			maxExpect: 1.6,
		},
		{
			name:      "2.5 years approximately",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2028, 7, 1, 0, 0, 0, 0, time.UTC),
			minExpect: 30.0,
			maxExpect: 30.1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.MonthsFloat()
			if result < tt.minExpect || result > tt.maxExpect {
				t.Errorf("MonthsFloat() = %f, want between %f and %f", result, tt.minExpect, tt.maxExpect)
			}
		})
	}
}

func TestDurationYearsFloat(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		minExpect float64
		maxExpect float64
	}{
		{
			name:      "1 year exactly",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			minExpect: 1.0,
			maxExpect: 1.0,
		},
		{
			name:      "1.5 years approximately",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2027, 7, 1, 0, 0, 0, 0, time.UTC),
			minExpect: 1.49,
			maxExpect: 1.51,
		},
		{
			name:      "2.5 years approximately",
			start:     time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:       time.Date(2028, 7, 1, 0, 0, 0, 0, time.UTC),
			minExpect: 2.49,
			maxExpect: 2.51,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.YearsFloat()
			if result < tt.minExpect || result > tt.maxExpect {
				t.Errorf("YearsFloat() = %f, want between %f and %f", result, tt.minExpect, tt.maxExpect)
			}
		})
	}
}

// TestDurationNegative verifies negative duration handling
func TestDurationNegative(t *testing.T) {
	start := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)

	if dur.Seconds() >= 0 {
		t.Error("Seconds() should be negative")
	}
	if dur.Days() >= 0 {
		t.Error("Days() should be negative")
	}
	if dur.Months() >= 0 {
		t.Error("Months() should be negative")
	}
	if dur.Years() >= 0 {
		t.Error("Years() should be negative")
	}
	if dur.MonthsFloat() >= 0 {
		t.Error("MonthsFloat() should be negative")
	}
	if dur.YearsFloat() >= 0 {
		t.Error("YearsFloat() should be negative")
	}
}

// TestDurationCrossBoundaries tests calculations across year boundaries
func TestDurationCrossBoundaries(t *testing.T) {
	tests := []struct {
		name   string
		start  time.Time
		end    time.Time
		months int
		years  int
	}{
		{
			name:   "cross one year boundary",
			start:  time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC),
			months: 3,
			years:  0,
		},
		{
			name:   "cross two year boundaries",
			start:  time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2027, 2, 1, 0, 0, 0, 0, time.UTC),
			months: 15,
			years:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			if dur.Months() != tt.months {
				t.Errorf("Months() = %d, want %d", dur.Months(), tt.months)
			}
			if dur.Years() != tt.years {
				t.Errorf("Years() = %d, want %d", dur.Years(), tt.years)
			}
		})
	}
}


// TestFloatPrecision verifies that float methods provide better precision
func TestFloatPrecision(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 2, 16, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)

	intMonths := dur.Months()
	floatMonths := dur.MonthsFloat()

	// Integer should be 1
	if intMonths != 1 {
		t.Errorf("Months() = %d, want 1", intMonths)
	}

	// Float should be more than 1 (around 1.5)
	if floatMonths <= 1.0 || floatMonths >= 2.0 {
		t.Errorf("MonthsFloat() = %f, expected between 1.0 and 2.0", floatMonths)
	}

	// Float should have fractional part
	if floatMonths == math.Floor(floatMonths) {
		t.Error("MonthsFloat() should have fractional part")
	}
}

func TestDurationHuman(t *testing.T) {
	tests := []struct {
		name     string
		start    time.Time
		end      time.Time
		lang     Lang
		expected string
	}{
		// English tests
		{
			name:     "10 months 16 days EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 11, 17, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "10 months, 16 days",
		},
		{
			name:     "2 days 5 hours EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 3, 5, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "2 days, 5 hours",
		},
		{
			name:     "3 hours 20 minutes EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 3, 20, 0, 0, time.UTC),
			lang:     EN,
			expected: "3 hours, 20 minutes",
		},
		{
			name:     "45 seconds EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 45, 0, time.UTC),
			lang:     EN,
			expected: "45 seconds",
		},
		{
			name:     "zero duration EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "0 seconds",
		},
		{
			name:     "1 year 2 months EN (singular/plural mix)",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 3, 1, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "1 year, 2 months",
		},
		{
			name:     "1 day EN (singular)",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "1 day",
		},

		// German tests
		{
			name:     "10 months 16 days DE",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 11, 17, 0, 0, 0, 0, time.UTC),
			lang:     DE,
			expected: "10 Monate, 16 Tage",
		},
		{
			name:     "2 days 5 hours DE",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 3, 5, 0, 0, 0, time.UTC),
			lang:     DE,
			expected: "2 Tage, 5 Stunden",
		},
		{
			name:     "3 hours 20 minutes DE",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 3, 20, 0, 0, time.UTC),
			lang:     DE,
			expected: "3 Stunden, 20 Minuten",
		},
		{
			name:     "45 seconds DE",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 45, 0, time.UTC),
			lang:     DE,
			expected: "45 Sekunden",
		},
		{
			name:     "zero duration DE",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			lang:     DE,
			expected: "0 Sekunden",
		},

		// Negative duration tests
		{
			name:     "negative 2 days EN",
			start:    time.Date(2026, 1, 3, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "-2 days",
		},
		{
			name:     "negative 1 month DE",
			start:    time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			lang:     DE,
			expected: "-1 Monat",
		},

		// Edge cases
		{
			name:     "exactly 1 year EN",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
			lang:     EN,
			expected: "1 year",
		},
		{
			name:     "1 minute 30 seconds EN (only shows 2 largest)",
			start:    time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2026, 1, 1, 0, 1, 30, 0, time.UTC),
			lang:     EN,
			expected: "1 minute, 30 seconds",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dur := Diff(tt.start, tt.end)
			result := dur.Human(tt.lang)
			if result != tt.expected {
				t.Errorf("Human(%v) = %q, want %q", tt.lang, result, tt.expected)
			}
		})
	}
}

func TestDurationHumanDefaultLanguage(t *testing.T) {
	// Test that Human() without argument defaults to English
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 3, 5, 0, 0, 0, time.UTC)
	dur := Diff(start, end)

	result := dur.Human()
	expected := "2 days, 5 hours"

	if result != expected {
		t.Errorf("Human() = %q, want %q (should default to English)", result, expected)
	}
}

func TestDurationHumanAdaptiveGranularity(t *testing.T) {
	// Test that only the 2 largest units are shown
	// 1 year, 2 months, 3 days should show "1 year, 2 months"
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2027, 3, 4, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)

	result := dur.Human(EN)
	// Should only show years and months, not days
	if !strings.Contains(result, "year") || !strings.Contains(result, "month") {
		t.Errorf("Human() = %q, should contain 'year' and 'month'", result)
	}
	if strings.Contains(result, "day") {
		t.Errorf("Human() = %q, should NOT contain 'day' (only 2 largest units)", result)
	}
}

