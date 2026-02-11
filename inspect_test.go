package quando

import (
	"testing"
	"time"
)

func TestWeekNumber(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		// Test all 7 days of week 7 in 2026 (Feb 9-15)
		{
			name: "2026-02-09 Monday week 7",
			date: time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-10 Tuesday week 7",
			date: time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-11 Wednesday week 7",
			date: time.Date(2026, 2, 11, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-12 Thursday week 7",
			date: time.Date(2026, 2, 12, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-13 Friday week 7",
			date: time.Date(2026, 2, 13, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-14 Saturday week 7",
			date: time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC),
			want: 7,
		},
		{
			name: "2026-02-15 Sunday week 7",
			date: time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC),
			want: 7,
		},

		// Year boundary cases
		{
			name: "2023-01-01 Sunday belongs to 2022 week 52",
			date: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			want: 52,
		},
		{
			name: "2023-01-02 Monday is week 1 of 2023",
			date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "2024-01-01 Monday is week 1",
			date: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "2025-01-01 Wednesday is week 1",
			date: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "2026-01-01 Thursday is week 1",
			date: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},

		// Week 53 scenarios
		{
			name: "2020-12-31 Thursday is week 53",
			date: time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			want: 53,
		},
		{
			name: "2026-12-31 Thursday is week 53",
			date: time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
			want: 53,
		},

		// Mid-year dates with known week numbers
		{
			name: "2026-06-15 Monday is week 25",
			date: time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC),
			want: 25,
		},
		{
			name: "2026-12-28 Monday is week 53",
			date: time.Date(2026, 12, 28, 0, 0, 0, 0, time.UTC),
			want: 53,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(tt.date)
			got := d.WeekNumber()
			if got != tt.want {
				t.Errorf("WeekNumber() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestQuarter(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		// Test all 12 months
		{"January Q1", time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC), 1},
		{"February Q1", time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC), 1},
		{"March Q1", time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC), 1},
		{"April Q2", time.Date(2026, 4, 15, 0, 0, 0, 0, time.UTC), 2},
		{"May Q2", time.Date(2026, 5, 15, 0, 0, 0, 0, time.UTC), 2},
		{"June Q2", time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC), 2},
		{"July Q3", time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC), 3},
		{"August Q3", time.Date(2026, 8, 15, 0, 0, 0, 0, time.UTC), 3},
		{"September Q3", time.Date(2026, 9, 15, 0, 0, 0, 0, time.UTC), 3},
		{"October Q4", time.Date(2026, 10, 15, 0, 0, 0, 0, time.UTC), 4},
		{"November Q4", time.Date(2026, 11, 15, 0, 0, 0, 0, time.UTC), 4},
		{"December Q4", time.Date(2026, 12, 15, 0, 0, 0, 0, time.UTC), 4},

		// Quarter boundaries
		{"March 31 last day of Q1", time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC), 1},
		{"April 1 first day of Q2", time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC), 2},
		{"June 30 last day of Q2", time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC), 2},
		{"July 1 first day of Q3", time.Date(2026, 7, 1, 0, 0, 0, 0, time.UTC), 3},
		{"September 30 last day of Q3", time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC), 3},
		{"October 1 first day of Q4", time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), 4},
		{"December 31 last day of Q4", time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC), 4},

		// Leap year Feb 29
		{"Feb 29 leap year Q1", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(tt.date)
			got := d.Quarter()
			if got != tt.want {
				t.Errorf("Quarter() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestDayOfYear(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		// Jan 1 is always day 1
		{"Jan 1 is day 1", time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), 1},

		// Dec 31 in non-leap year is day 365
		{"Dec 31 non-leap year is 365", time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC), 365},
		{"Dec 31 2025 is 365", time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC), 365},

		// Dec 31 in leap year is day 366
		{"Dec 31 leap year is 366", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 366},
		{"Dec 31 2020 is 366", time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC), 366},

		// Feb 29 in leap year is day 60
		{"Feb 29 leap year is 60", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), 60},
		{"Feb 29 2020 is 60", time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC), 60},

		// Random dates with known values
		{"Feb 9 2026 is 40", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC), 40},
		{"Mar 1 2026 is 60", time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC), 60},
		{"Mar 1 2024 leap year is 61", time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC), 61},
		{"Jun 15 2026 is 166", time.Date(2026, 6, 15, 0, 0, 0, 0, time.UTC), 166},
		{"Dec 25 2026 is 359", time.Date(2026, 12, 25, 0, 0, 0, 0, time.UTC), 359},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(tt.date)
			got := d.DayOfYear()
			if got != tt.want {
				t.Errorf("DayOfYear() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want bool
	}{
		// Test all 7 weekdays (Feb 9-15, 2026)
		{"Monday not weekend", time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC), false},
		{"Tuesday not weekend", time.Date(2026, 2, 10, 0, 0, 0, 0, time.UTC), false},
		{"Wednesday not weekend", time.Date(2026, 2, 11, 0, 0, 0, 0, time.UTC), false},
		{"Thursday not weekend", time.Date(2026, 2, 12, 0, 0, 0, 0, time.UTC), false},
		{"Friday not weekend", time.Date(2026, 2, 13, 0, 0, 0, 0, time.UTC), false},
		{"Saturday is weekend", time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC), true},
		{"Sunday is weekend", time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC), true},

		// Year boundaries on weekends
		{"Jan 1 2023 Sunday is weekend", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), true},
		{"Dec 31 2022 Saturday is weekend", time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC), true},
		{"Jan 1 2024 Monday not weekend", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), false},
		{"Dec 31 2024 Tuesday not weekend", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(tt.date)
			got := d.IsWeekend()
			if got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		// Regular leap years (divisible by 4, not by 100)
		{"2024 is leap year", 2024, true},
		{"2020 is leap year", 2020, true},
		{"2028 is leap year", 2028, true},
		{"2004 is leap year", 2004, true},

		// Non-leap years
		{"2026 not leap year", 2026, false},
		{"2025 not leap year", 2025, false},
		{"2023 not leap year", 2023, false},
		{"2001 not leap year", 2001, false},

		// Century rules (divisible by 100 but not 400)
		{"1900 not leap year (century)", 1900, false},
		{"2100 not leap year (century)", 2100, false},
		{"2200 not leap year (century)", 2200, false},
		{"2300 not leap year (century)", 2300, false},

		// Century rules (divisible by 400)
		{"2000 is leap year (400 rule)", 2000, true},
		{"2400 is leap year (400 rule)", 2400, true},
		{"1600 is leap year (400 rule)", 1600, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(time.Date(tt.year, 1, 1, 0, 0, 0, 0, time.UTC))
			got := d.IsLeapYear()
			if got != tt.want {
				t.Errorf("IsLeapYear() for year %d = %v, want %v", tt.year, got, tt.want)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want DateInfo
	}{
		{
			name: "2026-02-09 Monday",
			date: time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC),
			want: DateInfo{
				WeekNumber: 7,
				Quarter:    1,
				DayOfYear:  40,
				IsWeekend:  false,
				IsLeapYear: false,
				Unix:       time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC).Unix(),
			},
		},
		{
			name: "2024-02-29 leap year Thursday",
			date: time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
			want: DateInfo{
				WeekNumber: 9,
				Quarter:    1,
				DayOfYear:  60,
				IsWeekend:  false,
				IsLeapYear: true,
				Unix:       time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC).Unix(),
			},
		},
		{
			name: "2026-12-31 Thursday week 53",
			date: time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			want: DateInfo{
				WeekNumber: 53,
				Quarter:    4,
				DayOfYear:  365,
				IsWeekend:  false,
				IsLeapYear: false,
				Unix:       time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC).Unix(),
			},
		},
		{
			name: "2026-06-14 Sunday weekend Q2",
			date: time.Date(2026, 6, 14, 0, 0, 0, 0, time.UTC),
			want: DateInfo{
				WeekNumber: 24,
				Quarter:    2,
				DayOfYear:  165,
				IsWeekend:  true,
				IsLeapYear: false,
				Unix:       time.Date(2026, 6, 14, 0, 0, 0, 0, time.UTC).Unix(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := From(tt.date)
			got := d.Info()

			if got.WeekNumber != tt.want.WeekNumber {
				t.Errorf("Info().WeekNumber = %d, want %d", got.WeekNumber, tt.want.WeekNumber)
			}
			if got.Quarter != tt.want.Quarter {
				t.Errorf("Info().Quarter = %d, want %d", got.Quarter, tt.want.Quarter)
			}
			if got.DayOfYear != tt.want.DayOfYear {
				t.Errorf("Info().DayOfYear = %d, want %d", got.DayOfYear, tt.want.DayOfYear)
			}
			if got.IsWeekend != tt.want.IsWeekend {
				t.Errorf("Info().IsWeekend = %v, want %v", got.IsWeekend, tt.want.IsWeekend)
			}
			if got.IsLeapYear != tt.want.IsLeapYear {
				t.Errorf("Info().IsLeapYear = %v, want %v", got.IsLeapYear, tt.want.IsLeapYear)
			}
			if got.Unix != tt.want.Unix {
				t.Errorf("Info().Unix = %d, want %d", got.Unix, tt.want.Unix)
			}
		})
	}
}

