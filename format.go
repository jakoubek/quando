package quando

import (
	"fmt"
	"time"
)

// Format represents a preset date format for use with the Format method.
// Each format produces a different string representation of a date.
//
// Language Dependency:
//   - ISO, EU, US, RFC2822: Always language-independent
//   - Long: Uses the Date's Lang setting for month and weekday names
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
//	iso := date.Format(quando.ISO)       // "2026-02-09"
//	long := date.Format(quando.Long)     // "February 9, 2026" (EN)
//	longDE := date.WithLang(quando.DE).Format(quando.Long)  // "9. Februar 2026"
type Format int

const (
	// ISO represents the ISO 8601 date format: "2026-02-09" (YYYY-MM-DD).
	// This format is always language-independent and is the standard international format.
	ISO Format = iota

	// EU represents the European date format: "09.02.2026" (DD.MM.YYYY).
	// This format is always language-independent and uses dots as separators.
	// Common in Germany, Austria, Switzerland, and many other European countries.
	EU

	// US represents the US date format: "02/09/2026" (MM/DD/YYYY).
	// This format is always language-independent and uses slashes as separators.
	// Common in the United States and some other countries.
	US

	// Long represents a human-readable long format with full month name.
	// This format is language-dependent and uses the Date's Lang setting.
	//
	// Examples:
	//   - EN: "February 9, 2026"
	//   - DE: "9. Februar 2026"
	//
	// The format varies by language to match local conventions.
	Long

	// RFC2822 represents the RFC 2822 email date format.
	// Example: "Mon, 09 Feb 2026 12:30:45 +0000"
	// This format is always language-independent and includes time and timezone.
	RFC2822
)

// Format formats the date using the specified preset format.
//
// Supported formats:
//   - ISO: "2026-02-09" (YYYY-MM-DD)
//   - EU: "09.02.2026" (DD.MM.YYYY)
//   - US: "02/09/2026" (MM/DD/YYYY)
//   - Long: "February 9, 2026" (language-dependent)
//   - RFC2822: "Mon, 09 Feb 2026 12:30:45 +0000"
//
// The Long format respects the Date's Lang setting:
//   - EN: "February 9, 2026"
//   - DE: "9. Februar 2026"
//
// All other formats are language-independent.
//
// Example:
//
//	date := quando.From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
//	fmt.Println(date.Format(quando.ISO))     // "2026-02-09"
//	fmt.Println(date.Format(quando.EU))      // "09.02.2026"
//	fmt.Println(date.Format(quando.US))      // "02/09/2026"
//	fmt.Println(date.Format(quando.Long))    // "February 9, 2026"
//	fmt.Println(date.Format(quando.RFC2822)) // "Mon, 09 Feb 2026 12:30:45 +0000"
func (d Date) Format(format Format) string {
	t := d.t

	switch format {
	case ISO:
		// ISO 8601 format: YYYY-MM-DD
		return t.Format("2006-01-02")

	case EU:
		// European format: DD.MM.YYYY
		return t.Format("02.01.2006")

	case US:
		// US format: MM/DD/YYYY
		return t.Format("01/02/2006")

	case Long:
		// Long format with full month name (language-dependent)
		// EN: "February 9, 2026"
		// DE: "9. Februar 2026"
		return d.formatLong()

	case RFC2822:
		// RFC 2822 email format
		return t.Format(time.RFC1123Z)

	default:
		// Fallback to ISO format for unknown formats
		return t.Format("2006-01-02")
	}
}

// formatLong formats the date in long format with language-specific conventions.
// This is a helper method for Format(Long).
func (d Date) formatLong() string {
	t := d.t
	lang := d.lang
	if lang == "" {
		lang = EN // Default to English if no language set
	}

	// Get localized month name
	monthName := lang.MonthName(t.Month())

	// Different formats for different languages
	switch lang {
	case DE:
		// German format: "9. Februar 2026"
		// Pattern: day without leading zero + ". " + month + " " + year
		return fmt.Sprintf("%d. %s %d", t.Day(), monthName, t.Year())

	default:
		// English format (and fallback): "February 9, 2026"
		// Pattern: month + " " + day + ", " + year
		return fmt.Sprintf("%s %d, %d", monthName, t.Day(), t.Year())
	}
}

// String returns the string representation of the Format type.
// This is used for better test output and debugging.
func (f Format) String() string {
	switch f {
	case ISO:
		return "ISO"
	case EU:
		return "EU"
	case US:
		return "US"
	case Long:
		return "Long"
	case RFC2822:
		return "RFC2822"
	default:
		return "Unknown"
	}
}
