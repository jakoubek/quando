package quando

import "time"

// DateInfo contains aggregated metadata about a date.
// Returned by the Info() method.
type DateInfo struct {
	WeekNumber int   // ISO 8601 week number (1-53)
	Quarter    int   // Fiscal quarter (1-4)
	DayOfYear  int   // Day of year (1-366)
	IsWeekend  bool  // True if Saturday or Sunday
	IsLeapYear bool  // True if date is in a leap year
	Unix       int64 // Unix timestamp
}

// WeekNumber returns the ISO 8601 week number (1-53).
//
// ISO 8601 definition:
//   - Week 1 is the first week with a Thursday in it
//   - Weeks start on Monday and end on Sunday
//   - Dates in early January may belong to week 52/53 of previous year
//   - Dates in late December may belong to week 1 of next year
//
// Examples:
//   - Jan 1, 2026 (Thursday) is in week 1 of 2026
//   - Jan 1, 2025 (Wednesday) is in week 1 of 2025
//   - Jan 1, 2024 (Monday) is in week 1 of 2024
//   - Jan 1, 2023 (Sunday) is in week 52 of 2022
//
// Performance: < 1 µs, zero allocations
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
//	week := date.WeekNumber() // 7
func (d Date) WeekNumber() int {
	_, week := d.t.ISOWeek()
	return week
}

// Quarter returns the fiscal quarter (1-4) for the date.
//
// Quarter mapping:
//   - Q1: January, February, March
//   - Q2: April, May, June
//   - Q3: July, August, September
//   - Q4: October, November, December
//
// Performance: < 1 µs, zero allocations
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
//	q := date.Quarter() // 1 (Q1)
func (d Date) Quarter() int {
	month := d.t.Month()
	switch {
	case month >= 1 && month <= 3:
		return 1
	case month >= 4 && month <= 6:
		return 2
	case month >= 7 && month <= 9:
		return 3
	default: // month >= 10 && month <= 12
		return 4
	}
}

// DayOfYear returns the day of the year (1-366).
// Also known as "ordinal date" or "Julian day number".
//
// January 1 = 1, December 31 = 365 (or 366 in leap years)
//
// Performance: < 1 µs, zero allocations
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
//	day := date.DayOfYear() // 40 (Feb 9)
func (d Date) DayOfYear() int {
	return d.t.YearDay()
}

// IsWeekend returns true if the date falls on a weekend (Saturday or Sunday).
//
// Note: This uses the ISO convention where Saturday and Sunday are considered
// weekend days. This is not configurable in Phase 1.
//
// Performance: < 1 µs, zero allocations
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)) // Monday
//	isWeekend := date.IsWeekend() // false
func (d Date) IsWeekend() bool {
	weekday := d.t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsLeapYear returns true if the date's year is a leap year.
//
// Leap year rules:
//   - Divisible by 4: leap year (e.g., 2024)
//   - EXCEPT divisible by 100: not a leap year (e.g., 1900, 2100)
//   - EXCEPT divisible by 400: leap year (e.g., 2000, 2400)
//
// Performance: < 1 µs, zero allocations
//
// Example:
//
//	date := quando.From(time.Date(2024, 2, 9, 0, 0, 0, 0, time.UTC))
//	isLeap := date.IsLeapYear() // true (2024 is a leap year)
func (d Date) IsLeapYear() bool {
	year := d.t.Year()

	// Apply leap year rules
	if year%400 == 0 {
		return true // Divisible by 400: leap year
	}
	if year%100 == 0 {
		return false // Divisible by 100 but not 400: not leap year
	}
	if year%4 == 0 {
		return true // Divisible by 4 but not 100: leap year
	}
	return false // Not divisible by 4: not leap year
}

// Info returns aggregated metadata about the date.
//
// This is a convenience method that calls all inspection methods
// and packages the results into a single struct.
//
// Performance: < 1 µs (sum of all individual methods)
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
//	info := date.Info()
//	fmt.Printf("Week: %d, Quarter: %d\n", info.WeekNumber, info.Quarter)
func (d Date) Info() DateInfo {
	return DateInfo{
		WeekNumber: d.WeekNumber(),
		Quarter:    d.Quarter(),
		DayOfYear:  d.DayOfYear(),
		IsWeekend:  d.IsWeekend(),
		IsLeapYear: d.IsLeapYear(),
		Unix:       d.Unix(),
	}
}
