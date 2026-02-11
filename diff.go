package quando

import (
	"fmt"
	"time"
)

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

// Human returns a human-readable representation of the duration.
// It shows the two largest relevant time units for adaptive granularity.
//
// If no language is specified, English (EN) is used by default.
//
// Examples:
//
//	dur := quando.Diff(start, end)
//	dur.Human()           // "10 months, 16 days" (English default)
//	dur.Human(quando.DE)  // "10 Monate, 16 Tage" (German)
//
// Adaptive granularity examples:
//   - 10 months, 16 days → "10 months, 16 days"
//   - 2 days, 5 hours → "2 days, 5 hours"
//   - 3 hours, 20 minutes → "3 hours, 20 minutes"
//   - 45 seconds → "45 seconds"
//   - 0 → "0 seconds"
func (d Duration) Human(lang ...Lang) string {
	// Default to English if no language specified
	l := EN
	if len(lang) > 0 {
		l = lang[0]
	}

	// Handle negative durations
	negative := d.start.After(d.end)

	// Calculate all time components (using absolute values)
	totalSeconds := d.Seconds()
	if totalSeconds < 0 {
		totalSeconds = -totalSeconds
	}

	// Calculate months and years
	months := d.Months()
	if months < 0 {
		months = -months
	}
	years := months / 12
	remainingMonths := months % 12

	// Calculate remaining components after extracting larger units
	// After years and months, calculate remaining days
	// We need to subtract the time represented by years and months from total

	// Start from the beginning and add years + months
	baseTime := d.start
	if negative {
		baseTime = d.end
	}

	afterYearsMonths := baseTime.AddDate(years, remainingMonths, 0)

	// Calculate remaining time
	var remainingEnd time.Time
	if negative {
		remainingEnd = d.start
	} else {
		remainingEnd = d.end
	}

	remainingDuration := remainingEnd.Sub(afterYearsMonths)
	remainingDays := int(remainingDuration.Hours() / 24)
	remainingHours := int(remainingDuration.Hours()) % 24
	remainingMinutes := int(remainingDuration.Minutes()) % 60
	remainingSeconds := int(remainingDuration.Seconds()) % 60

	// Build component list with values
	type component struct {
		value int
		unit  string
	}

	components := []component{
		{years, "year"},
		{remainingMonths, "month"},
		{remainingDays, "day"},
		{remainingHours, "hour"},
		{remainingMinutes, "minute"},
		{remainingSeconds, "second"},
	}

	// Filter to non-zero components
	var nonZero []component
	for _, c := range components {
		if c.value > 0 {
			nonZero = append(nonZero, c)
		}
	}

	// Handle zero duration special case
	if len(nonZero) == 0 {
		return "0 " + l.DurationUnit("second", true)
	}

	// Take up to 2 largest units for adaptive granularity
	displayUnits := nonZero
	if len(displayUnits) > 2 {
		displayUnits = displayUnits[:2]
	}

	// Build the output string
	var parts []string
	for _, c := range displayUnits {
		unitName := l.DurationUnit(c.unit, c.value != 1)
		parts = append(parts, fmt.Sprintf("%d %s", c.value, unitName))
	}

	result := ""
	if len(parts) == 1 {
		result = parts[0]
	} else if len(parts) == 2 {
		result = parts[0] + ", " + parts[1]
	}

	// Add negative prefix if needed
	if negative {
		// Use minus sign for simplicity (could be localized in future)
		result = "-" + result
	}

	return result
}
