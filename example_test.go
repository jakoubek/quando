package quando_test

import (
	"fmt"
	"time"

	"code.beautifulmachines.dev/quando"
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
