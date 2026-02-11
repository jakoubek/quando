package quando

import (
	"testing"
	"time"
)

func TestStartOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string // Expected Monday
	}{
		{
			name:     "Monday stays Monday",
			date:     From(time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC)), // Monday
			expected: "2026-02-09 00:00:00",
		},
		{
			name:     "Tuesday goes to Monday",
			date:     From(time.Date(2026, 2, 10, 15, 30, 45, 0, time.UTC)), // Tuesday
			expected: "2026-02-09 00:00:00",
		},
		{
			name:     "Wednesday goes to Monday",
			date:     From(time.Date(2026, 2, 11, 15, 30, 45, 0, time.UTC)), // Wednesday
			expected: "2026-02-09 00:00:00",
		},
		{
			name:     "Sunday goes to Monday",
			date:     From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)), // Sunday
			expected: "2026-02-09 00:00:00",
		},
		{
			name:     "Saturday goes to Monday",
			date:     From(time.Date(2026, 2, 14, 15, 30, 45, 0, time.UTC)), // Saturday
			expected: "2026-02-09 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.StartOf(Weeks)
			if result.String() != tt.expected {
				t.Errorf("StartOf(Weeks) = %v, want %v", result, tt.expected)
			}

			// Verify it's Monday
			if result.Time().Weekday() != time.Monday {
				t.Errorf("StartOf(Weeks) weekday = %v, want Monday", result.Time().Weekday())
			}
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string // Expected Sunday
	}{
		{
			name:     "Monday goes to Sunday",
			date:     From(time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC)), // Monday
			expected: "2026-02-15 23:59:59",
		},
		{
			name:     "Sunday stays Sunday",
			date:     From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)), // Sunday
			expected: "2026-02-15 23:59:59",
		},
		{
			name:     "Saturday goes to Sunday",
			date:     From(time.Date(2026, 2, 14, 15, 30, 45, 0, time.UTC)), // Saturday
			expected: "2026-02-15 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.EndOf(Weeks)
			if result.String() != tt.expected {
				t.Errorf("EndOf(Weeks) = %v, want %v", result, tt.expected)
			}

			// Verify it's Sunday
			if result.Time().Weekday() != time.Sunday {
				t.Errorf("EndOf(Weeks) weekday = %v, want Sunday", result.Time().Weekday())
			}
		})
	}
}

func TestStartOfMonth(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "mid-month",
			date:     From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-02-01 00:00:00",
		},
		{
			name:     "first day",
			date:     From(time.Date(2026, 2, 1, 15, 30, 45, 0, time.UTC)),
			expected: "2026-02-01 00:00:00",
		},
		{
			name:     "last day",
			date:     From(time.Date(2026, 2, 28, 15, 30, 45, 0, time.UTC)),
			expected: "2026-02-01 00:00:00",
		},
		{
			name:     "31-day month",
			date:     From(time.Date(2026, 1, 31, 15, 30, 45, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.StartOf(Months)
			if result.String() != tt.expected {
				t.Errorf("StartOf(Months) = %v, want %v", result, tt.expected)
			}

			// Verify it's day 1
			if result.Time().Day() != 1 {
				t.Errorf("StartOf(Months) day = %d, want 1", result.Time().Day())
			}
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		name        string
		date        Date
		expectedDay int
	}{
		{
			name:        "February non-leap",
			date:        From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)),
			expectedDay: 28,
		},
		{
			name:        "February leap year",
			date:        From(time.Date(2024, 2, 15, 15, 30, 45, 0, time.UTC)),
			expectedDay: 29,
		},
		{
			name:        "30-day month (April)",
			date:        From(time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)),
			expectedDay: 30,
		},
		{
			name:        "31-day month (January)",
			date:        From(time.Date(2026, 1, 15, 15, 30, 45, 0, time.UTC)),
			expectedDay: 31,
		},
		{
			name:        "31-day month (December)",
			date:        From(time.Date(2026, 12, 15, 15, 30, 45, 0, time.UTC)),
			expectedDay: 31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.EndOf(Months)

			// Verify correct day
			if result.Time().Day() != tt.expectedDay {
				t.Errorf("EndOf(Months) day = %d, want %d", result.Time().Day(), tt.expectedDay)
			}

			// Verify time is 23:59:59
			if result.Time().Hour() != 23 || result.Time().Minute() != 59 || result.Time().Second() != 59 {
				t.Errorf("EndOf(Months) time = %02d:%02d:%02d, want 23:59:59",
					result.Time().Hour(), result.Time().Minute(), result.Time().Second())
			}
		})
	}
}

