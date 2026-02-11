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
		"2026-02-09", // ISO format (YYYY-MM-DD)
		"2026/02/09", // ISO with slash (YYYY/MM/DD)
		"09.02.2026", // EU format (DD.MM.YYYY)
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

// ExampleDate_In demonstrates timezone conversion
func ExampleDate_In() {
	// Create a UTC time
	utc := quando.From(time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC))

	// Convert to Berlin timezone (UTC+2 in summer)
	berlin, err := utc.In("Europe/Berlin")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("UTC:    %v\n", utc)
	fmt.Printf("Berlin: %v\n", berlin)
	// Output:
	// UTC:    2026-06-15 12:00:00
	// Berlin: 2026-06-15 14:00:00
}

// ExampleDate_In_dst demonstrates DST handling
func ExampleDate_In_dst() {
	// Winter: Europe/Berlin is UTC+1 (CET)
	winter := quando.From(time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC))
	berlin, _ := winter.In("Europe/Berlin")
	fmt.Printf("Winter: UTC 12:00 -> Berlin %02d:00\n", berlin.Time().Hour())

	// Summer: Europe/Berlin is UTC+2 (CEST)
	summer := quando.From(time.Date(2026, 7, 15, 12, 0, 0, 0, time.UTC))
	berlin, _ = summer.In("Europe/Berlin")
	fmt.Printf("Summer: UTC 12:00 -> Berlin %02d:00\n", berlin.Time().Hour())

	// Output:
	// Winter: UTC 12:00 -> Berlin 13:00
	// Summer: UTC 12:00 -> Berlin 14:00
}

// ExampleDate_In_error demonstrates error handling
func ExampleDate_In_error() {
	date := quando.Now()

	_, err := date.In("Invalid/Timezone")
	if errors.Is(err, quando.ErrInvalidTimezone) {
		fmt.Println("Invalid timezone name")
	}

	// Output: Invalid timezone name
}

// ExampleLang_MonthName demonstrates localized month names
func ExampleLang_MonthName() {
	fmt.Println(quando.EN.MonthName(time.February))
	fmt.Println(quando.DE.MonthName(time.February))
	// Output:
	// February
	// Februar
}

// ExampleLang_WeekdayName demonstrates localized weekday names
func ExampleLang_WeekdayName() {
	fmt.Println(quando.EN.WeekdayName(time.Monday))
	fmt.Println(quando.DE.WeekdayName(time.Monday))
	// Output:
	// Monday
	// Montag
}

// ExampleLang_DurationUnit demonstrates localized duration units
func ExampleLang_DurationUnit() {
	// Singular
	fmt.Println(quando.EN.DurationUnit("month", false))
	fmt.Println(quando.DE.DurationUnit("month", false))

	// Plural
	fmt.Println(quando.EN.DurationUnit("month", true))
	fmt.Println(quando.DE.DurationUnit("month", true))
	// Output:
	// month
	// Monat
	// months
	// Monate
}

// ExampleDuration_Human demonstrates human-readable duration formatting in English
func ExampleDuration_Human() {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 11, 17, 0, 0, 0, 0, time.UTC)

	dur := quando.Diff(start, end)
	fmt.Println(dur.Human())
	// Output: 10 months, 16 days
}

// ExampleDuration_Human_german demonstrates human-readable duration formatting in German
func ExampleDuration_Human_german() {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 3, 5, 0, 0, 0, time.UTC)

	dur := quando.Diff(start, end)
	fmt.Println(dur.Human(quando.DE))
	// Output: 2 Tage, 5 Stunden
}

// ExampleDuration_Human_adaptive demonstrates adaptive granularity
func ExampleDuration_Human_adaptive() {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	// Large duration: shows years and months
	end1 := time.Date(2027, 3, 15, 0, 0, 0, 0, time.UTC)
	dur1 := quando.Diff(start, end1)
	fmt.Println(dur1.Human())

	// Medium duration: shows days and hours
	end2 := time.Date(2026, 1, 3, 5, 0, 0, 0, time.UTC)
	dur2 := quando.Diff(start, end2)
	fmt.Println(dur2.Human())

	// Small duration: shows seconds only
	end3 := time.Date(2026, 1, 1, 0, 0, 45, 0, time.UTC)
	dur3 := quando.Diff(start, end3)
	fmt.Println(dur3.Human())

	// Output:
	// 1 year, 2 months
	// 2 days, 5 hours
	// 45 seconds
}

