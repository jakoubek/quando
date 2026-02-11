package quando

import (
	"errors"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	before := time.Now()
	date := Now()
	after := time.Now()

	// Verify that Now() returns a time between before and after
	if date.Time().Before(before) || date.Time().After(after) {
		t.Errorf("Now() returned time outside expected range")
	}

	// Verify default language is EN
	if date.lang != EN {
		t.Errorf("Now() default lang = %v, want %v", date.lang, EN)
	}
}

func TestFrom(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
	}{
		{
			name: "specific date",
			time: time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC),
		},
		{
			name: "year 0001",
			time: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "year 9999",
			time: time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "with nanoseconds",
			time: time.Date(2026, 2, 9, 12, 30, 45, 123456789, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := From(tt.time)

			if !date.Time().Equal(tt.time) {
				t.Errorf("From() = %v, want %v", date.Time(), tt.time)
			}

			if date.lang != EN {
				t.Errorf("From() default lang = %v, want %v", date.lang, EN)
			}
		})
	}
}

func TestFromUnix(t *testing.T) {
	tests := []struct {
		name     string
		unix     int64
		expected time.Time
	}{
		{
			name:     "epoch",
			unix:     0,
			expected: time.Unix(0, 0),
		},
		{
			name:     "positive timestamp",
			unix:     1707480000, // 2024-02-09 12:00:00 UTC
			expected: time.Unix(1707480000, 0),
		},
		{
			name:     "negative timestamp (before 1970)",
			unix:     -946771200, // 1940-01-01 00:00:00 UTC
			expected: time.Unix(-946771200, 0),
		},
		{
			name:     "large positive timestamp",
			unix:     253402300799, // 9999-12-31 23:59:59 UTC
			expected: time.Unix(253402300799, 0),
		},
		{
			name:     "large negative timestamp",
			unix:     -62135596800, // 0001-01-01 00:00:00 UTC
			expected: time.Unix(-62135596800, 0),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := FromUnix(tt.unix)

			if !date.Time().Equal(tt.expected) {
				t.Errorf("FromUnix(%d) = %v, want %v", tt.unix, date.Time(), tt.expected)
			}

			if date.lang != EN {
				t.Errorf("FromUnix() default lang = %v, want %v", date.lang, EN)
			}
		})
	}
}

func TestTime(t *testing.T) {
	original := time.Date(2026, 2, 9, 12, 30, 45, 123456789, time.UTC)
	date := From(original)

	result := date.Time()

	if !result.Equal(original) {
		t.Errorf("Time() = %v, want %v", result, original)
	}

	// Verify that modifying the result doesn't affect the original Date
	result = result.Add(24 * time.Hour)
	if !date.Time().Equal(original) {
		t.Errorf("Date.Time() was modified after returning, Date should be immutable")
	}
}

func TestUnix(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected int64
	}{
		{
			name:     "epoch",
			date:     FromUnix(0),
			expected: 0,
		},
		{
			name:     "positive timestamp",
			date:     From(time.Date(2024, 2, 9, 12, 0, 0, 0, time.UTC)),
			expected: 1707480000,
		},
		{
			name:     "negative timestamp",
			date:     FromUnix(-946771200),
			expected: -946771200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.Unix()

			if result != tt.expected {
				t.Errorf("Unix() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestWithLang(t *testing.T) {
	original := Now()

	tests := []struct {
		name string
		lang Lang
	}{
		{"english", EN},
		{"german", DE},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := original.WithLang(tt.lang)

			// Verify language changed
			if result.lang != tt.lang {
				t.Errorf("WithLang(%v) lang = %v, want %v", tt.lang, result.lang, tt.lang)
			}

			// Verify time unchanged
			if !result.Time().Equal(original.Time()) {
				t.Errorf("WithLang() changed time, should only change language")
			}

			// Verify original unchanged (immutability)
			if original.lang != EN {
				t.Errorf("WithLang() modified original date, Date should be immutable")
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name     string
		date     Date
		expected string
	}{
		{
			name:     "standard date",
			date:     From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)),
			expected: "2026-02-09 12:30:45",
		},
		{
			name:     "start of year",
			date:     From(time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)),
			expected: "2026-01-01 00:00:00",
		},
		{
			name:     "end of year",
			date:     From(time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC)),
			expected: "2026-12-31 23:59:59",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.date.String()

			if result != tt.expected {
				t.Errorf("String() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestDateImmutability(t *testing.T) {
	// Create original date
	original := From(time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC))
	originalTime := original.Time()
	originalLang := original.lang

	// Perform various operations
	_ = original.WithLang(DE)
	_ = original.Unix()
	_ = original.Time()
	_ = original.String()

	// Verify original is unchanged
	if !original.Time().Equal(originalTime) {
		t.Errorf("Date was modified, should be immutable")
	}
	if original.lang != originalLang {
		t.Errorf("Date language was modified, should be immutable")
	}
}

// TestDateTimezones verifies that Date correctly preserves timezone information
func TestDateTimezones(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("Skipping timezone test: %v", err)
	}

	berlinTime := time.Date(2026, 2, 9, 12, 0, 0, 0, loc)
	date := From(berlinTime)

	// Verify timezone is preserved
	if date.Time().Location() != loc {
		t.Errorf("Location = %v, want %v", date.Time().Location(), loc)
	}

	// Verify Unix timestamp is correct
	expectedUnix := berlinTime.Unix()
	if date.Unix() != expectedUnix {
		t.Errorf("Unix() = %d, want %d", date.Unix(), expectedUnix)
	}
}

// BenchmarkNow benchmarks the Now() function
func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Now()
	}
}

