package quando

import (
	"testing"
	"time"
)

// ============================================================================
// Arithmetic Benchmarks
// ============================================================================

// BenchmarkAddDays benchmarks Add with Days
func BenchmarkAddDays(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Days)
	}
}

// BenchmarkAddMonths benchmarks Add with Months
func BenchmarkAddMonths(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Months)
	}
}

// BenchmarkAddYears benchmarks Add with Years
func BenchmarkAddYears(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Years)
	}
}

// BenchmarkMethodChaining benchmarks chained operations
func BenchmarkMethodChaining(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Add(1, Months).Add(15, Days).Sub(2, Hours)
	}
}

// ============================================================================
// Clock Benchmarks
// ============================================================================

// BenchmarkDefaultClock_Now benchmarks DefaultClock.Now()
func BenchmarkDefaultClock_Now(b *testing.B) {
	clock := NewClock()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = clock.Now()
	}
}

// BenchmarkFixedClock_Now benchmarks FixedClock.Now()
func BenchmarkFixedClock_Now(b *testing.B) {
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(fixedTime)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = clock.Now()
	}
}

// BenchmarkDefaultClock_From benchmarks DefaultClock.From()
func BenchmarkDefaultClock_From(b *testing.B) {
	clock := NewClock()
	t := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = clock.From(t)
	}
}

// BenchmarkFixedClock_From benchmarks FixedClock.From()
func BenchmarkFixedClock_From(b *testing.B) {
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(fixedTime)
	t := time.Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = clock.From(t)
	}
}

// ============================================================================
// Date Conversion Benchmarks
// ============================================================================

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

// ============================================================================
// Diff/Duration Benchmarks
// ============================================================================

// BenchmarkDurationSeconds benchmarks Seconds()
func BenchmarkDurationSeconds(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.Seconds()
	}
}

// BenchmarkDurationDays benchmarks Days()
func BenchmarkDurationDays(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.Days()
	}
}

// BenchmarkDurationMonths benchmarks Months()
func BenchmarkDurationMonths(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.Months()
	}
}

// BenchmarkDurationMonthsFloat benchmarks MonthsFloat()
func BenchmarkDurationMonthsFloat(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC)
	dur := Diff(start, end)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.MonthsFloat()
	}
}

// BenchmarkDurationHuman benchmarks Human() with English
func BenchmarkDurationHuman(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 11, 17, 5, 30, 45, 0, time.UTC)
	dur := Diff(start, end)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.Human(EN)
	}
}

// BenchmarkDurationHumanGerman benchmarks Human() with German
func BenchmarkDurationHumanGerman(b *testing.B) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 11, 17, 5, 30, 45, 0, time.UTC)
	dur := Diff(start, end)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = dur.Human(DE)
	}
}

// ============================================================================
// Format Benchmarks
// ============================================================================

// BenchmarkFormat_ISO benchmarks Format with ISO format
func BenchmarkFormat_ISO(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(ISO)
	}
}

// BenchmarkFormat_EU benchmarks Format with EU format
func BenchmarkFormat_EU(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(EU)
	}
}

// BenchmarkFormat_US benchmarks Format with US format
func BenchmarkFormat_US(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(US)
	}
}

// BenchmarkFormat_Long_EN benchmarks Format with Long format (English)
func BenchmarkFormat_Long_EN(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)).WithLang(EN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(Long)
	}
}

// BenchmarkFormat_Long_DE benchmarks Format with Long format (German)
func BenchmarkFormat_Long_DE(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)).WithLang(DE)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(Long)
	}
}

// BenchmarkFormat_RFC2822 benchmarks Format with RFC2822 format
func BenchmarkFormat_RFC2822(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Format(RFC2822)
	}
}

// BenchmarkFormatLayout_EN_Simple benchmarks FormatLayout with English (simple)
func BenchmarkFormatLayout_EN_Simple(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)).WithLang(EN)
	layout := "Monday, January 2, 2006"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.FormatLayout(layout)
	}
}

// BenchmarkFormatLayout_EN_Numeric benchmarks FormatLayout with English (numeric)
func BenchmarkFormatLayout_EN_Numeric(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)).WithLang(EN)
	layout := "2006-01-02 15:04:05"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.FormatLayout(layout)
	}
}

// BenchmarkFormatLayout_DE_Simple benchmarks FormatLayout with German (simple)
func BenchmarkFormatLayout_DE_Simple(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)).WithLang(DE)
	layout := "Monday, January 2, 2006"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.FormatLayout(layout)
	}
}

// BenchmarkFormatLayout_DE_Complex benchmarks FormatLayout with German (complex)
func BenchmarkFormatLayout_DE_Complex(b *testing.B) {
	date := From(time.Date(2026, 2, 9, 14, 30, 45, 0, time.UTC)).WithLang(DE)
	layout := "Monday, January 2, 2006 at 15:04:05 MST (Mon, Jan)"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.FormatLayout(layout)
	}
}

