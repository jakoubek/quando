package quando

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	// Fixed date for testing: Feb 9, 2026 at 12:30:45 UTC
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))

	tests := []struct {
		name     string
		format   Format
		expected string
	}{
		// ISO format
		{
			name:     "ISO format",
			format:   ISO,
			expected: "2026-02-09",
		},

		// EU format
		{
			name:     "EU format",
			format:   EU,
			expected: "09.02.2026",
		},

		// US format
		{
			name:     "US format",
			format:   US,
			expected: "02/09/2026",
		},

		// Long format (default EN)
		{
			name:     "Long format (EN)",
			format:   Long,
			expected: "February 9, 2026",
		},

		// RFC2822 format
		{
			name:     "RFC2822 format",
			format:   RFC2822,
			expected: "Mon, 09 Feb 2026 12:30:45 +0000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := date.Format(tt.format)
			if result != tt.expected {
				t.Errorf("Format(%v) = %q, want %q", tt.format, result, tt.expected)
			}
		})
	}
}

func TestFormatLong_LanguageDependency(t *testing.T) {
	// Fixed date for testing: Feb 9, 2026
	baseDate := From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))

	tests := []struct {
		name     string
		lang     Lang
		expected string
	}{
		{
			name:     "Long format EN",
			lang:     EN,
			expected: "February 9, 2026",
		},
		{
			name:     "Long format DE",
			lang:     DE,
			expected: "9. Februar 2026",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := baseDate.WithLang(tt.lang)
			result := date.Format(Long)
			if result != tt.expected {
				t.Errorf("Format(Long) with lang=%v = %q, want %q", tt.lang, result, tt.expected)
			}
		})
	}
}

func TestFormat_LanguageIndependence(t *testing.T) {
	// Verify that ISO, EU, US, RFC2822 formats ignore Lang setting
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))

	formats := []struct {
		format   Format
		expected string
	}{
		{ISO, "2026-02-09"},
		{EU, "09.02.2026"},
		{US, "02/09/2026"},
		{RFC2822, "Mon, 09 Feb 2026 12:30:45 +0000"},
	}

	for _, tc := range formats {
		t.Run(tc.format.String(), func(t *testing.T) {
			// Test with EN
			resultEN := date.WithLang(EN).Format(tc.format)
			// Test with DE
			resultDE := date.WithLang(DE).Format(tc.format)

			// Both should be identical (language-independent)
			if resultEN != tc.expected {
				t.Errorf("Format(%v) with EN = %q, want %q", tc.format, resultEN, tc.expected)
			}
			if resultDE != tc.expected {
				t.Errorf("Format(%v) with DE = %q, want %q", tc.format, resultDE, tc.expected)
			}
			if resultEN != resultDE {
				t.Errorf("Format(%v) should be language-independent: EN=%q, DE=%q", tc.format, resultEN, resultDE)
			}
		})
	}
}

func TestFormat_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		format   Format
		expected string
	}{
		// Leap year
		{
			name:     "Leap year Feb 29",
			date:     From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)),
			format:   ISO,
			expected: "2024-02-29",
		},
		{
			name:     "Leap year Feb 29 (Long EN)",
			date:     From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)),
			format:   Long,
			expected: "February 29, 2024",
		},
		{
			name:     "Leap year Feb 29 (Long DE)",
			date:     From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)).WithLang(DE),
			format:   Long,
			expected: "29. Februar 2024",
		},

		// Year boundaries
		{
			name:     "New Year's Day",
			date:     From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)),
			format:   ISO,
			expected: "2026-01-01",
		},
		{
			name:     "Year end",
			date:     From(time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)),
			format:   ISO,
			expected: "2026-12-31",
		},

		// Month boundaries
		{
			name:     "Start of month",
			date:     From(time.Date(2026, 5, 1, 0, 0, 0, 0, time.UTC)),
			format:   EU,
			expected: "01.05.2026",
		},
		{
			name:     "End of month (30 days)",
			date:     From(time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)),
			format:   EU,
			expected: "30.06.2026",
		},
		{
			name:     "End of month (31 days)",
			date:     From(time.Date(2026, 7, 31, 0, 0, 0, 0, time.UTC)),
			format:   EU,
			expected: "31.07.2026",
		},

		// Different months (test all months)
		{
			name:     "January (Long EN)",
			date:     From(time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC)),
			format:   Long,
			expected: "January 15, 2026",
		},
		{
			name:     "December (Long DE)",
			date:     From(time.Date(2026, 12, 15, 0, 0, 0, 0, time.UTC)).WithLang(DE),
			format:   Long,
			expected: "15. Dezember 2026",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.Format(tt.format)
			if result != tt.expected {
				t.Errorf("Format(%v) = %q, want %q", tt.format, result, tt.expected)
			}
		})
	}
}

