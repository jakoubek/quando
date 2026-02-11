package quando

import (
	"testing"
	"time"
)

func TestAddSeconds(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"add 1 second", 1, "2026-01-01 12:00:01"},
		{"add 60 seconds", 60, "2026-01-01 12:01:00"},
		{"add 3600 seconds", 3600, "2026-01-01 13:00:00"},
		{"subtract 1 second", -1, "2026-01-01 11:59:59"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Add(tt.value, Seconds)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Seconds) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAddMinutes(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"add 1 minute", 1, "2026-01-01 12:01:00"},
		{"add 60 minutes", 60, "2026-01-01 13:00:00"},
		{"subtract 1 minute", -1, "2026-01-01 11:59:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Add(tt.value, Minutes)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Minutes) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAddHours(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"add 1 hour", 1, "2026-01-01 13:00:00"},
		{"add 24 hours", 24, "2026-01-02 12:00:00"},
		{"subtract 1 hour", -1, "2026-01-01 11:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Add(tt.value, Hours)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Hours) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAddDays(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"add 1 day", 1, "2026-01-02 12:00:00"},
		{"add 7 days", 7, "2026-01-08 12:00:00"},
		{"add 365 days", 365, "2027-01-01 12:00:00"},
		{"subtract 1 day", -1, "2025-12-31 12:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Add(tt.value, Days)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Days) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAddWeeks(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"add 1 week", 1, "2026-01-08 12:00:00"},
		{"add 4 weeks", 4, "2026-01-29 12:00:00"},
		{"subtract 1 week", -1, "2025-12-25 12:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Add(tt.value, Weeks)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Weeks) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestAddMonths(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		months   int
		expected string
	}{
		// Regular month addition (no overflow)
		{
			name:     "regular addition",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-02-15 12:00:00",
		},
		// Month-end overflow cases
		{
			name:     "Jan 31 + 1 month = Feb 28 (non-leap year)",
			date:     From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-02-28 12:00:00",
		},
		{
			name:     "Jan 31 + 1 month = Feb 29 (leap year)",
			date:     From(time.Date(2024, 1, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2024-02-29 12:00:00",
		},
		{
			name:     "May 31 + 1 month = Jun 30",
			date:     From(time.Date(2026, 5, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-06-30 12:00:00",
		},
		{
			name:     "Jul 31 + 1 month = Aug 31",
			date:     From(time.Date(2026, 7, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-08-31 12:00:00",
		},
		{
			name:     "Aug 31 + 1 month = Sep 30",
			date:     From(time.Date(2026, 8, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-09-30 12:00:00",
		},
		{
			name:     "Oct 31 + 1 month = Nov 30",
			date:     From(time.Date(2026, 10, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2026-11-30 12:00:00"},
		// Cross year boundary
		{
			name:     "Dec 15 + 1 month = Jan 15 next year",
			date:     From(time.Date(2026, 12, 15, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2027-01-15 12:00:00",
		},
		{
			name:     "Dec 31 + 1 month = Jan 31 next year",
			date:     From(time.Date(2026, 12, 31, 12, 0, 0, 0, time.UTC)),
			months:   1,
			expected: "2027-01-31 12:00:00",
		},
		// Multiple months
		{
			name:     "Jan 31 + 2 months = Mar 31",
			date:     From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC)),
			months:   2,
			expected: "2026-03-31 12:00:00",
		},
		{
			name:     "Jan 31 + 13 months = Feb 28 next year",
			date:     From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC)),
			months:   13,
			expected: "2027-02-28 12:00:00",
		},
		// Negative values (subtraction)
		{
			name:     "Mar 31 - 1 month = Feb 28",
			date:     From(time.Date(2026, 3, 31, 12, 0, 0, 0, time.UTC)),
			months:   -1,
			expected: "2026-02-28 12:00:00",
		},
		{
			name:     "Jan 15 - 1 month = Dec 15 prev year",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			months:   -1,
			expected: "2025-12-15 12:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.Add(tt.months, Months)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Months) = %v, want %v", tt.months, result, tt.expected)
			}
		})
	}
}

func TestAddQuarters(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		quarters int
		expected string
	}{
		{
			name:     "add 1 quarter",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			quarters: 1,
			expected: "2026-04-15 12:00:00",
		},
		{
			name:     "add 4 quarters (1 year)",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			quarters: 4,
			expected: "2027-01-15 12:00:00",
		},
		{
			name:     "quarter with month-end overflow",
			date:     From(time.Date(2026, 5, 31, 12, 0, 0, 0, time.UTC)),
			quarters: 1,
			expected: "2026-08-31 12:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.Add(tt.quarters, Quarters)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Quarters) = %v, want %v", tt.quarters, result, tt.expected)
			}
		})
	}
}

func TestAddYears(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		years    int
		expected string
	}{
		{
			name:     "add 1 year",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			years:    1,
			expected: "2027-01-15 12:00:00",
		},
		{
			name:     "add 10 years",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			years:    10,
			expected: "2036-01-15 12:00:00",
		},
		{
			name:     "leap year to non-leap year (Feb 29 -> Feb 28)",
			date:     From(time.Date(2024, 2, 29, 12, 0, 0, 0, time.UTC)),
			years:    1,
			expected: "2025-02-28 12:00:00",
		},
		{
			name:     "subtract 1 year",
			date:     From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC)),
			years:    -1,
			expected: "2025-01-15 12:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.Add(tt.years, Years)
			if result.String() != tt.expected {
				t.Errorf("Add(%d, Years) = %v, want %v", tt.years, result, tt.expected)
			}
		})
	}
}