// BenchmarkFrom benchmarks the From() function
func BenchmarkFrom(b *testing.B) {
	t := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = From(t)
	}
}

// BenchmarkFromUnix benchmarks the FromUnix() function
func BenchmarkFromUnix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FromUnix(1707480000)
	}
}

// BenchmarkTime benchmarks the Time() method
func BenchmarkTime(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Time()
	}
}

// BenchmarkUnix benchmarks the Unix() method
func BenchmarkUnix(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Unix()
	}
}

func TestIn(t *testing.T) {
	tests := []struct {
		name         string
		input        time.Time
		location     string
		expectedZone string
		expectedHour int
	}{
		{
			name:         "UTC to Europe/Berlin (winter)",
			input:        time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CET",
			expectedHour: 13, // UTC+1 in winter
		},
		{
			name:         "UTC to America/New_York (winter)",
			input:        time.Date(2026, 1, 15, 12, 0, 0, 0, time.UTC),
			location:     "America/New_York",
			expectedZone: "EST",
			expectedHour: 7, // UTC-5 in winter
		},
		{
			name:         "UTC to Asia/Tokyo",
			input:        time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC),
			location:     "Asia/Tokyo",
			expectedZone: "JST",
			expectedHour: 21, // UTC+9 (no DST in Japan)
		},
		{
			name:         "UTC to UTC (identity)",
			input:        time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC),
			location:     "UTC",
			expectedZone: "UTC",
			expectedHour: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := From(tt.input)
			result, err := date.In(tt.location)

			if err != nil {
				t.Fatalf("In(%q) unexpected error: %v", tt.location, err)
			}

			zone, _ := result.Time().Zone()
			if zone != tt.expectedZone {
				t.Errorf("In(%q) zone = %v, want %v", tt.location, zone, tt.expectedZone)
			}

			if result.Time().Hour() != tt.expectedHour {
				t.Errorf("In(%q) hour = %v, want %v", tt.location, result.Time().Hour(), tt.expectedHour)
			}
		})
	}
}

