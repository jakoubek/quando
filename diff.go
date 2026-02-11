package quando

import "time"

// Duration represents the difference between two dates.
// It provides methods to extract the duration in various units.
//
// Integer methods (Seconds, Minutes, Hours, Days, Weeks, Months, Years)
// return rounded-down values. For more precise calculations involving
// months and years, use MonthsFloat() and YearsFloat().
//
// Durations can be negative if the start date is after the end date.
type Duration struct {
	start time.Time
	end   time.Time
}

// Diff calculates the duration between two dates.
// If a is before b, the duration is positive. If a is after b, the duration is negative.
//
// Example:
//
//	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
//	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
//	duration := quando.Diff(start, end)
//	months := duration.Months() // 11 months
func Diff(a, b time.Time) Duration {
	return Duration{
		start: a,
		end:   b,
	}
}

// Seconds returns the number of seconds in the duration (rounded down).
func (d Duration) Seconds() int64 {
	return int64(d.end.Sub(d.start).Seconds())
}

// Minutes returns the number of minutes in the duration (rounded down).
func (d Duration) Minutes() int64 {
	return int64(d.end.Sub(d.start).Minutes())
}

// Hours returns the number of hours in the duration (rounded down).
func (d Duration) Hours() int64 {
	return int64(d.end.Sub(d.start).Hours())
}

// Days returns the number of days in the duration (rounded down).
// This calculates calendar days, not 24-hour periods.
func (d Duration) Days() int {
	start := d.start
	end := d.end
	negative := false

	// Handle negative duration
	if end.Before(start) {
		start, end = end, start
		negative = true
	}

	// Calculate days by counting calendar days
	days := 0
	for start.Before(end) {
		start = start.AddDate(0, 0, 1)
		days++
	}

	if negative {
		return -days
	}
	return days
}

// Weeks returns the number of weeks in the duration (rounded down).
func (d Duration) Weeks() int {
	return d.Days() / 7
}

// Months returns the number of months in the duration (rounded down).
// This handles month-end dates and leap years correctly.
func (d Duration) Months() int {
	start := d.start
	end := d.end
	negative := false

	// Handle negative duration
	if end.Before(start) {
		start, end = end, start
		negative = true
	}

	// Calculate months based on year and month difference
	years := end.Year() - start.Year()
	months := int(end.Month()) - int(start.Month())
	totalMonths := years*12 + months

	// Adjust if end day is before start day (not a full month yet)
	if end.Day() < start.Day() {
		totalMonths--
	}

	if negative {
		return -totalMonths
	}
	return totalMonths
}

// Years returns the number of years in the duration (rounded down).
func (d Duration) Years() int {
	return d.Months() / 12
}

// MonthsFloat returns the precise number of months in the duration as a float64.
// This provides more accurate calculations than the integer Months() method.
func (d Duration) MonthsFloat() float64 {
	start := d.start
	end := d.end
	negative := false

	// Handle negative duration
	if end.Before(start) {
		start, end = end, start
		negative = true
	}

	// Get integer months first
	fullMonths := float64(d.Months())
	if negative {
		fullMonths = -fullMonths
	}

	// Calculate fractional month based on days
	// Create a date at the same day in the month after start + full months
	baseDate := start.AddDate(0, int(fullMonths), 0)

	// Days remaining after full months
	daysRemaining := end.Sub(baseDate).Hours() / 24

	// Days in the month we're currently in
	nextMonth := baseDate.AddDate(0, 1, 0)
	daysInMonth := nextMonth.Sub(baseDate).Hours() / 24

	fractionalMonth := daysRemaining / daysInMonth

	result := fullMonths + fractionalMonth
	if negative {
		return -result
	}
	return result
}

// YearsFloat returns the precise number of years in the duration as a float64.
// This provides more accurate calculations than the integer Years() method.
func (d Duration) YearsFloat() float64 {
	return d.MonthsFloat() / 12.0
}