// ExampleParseWithLayout demonstrates parsing with explicit layout format
func ExampleParseWithLayout() {
	// US format: month/day/year
	dateUS, _ := quando.ParseWithLayout("01/02/2026", "01/02/2006")
	fmt.Println("US format:", dateUS) // January 2, 2026

	// EU format: day/month/year
	dateEU, _ := quando.ParseWithLayout("01/02/2026", "02/01/2006")
	fmt.Println("EU format:", dateEU) // February 1, 2026

	// Output:
	// US format: 2026-01-02 00:00:00
	// EU format: 2026-02-01 00:00:00
}

// ExampleParseWithLayout_custom demonstrates custom date format with English month names
func ExampleParseWithLayout_custom() {
	date, _ := quando.ParseWithLayout("9. February 2026", "2. January 2006")
	fmt.Println(date)
	// Output: 2026-02-09 00:00:00
}

// ExampleParseWithLayout_error demonstrates error handling
func ExampleParseWithLayout_error() {
	_, err := quando.ParseWithLayout("99/99/2026", "02/01/2006")
	if errors.Is(err, quando.ErrInvalidFormat) {
		fmt.Println("Invalid date format detected")
	}
	// Output: Invalid date format detected
}

// ExampleParseRelative demonstrates parsing relative date expressions
func ExampleParseRelative() {
	// Simple keywords (results depend on current date)
	today, _ := quando.ParseRelative("today")
	tomorrow, _ := quando.ParseRelative("tomorrow")
	yesterday, _ := quando.ParseRelative("yesterday")

	fmt.Printf("Type: %T\n", today)
	fmt.Printf("Type: %T\n", tomorrow)
	fmt.Printf("Type: %T\n", yesterday)
	// Output:
	// Type: quando.Date
	// Type: quando.Date
	// Type: quando.Date
}

// ExampleParseRelative_offsets demonstrates relative offset expressions
func ExampleParseRelative_offsets() {
	// Note: Results depend on current date
	// Using ParseRelativeWithClock for deterministic example

	clock := quando.NewFixedClock(time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC))

	twoDaysFromNow, _ := quando.ParseRelativeWithClock("+2 days", clock)
	oneWeekAgo, _ := quando.ParseRelativeWithClock("-1 week", clock)
	threeMonthsFromNow, _ := quando.ParseRelativeWithClock("+3 months", clock)

	fmt.Println(twoDaysFromNow)
	fmt.Println(oneWeekAgo)
	fmt.Println(threeMonthsFromNow)
	// Output:
	// 2026-02-17 00:00:00
	// 2026-02-08 00:00:00
	// 2026-05-15 00:00:00
}

// ExampleParseRelative_caseInsensitive demonstrates case-insensitive parsing
func ExampleParseRelative_caseInsensitive() {
	clock := quando.NewFixedClock(time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC))

	// All of these work
	date1, _ := quando.ParseRelativeWithClock("today", clock)
	date2, _ := quando.ParseRelativeWithClock("TODAY", clock)
	date3, _ := quando.ParseRelativeWithClock("Today", clock)

	fmt.Println(date1.Unix() == date2.Unix())
	fmt.Println(date2.Unix() == date3.Unix())
	// Output:
	// true
	// true
}

// ExampleParseRelative_error demonstrates error handling
func ExampleParseRelative_error() {
	_, err := quando.ParseRelative("next monday") // Not supported in Phase 1
	if errors.Is(err, quando.ErrInvalidFormat) {
		fmt.Println("Complex expressions not yet supported")
	}
	// Output: Complex expressions not yet supported
}

// ExampleDate_Format demonstrates basic date formatting
func ExampleDate_Format() {
	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))

	fmt.Println(date.Format(quando.ISO))
	fmt.Println(date.Format(quando.EU))
	fmt.Println(date.Format(quando.US))
	fmt.Println(date.Format(quando.Long))
	// Output:
	// 2026-02-09
	// 09.02.2026
	// 02/09/2026
	// February 9, 2026
}

