package quando

import (
	"testing"
	"time"
)

func TestMonthName(t *testing.T) {
	tests := []struct {
		lang     Lang
		month    time.Month
		expected string
	}{
		// English months
		{EN, time.January, "January"},
		{EN, time.February, "February"},
		{EN, time.March, "March"},
		{EN, time.April, "April"},
		{EN, time.May, "May"},
		{EN, time.June, "June"},
		{EN, time.July, "July"},
		{EN, time.August, "August"},
		{EN, time.September, "September"},
		{EN, time.October, "October"},
		{EN, time.November, "November"},
		{EN, time.December, "December"},

		// German months
		{DE, time.January, "Januar"},
		{DE, time.February, "Februar"},
		{DE, time.March, "März"},
		{DE, time.April, "April"},
		{DE, time.May, "Mai"},
		{DE, time.June, "Juni"},
		{DE, time.July, "Juli"},
		{DE, time.August, "August"},
		{DE, time.September, "September"},
		{DE, time.October, "Oktober"},
		{DE, time.November, "November"},
		{DE, time.December, "Dezember"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.lang.MonthName(tt.month)
			if result != tt.expected {
				t.Errorf("MonthName(%v, %v) = %v, want %v", tt.lang, tt.month, result, tt.expected)
			}
		})
	}
}

func TestMonthNameShort(t *testing.T) {
	tests := []struct {
		lang     Lang
		month    time.Month
		expected string
	}{
		// English short months
		{EN, time.January, "Jan"},
		{EN, time.February, "Feb"},
		{EN, time.March, "Mar"},
		{EN, time.April, "Apr"},
		{EN, time.May, "May"},
		{EN, time.June, "Jun"},
		{EN, time.July, "Jul"},
		{EN, time.August, "Aug"},
		{EN, time.September, "Sep"},
		{EN, time.October, "Oct"},
		{EN, time.November, "Nov"},
		{EN, time.December, "Dec"},

		// German short months
		{DE, time.January, "Jan"},
		{DE, time.February, "Feb"},
		{DE, time.March, "Mär"},
		{DE, time.April, "Apr"},
		{DE, time.May, "Mai"},
		{DE, time.June, "Jun"},
		{DE, time.July, "Jul"},
		{DE, time.August, "Aug"},
		{DE, time.September, "Sep"},
		{DE, time.October, "Okt"},
		{DE, time.November, "Nov"},
		{DE, time.December, "Dez"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.lang.MonthNameShort(tt.month)
			if result != tt.expected {
				t.Errorf("MonthNameShort(%v, %v) = %v, want %v", tt.lang, tt.month, result, tt.expected)
			}
		})
	}
}

func TestWeekdayName(t *testing.T) {
	tests := []struct {
		lang     Lang
		weekday  time.Weekday
		expected string
	}{
		// English weekdays
		{EN, time.Sunday, "Sunday"},
		{EN, time.Monday, "Monday"},
		{EN, time.Tuesday, "Tuesday"},
		{EN, time.Wednesday, "Wednesday"},
		{EN, time.Thursday, "Thursday"},
		{EN, time.Friday, "Friday"},
		{EN, time.Saturday, "Saturday"},

		// German weekdays
		{DE, time.Sunday, "Sonntag"},
		{DE, time.Monday, "Montag"},
		{DE, time.Tuesday, "Dienstag"},
		{DE, time.Wednesday, "Mittwoch"},
		{DE, time.Thursday, "Donnerstag"},
		{DE, time.Friday, "Freitag"},
		{DE, time.Saturday, "Samstag"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.lang.WeekdayName(tt.weekday)
			if result != tt.expected {
				t.Errorf("WeekdayName(%v, %v) = %v, want %v", tt.lang, tt.weekday, result, tt.expected)
			}
		})
	}
}

func TestWeekdayNameShort(t *testing.T) {
	tests := []struct {
		lang     Lang
		weekday  time.Weekday
		expected string
	}{
		// English short weekdays
		{EN, time.Sunday, "Sun"},
		{EN, time.Monday, "Mon"},
		{EN, time.Tuesday, "Tue"},
		{EN, time.Wednesday, "Wed"},
		{EN, time.Thursday, "Thu"},
		{EN, time.Friday, "Fri"},
		{EN, time.Saturday, "Sat"},

		// German short weekdays
		{DE, time.Sunday, "So"},
		{DE, time.Monday, "Mo"},
		{DE, time.Tuesday, "Di"},
		{DE, time.Wednesday, "Mi"},
		{DE, time.Thursday, "Do"},
		{DE, time.Friday, "Fr"},
		{DE, time.Saturday, "Sa"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.lang.WeekdayNameShort(tt.weekday)
			if result != tt.expected {
				t.Errorf("WeekdayNameShort(%v, %v) = %v, want %v", tt.lang, tt.weekday, result, tt.expected)
			}
		})
	}
}

