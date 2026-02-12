package quando

import (
	"fmt"
	"strings"
	"time"
)

// Parse automatically detects and parses common date formats.
//
// Supported formats (automatic detection):
//   - ISO format: "2026-02-09" (YYYY-MM-DD)
//   - ISO with slash: "2026/02/09" (YYYY/MM/DD)
//   - EU format: "09.02.2026" (DD.MM.YYYY)
//   - RFC2822: "Mon, 09 Feb 2026 00:00:00 +0000"
//
// Ambiguous format detection:
//
// Slash formats without year prefix are ambiguous and will return an error:
//   - "01/02/2026" - ERROR (could be US: Jan 2 or EU: Feb 1)
//   - "31/12/2024" - ERROR (ambiguous format)
//
// Use ParseWithLayout() for explicit format handling when needed.
//
// The parsed date uses UTC timezone by default (for formats without timezone info).
// The language is set to EN for formatting operations.
//
// Example:
//
//	date, err := quando.Parse("2026-02-09")
//	if err != nil {
//	    return err
//	}
//	fmt.Println(date) // 2026-02-09 00:00:00
//
// Example with error handling:
//
//	date, err := quando.Parse("01/02/2026")
//	if errors.Is(err, quando.ErrInvalidFormat) {
//	    // Ambiguous format - use ParseWithLayout instead
//	}
func Parse(s string) (Date, error) {
	// Trim whitespace
	s = strings.TrimSpace(s)

	// Check for empty string
	if s == "" {
		return Date{}, fmt.Errorf("parsing date %q: empty string: %w", s, ErrInvalidFormat)
	}

	// Check for ambiguous slash format (DD/MM/YYYY or MM/DD/YYYY without year prefix)
	// Pattern: exactly 10 chars, two slashes at positions 2 and 5
	if len(s) == 10 && s[2] == '/' && s[5] == '/' && strings.Count(s, "/") == 2 {
		// Check if it's NOT the ISO format (YYYY/MM/DD)
		// ISO format has year prefix, so first 4 chars should be digits representing year >= 1000
		if !isYearPrefix(s[:4]) {
			return Date{}, fmt.Errorf("parsing date %q: ambiguous format (use ParseWithLayout for slash dates without year prefix): %w", s, ErrInvalidFormat)
		}
	}

	// Try parsing with each supported format
	layouts := []struct {
		layout    string
		validator func(string) bool
	}{
		// ISO format: YYYY-MM-DD
		{
			layout: "2006-01-02",
			validator: func(s string) bool {
				return len(s) == 10 && s[4] == '-' && s[7] == '-' && strings.Count(s, "-") == 2
			},
		},
		// ISO with slash: YYYY/MM/DD
		{
			layout: "2006/01/02",
			validator: func(s string) bool {
				return len(s) == 10 && s[4] == '/' && s[7] == '/' && strings.Count(s, "/") == 2 && isYearPrefix(s[:4])
			},
		},
		// EU format: DD.MM.YYYY
		{
			layout: "02.01.2006",
			validator: func(s string) bool {
				return len(s) == 10 && s[2] == '.' && s[5] == '.' && strings.Count(s, ".") == 2
			},
		},
		// RFC2822 / RFC1123Z format
		{
			layout: time.RFC1123Z,
			validator: func(s string) bool {
				// RFC1123Z is longer and contains commas
				return strings.Contains(s, ",") && len(s) > 20
			},
		},
		// RFC1123 (without timezone)
		{
			layout: time.RFC1123,
			validator: func(s string) bool {
				return strings.Contains(s, ",") && len(s) > 20
			},
		},
	}

	var lastErr error
	for _, lt := range layouts {
		// Quick validation before attempting parse
		if !lt.validator(s) {
			continue
		}

		// Attempt to parse
		t, err := time.Parse(lt.layout, s)
		if err == nil {
			// Successfully parsed
			return Date{
				t:    t,
				lang: EN,
			}, nil
		}
		lastErr = err
	}

	// If we got here, none of the formats worked
	if lastErr != nil {
		return Date{}, fmt.Errorf("parsing date %q: %w", s, ErrInvalidFormat)
	}

	return Date{}, fmt.Errorf("parsing date %q: no matching format: %w", s, ErrInvalidFormat)
}

