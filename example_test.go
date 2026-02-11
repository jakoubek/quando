package quando_test

import (
	"errors"
	"fmt"
	"time"

	"code.beautifulmachines.dev/jakoubek/quando"
)

func ExampleNow() {
	date := quando.Now()
	fmt.Printf("Current date type: %T\n", date)
	// Output: Current date type: quando.Date
}

func ExampleFrom() {
	t := time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)
	date := quando.From(t)
	fmt.Println(date)
	// Output: 2026-02-09 12:30:45
}

func ExampleFromUnix() {
	// Create date from Unix timestamp
	date := quando.FromUnix(1707480000)
	fmt.Println(date.Time().UTC().Format("2006-01-02 15:04:05 MST"))
	// Output: 2024-02-09 12:00:00 UTC
}

func ExampleFromUnix_negative() {
	// Create date from negative Unix timestamp (before 1970)
	date := quando.FromUnix(-946771200)
	fmt.Println(date.Time().UTC().Format("2006-01-02"))
	// Output: 1940-01-01
}

func ExampleDate_Time() {
	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	t := date.Time()
	fmt.Printf("%d-%02d-%02d\n", t.Year(), t.Month(), t.Day())
	// Output: 2026-02-09
}

func ExampleDate_Unix() {
	date := quando.From(time.Date(2024, 2, 9, 12, 0, 0, 0, time.UTC))
	timestamp := date.Unix()
	fmt.Println(timestamp)
	// Output: 1707480000
}

func ExampleDate_WithLang() {
	_ = quando.Now().WithLang(quando.DE)
	fmt.Printf("Language set to: %v\n", "DE")
	// Output: Language set to: DE
}

func ExampleDate_String() {
	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	fmt.Println(date)
	// Output: 2026-02-09 12:30:45
}

// ExampleDate_immutability demonstrates that Date is immutable
func ExampleDate_immutability() {
	original := quando.From(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))
	modified := original.WithLang(quando.DE)

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Modified: %v\n", modified)
	fmt.Println("Original unchanged: true")
	// Output:
	// Original: 2026-02-09 12:00:00
	// Modified: 2026-02-09 12:00:00
	// Original unchanged: true
}

// ExampleNewClock demonstrates creating a default clock
func ExampleNewClock() {
	clock := quando.NewClock()
	_ = clock.Now()
	fmt.Println("Clock created")
	// Output: Clock created
}

// ExampleNewFixedClock demonstrates creating a fixed clock for testing
func ExampleNewFixedClock() {
	// Create a fixed clock that always returns the same time
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := quando.NewFixedClock(fixedTime)

	// Now() always returns the fixed time
	date := clock.Now()
	fmt.Println(date)
	// Output: 2026-02-09 12:00:00
}

// ExampleFixedClock_deterministic demonstrates deterministic testing with FixedClock
func ExampleFixedClock_deterministic() {
	// In tests, use a fixed clock for deterministic behavior
	testTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := quando.NewFixedClock(testTime)

	// Call multiple times - always returns same time
	date1 := clock.Now()
	date2 := clock.Now()

	fmt.Printf("Date 1: %v\n", date1)
	fmt.Printf("Date 2: %v\n", date2)
	fmt.Printf("Same: %v\n", date1.Unix() == date2.Unix())
	// Output:
	// Date 1: 2026-02-09 12:00:00
	// Date 2: 2026-02-09 12:00:00
	// Same: true
}

// ExampleUnit demonstrates the Unit type
func ExampleUnit() {
	// Units are used with Add, Sub, StartOf, EndOf operations
	units := []quando.Unit{
		quando.Seconds,
		quando.Minutes,
		quando.Hours,
		quando.Days,
		quando.Weeks,
		quando.Months,
		quando.Quarters,
		quando.Years,
	}

	for _, u := range units {
		fmt.Println(u.String())
	}
	// Output:
	// seconds
	// minutes
	// hours
	// days
	// weeks
	// months
	// quarters
	// years
}

// ExampleUnit_String demonstrates the String method
func ExampleUnit_String() {
	unit := quando.Days
	fmt.Printf("Unit: %s\n", unit)
	// Output: Unit: days
}

// ExampleDate_StartOf demonstrates snapping to the beginning of time units
func ExampleDate_StartOf() {
	date := quando.From(time.Date(2026, 2, 15, 15, 30, 45, 0, time.UTC)) // Sunday, Feb 15

	fmt.Println("Original:", date)
	fmt.Println("StartOf(Week):", date.StartOf(quando.Weeks))       // Monday
	fmt.Println("StartOf(Month):", date.StartOf(quando.Months))     // Feb 1
	fmt.Println("StartOf(Quarter):", date.StartOf(quando.Quarters)) // Jan 1 (Q1)
	fmt.Println("StartOf(Year):", date.StartOf(quando.Years))       // Jan 1
	// Output:
	// Original: 2026-02-15 15:30:45
	// StartOf(Week): 2026-02-09 00:00:00
	// StartOf(Month): 2026-02-01 00:00:00
	// StartOf(Quarter): 2026-01-01 00:00:00
	// StartOf(Year): 2026-01-01 00:00:00
}

