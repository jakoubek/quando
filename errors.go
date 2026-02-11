package quando

import "errors"

// Error Handling Policy
//
// The quando library follows a strict no-panic policy for all normal operations.
// Errors are always returned as values, allowing calling code to handle them
// appropriately.
//
// Exception: Must* functions (e.g., MustParse) are provided for convenience
// in tests and initialization code. These functions panic on error and should
// never be used in production code where errors need to be handled gracefully.
//
// Error Types
//
// The library defines sentinel errors for common error conditions. Use errors.Is()
// to check for specific error types:
//
//	date, err := quando.Parse("invalid")
//	if errors.Is(err, quando.ErrInvalidFormat) {
//	    // Handle invalid format
//	}
//
// For errors with context, use fmt.Errorf with %w to wrap errors:
//
//	return Date{}, fmt.Errorf("parsing date %q: %w", input, quando.ErrInvalidFormat)

// ErrInvalidFormat indicates that a date string could not be parsed.
//
// This error is returned when:
//   - The input string doesn't match any supported format
//   - The format is ambiguous (e.g., "01/02/2026" without year prefix)
//   - The date components are invalid (e.g., "2026-13-01" for month 13)
//
// Example:
//
//	date, err := quando.Parse("not-a-date")
//	if errors.Is(err, quando.ErrInvalidFormat) {
//	    log.Printf("Invalid date format: %v", err)
//	}
var ErrInvalidFormat = errors.New("invalid date format")

// ErrInvalidTimezone indicates that a timezone name is not recognized.
//
// This error is returned when:
//   - The IANA timezone name is not found in the system timezone database
//   - The timezone string is malformed
//
// Valid timezone names include "UTC", "America/New_York", "Europe/Berlin", etc.
// See the IANA Time Zone Database for a complete list.
//
// Example:
//
//	date, err := quando.Now().InTimezone("Invalid/Timezone")
//	if errors.Is(err, quando.ErrInvalidTimezone) {
//	    log.Printf("Unknown timezone: %v", err)
//	}
var ErrInvalidTimezone = errors.New("invalid timezone")

// ErrOverflow indicates that a date arithmetic operation would result in
// a date outside the representable range.
//
// This error is returned when:
//   - Adding/subtracting would exceed Go's time.Time range (approximately year 0 to 9999)
//   - The resulting date would be before the minimum or after the maximum representable time
//
// Example:
//
//	date := quando.From(time.Date(9999, 12, 31, 0, 0, 0, 0, time.UTC))
//	result, err := date.Add(1, quando.Years)
//	if errors.Is(err, quando.ErrOverflow) {
//	    log.Printf("Date overflow: %v", err)
//	}
//
// Note: In the current implementation, arithmetic operations that would overflow
// are handled by Go's time.Time, which has its own overflow behavior. This error
// is reserved for future use when explicit overflow detection is added.
var ErrOverflow = errors.New("date overflow")