func TestStartOfQuarter(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "Q1 January",
			date:     From(time.Date(2026, 1, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "Q1 February",
			date:     From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "Q1 March",
			date:     From(time.Date(2026, 3, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "Q2 April",
			date:     From(time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-04-01 00:00:00",
		},
		{
			name:     "Q2 May",
			date:     From(time.Date(2026, 5, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-04-01 00:00:00",
		},
		{
			name:     "Q2 June",
			date:     From(time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-04-01 00:00:00",
		},
		{
			name:     "Q3 July",
			date:     From(time.Date(2026, 7, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-07-01 00:00:00",
		},
		{
			name:     "Q3 August",
			date:     From(time.Date(2026, 8, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-07-01 00:00:00",
		},
		{
			name:     "Q3 September",
			date:     From(time.Date(2026, 9, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-07-01 00:00:00",
		},
		{
			name:     "Q4 October",
			date:     From(time.Date(2026, 10, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-10-01 00:00:00",
		},
		{
			name:     "Q4 November",
			date:     From(time.Date(2026, 11, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-10-01 00:00:00",
		},
		{
			name:     "Q4 December",
			date:     From(time.Date(2026, 12, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-10-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.StartOf(Quarters)
			if result.String() != tt.expected {
				t.Errorf("StartOf(Quarters) = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEndOfQuarter(t *testing.T) {
	tests := []struct {
		name          string
		date          Date
		expectedMonth time.Month
		expectedDay   int
	}{
		{
			name:          "Q1 January",
			date:          From(time.Date(2026, 1, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.March,
			expectedDay:   31,
		},
		{
			name:          "Q1 February",
			date:          From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.March,
			expectedDay:   31,
		},
		{
			name:          "Q1 March",
			date:          From(time.Date(2026, 3, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.March,
			expectedDay:   31,
		},
		{
			name:          "Q2 April",
			date:          From(time.Date(2026, 4, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.June,
			expectedDay:   30,
		},
		{
			name:          "Q2 June",
			date:          From(time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.June,
			expectedDay:   30,
		},
		{
			name:          "Q3 July",
			date:          From(time.Date(2026, 7, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.September,
			expectedDay:   30,
		},
		{
			name:          "Q3 September",
			date:          From(time.Date(2026, 9, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.September,
			expectedDay:   30,
		},
		{
			name:          "Q4 October",
			date:          From(time.Date(2026, 10, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.December,
			expectedDay:   31,
		},
		{
			name:          "Q4 December",
			date:          From(time.Date(2026, 12, 15, 15, 30, 45, 0, time.UTC)),
			expectedMonth: time.December,
			expectedDay:   31,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.EndOf(Quarters)

			// Verify correct month and day
			if result.Time().Month() != tt.expectedMonth {
				t.Errorf("EndOf(Quarters) month = %v, want %v", result.Time().Month(), tt.expectedMonth)
			}
			if result.Time().Day() != tt.expectedDay {
				t.Errorf("EndOf(Quarters) day = %d, want %d", result.Time().Day(), tt.expectedDay)
			}

			// Verify time is 23:59:59
			if result.Time().Hour() != 23 || result.Time().Minute() != 59 || result.Time().Second() != 59 {
				t.Errorf("EndOf(Quarters) time = %02d:%02d:%02d, want 23:59:59",
					result.Time().Hour(), result.Time().Minute(), result.Time().Second())
			}
		})
	}
}

func TestStartOfYear(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "mid-year",
			date:     From(time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "start of year",
			date:     From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "end of year",
			date:     From(time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.StartOf(Years)
			if result.String() != tt.expected {
				t.Errorf("StartOf(Years) = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEndOfYear(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "mid-year",
			date:     From(time.Date(2026, 6, 15, 15, 30, 45, 0, time.UTC)),
			expected: "2026-12-31 23:59:59",
		},
		{
			name:     "start of year",
			date:     From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)),
			expected: "2026-12-31 23:59:59",
		},
		{
			name:     "end of year",
			date:     From(time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)),
			expected: "2026-12-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.EndOf(Years)
			if result.String() != tt.expected {
				t.Errorf("EndOf(Years) = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestSnapImmutability verifies that snap operations don't modify the original date
func TestSnapImmutability(t *testing.T) {
	original := From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC))
	originalTime := original.Time()

	_ = original.StartOf(Weeks)
	_ = original.EndOf(Weeks)
	_ = original.StartOf(Months)
	_ = original.EndOf(Months)
	_ = original.StartOf(Quarters)
	_ = original.EndOf(Quarters)
	_ = original.StartOf(Years)
	_ = original.EndOf(Years)

	if !original.Time().Equal(originalTime) {
		t.Error("Snap operations modified the original date")
	}
}

// TestSnapTimezones verifies that snap operations preserve timezone
func TestSnapTimezones(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("Skipping timezone test: %v", err)
	}

	berlinTime := time.Date(2026, 2, 15, 15, 30, 45, 0, loc)
	date := From(berlinTime)

	operations := []struct {
		name   string
		result Date
	}{
		{"StartOf(Weeks)", date.StartOf(Weeks)},
		{"EndOf(Weeks)", date.EndOf(Weeks)},
		{"StartOf(Months)", date.StartOf(Months)},
		{"EndOf(Months)", date.EndOf(Months)},
		{"StartOf(Quarters)", date.StartOf(Quarters)},
		{"EndOf(Quarters)", date.EndOf(Quarters)},
		{"StartOf(Years)", date.StartOf(Years)},
		{"EndOf(Years)", date.EndOf(Years)},
	}

	for _, op := range operations {
		t.Run(op.name, func(t *testing.T) {
			if op.result.Time().Location() != loc {
				t.Errorf("%s location = %v, want %v", op.name, op.result.Time().Location(), loc)
			}
		})
	}
}

// BenchmarkStartOfWeek benchmarks StartOf(Weeks)
func BenchmarkStartOfWeek(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Weeks)
	}
}

// BenchmarkEndOfWeek benchmarks EndOf(Weeks)
func BenchmarkEndOfWeek(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Weeks)
	}
}

// BenchmarkStartOfMonth benchmarks StartOf(Months)
func BenchmarkStartOfMonth(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Months)
	}
}

// BenchmarkEndOfMonth benchmarks EndOf(Months)
func BenchmarkEndOfMonth(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Months)
	}
}

// BenchmarkStartOfQuarter benchmarks StartOf(Quarters)
func BenchmarkStartOfQuarter(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Quarters)
	}
}

// BenchmarkEndOfQuarter benchmarks EndOf(Quarters)
func BenchmarkEndOfQuarter(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Quarters)
	}
}

// BenchmarkStartOfYear benchmarks StartOf(Years)
func BenchmarkStartOfYear(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Years)
	}
}

// BenchmarkEndOfYear benchmarks EndOf(Years)
func BenchmarkEndOfYear(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Years)
	}
}
