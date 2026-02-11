# quando

> Intuitive and idiomatic date calculations for Go

**quando** is a standalone Go library for complex date operations that are cumbersome or impossible with the Go standard library alone. It provides a fluent API for date arithmetic, parsing, formatting, and timezone-aware calculations.

## Features

- **Fluent API**: Chain operations naturally: `quando.Now().Add(2, Months).StartOf(Week)`
- **Month-End Aware**: Handles edge cases like `Jan 31 + 1 month = Feb 28`
- **DST Safe**: Calendar-based arithmetic (not clock-based)
- **Zero Dependencies**: Only Go stdlib
- **Immutable**: Thread-safe by design
- **i18n Ready**: Multilingual formatting (EN, DE in Phase 1)
- **Testable**: Built-in Clock abstraction for deterministic tests

## Installation

```bash
go get code.beautifulmachines.dev/quando
```

## Quick Start

```go
import "code.beautifulmachines.dev/quando"

// Get current date
now := quando.Now()

// Date arithmetic with month-end handling
date := quando.From(time.Date(2026, 1, 31, 0, 0, 0, 0, time.UTC))
result := date.Add(1, quando.Months) // Feb 28, 2026 (not overflow)

// Snap to boundaries
monday := quando.Now().StartOf(quando.Week)     // This week's Monday 00:00
endOfMonth := quando.Now().EndOf(quando.Month)  // Last day of month 23:59:59

// Next/Previous dates
nextFriday := quando.Now().Next(time.Friday)
prevMonday := quando.Now().Prev(time.Monday)

// Differences
duration := quando.Diff(startDate, endDate)
months := duration.Months()
humanReadable := duration.Human() // "2 years, 3 months, 5 days"

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

## Requirements

- Go 1.22 or later
- No external dependencies

## Documentation

Full documentation available at: [pkg.go.dev/code.beautifulmachines.dev/quando](https://pkg.go.dev/code.beautifulmachines.dev/quando)

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Contributing

Contributions welcome! Please ensure:
- Tests pass (`go test ./...`)
- Code is formatted (`go fmt ./...`)
- No vet warnings (`go vet ./...`)
- Coverage ≥95% for new code

## Roadmap

### Phase 1 (Current)
- ✅ Project setup
- 🚧 Core date operations
- 🚧 Parsing and formatting
- 🚧 Timezone handling
- 🚧 i18n (EN, DE)

### Phase 2
- Date ranges and series
- Batch operations
- Performance optimizations

### Phase 3
- Holiday calendars
- Business day calculations
- Extended language support

## Acknowledgments

Inspired by [Moment.js](https://momentjs.com/), [Carbon](https://carbon.nesbot.com/), and [date-fns](https://date-fns.org/), but designed to be idiomatic Go.
