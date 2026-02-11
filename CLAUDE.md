# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**quando** is a standalone Go library for intuitive and idiomatic date calculations. It provides a fluent API for complex date operations that are cumbersome or impossible with the Go standard library alone.

**Vision**: The preferred Go library for date calculations – as natural and intuitive as Moment.js or Carbon, but Go-idiomatic and without external dependencies.

**Target**: Go 1.22+, zero dependencies (stdlib only), MIT License

## Development Commands

```bash
# Run tests
go test ./... -v

# Run tests with coverage
go test ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out

# Run a single test
go test -v -run TestFunctionName

# Run benchmarks
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkAdd -benchmem

# Format code
go fmt ./...

# Lint (when golangci-lint is configured)
golangci-lint run

# Vet code
go vet ./...

# Build (library - no binary)
go build ./...

# Tidy dependencies
go mod tidy
```

## Architecture Overview

### Core Design Principles

1. **Fluent API**: All operations return `quando.Date` for method chaining
2. **Immutability**: Every operation returns a new `Date` instance (thread-safe by design)
3. **Zero Dependencies**: Only Go stdlib (except optional i18n extensions)
4. **Stdlib Delegation**: Wrap `time.Time`, don't reimplement it
5. **No Panics**: All errors via return values (except `Must*` variants for tests)

### Package Structure (Flat Layout)

```
quando/
├── quando.go       # Package-level functions (Now, From, Parse, Diff)
├── date.go         # Date type, core methods, conversions
├── arithmetic.go   # Add, Sub (includes month-end overflow logic)
├── snap.go         # StartOf, EndOf, Next, Prev
├── diff.go         # Duration type, difference calculations
├── inspect.go      # WeekNumber, Quarter, DayOfYear, etc.
├── format.go       # Format, FormatLayout, preset constants
├── parse.go        # Parse, ParseWithLayout, ParseRelative
├── clock.go        # Clock interface for testability
├── i18n.go         # Internationalization (EN, DE in Phase 1)
├── errors.go       # Custom error types
└── internal/
    └── calc/       # Non-exported helper functions
```

**Rationale**: Flat package structure keeps all functions under `quando.*` namespace, avoiding complex imports and cyclic dependencies.

### Core Types

```go
// Date wraps time.Time and provides fluent API
type Date struct {
    t    time.Time
    lang Lang  // optional, for formatting
}

// Unit represents time units for arithmetic
type Unit int
const (
    Seconds Unit = iota
    Minutes
    Hours
    Days
    Weeks
    Months
    Quarters
    Years
)

// Duration represents time differences
type Duration struct {
    // Methods: Days(), Months(), Years(), Human(), etc.
}

// Clock interface for testability
type Clock interface {
    Now() Date
    From(t time.Time) Date
}
```

## Critical Implementation Details

### Month-End Overflow Behavior

When adding months, if the target date doesn't exist, snap to month end:
- `2026-01-31` + 1 month = `2026-02-28` (February end)
- `2026-01-24` + 1 month = `2026-02-24` (regular)
- `2026-05-31` + 1 month = `2026-06-30` (June has 30 days)

This is a **must-have feature** and requires extensive edge case testing.

### DST Handling

`Add(1, Days)` means "same time on next calendar day", NOT "24 hours later":
- Example: `2026-03-31 02:00 CET` + 1 Day = `2026-04-01 02:00 CEST` (only 23 actual hours due to DST)
- Rationale: Humans think in calendar days, not hour deltas

### Week Start Convention

- **Default**: Monday (ISO 8601)
- **StartOf(Week)**: Returns Monday 00:00:00
- **EndOf(Week)**: Returns Sunday 23:59:59
- **WeekNumber**: ISO 8601 (Week 1 = first week containing Thursday)

### Next/Prev Behavior

- `Next(Monday)`: Always NEXT Monday, never today (even if today is Monday)
- `Prev(Friday)`: Always PREVIOUS Friday, never today

### Parsing Ambiguity Rules

Slash formats without year prefix are **ambiguous** and must error:
- `2026-02-01` ✅ ISO, unambiguous
- `01.02.2026` ✅ EU (dot separator indicates EU convention)
- `2026/02/09` ✅ ISO with slash (year prefix is unambiguous)
- `01/02/2026` ❌ ERROR (ambiguous: US vs EU)

Use `ParseWithLayout()` for explicit format handling.

## Testing Requirements

### Coverage Target

**Minimum 95%** for all calculation functions.

### Critical Test Scenarios

1. **Month arithmetic**: Overflow edge cases, leap years
2. **Snap operations**: All units (Week, Month, Quarter, Year)
3. **Next/Prev**: Same weekday edge case (must skip)
4. **Diff calculations**: Cross year boundaries, leap years, negative diffs
5. **DST handling**: Add operations across DST transitions
6. **Parsing**: All formats, ambiguous inputs, invalid inputs
7. **WeekNumber**: ISO 8601 compliance (Week 1 contains Thursday)