func TestDurationUnit(t *testing.T) {
	tests := []struct {
		lang     Lang
		unit     string
		plural   bool
		expected string
	}{
		// English singular
		{EN, "year", false, "year"},
		{EN, "month", false, "month"},
		{EN, "week", false, "week"},
		{EN, "day", false, "day"},
		{EN, "hour", false, "hour"},
		{EN, "minute", false, "minute"},
		{EN, "second", false, "second"},

		// English plural
		{EN, "year", true, "years"},
		{EN, "month", true, "months"},
		{EN, "week", true, "weeks"},
		{EN, "day", true, "days"},
		{EN, "hour", true, "hours"},
		{EN, "minute", true, "minutes"},
		{EN, "second", true, "seconds"},

		// German singular
		{DE, "year", false, "Jahr"},
		{DE, "month", false, "Monat"},
		{DE, "week", false, "Woche"},
		{DE, "day", false, "Tag"},
		{DE, "hour", false, "Stunde"},
		{DE, "minute", false, "Minute"},
		{DE, "second", false, "Sekunde"},

		// German plural
		{DE, "year", true, "Jahre"},
		{DE, "month", true, "Monate"},
		{DE, "week", true, "Wochen"},
		{DE, "day", true, "Tage"},
		{DE, "hour", true, "Stunden"},
		{DE, "minute", true, "Minuten"},
		{DE, "second", true, "Sekunden"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.lang.DurationUnit(tt.unit, tt.plural)
			if result != tt.expected {
				t.Errorf("DurationUnit(%v, %v, %v) = %v, want %v", tt.lang, tt.unit, tt.plural, result, tt.expected)
			}
		})
	}
}

func TestLanguageFallback(t *testing.T) {
	// Unknown language should fallback to English
	unknownLang := Lang("unknown")

	t.Run("MonthName fallback", func(t *testing.T) {
		result := unknownLang.MonthName(time.January)
		if result != "January" {
			t.Errorf("Unknown language should fallback to English, got %v", result)
		}
	})

	t.Run("MonthNameShort fallback", func(t *testing.T) {
		result := unknownLang.MonthNameShort(time.February)
		if result != "Feb" {
			t.Errorf("Unknown language should fallback to English, got %v", result)
		}
	})

	t.Run("WeekdayName fallback", func(t *testing.T) {
		result := unknownLang.WeekdayName(time.Monday)
		if result != "Monday" {
			t.Errorf("Unknown language should fallback to English, got %v", result)
		}
	})

	t.Run("WeekdayNameShort fallback", func(t *testing.T) {
		result := unknownLang.WeekdayNameShort(time.Friday)
		if result != "Fri" {
			t.Errorf("Unknown language should fallback to English, got %v", result)
		}
	})

	t.Run("DurationUnit fallback", func(t *testing.T) {
		result := unknownLang.DurationUnit("month", true)
		if result != "months" {
			t.Errorf("Unknown language should fallback to English, got %v", result)
		}
	})
}

func TestDurationUnitUnknownUnit(t *testing.T) {
	// Unknown unit should return the unit string itself
	result := EN.DurationUnit("unknown", false)
	if result != "unknown" {
		t.Errorf("Unknown unit should return unit string, got %v", result)
	}

	result = DE.DurationUnit("unknown", true)
	if result != "unknown" {
		t.Errorf("Unknown unit should return unit string, got %v", result)
	}
}

func TestAllMonthsPresent(t *testing.T) {
	// Ensure all languages have all 12 months
	t.Run("monthNames", func(t *testing.T) {
		for lang, months := range monthNames {
			if len(months) != 12 {
				t.Errorf("Language %v has %d months in monthNames, expected 12", lang, len(months))
			}
			// Check for empty strings
			for i, name := range months {
				if name == "" {
					t.Errorf("Language %v has empty month name at index %d", lang, i)
				}
			}
		}
	})

	t.Run("monthNamesShort", func(t *testing.T) {
		for lang, months := range monthNamesShort {
			if len(months) != 12 {
				t.Errorf("Language %v has %d months in monthNamesShort, expected 12", lang, len(months))
			}
			// Check for empty strings
			for i, name := range months {
				if name == "" {
					t.Errorf("Language %v has empty short month name at index %d", lang, i)
				}
			}
		}
	})
}

func TestAllWeekdaysPresent(t *testing.T) {
	// Ensure all languages have all 7 weekdays
	t.Run("weekdayNames", func(t *testing.T) {
		for lang, weekdays := range weekdayNames {
			if len(weekdays) != 7 {
				t.Errorf("Language %v has %d weekdays in weekdayNames, expected 7", lang, len(weekdays))
			}
			// Check for empty strings
			for i, name := range weekdays {
				if name == "" {
					t.Errorf("Language %v has empty weekday name at index %d", lang, i)
				}
			}
		}
	})

	t.Run("weekdayNamesShort", func(t *testing.T) {
		for lang, weekdays := range weekdayNamesShort {
			if len(weekdays) != 7 {
				t.Errorf("Language %v has %d weekdays in weekdayNamesShort, expected 7", lang, len(weekdays))
			}
			// Check for empty strings
			for i, name := range weekdays {
				if name == "" {
					t.Errorf("Language %v has empty short weekday name at index %d", lang, i)
				}
			}
		}
	})
}

func TestAllDurationUnitsPresent(t *testing.T) {
	// Ensure all languages have all expected duration units
	expectedUnits := []string{"year", "month", "week", "day", "hour", "minute", "second"}

	for lang, units := range durationUnits {
		for _, unit := range expectedUnits {
			forms, ok := units[unit]
			if !ok {
				t.Errorf("Language %v missing duration unit %v", lang, unit)
				continue
			}
			// Check singular
			if forms[0] == "" {
				t.Errorf("Language %v has empty singular form for unit %v", lang, unit)
			}
			// Check plural
			if forms[1] == "" {
				t.Errorf("Language %v has empty plural form for unit %v", lang, unit)
			}
		}
	}
}

func TestGermanSpecialCharacters(t *testing.T) {
	// Verify German special characters are correctly stored
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{"March", DE.MonthName(time.March), "März"},
		{"March short", DE.MonthNameShort(time.March), "Mär"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, tt.value)
			}
		})
	}
}
