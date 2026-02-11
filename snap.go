package quando

import "time"

// StartOf returns a new Date snapped to the beginning of the specified unit.
// Time is set to 00:00:00.000 unless otherwise specified.
//
// Supported units:
//   - Week: Returns Monday 00:00:00 (ISO 8601 convention)
//   - Month: Returns 1st day of month, 00:00:00
//   - Quarter: Returns first day of quarter (Q1=Jan 1, Q2=Apr 1, Q3=Jul 1, Q4=Oct 1)
//   - Year: Returns Jan 1, 00:00:00
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC))
//	monday := date.StartOf(quando.Week)     // Feb 9, 2026 00:00:00 (Monday)
//	month := date.StartOf(quando.Month)     // Feb 1, 2026 00:00:00
//	quarter := date.StartOf(quando.Quarter) // Jan 1, 2026 00:00:00 (Q1)
//	year := date.StartOf(quando.Year)       // Jan 1, 2026 00:00:00
func (d Date) StartOf(unit Unit) Date {
	t := d.t
	loc := t.Location()

	switch unit {
	case Weeks:
		// Find Monday of current week (ISO 8601: Monday is day 1)
		// time.Weekday: Sunday=0, Monday=1, ..., Saturday=6
		weekday := int(t.Weekday())
		if weekday == 0 { // Sunday
			weekday = 7 // Treat Sunday as day 7 for ISO 8601
		}
		daysToMonday := weekday - 1
		mondayDate := t.AddDate(0, 0, -daysToMonday)
		result := time.Date(mondayDate.Year(), mondayDate.Month(), mondayDate.Day(), 0, 0, 0, 0, loc)
		return Date{t: result, lang: d.lang}

	case Months:
		// First day of month, 00:00:00
		result := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, loc)
		return Date{t: result, lang: d.lang}

	case Quarters:
		// Q1=Jan-Mar (start: Jan 1), Q2=Apr-Jun (start: Apr 1),
		// Q3=Jul-Sep (start: Jul 1), Q4=Oct-Dec (start: Oct 1)
		month := t.Month()
		var quarterStart time.Month
		switch {
		case month >= 1 && month <= 3:
			quarterStart = time.January
		case month >= 4 && month <= 6:
			quarterStart = time.April
		case month >= 7 && month <= 9:
			quarterStart = time.July
		default: // month >= 10 && month <= 12
			quarterStart = time.October
		}
		result := time.Date(t.Year(), quarterStart, 1, 0, 0, 0, 0, loc)
		return Date{t: result, lang: d.lang}

	case Years:
		// Jan 1, 00:00:00
		result := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, loc)
		return Date{t: result, lang: d.lang}

	default:
		// For other units, return the date unchanged
		return d
	}
}

