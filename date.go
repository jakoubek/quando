package quando

import (
	"fmt"
	"time"
)

// Lang represents a language for internationalization (i18n) in formatting.
//
// Phase 1 supports English (EN) and German (DE). Future phases will expand to
// include 21 additional languages: FR, ES, IT, PT, NL, PL, RU, JA, ZH, KO, AR,
// HI, TR, SV, NO, DA, FI, CS, HU, RO, UK, EL.
//
// Language affects:
//   - Format(Long): month and weekday names
//   - FormatLayout: custom layouts with month/weekday names
//   - Duration.Human(): time unit names
//
// Language does NOT affect:
//   - ISO, EU, US, RFC2822 formats (always language-independent)
//   - Numeric outputs (WeekNumber, Quarter, DayOfYear)
//
// See i18n.go for translation data and helper methods.
type Lang string

const (
	// EN represents English language.
	EN Lang = "en"
	// DE represents German (Deutsch) language.
	DE Lang = "de"
)

// Date wraps time.Time and provides a fluent API for date operations.
// All operations return new Date instances, making Date immutable and thread-safe.
//
// The Date type supports the full range of Go's time.Time (approximately
// year 0001 to year 9999, with extensions beyond that range).
type Date struct {
	t    time.Time
	lang Lang
}

// Now returns a Date representing the current moment in time.
// The Date uses the local timezone by default.
//
// Example:
//
//	now := quando.Now()
func Now() Date {
	return Date{
		t:    time.Now(),
		lang: EN, // Default language
	}
}

// From converts a time.Time to a Date.
// This is the primary way to create a Date from an existing time value.
//
// Example:
//
//	t := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
//	date := quando.From(t)
func From(t time.Time) Date {
	return Date{
		t:    t,
		lang: EN, // Default language
	}
}

// FromUnix creates a Date from a Unix timestamp (seconds since January 1, 1970 UTC).
// Supports negative timestamps for dates before 1970.
//
// Example:
//
//	date := quando.FromUnix(1707480000) // Feb 9, 2024
//	past := quando.FromUnix(-946684800)  // Jan 1, 1940
func FromUnix(sec int64) Date {
	return Date{
		t:    time.Unix(sec, 0),
		lang: EN, // Default language
	}
}

// Time returns the underlying time.Time value.
// Use this to convert back to standard library time when needed.
//
// Example:
//
//	date := quando.Now()
//	t := date.Time()
func (d Date) Time() time.Time {
	return d.t
}

// Unix returns the Unix timestamp (seconds since January 1, 1970 UTC).
// The value may be negative for dates before 1970.
//
// Example:
//
//	date := quando.Now()
//	timestamp := date.Unix()
func (d Date) Unix() int64 {
	return d.t.Unix()
}

// WithLang returns a new Date with the specified language for formatting.
// This does not modify the date or time, only the language used for formatting operations.
//
// Example:
//
//	date := quando.Now().WithLang(quando.DE)
func (d Date) WithLang(lang Lang) Date {
	return Date{
		t:    d.t,
		lang: lang,
	}
}

// In converts the date to the specified IANA timezone.
// Returns error for invalid timezone names. Never panics.
//
// The method uses the IANA Timezone Database (e.g., "America/New_York", "Europe/Berlin", "UTC").
// Daylight Saving Time (DST) transitions are handled automatically by the timezone database.
//
// When combined with arithmetic operations, DST-safe behavior is preserved:
// Add(1, Days) means "same wall clock time on next calendar day", not "24 hours later".
//
// Example:
//
//	utc := quando.From(time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC))
//	berlin, err := utc.In("Europe/Berlin")
//	// berlin is 2026-06-15 14:00:00 CEST (UTC+2 in summer)
//
// For a list of valid timezone names, see: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
func (d Date) In(location string) (Date, error) {
	// Validate input
	if location == "" {
		return Date{}, fmt.Errorf("timezone location is empty: %w", ErrInvalidTimezone)
	}

	// Load timezone from IANA database
	loc, err := time.LoadLocation(location)
	if err != nil {
		return Date{}, fmt.Errorf("loading timezone %q: %w", location, ErrInvalidTimezone)
	}

	// Convert time to new timezone
	converted := d.t.In(loc)

	// Return new Date with converted time, preserving language
	return Date{
		t:    converted,
		lang: d.lang,
	}, nil
}

// String returns the ISO 8601 representation of the date (YYYY-MM-DD HH:MM:SS).
// This method is called automatically by fmt.Println and similar functions.
func (d Date) String() string {
	return d.t.Format("2006-01-02 15:04:05")
}
