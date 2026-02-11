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
