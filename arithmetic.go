package quando

import "time"

// Add adds the specified number of units to the date and returns a new Date.
// The original date is not modified (immutability).
//
// Supported units: Seconds, Minutes, Hours, Days, Weeks, Months, Quarters, Years
//
// Month-End Overflow Behavior:
//
// When adding months (or quarters/years which add months internally), if the target
// day doesn't exist in the destination month, the date is snapped to the last day
// of that month instead of overflowing into the next month.
//
// Examples:
//   - 2026-01-31 + 1 month  = 2026-02-28 (February has only 28 days in 2026)
//   - 2026-01-24 + 1 month  = 2026-02-24 (regular addition, day exists)
//   - 2026-05-31 + 1 month  = 2026-06-30 (June has only 30 days)
//   - 2024-01-31 + 1 month  = 2024-02-29 (leap year, February has 29 days)
//
// DST Handling:
//
// Adding Days means "same time on the next calendar day", not "24 hours later".
// This ensures that operations work intuitively across DST transitions.
//
// Example:
//   - 2026-03-31 02:00 CET + 1 Day = 2026-04-01 02:00 CEST (not 03:00)
//
// Negative Values:
//
// Negative values are supported and equivalent to subtraction:
//   - Add(-1, Months) is the same as Sub(1, Months)
//
// Example:
//
//	date := quando.From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC))
//	result := date.Add(1, quando.Months) // 2026-02-28 12:00:00
func (d Date) Add(value int, unit Unit) Date {
	t := d.t

	switch unit {
	case Seconds:
		t = t.Add(time.Duration(value) * time.Second)

	case Minutes:
		t = t.Add(time.Duration(value) * time.Minute)

	case Hours:
		t = t.Add(time.Duration(value) * time.Hour)

	case Days:
		// Use AddDate for calendar days (DST-safe)
		t = t.AddDate(0, 0, value)

	case Weeks:
		// 1 week = 7 days
		t = t.AddDate(0, 0, value*7)

	case Months:
		// Add months with month-end overflow handling
		t = addMonthsWithOverflow(t, value)

	case Quarters:
		// 1 quarter = 3 months
		t = addMonthsWithOverflow(t, value*3)

	case Years:
		// 1 year = 12 months
		t = addMonthsWithOverflow(t, value*12)
	}

	return Date{t: t, lang: d.lang}
}

// Sub subtracts the specified number of units from the date and returns a new Date.
// This is equivalent to Add with a negative value.
//
// Example:
//
//	date := quando.From(time.Date(2026, 3, 31, 12, 0, 0, 0, time.UTC))
//	result := date.Sub(1, quando.Months) // 2026-02-28 12:00:00 (month-end snap)
func (d Date) Sub(value int, unit Unit) Date {
	return d.Add(-value, unit)
}

// addMonthsWithOverflow adds months to a time.Time with month-end overflow handling.
// If the target day doesn't exist in the destination month, it snaps to the last day.
func addMonthsWithOverflow(t time.Time, months int) time.Time {
	// Calculate target year and month
	year := t.Year()
	month := int(t.Month()) + months
	day := t.Day()

	// Handle year overflow/underflow
	for month > 12 {
		year++
		month -= 12
	}
	for month < 1 {
		year--
		month += 12
	}

	targetMonth := time.Month(month)

	// Get the last day of the target month
	// Strategy: First day of next month minus 1 day
	firstOfNextMonth := time.Date(year, targetMonth+1, 1, 0, 0, 0, 0, t.Location())
	lastDayOfTargetMonth := firstOfNextMonth.AddDate(0, 0, -1).Day()

	// If target day exceeds last day of month, snap to last day
	if day > lastDayOfTargetMonth {
		day = lastDayOfTargetMonth
	}

	// Construct the result date with the same time of day
	return time.Date(year, targetMonth, day,
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}