// ============================================================================
// Inspection Benchmarks
// ============================================================================

// BenchmarkWeekNumber benchmarks WeekNumber()
func BenchmarkWeekNumber(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.WeekNumber()
	}
}

// BenchmarkQuarter benchmarks Quarter()
func BenchmarkQuarter(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.Quarter()
	}
}

// BenchmarkDayOfYear benchmarks DayOfYear()
func BenchmarkDayOfYear(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.DayOfYear()
	}
}

// BenchmarkIsWeekend benchmarks IsWeekend()
func BenchmarkIsWeekend(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.IsWeekend()
	}
}

// BenchmarkIsLeapYear benchmarks IsLeapYear()
func BenchmarkIsLeapYear(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.IsLeapYear()
	}
}

// BenchmarkInfo benchmarks Info()
func BenchmarkInfo(b *testing.B) {
	d := From(time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = d.Info()
	}
}

// ============================================================================
// Parse Benchmarks
// ============================================================================

// BenchmarkParse benchmarks the Parse function with different formats
func BenchmarkParse(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{"ISO format", "2026-02-09"},
		{"ISO slash", "2026/02/09"},
		{"EU format", "09.02.2026"},
		{"RFC2822", "Mon, 09 Feb 2026 00:00:00 +0000"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_, err := Parse(bm.input)
				if err != nil {
					b.Fatalf("Parse failed: %v", err)
				}
			}
		})
	}
}

// BenchmarkParseError benchmarks error case (ambiguous format)
func BenchmarkParseError(b *testing.B) {
	input := "01/02/2026" // Ambiguous format

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := Parse(input)
		if err == nil {
			b.Fatal("Expected error for ambiguous format")
		}
	}
}

// BenchmarkParseWithLayout benchmarks ParseWithLayout
func BenchmarkParseWithLayout(b *testing.B) {
	layout := "02/01/2006"
	input := "09/02/2026"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseWithLayout(input, layout)
	}
}

// BenchmarkParseWithLayoutCustom benchmarks ParseWithLayout with custom layout
func BenchmarkParseWithLayoutCustom(b *testing.B) {
	layout := "2. January 2006"
	input := "9. February 2026"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseWithLayout(input, layout)
	}
}

// BenchmarkParseRelativeKeyword benchmarks ParseRelative with keyword
func BenchmarkParseRelativeKeyword(b *testing.B) {
	clock := NewFixedClock(time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseRelativeWithClock("today", clock)
	}
}

// BenchmarkParseRelativeOffset benchmarks ParseRelative with offset
func BenchmarkParseRelativeOffset(b *testing.B) {
	clock := NewFixedClock(time.Date(2026, 2, 15, 12, 0, 0, 0, time.UTC))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParseRelativeWithClock("+2 days", clock)
	}
}

// ============================================================================
// Snap Benchmarks
// ============================================================================

// BenchmarkStartOfWeek benchmarks StartOf(Weeks)
func BenchmarkStartOfWeek(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Weeks)
	}
}

// BenchmarkEndOfWeek benchmarks EndOf(Weeks)
func BenchmarkEndOfWeek(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Weeks)
	}
}

// BenchmarkStartOfMonth benchmarks StartOf(Months)
func BenchmarkStartOfMonth(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Months)
	}
}

// BenchmarkEndOfMonth benchmarks EndOf(Months)
func BenchmarkEndOfMonth(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Months)
	}
}

// BenchmarkStartOfQuarter benchmarks StartOf(Quarters)
func BenchmarkStartOfQuarter(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Quarters)
	}
}

// BenchmarkEndOfQuarter benchmarks EndOf(Quarters)
func BenchmarkEndOfQuarter(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Quarters)
	}
}

// BenchmarkStartOfYear benchmarks StartOf(Years)
func BenchmarkStartOfYear(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.StartOf(Years)
	}
}

// BenchmarkEndOfYear benchmarks EndOf(Years)
func BenchmarkEndOfYear(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.EndOf(Years)
	}
}

// BenchmarkNext benchmarks the Next() method
func BenchmarkNext(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Next(time.Friday)
	}
}

// BenchmarkPrev benchmarks the Prev() method
func BenchmarkPrev(b *testing.B) {
	date := Now()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = date.Prev(time.Friday)
	}
}

// ============================================================================
// Unit Benchmarks
// ============================================================================

// BenchmarkUnitString benchmarks the String() method
func BenchmarkUnitString(b *testing.B) {
	units := []Unit{Seconds, Minutes, Hours, Days, Weeks, Months, Quarters, Years}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u := units[i%len(units)]
		_ = u.String()
	}
}