// TestInfo_ConsistentWithIndividualMethods verifies that Info() returns
// the same values as calling each method individually
func TestInfo_ConsistentWithIndividualMethods(t *testing.T) {
	dates := []time.Time{
		time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC),
		time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
	}

	for _, date := range dates {
		t.Run(date.Format("2006-01-02"), func(t *testing.T) {
			d := From(date)
			info := d.Info()

			if info.WeekNumber != d.WeekNumber() {
				t.Errorf("Info().WeekNumber inconsistent: got %d, individual method %d", info.WeekNumber, d.WeekNumber())
			}
			if info.Quarter != d.Quarter() {
				t.Errorf("Info().Quarter inconsistent: got %d, individual method %d", info.Quarter, d.Quarter())
			}
			if info.DayOfYear != d.DayOfYear() {
				t.Errorf("Info().DayOfYear inconsistent: got %d, individual method %d", info.DayOfYear, d.DayOfYear())
			}
			if info.IsWeekend != d.IsWeekend() {
				t.Errorf("Info().IsWeekend inconsistent: got %v, individual method %v", info.IsWeekend, d.IsWeekend())
			}
			if info.IsLeapYear != d.IsLeapYear() {
				t.Errorf("Info().IsLeapYear inconsistent: got %v, individual method %v", info.IsLeapYear, d.IsLeapYear())
			}
			if info.Unix != d.Unix() {
				t.Errorf("Info().Unix inconsistent: got %d, individual method %d", info.Unix, d.Unix())
			}
		})
	}
}