// MustParse parses a date string using automatic format detection and panics on error.
//
// This is a convenience wrapper around Parse() for use in tests and initialization code
// where the input is known to be valid. If parsing fails, it panics with the parse error.
//
// IMPORTANT: This function panics on error. For production code that needs to handle
// errors gracefully, use Parse() instead.
//
// Supported formats (same as Parse):
//   - ISO format: "2026-02-09" (YYYY-MM-DD)
//   - ISO with slash: "2026/02/09" (YYYY/MM/DD)
//   - EU format: "09.02.2026" (DD.MM.YYYY)
//   - RFC2822: "Mon, 09 Feb 2026 00:00:00 +0000"
//
// Use cases:
//   - Test fixtures: date := quando.MustParse("2026-02-09")
//   - Static initialization: var StartDate = quando.MustParse("2026-01-01")
//   - Configuration where dates are validated at load time
//
// Example (test fixture):
//
//	func TestSomething(t *testing.T) {
//	    date := quando.MustParse("2026-02-09")
//	    // Use date without error handling...
//	}
//
// Example (static initialization):
//
//	var (
//	    EpochDate = quando.MustParse("1970-01-01")
//	    Y2K       = quando.MustParse("2000-01-01")
//	)
func MustParse(s string) Date {
	date, err := Parse(s)
	if err != nil {
		panic(fmt.Sprintf("quando.MustParse(%q): %v", s, err))
	}
	return date
}

// isYearPrefix checks if the first 4 characters represent a valid year (>= 1000).
// This helps distinguish YYYY/MM/DD from DD/MM/YYYY or MM/DD/YYYY.
func isYearPrefix(s string) bool {
	if len(s) != 4 {
		return false
	}

	// Check if all characters are digits
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	// Check if it's a plausible year (>= 1000)
	// Years before 1000 are unlikely in modern applications
	if s[0] == '0' {
		return false
	}

	return true
}

// ParseWithLayout parses a date string using an explicit Go layout format.
// This is useful for disambiguating ambiguous formats or parsing custom formats
// that cannot be automatically detected.
//
// Layout Format:
// Go uses a reference date approach. The layout string must use the reference date:
//   Mon Jan 2 15:04:05 MST 2006
//
// Components you can use in your layout:
//   Year:    2006 (4-digit), 06 (2-digit)
//   Month:   01 (2-digit), 1 (1-digit), Jan (short), January (long)
//   Day:     02 (2-digit), 2 (1-digit), _2 (space-padded)
//   Weekday: Mon (short), Monday (long)
//   Hour:    15 (24-hour), 03 (12-hour), 3 (1-digit 12-hour)
//   Minute:  04 (2-digit), 4 (1-digit)
//   Second:  05 (2-digit), 5 (1-digit)
//   AM/PM:   PM
//   Timezone: MST (abbrev), -0700 (offset), Z07:00 (ISO 8601)
//
// Note: Month and weekday names must be in English (Go limitation).
//
// Examples:
//
//	// Disambiguate US vs EU slash format
//	ParseWithLayout("01/02/2026", "01/02/2006")  // US: January 2, 2026
//	ParseWithLayout("01/02/2026", "02/01/2006")  // EU: February 1, 2026
//
//	// Custom format with text month
//	ParseWithLayout("9. February 2026", "2. January 2006")  // February 9, 2026
//
//	// ISO 8601 with time
//	ParseWithLayout("2026-02-09T14:30:00", "2006-01-02T15:04:05")
//
// If the string cannot be parsed with the given layout, returns an error
// wrapping ErrInvalidFormat. The returned Date uses UTC timezone unless
// the layout and input include timezone information.
func ParseWithLayout(s, layout string) (Date, error) {
	// Trim whitespace from input
	s = strings.TrimSpace(s)

	// Empty input check
	if s == "" {
		return Date{}, fmt.Errorf("parsing date with layout %q: empty input: %w", layout, ErrInvalidFormat)
	}

	// Parse using time.Parse with the provided layout
	t, err := time.Parse(layout, s)
	if err != nil {
		return Date{}, fmt.Errorf("parsing date %q with layout %q: %w", s, layout, ErrInvalidFormat)
	}

	// Wrap in quando.Date with default language
	return Date{t: t, lang: EN}, nil
}

