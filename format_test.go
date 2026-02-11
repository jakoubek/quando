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

// Benchmarks
func BenchmarkFormat_ISO(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(ISO)
	}
}

func BenchmarkFormat_EU(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(EU)
	}
}

func BenchmarkFormat_US(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(US)
	}
}

func BenchmarkFormat_Long_EN(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)).WithLang(EN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(Long)
	}
}

func BenchmarkFormat_Long_DE(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)).WithLang(DE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(Long)
	}
}

func BenchmarkFormat_RFC2822(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(RFC2822)
	}
}