func TestFormat_Immutability(t *testing.T) {
	// Format should not modify the original Date
	original := From(time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC))

	// Format multiple times
	_ = original.Format(ISO)
	_ = original.Format(EU)
	_ = original.Format(Long)

	// Verify original is unchanged
	expected := time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC)
	if !original.Time().Equal(expected) {
		t.Errorf("Format modified the original date: got %v, want %v", original.Time(), expected)
	}
}


// TestFormatLayout tests basic FormatLayout functionality
func TestFormatLayout(t *testing.T) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC))

	tests := []struct {
		name     string
		lang     Lang
		layout   string
		expected string
	}{
		// English (default)
		{
			name:     "EN: Full month and weekday",
			lang:     EN,
			layout:   "Monday, 2. January 2006",
			expected: "Monday, 9. February 2026",
		},
		{
			name:     "EN: Short month and weekday",
			lang:     EN,
			layout:   "Mon, 02 Jan 2006",
			expected: "Mon, 09 Feb 2026",
		},
		{
			name:     "EN: With time",
			lang:     EN,
			layout:   "Monday, January 2, 2006 at 15:04",
			expected: "Monday, February 9, 2026 at 14:30",
		},

		// German
		{
			name:     "DE: Full month and weekday",
			lang:     DE,
			layout:   "Monday, 2. January 2006",
			expected: "Montag, 9. Februar 2026",
		},
		{
			name:     "DE: Short month and weekday",
			lang:     DE,
			layout:   "Mon, 02 Jan 2006",
			expected: "Mo, 09 Feb 2026",
		},
		{
			name:     "DE: With time",
			lang:     DE,
			layout:   "Monday, 2. January 2006 um 15:04 Uhr",
			expected: "Montag, 9. Februar 2026 um 14:30 Uhr",
		},

		// Empty lang defaults to EN
		{
			name:     "Empty lang defaults to EN",
			lang:     "",
			layout:   "Monday, January 2, 2006",
			expected: "Monday, February 9, 2026",
		},

		// Numeric only (language-independent)
		{
			name:     "Numeric only layout",
			lang:     DE,
			layout:   "2006-01-02 15:04:05",
			expected: "2026-02-09 14:30:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDate := date
			if tt.lang != "" {
				testDate = date.WithLang(tt.lang)
			}
			result := testDate.FormatLayout(tt.layout)
			if result != tt.expected {
				t.Errorf("FormatLayout(%q) with lang=%v = %q, want %q", tt.layout, tt.lang, result, tt.expected)
			}
		})
	}
}

// TestFormatLayout_EdgeCases tests edge cases including all months, all weekdays, and substring collisions
func TestFormatLayout_EdgeCases(t *testing.T) {
	// Test all 12 months - full names
	t.Run("All months full names DE", func(t *testing.T) {
		for m := time.January; m <= time.December; m++ {
			date := From(time.Date(2026, m, 15, 0, 0, 0, 0, time.UTC)).WithLang(DE)
			result := date.FormatLayout("January 2006")
			expectedMonth := DE.MonthName(m)
			expected := expectedMonth + " 2026"
			if result != expected {
				t.Errorf("Month %d: FormatLayout = %q, want %q", m, result, expected)
			}
		}
	})

	// Test all 12 months - short names
	t.Run("All months short names DE", func(t *testing.T) {
		for m := time.January; m <= time.December; m++ {
			date := From(time.Date(2026, m, 15, 0, 0, 0, 0, time.UTC)).WithLang(DE)
			result := date.FormatLayout("Jan 2006")
			expectedMonth := DE.MonthNameShort(m)
			expected := expectedMonth + " 2026"
			if result != expected {
				t.Errorf("Month %d: FormatLayout = %q, want %q", m, result, expected)
			}
		}
	})

	// Test all 7 weekdays - full names
	t.Run("All weekdays full names DE", func(t *testing.T) {
		// Start with Monday, Feb 9, 2026
		for i := 0; i < 7; i++ {
			date := From(time.Date(2026, 2, 9+i, 0, 0, 0, 0, time.UTC)).WithLang(DE)
			weekday := date.Time().Weekday()
			result := date.FormatLayout("Monday")
			expected := DE.WeekdayName(weekday)
			if result != expected {
				t.Errorf("Weekday %v: FormatLayout = %q, want %q", weekday, result, expected)
			}
		}
	})

	// Test all 7 weekdays - short names
	t.Run("All weekdays short names DE", func(t *testing.T) {
		// Start with Monday, Feb 9, 2026
		for i := 0; i < 7; i++ {
			date := From(time.Date(2026, 2, 9+i, 0, 0, 0, 0, time.UTC)).WithLang(DE)
			weekday := date.Time().Weekday()
			result := date.FormatLayout("Mon")
			expected := DE.WeekdayNameShort(weekday)
			if result != expected {
				t.Errorf("Weekday %v: FormatLayout = %q, want %q", weekday, result, expected)
			}
		}
	})

	// Test substring collision: "March" vs "Mar"
	t.Run("Substring collision March/Mar", func(t *testing.T) {
		date := From(time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC)).WithLang(DE)
		// Layout contains both full and short month name
		result := date.FormatLayout("March and Mar")
		expected := "März and Mär"
		if result != expected {
			t.Errorf("Substring collision: FormatLayout = %q, want %q", result, expected)
		}
	})

	// Test complex layout with multiple components
	t.Run("Complex layout with multiple components", func(t *testing.T) {
		date := From(time.Date(2026, 12, 25, 15, 30, 45, 0, time.UTC)).WithLang(DE)
		result := date.FormatLayout("Monday, January 2, 2006 at 15:04:05 (Mon, Jan)")
		expected := "Freitag, Dezember 25, 2026 at 15:30:45 (Fr, Dez)" // Dec 25, 2026 is Friday
		if result != expected {
			t.Errorf("Complex layout: FormatLayout = %q, want %q", result, expected)
		}
	})

	// Test leap year edge case
	t.Run("Leap year Feb 29", func(t *testing.T) {
		date := From(time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)).WithLang(DE)
		result := date.FormatLayout("Monday, January 2, 2006")
		expected := "Donnerstag, Februar 29, 2024"
		if result != expected {
			t.Errorf("Leap year: FormatLayout = %q, want %q", result, expected)
		}
	})
}