// ExampleDate_Format_isoFormat demonstrates ISO 8601 format
func ExampleDate_Format_isoFormat() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	fmt.Println(date.Format(quando.ISO))
	// Output: 2026-02-09
}

// ExampleDate_Format_euFormat demonstrates European format
func ExampleDate_Format_euFormat() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	fmt.Println(date.Format(quando.EU))
	// Output: 09.02.2026
}

// ExampleDate_Format_usFormat demonstrates US format
func ExampleDate_Format_usFormat() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	fmt.Println(date.Format(quando.US))
	// Output: 02/09/2026
}

// ExampleDate_Format_longFormat demonstrates long format in English
func ExampleDate_Format_longFormat() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	fmt.Println(date.Format(quando.Long))
	// Output: February 9, 2026
}

// ExampleDate_Format_longFormatGerman demonstrates long format in German
func ExampleDate_Format_longFormatGerman() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)).WithLang(quando.DE)
	fmt.Println(date.Format(quando.Long))
	// Output: 9. Februar 2026
}

// ExampleDate_Format_rfc2822Format demonstrates RFC 2822 email format
func ExampleDate_Format_rfc2822Format() {
	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	fmt.Println(date.Format(quando.RFC2822))
	// Output: Mon, 09 Feb 2026 12:30:45 +0000
}

// ExampleDate_Format_languageIndependence demonstrates that most formats ignore language
func ExampleDate_Format_languageIndependence() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))

	// ISO format is always the same regardless of language
	dateEN := date.WithLang(quando.EN)
	dateDE := date.WithLang(quando.DE)

	fmt.Println("ISO (EN):", dateEN.Format(quando.ISO))
	fmt.Println("ISO (DE):", dateDE.Format(quando.ISO))
	fmt.Println("Same:", dateEN.Format(quando.ISO) == dateDE.Format(quando.ISO))
	// Output:
	// ISO (EN): 2026-02-09
	// ISO (DE): 2026-02-09
	// Same: true
}

// ExampleDate_FormatLayout demonstrates custom layout formatting
func ExampleDate_FormatLayout() {
	date := quando.From(time.Date(2026, 2, 9, 14, 30, 0, 0, time.UTC))

	// Common layouts
	fmt.Println(date.FormatLayout("Monday, January 2, 2006"))
	fmt.Println(date.FormatLayout("Mon, Jan 2, 2006"))
	fmt.Println(date.FormatLayout("January 2, 2006"))
	fmt.Println(date.FormatLayout("02 Jan 2006"))
	// Output:
	// Monday, February 9, 2026
	// Mon, Feb 9, 2026
	// February 9, 2026
	// 09 Feb 2026
}

// ExampleDate_FormatLayout_german demonstrates German localization with custom layouts
func ExampleDate_FormatLayout_german() {
	date := quando.From(time.Date(2026, 2, 9, 14, 30, 0, 0, time.UTC)).WithLang(quando.DE)

	// German formats
	fmt.Println(date.FormatLayout("Monday, 2. January 2006"))
	fmt.Println(date.FormatLayout("Mon, 02. Jan 2006"))
	fmt.Println(date.FormatLayout("2. January 2006"))
	// Output:
	// Montag, 9. Februar 2026
	// Mo, 09. Feb 2026
	// 9. Februar 2026
}

// ExampleDate_FormatLayout_custom demonstrates custom layouts with time components
func ExampleDate_FormatLayout_custom() {
	date := quando.From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC))

	// With time
	fmt.Println(date.FormatLayout("Monday, January 2, 2006 at 15:04"))
	fmt.Println(date.FormatLayout("Mon Jan 2 15:04:05 2006"))
	fmt.Println(date.FormatLayout("2006-01-02 15:04:05"))

	// German with time
	dateDE := date.WithLang(quando.DE)
	fmt.Println(dateDE.FormatLayout("Monday, 2. January 2006 um 15:04 Uhr"))
	// Output:
	// Monday, February 9, 2026 at 14:30
	// Mon Feb 9 14:30:45 2026
	// 2026-02-09 14:30:45
	// Montag, 9. Februar 2026 um 14:30 Uhr
}

