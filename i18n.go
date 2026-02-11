package quando

import "time"

// i18n.go - Internationalization support for quando
//
// This file contains translations for month names, weekday names, and
// duration units used in formatting operations.
//
// Phase 1 Languages:
//   - EN (English) - Default
//   - DE (Deutsch/German)
//
// Future expansion will add 21 more languages including:
//   FR, ES, IT, PT, NL, PL, RU, JA, ZH, KO, AR, HI, TR, SV, NO, DA, FI, CS, HU, RO, UK, EL
//
// i18n applies to:
//   - Format(Long): "February 9, 2026" vs "9. Februar 2026"
//   - FormatLayout with month/weekday names
//   - Duration.Human(): "10 months, 16 days" vs "10 Monate, 16 Tage"
//
// i18n does NOT apply to:
//   - ISO, EU, US, RFC2822 formats (always language-independent)
//   - Numeric outputs (WeekNumber, Quarter, DayOfYear)

// monthNames contains full month name translations.
var monthNames = map[Lang][12]string{
	EN: {
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	},
	DE: {
		"Januar", "Februar", "März", "April", "Mai", "Juni",
		"Juli", "August", "September", "Oktober", "November", "Dezember",
	},
}

// monthNamesShort contains short (3-letter) month name translations.
var monthNamesShort = map[Lang][12]string{
	EN: {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
	DE: {"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
}

// weekdayNames contains full weekday name translations.
// Index: Sunday = 0, Monday = 1, ..., Saturday = 6
var weekdayNames = map[Lang][7]string{
	EN: {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	DE: {"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
}

// weekdayNamesShort contains short (3-letter) weekday name translations.
// Index: Sunday = 0, Monday = 1, ..., Saturday = 6
var weekdayNamesShort = map[Lang][7]string{
	EN: {"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
	DE: {"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
}

// durationUnits contains duration unit translations for Human() formatting.
// Each unit has singular and plural forms: [0] = singular, [1] = plural
var durationUnits = map[Lang]map[string][2]string{
	EN: {
		"year":   {"year", "years"},
		"month":  {"month", "months"},
		"week":   {"week", "weeks"},
		"day":    {"day", "days"},
		"hour":   {"hour", "hours"},
		"minute": {"minute", "minutes"},
		"second": {"second", "seconds"},
	},
	DE: {
		"year":   {"Jahr", "Jahre"},
		"month":  {"Monat", "Monate"},
		"week":   {"Woche", "Wochen"},
		"day":    {"Tag", "Tage"},
		"hour":   {"Stunde", "Stunden"},
		"minute": {"Minute", "Minuten"},
		"second": {"Sekunde", "Sekunden"},
	},
}

// MonthName returns the localized month name for the given language.
// Returns English name if language not found.
func (l Lang) MonthName(month time.Month) string {
	if names, ok := monthNames[l]; ok {
		return names[month-1]
	}
	// Fallback to English
	return monthNames[EN][month-1]
}

// MonthNameShort returns the short (3-letter) localized month name.
// Returns English abbreviation if language not found.
func (l Lang) MonthNameShort(month time.Month) string {
	if names, ok := monthNamesShort[l]; ok {
		return names[month-1]
	}
	return monthNamesShort[EN][month-1]
}

// WeekdayName returns the localized weekday name for the given language.
// Returns English name if language not found.
func (l Lang) WeekdayName(weekday time.Weekday) string {
	if names, ok := weekdayNames[l]; ok {
		return names[weekday]
	}
	return weekdayNames[EN][weekday]
}

// WeekdayNameShort returns the short (3-letter) localized weekday name.
// Returns English abbreviation if language not found.
func (l Lang) WeekdayNameShort(weekday time.Weekday) string {
	if names, ok := weekdayNamesShort[l]; ok {
		return names[weekday]
	}
	return weekdayNamesShort[EN][weekday]
}

// DurationUnit returns the localized duration unit name (singular or plural).
// The plural parameter determines which form to use.
// Returns English name if language not found.
func (l Lang) DurationUnit(unit string, plural bool) string {
	if units, ok := durationUnits[l]; ok {
		if forms, ok := units[unit]; ok {
			if plural {
				return forms[1]
			}
			return forms[0]
		}
	}
	// Fallback to English
	if forms, ok := durationUnits[EN][unit]; ok {
		if plural {
			return forms[1]
		}
		return forms[0]
	}
	return unit
}