// ParseRelative parses relative date expressions and returns a Date.
//
// Supported expressions:
//   - Keywords: "today", "tomorrow", "yesterday"
//   - Relative offsets: "+N <unit>", "-N <unit>"
//
// Supported units (singular and plural):
//   - day, days
//   - week, weeks
//   - month, months
//   - quarter, quarters
//   - year, years
//
// Examples:
//
//	ParseRelative("today")       // Today at 00:00:00
//	ParseRelative("tomorrow")    // Tomorrow at 00:00:00
//	ParseRelative("yesterday")   // Yesterday at 00:00:00
//	ParseRelative("+2 days")     // Two days from today
//	ParseRelative("-1 week")     // One week ago
//	ParseRelative("+3 months")   // Three months from today
//
// All keywords and unit names are case-insensitive.
// Results are always at 00:00:00 in the local timezone.
//
// Note: Complex expressions like "next monday" or "start of month" are not
// yet supported. Use ParseRelative("+7 days") and StartOf(Months) instead.
//
// Returns an error wrapping ErrInvalidFormat if the expression cannot be parsed.
func ParseRelative(s string) (Date, error) {
	clock := NewClock()
	return ParseRelativeWithClock(s, clock)
}

// ParseRelativeWithClock parses relative date expressions using a specific Clock.
// This is the testable version of ParseRelative that accepts a Clock parameter.
//
// See ParseRelative for supported expressions and usage examples.
func ParseRelativeWithClock(s string, clock Clock) (Date, error) {
	// Trim whitespace and convert to lowercase for case-insensitive matching
	s = strings.TrimSpace(s)
	sLower := strings.ToLower(s)

	// Empty input check
	if s == "" {
		return Date{}, fmt.Errorf("parsing relative date: empty input: %w", ErrInvalidFormat)
	}

	// Get base date (today at 00:00:00 in UTC timezone)
	now := clock.Now()
	t := now.Time()
	loc := t.Location()
	today := Date{
		t:    time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc),
		lang: EN,
	}

	// Handle simple keywords
	switch sLower {
	case "today":
		return today, nil
	case "tomorrow":
		return today.Add(1, Days), nil
	case "yesterday":
		return today.Add(-1, Days), nil
	}

	// Handle relative offset pattern: "+N unit" or "-N unit"
	// Examples: "+2 days", "-1 week", "+3 months"

	// Split on whitespace
	parts := strings.Fields(s)
	if len(parts) != 2 {
		return Date{}, fmt.Errorf("parsing relative date %q: invalid format (expected \"today\", \"tomorrow\", \"yesterday\", or \"+/-N unit\"): %w", s, ErrInvalidFormat)
	}

	offsetStr := parts[0]
	unitStr := strings.ToLower(parts[1])

	// Parse offset (must start with + or -)
	if len(offsetStr) < 2 || (offsetStr[0] != '+' && offsetStr[0] != '-') {
		return Date{}, fmt.Errorf("parsing relative date %q: offset must start with + or - (e.g., \"+2\" or \"-1\"): %w", s, ErrInvalidFormat)
	}

	// Check for invalid characters (like decimal points)
	if strings.Contains(offsetStr, ".") {
		return Date{}, fmt.Errorf("parsing relative date %q: offset must be an integer, not a float: %w", s, ErrInvalidFormat)
	}

	// Parse the number part
	var offset int
	_, err := fmt.Sscanf(offsetStr, "%d", &offset)
	if err != nil {
		return Date{}, fmt.Errorf("parsing relative date %q: invalid offset number %q: %w", s, offsetStr, ErrInvalidFormat)
	}

	// Map unit string to Unit constant
	unit, err := parseUnitString(unitStr)
	if err != nil {
		return Date{}, fmt.Errorf("parsing relative date %q: %w", s, err)
	}

	// Apply the offset
	return today.Add(offset, unit), nil
}

// parseUnitString maps unit name strings to Unit constants.
// Supports both singular and plural forms, case-insensitive.
func parseUnitString(s string) (Unit, error) {
	switch s {
	case "day", "days":
		return Days, nil
	case "week", "weeks":
		return Weeks, nil
	case "month", "months":
		return Months, nil
	case "quarter", "quarters":
		return Quarters, nil
	case "year", "years":
		return Years, nil
	default:
		return 0, fmt.Errorf("unknown unit %q (supported: day, week, month, quarter, year): %w", s, ErrInvalidFormat)
	}
}