// EndOf returns a new Date snapped to the end of the specified unit.
// Time is set to 23:59:59.999999999.
//
// Supported units:
//   - Week: Returns Sunday 23:59:59 (ISO 8601 convention)
//   - Month: Returns last day of month, 23:59:59 (handles all month lengths)
//   - Quarter: Returns last day of quarter, 23:59:59
//   - Year: Returns Dec 31, 23:59:59
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC))
//	sunday := date.EndOf(quando.Week)       // Feb 15, 2026 23:59:59 (Sunday)
//	monthEnd := date.EndOf(quando.Month)    // Feb 28, 2026 23:59:59
//	quarterEnd := date.EndOf(quando.Quarter) // Mar 31, 2026 23:59:59 (Q1)
//	yearEnd := date.EndOf(quando.Year)      // Dec 31, 2026 23:59:59
func (d Date) EndOf(unit Unit) Date {
	t := d.t
	loc := t.Location()

	switch unit {
	case Weeks:
		// Find Sunday of current week (ISO 8601: Sunday is day 7)
		// time.Weekday: Sunday=0, Monday=1, ..., Saturday=6
		weekday := int(t.Weekday())
		if weekday == 0 { // Sunday
			weekday = 7
		}
		daysToSunday := 7 - weekday
		sundayDate := t.AddDate(0, 0, daysToSunday)
		result := time.Date(sundayDate.Year(), sundayDate.Month(), sundayDate.Day(), 23, 59, 59, 999999999, loc)
		return Date{t: result, lang: d.lang}

	case Months:
		// Last day of month, 23:59:59
		// Strategy: Go to first day of next month, then subtract one day
		firstOfNextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, loc)
		lastOfMonth := firstOfNextMonth.AddDate(0, 0, -1)
		result := time.Date(lastOfMonth.Year(), lastOfMonth.Month(), lastOfMonth.Day(), 23, 59, 59, 999999999, loc)
		return Date{t: result, lang: d.lang}

	case Quarters:
		// Q1=Jan-Mar (end: Mar 31), Q2=Apr-Jun (end: Jun 30),
		// Q3=Jul-Sep (end: Sep 30), Q4=Oct-Dec (end: Dec 31)
		month := t.Month()
		var quarterEnd time.Month
		switch {
		case month >= 1 && month <= 3:
			quarterEnd = time.March
		case month >= 4 && month <= 6:
			quarterEnd = time.June
		case month >= 7 && month <= 9:
			quarterEnd = time.September
		default: // month >= 10 && month <= 12
			quarterEnd = time.December
		}
		// Get last day of quarter end month
		firstOfNextMonth := time.Date(t.Year(), quarterEnd+1, 1, 0, 0, 0, 0, loc)
		lastOfQuarter := firstOfNextMonth.AddDate(0, 0, -1)
		result := time.Date(lastOfQuarter.Year(), lastOfQuarter.Month(), lastOfQuarter.Day(), 23, 59, 59, 999999999, loc)
		return Date{t: result, lang: d.lang}

	case Years:
		// Dec 31, 23:59:59
		result := time.Date(t.Year(), time.December, 31, 23, 59, 59, 999999999, loc)
		return Date{t: result, lang: d.lang}

	default:
		// For other units, return the date unchanged
		return d
	}
}

// Next returns a new Date representing the next occurrence of the specified weekday.
// The time of day is preserved from the source date.
//
// IMPORTANT: Next ALWAYS returns a future date, never today. If today is the
// specified weekday, Next returns the same weekday next week (7 days later).
//
// Example:
//
//	// On Monday, Feb 9, 2026
//	date := quando.From(time.Date(2026, 2, 9, 15, 30, 0, 0, time.UTC)) // Monday
//	nextMonday := date.Next(time.Monday)    // Feb 16, 2026 15:30 (next Monday)
//	nextFriday := date.Next(time.Friday)    // Feb 13, 2026 15:30 (this Friday)
func (d Date) Next(weekday time.Weekday) Date {
	t := d.t
	currentWeekday := t.Weekday()

	// Calculate days until target weekday
	daysUntil := int(weekday - currentWeekday)
	if daysUntil <= 0 {
		// If target is today or in the past this week, jump to next week
		daysUntil += 7
	}

	result := t.AddDate(0, 0, daysUntil)
	return Date{t: result, lang: d.lang}
}

// Prev returns a new Date representing the previous occurrence of the specified weekday.
// The time of day is preserved from the source date.
//
// IMPORTANT: Prev ALWAYS returns a past date, never today. If today is the
// specified weekday, Prev returns the same weekday last week (7 days earlier).
//
// Example:
//
//	// On Monday, Feb 9, 2026
//	date := quando.From(time.Date(2026, 2, 9, 15, 30, 0, 0, time.UTC)) // Monday
//	prevMonday := date.Prev(time.Monday)    // Feb 2, 2026 15:30 (last Monday)
//	prevFriday := date.Prev(time.Friday)    // Feb 6, 2026 15:30 (last Friday)
func (d Date) Prev(weekday time.Weekday) Date {
	t := d.t
	currentWeekday := t.Weekday()

	// Calculate days until target weekday (going backwards)
	daysUntil := int(currentWeekday - weekday)
	if daysUntil <= 0 {
		// If target is today or in the future this week, jump to previous week
		daysUntil += 7
	}

	result := t.AddDate(0, 0, -daysUntil)
	return Date{t: result, lang: d.lang}
}
