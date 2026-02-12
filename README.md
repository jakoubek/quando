# quando

[![Mirror on GitHub](https://img.shields.io/badge/mirror-GitHub-blue)](https://github.com/jakoubek/quando)
[![Go Reference](https://pkg.go.dev/badge/code.beautifulmachines.dev/jakoubek/quando.svg)](https://pkg.go.dev/code.beautifulmachines.dev/jakoubek/quando)
[![Go Report Card](https://goreportcard.com/badge/code.beautifulmachines.dev/jakoubek/quando)](https://goreportcard.com/report/code.beautifulmachines.dev/jakoubek/quando)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> **Primary repository:** [code.beautifulmachines.dev/jakoubek/quando](https://code.beautifulmachines.dev/jakoubek/quando) · GitHub is a read-only mirror.

> Intuitive and idiomatic date calculations for Go

**quando** is a standalone Go library for complex date operations that are cumbersome or impossible with the Go standard library alone. It provides a fluent API for date arithmetic, parsing, formatting, and timezone-aware calculations.

## Features

- **Fluent API**: Chain operations naturally: `quando.Now().Add(2, quando.Months).StartOf(quando.Weeks)`
- **Month-End Aware**: Handles edge cases like `Jan 31 + 1 month = Feb 28`
- **DST Safe**: Calendar-based arithmetic (not clock-based)
- **Zero Dependencies**: Only Go stdlib
- **Immutable**: Thread-safe by design
- **i18n Ready**: Multilingual formatting (EN, DE in Phase 1)
- **Testable**: Built-in Clock abstraction for deterministic tests

## Installation

```bash
go get code.beautifulmachines.dev/jakoubek/quando
```

## Quick Start

```go
import "code.beautifulmachines.dev/jakoubek/quando"

// Get current date
now := quando.Now()

// Date arithmetic with month-end handling
date := quando.From(time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC))
result := date.Add(1, quando.Months) // Feb 28, 2026 (not overflow)

// Snap to boundaries
monday := quando.Now().StartOf(quando.Weeks)     // This week's Monday 00:00
endOfMonth := quando.Now().EndOf(quando.Months)  // Last day of month 23:59:59

// Next/Previous dates
nextFriday := quando.Now().Next(time.Friday)
prevMonday := quando.Now().Prev(time.Monday)

// Human-readable differences
start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2028, 3, 15, 0, 0, 0, 0, time.UTC)
duration := quando.Diff(start, end)
fmt.Println(duration.Human()) // "2 years, 2 months, 14 days"

// Parsing (automatic format detection)
date, err := quando.Parse("2026-02-09")
date, err := quando.Parse("09.02.2026")       // EU format
date, err := quando.Parse("3 days ago")       // Relative

// Formatting
date.Format(quando.ISO)      // "2026-02-09"
date.Format(quando.Long)     // "February 9, 2026"
date.FormatLayout("02 Jan")  // "09 Feb"

// Timezone conversion
utc := date.InTimezone("UTC")
berlin := date.InTimezone("Europe/Berlin")

// Date inspection
week := date.WeekNumber()       // ISO 8601 week number
quarter := date.Quarter()       // 1-4
dayOfYear := date.DayOfYear()   // 1-366

// Complex chaining for real-world scenarios
reportDeadline := quando.Now().
    Add(1, quando.Quarters).     // Next quarter
    EndOf(quando.Quarters).      // Last day of that quarter
    StartOf(quando.Weeks).       // Monday of that week
    Add(-1, quando.Weeks)        // One week before

// Multilingual formatting
dateEN := quando.Now().WithLang(quando.EN)
dateDE := quando.Now().WithLang(quando.DE)
fmt.Println(dateEN.Format(quando.Long)) // "February 9, 2026"
fmt.Println(dateDE.Format(quando.Long)) // "9. Februar 2026"
```

## Core Concepts

### Month-End Overflow Handling

When adding months, if the target day doesn't exist, quando snaps to the last valid day:

```go
jan31 := quando.From(time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC))
feb28 := jan31.Add(1, quando.Months) // Feb 28, not March 3
```

### DST-Aware Arithmetic

Adding days means "same time on next calendar day", not "24 hours later":

```go
// During DST transition (23-hour day)
date := quando.From(time.Date(2026, 3, 31, 2, 0, 0, 0, cetLocation))
next := date.Add(1, quando.Days) // April 1, 2:00 CEST (not 3:00)
```

### Immutability

All operations return new instances. Original values are never modified:

```go
original := quando.Now()
modified := original.Add(1, quando.Days)
// original is unchanged
```

## Why quando? Comparison to time.Time

### When to Use quando

Use **quando** when you need:
- **Month-aware arithmetic**: `Add(1, Months)` handles month-end overflow
- **Business logic**: "Next Friday", "End of Quarter", ISO week numbers
- **Human-readable durations**: "2 years, 3 months, 5 days"
- **Fluent API**: Method chaining for complex date calculations
- **Automatic parsing**: Detect ISO, EU, US formats automatically
- **i18n formatting**: Multilingual date formatting (EN, DE, more coming)

Use **time.Time** when you need:
- Simple clock arithmetic (add 24 hours)
- High-precision timestamps (nanoseconds matter)
- Minimal dependencies (quando is stdlib-only but adds abstraction)
- Low-level system operations

### Side-by-Side Comparison

#### Month Arithmetic with Overflow

```go
// stdlib: Complex and error-prone
t := time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC)
// Add 1 month manually - need to handle overflow
nextMonth := t.AddDate(0, 1, 0) // March 3! ❌ Unexpected

// quando: Intuitive and correct
date := quando.From(time.Date(2026, 1, 31, 12, 0, 0, 0, time.UTC))
nextMonth := date.Add(1, quando.Months) // Feb 28 ✅ Expected
```

#### Finding "Next Friday"

```go
// stdlib: Manual calculation required
t := time.Now()
daysUntilFriday := (int(time.Friday) - int(t.Weekday()) + 7) % 7
if daysUntilFriday == 0 {
    daysUntilFriday = 7 // Never return today
}
nextFriday := t.AddDate(0, 0, daysUntilFriday)

// quando: One method call
nextFriday := quando.Now().Next(time.Friday)
```

#### Human-Readable Duration

```go
// stdlib: No built-in solution, must implement yourself
start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2028, 3, 15, 0, 0, 0, 0, time.UTC)
duration := end.Sub(start)
// duration is just time.Duration (2y3m15d shown as "19480h0m0s") ❌

// quando: Built-in human formatting
duration := quando.Diff(start, end)
fmt.Println(duration.Human()) // "2 years, 2 months, 14 days" ✅
```

#### Start of Week (Monday)

```go
// stdlib: Manual calculation
t := time.Now()
weekday := int(t.Weekday())
if weekday == 0 { // Sunday
    weekday = 7
}
daysToMonday := weekday - 1
startOfWeek := t.AddDate(0, 0, -daysToMonday)
startOfWeek = time.Date(startOfWeek.Year(), startOfWeek.Month(),
    startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())

// quando: One method call
startOfWeek := quando.Now().StartOf(quando.Weeks)
```

### Feature Comparison

Feature | time.Time | quando
--------|-----------|--------
Basic date/time | ✅ | ✅
Add/subtract duration | ✅ | ✅
Add/subtract months (overflow-safe) | ❌ | ✅
Snap to start/end of period | ❌ | ✅
Next/Previous weekday | ❌ | ✅
ISO 8601 week number | ❌ | ✅
Quarter calculation | ❌ | ✅
Human-readable duration | ❌ | ✅
Automatic format parsing | ❌ | ✅
Relative parsing ("tomorrow") | ❌ | ✅
i18n formatting | ❌ | ✅
Fluent API / chaining | ❌ | ✅
Immutability guarantee | ⚠️ (manual) | ✅
Testing (Clock abstraction) | ❌ | ✅

## Testing

quando provides a `Clock` interface for deterministic tests:

```go
// In production
date := quando.Now()

// In tests
fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
clock := quando.NewFixedClock(fixedTime)
date := clock.Now() // Always returns Feb 9, 2026
```

## Performance

quando is designed for high performance with zero allocations in hot paths:

### Benchmark Results

Operation | Target | Actual | Status
----------|--------|--------|--------
Add/Sub (Days) | < 1µs | 37 ns | ✅ 27x faster
Add/Sub (Months) | < 1µs | 181 ns | ✅ 5.5x faster
Diff (integer) | < 1µs | 51 ns | ✅ 20x faster
Diff (float) | < 2µs | 155 ns | ✅ 13x faster
Format (ISO/EU/US) | < 5µs | 91 ns | ✅ 55x faster
Format (Long, i18n) | < 10µs | 267 ns | ✅ 37x faster
Parse (automatic) | < 10µs | 106 ns | ✅ 94x faster
Parse (relative) | < 20µs | 581 ns | ✅ 34x faster

**Key Performance Features:**
- Zero allocations for arithmetic operations (Add, Sub)
- Zero allocations for snap operations (StartOf, EndOf, Next, Prev)
- Zero allocations for date inspection (WeekNumber, Quarter, etc.)
- Minimal allocations for formatting (1-3 per operation)
- Immutable design enables safe concurrent use

**Run benchmarks:**
```bash
go test -bench=. -benchmem
```

## Requirements

- Go 1.22 or later
- No external dependencies

## Documentation

Full documentation available at: [pkg.go.dev/code.beautifulmachines.dev/jakoubek/quando](https://pkg.go.dev/code.beautifulmachines.dev/jakoubek/quando)

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions welcome! Please ensure:
- Tests pass (`go test ./...`)
- Code is formatted (`go fmt ./...`)
- No vet warnings (`go vet ./...`)
- Coverage ≥95% for new code

## Roadmap

### Phase 1 ✅ COMPLETE
- ✅ Project setup
- ✅ Core date operations (Add, Sub, StartOf, EndOf, Next, Prev, Diff)
- ✅ Parsing and formatting (ISO, EU, US, Long, RFC2822, relative)
- ✅ Timezone handling (IANA database support)
- ✅ i18n (EN, DE)
- ✅ Comprehensive test suite (99.5% coverage)
- ✅ Complete documentation and examples

### Phase 2 (Planned)
- Date ranges and series
- Batch operations
- Performance optimizations

### Phase 3
- Holiday calendars
- Business day calculations
- Extended language support

## Acknowledgments

Inspired by [Moment.js](https://momentjs.com/), [Carbon](https://carbon.nesbot.com/), and [date-fns](https://date-fns.org/), but designed to be idiomatic Go.