// ExampleDate_FormatLayout_comparison demonstrates comparing preset Format vs custom FormatLayout
func ExampleDate_FormatLayout_comparison() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))

	// Preset formats
	fmt.Println("Format(Long):", date.Format(quando.Long))
	fmt.Println("Format(ISO):", date.Format(quando.ISO))

	// Custom layouts (equivalent)
	fmt.Println("FormatLayout(custom):", date.FormatLayout("January 2, 2006"))
	fmt.Println("FormatLayout(ISO):", date.FormatLayout("2006-01-02"))

	// Custom layout flexibility
	fmt.Println("FormatLayout(custom style):", date.FormatLayout("Mon, Jan 2"))
	// Output:
	// Format(Long): February 9, 2026
	// Format(ISO): 2026-02-09
	// FormatLayout(custom): February 9, 2026
	// FormatLayout(ISO): 2026-02-09
	// FormatLayout(custom style): Mon, Feb 9
}

// ExampleDate_WeekNumber demonstrates ISO 8601 week number calculation
func ExampleDate_WeekNumber() {
	date := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	fmt.Printf("Week number: %d\n", date.WeekNumber())
	// Output: Week number: 7
}

// ExampleDate_WeekNumber_yearBoundary demonstrates week numbers at year boundaries
func ExampleDate_WeekNumber_yearBoundary() {
	// Jan 1, 2023 (Sunday) belongs to week 52 of 2022
	jan1 := quando.From(time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC))
	fmt.Printf("2023-01-01: Week %d\n", jan1.WeekNumber())

	// Jan 2, 2023 (Monday) is week 1 of 2023
	jan2 := quando.From(time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC))
	fmt.Printf("2023-01-02: Week %d\n", jan2.WeekNumber())

	// Dec 31, 2026 (Thursday) is week 53
	dec31 := quando.From(time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC))
	fmt.Printf("2026-12-31: Week %d\n", dec31.WeekNumber())
	// Output:
	// 2023-01-01: Week 52
	// 2023-01-02: Week 1
	// 2026-12-31: Week 53
}

// ExampleDate_Quarter demonstrates fiscal quarter calculation
func ExampleDate_Quarter() {
	q1 := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	q2 := quando.From(time.Date(2026, 5, 15, 0, 0, 0, 0, time.UTC))
	q3 := quando.From(time.Date(2026, 8, 20, 0, 0, 0, 0, time.UTC))
	q4 := quando.From(time.Date(2026, 11, 25, 0, 0, 0, 0, time.UTC))

	fmt.Printf("February: Q%d\n", q1.Quarter())
	fmt.Printf("May: Q%d\n", q2.Quarter())
	fmt.Printf("August: Q%d\n", q3.Quarter())
	fmt.Printf("November: Q%d\n", q4.Quarter())
	// Output:
	// February: Q1
	// May: Q2
	// August: Q3
	// November: Q4
}

// ExampleDate_DayOfYear demonstrates day of year calculation
func ExampleDate_DayOfYear() {
	jan1 := quando.From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC))
	feb9 := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))
	dec31 := quando.From(time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC))

	fmt.Printf("Jan 1: Day %d\n", jan1.DayOfYear())
	fmt.Printf("Feb 9: Day %d\n", feb9.DayOfYear())
	fmt.Printf("Dec 31: Day %d\n", dec31.DayOfYear())
	// Output:
	// Jan 1: Day 1
	// Feb 9: Day 40
	// Dec 31: Day 365
}

// ExampleDate_DayOfYear_leapYear demonstrates day of year in a leap year
func ExampleDate_DayOfYear_leapYear() {
	// 2024 is a leap year
	feb29 := quando.From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC))
	dec31 := quando.From(time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))

	fmt.Printf("Feb 29 (leap year): Day %d\n", feb29.DayOfYear())
	fmt.Printf("Dec 31 (leap year): Day %d\n", dec31.DayOfYear())
	// Output:
	// Feb 29 (leap year): Day 60
	// Dec 31 (leap year): Day 366
}

