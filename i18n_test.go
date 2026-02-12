package quando

import (
	"testing"
	"time"
	"unicode/utf8"
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

func TestNewLanguagesMonthNames(t *testing.T) {
	// Representative tests for new languages (Jan, Jun, Dec)
	tests := []struct {
		lang     Lang
		month    time.Month
		expected string
	}{
		// Spanish
		{ES, time.January, "enero"},
		{ES, time.June, "junio"},
		{ES, time.December, "diciembre"},
		// French
		{FR, time.January, "janvier"},
		{FR, time.June, "juin"},
		{FR, time.December, "décembre"},
		// Italian
		{IT, time.January, "gennaio"},
		{IT, time.June, "giugno"},
		{IT, time.December, "dicembre"},
		// Portuguese
		{PT, time.January, "janeiro"},
		{PT, time.June, "junho"},
		{PT, time.December, "dezembro"},
		// Dutch
		{NL, time.January, "januari"},
		{NL, time.June, "juni"},
		{NL, time.December, "december"},
		// Polish
		{PL, time.January, "styczeń"},
		{PL, time.June, "czerwiec"},
		{PL, time.December, "grudzień"},
		// Russian
		{RU, time.January, "январь"},
		{RU, time.June, "июнь"},
		{RU, time.December, "декабрь"},
		// Turkish
		{TR, time.January, "Ocak"},
		{TR, time.June, "Haziran"},
		{TR, time.December, "Aralık"},
		// Vietnamese
		{VI, time.January, "Tháng 1"},
		{VI, time.June, "Tháng 6"},
		{VI, time.December, "Tháng 12"},
		// Japanese
		{JA, time.January, "1月"},
		{JA, time.June, "6月"},
		{JA, time.December, "12月"},
		// Korean
		{KO, time.January, "1월"},
		{KO, time.June, "6월"},
		{KO, time.December, "12월"},
		// Chinese Simplified
		{ZhCN, time.January, "一月"},
		{ZhCN, time.June, "六月"},
		{ZhCN, time.December, "十二月"},
		// Chinese Traditional
		{ZhTW, time.January, "一月"},
		{ZhTW, time.June, "六月"},
		{ZhTW, time.December, "十二月"},
		// Hindi
		{HI, time.January, "जनवरी"},
		{HI, time.June, "जून"},
		{HI, time.December, "दिसंबर"},
		// Thai
		{TH, time.January, "มกราคม"},
		{TH, time.June, "มิถุนายน"},
		{TH, time.December, "ธันวาคม"},
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

func TestNewLanguagesWeekdayNames(t *testing.T) {
	// Representative tests for new languages (Mon, Sat)
	tests := []struct {
		lang     Lang
		weekday  time.Weekday
		expected string
	}{
		// Spanish
		{ES, time.Monday, "lunes"},
		{ES, time.Saturday, "sábado"},
		// French
		{FR, time.Monday, "lundi"},
		{FR, time.Saturday, "samedi"},
		// Italian
		{IT, time.Monday, "lunedì"},
		{IT, time.Saturday, "sabato"},
		// Portuguese
		{PT, time.Monday, "segunda-feira"},
		{PT, time.Saturday, "sábado"},
		// Dutch
		{NL, time.Monday, "maandag"},
		{NL, time.Saturday, "zaterdag"},
		// Polish
		{PL, time.Monday, "poniedziałek"},
		{PL, time.Saturday, "sobota"},
		// Russian
		{RU, time.Monday, "понедельник"},
		{RU, time.Saturday, "суббота"},
		// Turkish
		{TR, time.Monday, "Pazartesi"},
		{TR, time.Saturday, "Cumartesi"},
		// Vietnamese
		{VI, time.Monday, "Thứ Hai"},
		{VI, time.Saturday, "Thứ Bảy"},
		// Japanese
		{JA, time.Monday, "月曜日"},
		{JA, time.Saturday, "土曜日"},
		// Korean
		{KO, time.Monday, "월요일"},
		{KO, time.Saturday, "토요일"},
		// Chinese Simplified
		{ZhCN, time.Monday, "星期一"},
		{ZhCN, time.Saturday, "星期六"},
		// Chinese Traditional
		{ZhTW, time.Monday, "星期一"},
		{ZhTW, time.Saturday, "星期六"},
		// Hindi
		{HI, time.Monday, "सोमवार"},
		{HI, time.Saturday, "शनिवार"},
		// Thai
		{TH, time.Monday, "วันจันทร์"},
		{TH, time.Saturday, "วันเสาร์"},
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

func TestNewLanguagesDurationUnits(t *testing.T) {
	// Representative tests for new languages (year and day, singular/plural)
	tests := []struct {
		lang     Lang
		unit     string
		plural   bool
		expected string
	}{
		// Spanish
		{ES, "year", false, "año"},
		{ES, "year", true, "años"},
		{ES, "day", false, "día"},
		{ES, "day", true, "días"},
		// French
		{FR, "year", false, "an"},
		{FR, "year", true, "ans"},
		{FR, "month", false, "mois"},
		{FR, "month", true, "mois"},
		// Italian
		{IT, "year", false, "anno"},
		{IT, "year", true, "anni"},
		{IT, "hour", false, "ora"},
		{IT, "hour", true, "ore"},
		// Portuguese
		{PT, "year", false, "ano"},
		{PT, "year", true, "anos"},
		{PT, "day", false, "dia"},
		{PT, "day", true, "dias"},
		// Dutch
		{NL, "year", false, "jaar"},
		{NL, "year", true, "jaar"},
		{NL, "month", false, "maand"},
		{NL, "month", true, "maanden"},
		// Polish (using common plural form)
		{PL, "year", false, "rok"},
		{PL, "year", true, "lata"},
		{PL, "month", false, "miesiąc"},
		{PL, "month", true, "miesiące"},
		// Russian (using common plural form)
		{RU, "year", false, "год"},
		{RU, "year", true, "года"},
		{RU, "month", false, "месяц"},
		{RU, "month", true, "месяца"},
		// Turkish
		{TR, "year", false, "yıl"},
		{TR, "year", true, "yıl"},
		{TR, "day", false, "gün"},
		{TR, "day", true, "gün"},
		// Vietnamese
		{VI, "year", false, "năm"},
		{VI, "year", true, "năm"},
		{VI, "day", false, "ngày"},
		{VI, "day", true, "ngày"},
		// Japanese
		{JA, "year", false, "年"},
		{JA, "year", true, "年"},
		{JA, "month", false, "月"},
		{JA, "month", true, "月"},
		// Korean
		{KO, "year", false, "년"},
		{KO, "year", true, "년"},
		{KO, "month", false, "월"},
		{KO, "month", true, "월"},
		// Chinese Simplified
		{ZhCN, "year", false, "年"},
		{ZhCN, "year", true, "年"},
		{ZhCN, "hour", false, "小时"},
		{ZhCN, "hour", true, "小时"},
		// Chinese Traditional
		{ZhTW, "year", false, "年"},
		{ZhTW, "year", true, "年"},
		{ZhTW, "hour", false, "小時"},
		{ZhTW, "hour", true, "小時"},
		// Hindi
		{HI, "year", false, "वर्ष"},
		{HI, "year", true, "वर्ष"},
		{HI, "month", false, "महीना"},
		{HI, "month", true, "महीने"},
		// Thai
		{TH, "year", false, "ปี"},
		{TH, "year", true, "ปี"},
		{TH, "day", false, "วัน"},
		{TH, "day", true, "วัน"},
	}

	for _, tt := range tests {
		pluralStr := "singular"
		if tt.plural {
			pluralStr = "plural"
		}
		name := string(tt.lang) + "_" + tt.unit + "_" + pluralStr
		t.Run(name, func(t *testing.T) {
			result := tt.lang.DurationUnit(tt.unit, tt.plural)
			if result != tt.expected {
				t.Errorf("DurationUnit(%v, %v, %v) = %v, want %v", tt.lang, tt.unit, tt.plural, result, tt.expected)
			}
		})
	}
}

func TestUTF8Validity(t *testing.T) {
	// Verify UTF-8 encoding for non-Latin scripts
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		// Russian (Cyrillic)
		{"Russian January", RU.MonthName(time.January), "январь"},
		{"Russian Monday", RU.WeekdayName(time.Monday), "понедельник"},
		// Japanese (Kanji)
		{"Japanese January", JA.MonthName(time.January), "1月"},
		{"Japanese Monday", JA.WeekdayName(time.Monday), "月曜日"},
		{"Japanese weekday short", JA.WeekdayNameShort(time.Monday), "月"},
		// Korean (Hangul)
		{"Korean January", KO.MonthName(time.January), "1월"},
		{"Korean Monday", KO.WeekdayName(time.Monday), "월요일"},
		// Chinese Simplified
		{"Chinese Simplified January", ZhCN.MonthName(time.January), "一月"},
		{"Chinese Simplified Monday", ZhCN.WeekdayName(time.Monday), "星期一"},
		// Chinese Traditional
		{"Chinese Traditional January", ZhTW.MonthName(time.January), "一月"},
		{"Chinese Traditional Monday", ZhTW.WeekdayName(time.Monday), "星期一"},
		// Hindi (Devanagari)
		{"Hindi January", HI.MonthName(time.January), "जनवरी"},
		{"Hindi Monday", HI.WeekdayName(time.Monday), "सोमवार"},
		// Thai
		{"Thai January", TH.MonthName(time.January), "มกราคม"},
		{"Thai Monday", TH.WeekdayName(time.Monday), "วันจันทร์"},
		// Polish special characters
		{"Polish special char", PL.MonthName(time.January), "styczeń"},
		{"Polish weekday", PL.WeekdayName(time.Wednesday), "środa"},
		// Turkish special characters
		{"Turkish special char", TR.MonthName(time.February), "Şubat"},
		{"Turkish weekday", TR.WeekdayName(time.Wednesday), "Çarşamba"},
		// Spanish special characters
		{"Spanish special char", ES.WeekdayName(time.Wednesday), "miércoles"},
		// French special characters
		{"French special char", FR.MonthName(time.February), "février"},
		{"French December", FR.MonthNameShort(time.December), "déc."},
		// Portuguese special characters
		{"Portuguese March", PT.MonthName(time.March), "março"},
		{"Portuguese month", PT.DurationUnit("month", false), "mês"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, tt.value)
			}
			// Verify string is valid UTF-8
			if !isValidUTF8(tt.value) {
				t.Errorf("String %v is not valid UTF-8", tt.value)
			}
		})
	}
}

// isValidUTF8 checks if a string is valid UTF-8
func isValidUTF8(s string) bool {
	return utf8.ValidString(s)
}