func TestSub(t *testing.T) {
	date := From(time.Date(2026, 3, 31, 12, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		value    int
		unit     Unit
		expected string
	}{
		{"subtract 1 day", 1, Days, "2026-03-30 12:00:00"},
		{"subtract 1 month", 1, Months, "2026-02-28 12:00:00"},
		{"subtract 1 year", 1, Years, "2025-03-31 12:00:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Sub(tt.value, tt.unit)
			if result.String() != tt.expected {
				t.Errorf("Sub(%d, %v) = %v, want %v", tt.value, tt.unit, result, tt.expected)
			}
		})
	}
}

func TestAddSubEquivalence(t *testing.T) {
	date := From(time.Date(2026, 3, 15, 12, 0, 0, 0, time.UTC))

	// Sub(n) should equal Add(-n)
	units := []Unit{Seconds, Minutes, Hours, Days, Weeks, Months, Quarters, Years}

	for _, unit := range units {
		result1 := date.Sub(5, unit)
		result2 := date.Add(-5, unit)

		if !result1.Time().Equal(result2.Time()) {
			t.Errorf("Sub(5, %v) != Add(-5, %v): %v != %v", unit, unit, result1, result2)
		}
	}
}

func TestMethodChaining(t *testing.T) {
	date := From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	// Chain multiple operations
	result := date.
		Add(1, Months).
		Add(15, Days).
		Sub(2, Hours)

	expected := "2026-02-16 10:00:00"
	if result.String() != expected {
		t.Errorf("Chained operations = %v, want %v", result, expected)
	}
}

func TestImmutability(t *testing.T) {
	original := From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC))
	originalTime := original.Time()

	// Perform various operations
	_ = original.Add(1, Days)
	_ = original.Add(1, Months)
	_ = original.Sub(1, Years)

	// Verify original is unchanged
	if !original.Time().Equal(originalTime) {
		t.Error("Add/Sub operations modified the original date")
	}
}

// TestTimezonePreservation verifies that arithmetic preserves timezone
func TestTimezonePreservation(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("Skipping timezone test: %v", err)
	}

	berlinTime := time.Date(2026, 1, 15, 12, 0, 0, 0, loc)
	date := From(berlinTime)

	result := date.Add(1, Months)

	if result.Time().Location() != loc {
		t.Errorf("Add() location = %v, want %v", result.Time().Location(), loc)
	}
}

// BenchmarkAddDays benchmarks Add with Days
func BenchmarkAddDays(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Days)
	}
}

// BenchmarkAddMonths benchmarks Add with Months
func BenchmarkAddMonths(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Months)
	}
}

// BenchmarkAddYears benchmarks Add with Years
func BenchmarkAddYears(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Years)
	}
}

// BenchmarkMethodChaining benchmarks chained operations
func BenchmarkMethodChaining(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Months).Add(15, Days).Sub(2, Hours)
	}
}