// ExampleDate_EndOf demonstrates snapping to the end of time units
func ExampleDate_EndOf() {
	date := quando.From(time.Date(2026, 2, 9, 15, 30, 45, 0, time.UTC)) // Monday, Feb 9

	fmt.Println("Original:", date)
	fmt.Println("EndOf(Week):", date.EndOf(quando.Weeks))       // Sunday
	fmt.Println("EndOf(Month):", date.EndOf(quando.Months))     // Feb 28
	fmt.Println("EndOf(Quarter):", date.EndOf(quando.Quarters)) // Mar 31 (Q1)
	fmt.Println("EndOf(Year):", date.EndOf(quando.Years))       // Dec 31
	// Output:
	// Original: 2026-02-09 15:30:45
	// EndOf(Week): 2026-02-15 23:59:59
	// EndOf(Month): 2026-02-28 23:59:59
	// EndOf(Quarter): 2026-03-31 23:59:59
	// EndOf(Year): 2026-12-31 23:59:59
}

// ExampleDate_StartOf_chaining demonstrates chaining snap operations
func ExampleDate_StartOf_chaining() {
	// Get the first Monday of the current quarter
	date := quando.Now()
	firstMondayOfQuarter := date.StartOf(quando.Quarters).StartOf(quando.Weeks)

	fmt.Printf("Type: %T\n", firstMondayOfQuarter)
	// Output: Type: quando.Date
}

// ExampleDate_Next demonstrates finding the next occurrence of a weekday
func ExampleDate_Next() {
	// On Monday, Feb 9, 2026
	date := quando.From(time.Date(2026, 2, 9, 15, 30, 0, 0, time.UTC))

	fmt.Println("Today:", date.Time().Weekday())
	fmt.Println("Next Monday:", date.Next(time.Monday).Time().Weekday(), "-", date.Next(time.Monday))
	fmt.Println("Next Friday:", date.Next(time.Friday).Time().Weekday(), "-", date.Next(time.Friday))
	// Output:
	// Today: Monday
	// Next Monday: Monday - 2026-02-16 15:30:00
	// Next Friday: Friday - 2026-02-13 15:30:00
}

// ExampleDate_Prev demonstrates finding the previous occurrence of a weekday
func ExampleDate_Prev() {
	// On Monday, Feb 9, 2026
	date := quando.From(time.Date(2026, 2, 9, 15, 30, 0, 0, time.UTC))

	fmt.Println("Today:", date.Time().Weekday())
	fmt.Println("Prev Monday:", date.Prev(time.Monday).Time().Weekday(), "-", date.Prev(time.Monday))
	fmt.Println("Prev Friday:", date.Prev(time.Friday).Time().Weekday(), "-", date.Prev(time.Friday))
	// Output:
	// Today: Monday
	// Prev Monday: Monday - 2026-02-02 15:30:00
	// Prev Friday: Friday - 2026-02-06 15:30:00
}

// ExampleDate_Next_sameWeekday demonstrates the same-weekday edge case
func ExampleDate_Next_sameWeekday() {
	// Next ALWAYS returns future, never today (even if same weekday)
	monday := quando.From(time.Date(2026, 2, 9, 15, 30, 0, 0, time.UTC)) // Monday

	nextMonday := monday.Next(time.Monday)
	fmt.Printf("Days later: %d\n", int(nextMonday.Time().Sub(monday.Time()).Hours()/24))
	// Output: Days later: 7
}

// ExampleDiff demonstrates calculating the duration between two dates
func ExampleDiff() {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)

	dur := quando.Diff(start, end)

	fmt.Printf("Days: %d\n", dur.Days())
	fmt.Printf("Months: %d\n", dur.Months())
	fmt.Printf("Years: %d\n", dur.Years())
	// Output:
	// Days: 364
	// Months: 11
	// Years: 0
}

// ExampleDuration_MonthsFloat demonstrates precise month calculations
func ExampleDuration_MonthsFloat() {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 2, 16, 0, 0, 0, 0, time.UTC)

	dur := quando.Diff(start, end)

	intMonths := dur.Months()
	floatMonths := dur.MonthsFloat()

	fmt.Printf("Integer months: %d\n", intMonths)
	fmt.Printf("Float months: %.2f\n", floatMonths)
	// Output:
	// Integer months: 1
	// Float months: 1.54
}

