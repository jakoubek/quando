// Package quando provides intuitive and idiomatic date calculations for Go.
//
// quando wraps the standard library's time.Time and provides a fluent API
// for complex date operations that are cumbersome with the standard library alone.
//
// Key features:
//   - Fluent API for method chaining
//   - Month-end aware arithmetic (Jan 31 + 1 month = Feb 28)
//   - DST-safe calendar-based operations
//   - Zero dependencies (stdlib only)
//   - Immutable and thread-safe
//   - Internationalization support (EN, DE in Phase 1)
//   - Clock abstraction for testable code
//
// # Quick Start
//
// Get the current date:
//
//	now := quando.Now()
//
// Create from time.Time:
//
//	date := quando.From(time.Now())
//
// Create from Unix timestamp:
//
//	date := quando.FromUnix(1707480000)
//
// Convert back to time.Time:
//
//	t := date.Time()
//
// # Design Principles
//
// Immutability: All operations return new Date instances. Original values
// are never modified, making Date thread-safe by design.
//
// Fluent API: Methods can be chained naturally:
//
//	result := quando.Now().Add(2, Months).StartOf(Week)
//
// Stdlib Delegation: quando wraps time.Time rather than reimplementing
// time calculations, ensuring correctness and compatibility.
//
// No Panics: All errors are returned as values (except Must* variants
// intended for tests/initialization).
package quando

// Version is the current version of the quando library.
const Version = "0.1.0"