// TestFormatLayout_NumericFormatsLanguageIndependent verifies numeric layouts are language-independent
func TestFormatLayout_NumericFormatsLanguageIndependent(t *testing.T) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC))

	numericLayouts := []string{
		"2006-01-02",
		"02.01.2006",
		"01/02/2006",
		"2006-01-02 15:04:05",
		"15:04:05",
		"2006",
		"01",
		"02",
	}

	for _, layout := range numericLayouts {
		t.Run(layout, func(t *testing.T) {
			resultEN := date.WithLang(EN).FormatLayout(layout)
			resultDE := date.WithLang(DE).FormatLayout(layout)

			if resultEN != resultDE {
				t.Errorf("Numeric layout %q should be language-independent: EN=%q, DE=%q", layout, resultEN, resultDE)
			}
		})
	}
}

// TestFormatLayout_Immutability verifies FormatLayout doesn't modify the original Date
func TestFormatLayout_Immutability(t *testing.T) {
	original := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)).WithLang(DE)

	// Call FormatLayout multiple times
	_ = original.FormatLayout("Monday, January 2, 2006")
	_ = original.FormatLayout("Mon, 02 Jan 2006")
	_ = original.FormatLayout("2006-01-02")

	// Verify original is unchanged
	expected := time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)
	if !original.Time().Equal(expected) {
		t.Errorf("FormatLayout modified the original date: got %v, want %v", original.Time(), expected)
	}

	// Verify lang is unchanged
	if original.lang != DE {
		t.Errorf("FormatLayout modified the lang: got %v, want %v", original.lang, DE)
	}
}

// TestFormat_UnknownFormatType tests Format() with an unknown format type
func TestFormat_UnknownFormatType(t *testing.T) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))

	// Cast an invalid int to Format type to test fallback
	unknownFormat := Format(999)

	result := date.Format(unknownFormat)

	// Should fallback to ISO format (documented behavior)
	expected := "2026-02-09"
	if result != expected {
		t.Errorf("Format(unknown) = %q, want %q (ISO fallback)", result, expected)
	}
}

// TestFormatLong_EmptyLang tests formatLong() with empty lang
func TestFormatLong_EmptyLang(t *testing.T) {
	// Create date with empty lang (should default to EN)
	date := Date{
		t:    time.Date(2026, 2, 9, 0, 0, 0, 0, time.UTC),
		lang: "", // Empty lang
	}

	result := date.formatLong()

	// Should default to English format
	expected := "February 9, 2026"
	if result != expected {
		t.Errorf("formatLong() with empty lang = %q, want %q", result, expected)
	}
}

// TestFormatString_Unknown tests String() for Format type with unknown value
func TestFormatString_Unknown(t *testing.T) {
	// Create an unknown format value
	unknownFormat := Format(999)

	result := unknownFormat.String()

	// Should return "unknown" or similar
	if result == "" {
		t.Error("Format.String() for unknown format should not be empty")
	}
}