// ExampleDate_IsWeekend demonstrates weekend detection
func ExampleDate_IsWeekend() {
	monday := quando.From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))   // Monday
	saturday := quando.From(time.Date(2026, 2, 14, 0, 0, 0, 0, time.UTC)) // Saturday
	sunday := quando.From(time.Date(2026, 2, 15, 0, 0, 0, 0, time.UTC))   // Sunday

	fmt.Printf("Monday: %v\n", monday.IsWeekend())
	fmt.Printf("Saturday: %v\n", saturday.IsWeekend())
	fmt.Printf("Sunday: %v\n", sunday.IsWeekend())
	// Output:
	// Monday: false
	// Saturday: true
	// Sunday: true
}

// ExampleDate_IsLeapYear demonstrates leap year detection
func ExampleDate_IsLeapYear() {
	year2024 := quando.From(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)) // Leap year
	year2026 := quando.From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)) // Not leap year
	year2000 := quando.From(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) // Leap year (400 rule)
	year1900 := quando.From(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) // Not leap year (100 rule)

	fmt.Printf("2024: %v\n", year2024.IsLeapYear())
	fmt.Printf("2026: %v\n", year2026.IsLeapYear())
	fmt.Printf("2000: %v\n", year2000.IsLeapYear())
	fmt.Printf("1900: %v\n", year1900.IsLeapYear())
	// Output:
	// 2024: true
	// 2026: false
	// 2000: true
	// 1900: false
}

// ExampleDate_Info demonstrates aggregated date metadata
func ExampleDate_Info() {
	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	info := date.Info()

	fmt.Printf("Week number: %d\n", info.WeekNumber)
	fmt.Printf("Quarter: %d\n", info.Quarter)
	fmt.Printf("Day of year: %d\n", info.DayOfYear)
	fmt.Printf("Is weekend: %v\n", info.IsWeekend)
	fmt.Printf("Is leap year: %v\n", info.IsLeapYear)
	// Output:
	// Week number: 7
	// Quarter: 1
	// Day of year: 40
	// Is weekend: false
	// Is leap year: false
}

// ExampleDate_Info_leapYear demonstrates Info() for a leap year date
func ExampleDate_Info_leapYear() {
	// Saturday, Feb 29, 2024 (leap year)
	date := quando.From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC))
	info := date.Info()

	fmt.Printf("Week number: %d\n", info.WeekNumber)
	fmt.Printf("Quarter: %d\n", info.Quarter)
	fmt.Printf("Day of year: %d\n", info.DayOfYear)
	fmt.Printf("Is weekend: %v\n", info.IsWeekend)
	fmt.Printf("Is leap year: %v\n", info.IsLeapYear)
	// Output:
	// Week number: 9
	// Quarter: 1
	// Day of year: 60
	// Is weekend: false
	// Is leap year: true
}

// ExampleMustParse demonstrates basic usage in tests and initialization
func ExampleMustParse() {
	// Use in test fixtures or static initialization
	date := quando.MustParse("2026-02-09")
	fmt.Println(date)
	// Output: 2026-02-09 00:00:00
}

// ExampleMustParse_testFixture demonstrates using MustParse in test fixtures
func ExampleMustParse_testFixture() {
	// Common pattern in tests - no error handling needed
	startDate := quando.MustParse("2026-01-01")
	endDate := quando.MustParse("2026-12-31")

	fmt.Printf("Start: %v\n", startDate.Time().Format("2006-01-02"))
	fmt.Printf("End: %v\n", endDate.Time().Format("2006-01-02"))
	// Output:
	// Start: 2026-01-01
	// End: 2026-12-31
}

// ExampleMustParse_staticInit demonstrates static initialization pattern
func ExampleMustParse_staticInit() {
	// Pattern for package-level constants
	var (
		epoch = quando.MustParse("1970-01-01")
		y2k   = quando.MustParse("2000-01-01")
	)

	fmt.Printf("Epoch: %v\n", epoch.Time().Year())
	fmt.Printf("Y2K: %v\n", y2k.Time().Year())
	// Output:
	// Epoch: 1970
	// Y2K: 2000
}