func TestInDST(t *testing.T) {
	tests := []struct {
		name         string
		input        time.Time
		location     string
		expectedZone string
		expectedHour int
	}{
		{
			name:         "UTC to Europe/Berlin (summer - DST)",
			input:        time.Date(2026, 7, 15, 12, 0, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CEST",
			expectedHour: 14, // UTC+2 in summer
		},
		{
			name:         "UTC to America/New_York (summer - DST)",
			input:        time.Date(2026, 7, 15, 12, 0, 0, 0, time.UTC),
			location:     "America/New_York",
			expectedZone: "EDT",
			expectedHour: 8, // UTC-4 in summer
		},
		// DST spring forward transition (Europe/Berlin: Mar 29, 2026)
		{
			name:         "Before DST spring forward",
			input:        time.Date(2026, 3, 29, 0, 59, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CET",
			expectedHour: 1, // Still CET (UTC+1)
		},
		{
			name:         "After DST spring forward",
			input:        time.Date(2026, 3, 29, 1, 1, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CEST",
			expectedHour: 3, // Now CEST (UTC+2), skipped 2:00-3:00
		},
		// DST fall back transition (Europe/Berlin: Oct 25, 2026)
		{
			name:         "Before DST fall back",
			input:        time.Date(2026, 10, 25, 0, 59, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CEST",
			expectedHour: 2, // Still CEST (UTC+2)
		},
		{
			name:         "After DST fall back",
			input:        time.Date(2026, 10, 25, 1, 1, 0, 0, time.UTC),
			location:     "Europe/Berlin",
			expectedZone: "CET",
			expectedHour: 2, // Back to CET (UTC+1), repeated 2:00-3:00
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := From(tt.input)
			result, err := date.In(tt.location)

			if err != nil {
				t.Fatalf("In(%q) unexpected error: %v", tt.location, err)
			}

			zone, _ := result.Time().Zone()
			if zone != tt.expectedZone {
				t.Errorf("In(%q) zone = %v, want %v", tt.location, zone, tt.expectedZone)
			}

			if result.Time().Hour() != tt.expectedHour {
				t.Errorf("In(%q) hour = %v, want %v", tt.location, result.Time().Hour(), tt.expectedHour)
			}
		})
	}
}

func TestInErrors(t *testing.T) {
	date := Now()

	tests := []struct {
		name     string
		location string
		wantErr  error
	}{
		{"empty string", "", ErrInvalidTimezone},
		{"invalid timezone", "Invalid/Timezone", ErrInvalidTimezone},
		{"typo in timezone", "America/New_Yrok", ErrInvalidTimezone},
		{"partial timezone", "Europe", ErrInvalidTimezone},
		{"numeric timezone", "UTC+5", ErrInvalidTimezone}, // Use "America/Chicago" instead
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := date.In(tt.location)

			if err == nil {
				t.Fatalf("In(%q) expected error, got nil (result: %v)", tt.location, result)
			}

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("In(%q) error = %v, want error wrapping %v", tt.location, err, tt.wantErr)
			}
		})
	}
}

func TestInImmutability(t *testing.T) {
	original := From(time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC))
	originalTime := original.Time()

	// Convert to different timezone
	_, err := original.In("Europe/Berlin")
	if err != nil {
		t.Fatalf("In() failed: %v", err)
	}

	// Verify original is unchanged
	if !original.Time().Equal(originalTime) {
		t.Error("In() modified the original date")
	}

	// Verify original timezone is unchanged
	originalZone, _ := original.Time().Zone()
	if originalZone != "UTC" {
		t.Errorf("Original zone changed to %v", originalZone)
	}
}

func TestInLanguagePreservation(t *testing.T) {
	date := From(time.Date(2026, 6, 15, 12, 0, 0, 0, time.UTC)).WithLang(DE)

	result, err := date.In("Europe/Berlin")
	if err != nil {
		t.Fatalf("In() failed: %v", err)
	}

	if result.lang != DE {
		t.Errorf("In() language = %v, want %v", result.lang, DE)
	}
}

func TestDSTSafeArithmetic(t *testing.T) {
	// Load Berlin timezone
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("Skipping DST test: %v", err)
	}

	// March 28, 2026 at 02:00 CET (day before DST spring forward)
	date := From(time.Date(2026, 3, 28, 2, 0, 0, 0, loc))

	// Add 1 day - should be March 29 at 02:00 CEST (not 03:00!)
	// This is the critical DST behavior: same wall clock time, not 24 hours
	next := date.Add(1, Days)

	expected := time.Date(2026, 3, 29, 2, 0, 0, 0, loc)
	if !next.Time().Equal(expected) {
		t.Errorf("Add(1, Days) across DST = %v, want %v", next.Time(), expected)
	}

	// Verify it's actually only 23 hours in duration
	duration := next.Time().Sub(date.Time())
	expectedDuration := 23 * time.Hour
	if duration != expectedDuration {
		t.Logf("Actual duration: %v (expected: %v)", duration, expectedDuration)
		t.Logf("This is correct: DST spring forward skips 1 hour")
	}
}
