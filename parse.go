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