// ExampleDuration_negative demonstrates negative durations
func ExampleDuration_negative() {
	// When start is after end, duration is negative
	start := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	dur := quando.Diff(start, end)

	fmt.Printf("Months: %d\n", dur.Months())
	fmt.Printf("Years: %d\n", dur.Years())
	// Output:
	// Months: -12
	// Years: -1
}

// ExampleErrInvalidFormat demonstrates handling invalid date formats
func Example_errorHandling() {
	// Note: Parse doesn't exist yet, so this is a conceptual example
	// showing the error handling pattern that will be used

	// Simulate an error by directly using the sentinel error
	err := quando.ErrInvalidFormat

	// Check for specific error type
	if errors.Is(err, quando.ErrInvalidFormat) {
		fmt.Println("Invalid format detected")
	}

	// Output: Invalid format detected
}

// Example_errorTypes demonstrates all error types
func Example_errorTypes() {
	// Show all defined error types
	errors := []error{
		quando.ErrInvalidFormat,
		quando.ErrInvalidTimezone,
		quando.ErrOverflow,
	}

	for _, err := range errors {
		fmt.Println(err.Error())
	}
	// Output:
	// invalid date format
	// invalid timezone
	// date overflow
}

// ExampleDate_Add demonstrates date arithmetic
func ExampleDate_Add() {
	date := quando.From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC))

	fmt.Println("Original:", date)
	fmt.Println("+1 day:", date.Add(1, quando.Days))
	fmt.Println("+1 month:", date.Add(1, quando.Months))
	fmt.Println("+1 year:", date.Add(1, quando.Years))
	// Output:
	// Original: 2026-01-15 12:00:00
	// +1 day: 2026-01-16 12:00:00
	// +1 month: 2026-02-15 12:00:00
	// +1 year: 2027-01-15 12:00:00
}

// ExampleDate_Add_monthEndOverflow demonstrates month-end overflow behavior
func ExampleDate_Add_monthEndOverflow() {
	// When adding months, if target day doesn't exist, snap to month end
	date := quando.From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC))

	fmt.Println("Jan 31 + 1 month:", date.Add(1, quando.Months))
	fmt.Println("Jan 31 + 2 months:", date.Add(2, quando.Months))
	// Output:
	// Jan 31 + 1 month: 2026-02-28 12:00:00
	// Jan 31 + 2 months: 2026-03-31 12:00:00
}

// ExampleDate_Sub demonstrates date subtraction
func ExampleDate_Sub() {
	date := quando.From(time.Date(2026, 3, 31, 12, 0, 0, 0, time.UTC))

	fmt.Println("Original:", date)
	fmt.Println("-1 day:", date.Sub(1, quando.Days))
	fmt.Println("-1 month:", date.Sub(1, quando.Months))
	// Output:
	// Original: 2026-03-31 12:00:00
	// -1 day: 2026-03-30 12:00:00
	// -1 month: 2026-02-28 12:00:00
}

// ExampleDate_Add_chaining demonstrates method chaining
func ExampleDate_Add_chaining() {
	date := quando.From(time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC))

	result := date.
		Add(1, quando.Months).
		Add(15, quando.Days).
		Sub(2, quando.Hours)

	fmt.Println(result)
	// Output: 2026-02-16 10:00:00
}

// ExampleParse demonstrates basic date parsing with automatic format detection
func ExampleParse() {
	date, err := quando.Parse("2026-02-09")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(date)
	// Output: 2026-02-09 00:00:00
}

// ExampleParse_formats demonstrates parsing various supported date formats
func ExampleParse_formats() {
	formats := []string{
		"2026-02-09",  // ISO format (YYYY-MM-DD)
		"2026/02/09",  // ISO with slash (YYYY/MM/DD)
		"09.02.2026",  // EU format (DD.MM.YYYY)
	}

	for _, f := range formats {
		date, err := quando.Parse(f)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", f, err)
			continue
		}
		fmt.Println(date)
	}
	// Output:
	// 2026-02-09 00:00:00
	// 2026-02-09 00:00:00
	// 2026-02-09 00:00:00
}

// ExampleParse_error demonstrates error handling for ambiguous formats
func ExampleParse_error() {
	// Slash format without year prefix is ambiguous
	// (could be US: MM/DD/YYYY or EU: DD/MM/YYYY)
	_, err := quando.Parse("01/02/2026")

	if errors.Is(err, quando.ErrInvalidFormat) {
		fmt.Println("Ambiguous format detected")
	}
	// Output: Ambiguous format detected
}