### Test Organization

```go
// Use table-driven tests for edge cases
func TestAddMonths(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        months   int
        expected string
    }{
        {"month-end overflow", "2026-01-31", 1, "2026-02-28"},
        {"regular addition", "2026-01-24", 1, "2026-02-24"},
        // ... more cases
    }
    // ...
}
```

### Testability Pattern

Use `Clock` interface for deterministic tests:

```go
// Production code
date := quando.Now()

// Test code
clock := quando.NewFixedClock(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))
date := clock.Now()
```

## Performance Goals

All benchmarks must meet these targets:

- **Add/Sub operations**: < 1 µs
- **Diff calculations**: < 2 µs (< 1 µs for integer variants)
- **Formatting**: < 10 µs (< 5 µs without i18n)
- **Parsing**: < 10 µs (automatic), < 20 µs (relative)
- **Memory**: Zero allocations for chained operations (except final result)

Run benchmarks with:
```bash
go test -bench=. -benchmem -benchtime=10s
```

## Error Handling

### Sentinel Errors

```go
var (
    ErrInvalidFormat   = errors.New("invalid date format")
    ErrInvalidTimezone = errors.New("invalid timezone")
    ErrOverflow        = errors.New("date overflow")
)
```

### Error Wrapping

```go
return Date{}, fmt.Errorf("parsing date: %w", err)
```

### Panic Policy

**NEVER panic** except in `Must*` variants:
- `MustParse(s string) Date` - for tests/initialization only
- Document clearly that `Must*` functions panic on error

## Internationalization

### Phase 1 Languages

- **EN** (English) - default
- **DE** (Deutsch) - must-have for Phase 1

### i18n Applies To

- `Format(Long)` - "February 9, 2026" vs "9. Februar 2026"
- Custom layouts with month/weekday names
- `Human()` - "10 months, 16 days" vs "10 Monate, 16 Tage"

### i18n Does NOT Apply To

- ISO, EU, US, RFC2822 formats (always language-independent)
- Numeric outputs (WeekNumber, Quarter, etc.)

## Code Style

### Go Idioms

- Follow standard Go conventions (go fmt, go vet)
- Use godoc comments for all exported types/functions
- Prefer composition over inheritance
- Keep functions focused and single-purpose

### Documentation

Every exported function needs:
1. Godoc comment starting with function name
2. Example test in `example_test.go` for key functions
3. Edge case documentation where applicable

```go
// Add adds the specified number of units to the date.
// When adding months, if the target day doesn't exist, it snaps to month end.
//
// Example:
//   quando.From(date).Add(2, quando.Months)
func (d Date) Add(value int, unit Unit) Date {
    // ...
}
```

## Beads Workflow

This project uses `bd` (beads) for issue tracking:

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --status in_progress  # Claim work
bd close <id>         # Complete work
bd sync               # Sync with git
```

See AGENTS.md for detailed workflow including the two-phase implement/finalize process.

## Definition of Done

A feature is complete when:

1. **Implementation**: Code follows Go idioms, all exports documented
2. **Tests**: Min 95% coverage, edge cases covered, benchmarks meet goals
3. **Documentation**: Godoc complete, README updated, example tests added
4. **CI/CD**: All tests pass, linting clean (when CI is configured)

## Common Pitfalls

1. **Month arithmetic**: Don't forget month-end overflow logic - this is complex
2. **DST transitions**: Test across DST boundaries in multiple timezones
3. **ISO 8601 week numbers**: Week 1 is the first week with Thursday, not January 1st
4. **Immutability**: Every method must return new Date, never modify receiver
5. **Time zones**: Always handle invalid IANA names with errors, never panic
6. **Parsing ambiguity**: Slash dates without year prefix must error

## Development Priorities (Phase 1)

1. Core infrastructure (Date type, conversions, Clock abstraction)
2. Arithmetic operations (Add, Sub with month-end overflow)
3. Snap operations (StartOf, EndOf, Next, Prev)
4. Difference calculations (Diff, Duration, Human format)
5. Parsing (automatic, explicit, relative)
6. Formatting (presets, custom layouts, i18n)
7. Date inspection (WeekNumber, Quarter, DayOfYear, etc.)
8. Timezone support and DST handling

Total estimated timeline: 14 weeks (~3.5 months) per PRD.

## Future Phases (Out of Scope for Phase 1)

- HTTP/API layer (separate web server)
- Holidays & business days (Phase 3)
- Date series/ranges (Phase 2)
- Batch operations (Phase 2)
- Additional languages beyond EN/DE (22 more languages planned)
