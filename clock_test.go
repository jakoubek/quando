package quando

import (
	"testing"
	"time"
)

func TestNewClock(t *testing.T) {
	clock := NewClock()

	if clock == nil {
		t.Fatal("NewClock() returned nil")
	}

	// Verify it's a DefaultClock
	if _, ok := clock.(*DefaultClock); !ok {
		t.Errorf("NewClock() returned %T, want *DefaultClock", clock)
	}
}

func TestDefaultClock_Now(t *testing.T) {
	clock := NewClock()

	before := time.Now()
	date := clock.Now()
	after := time.Now()

	// Verify that Now() returns a time between before and after
	if date.Time().Before(before) || date.Time().After(after) {
		t.Errorf("DefaultClock.Now() returned time outside expected range")
	}
}

func TestDefaultClock_From(t *testing.T) {
	clock := NewClock()
	testTime := time.Date(2026, 2, 9, 12, 30, 45, 0, time.UTC)

	date := clock.From(testTime)

	if !date.Time().Equal(testTime) {
		t.Errorf("DefaultClock.From() = %v, want %v", date.Time(), testTime)
	}
}

func TestNewFixedClock(t *testing.T) {
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(fixedTime)

	if clock == nil {
		t.Fatal("NewFixedClock() returned nil")
	}

	// Verify it's a FixedClock
	if _, ok := clock.(*FixedClock); !ok {
		t.Errorf("NewFixedClock() returned %T, want *FixedClock", clock)
	}
}

func TestFixedClock_Now(t *testing.T) {
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(fixedTime)

	// Call Now() multiple times to verify it always returns the same time
	date1 := clock.Now()
	time.Sleep(1 * time.Millisecond)
	date2 := clock.Now()
	time.Sleep(1 * time.Millisecond)
	date3 := clock.Now()

	// All should return the same fixed time
	if !date1.Time().Equal(fixedTime) {
		t.Errorf("FixedClock.Now() (call 1) = %v, want %v", date1.Time(), fixedTime)
	}
	if !date2.Time().Equal(fixedTime) {
		t.Errorf("FixedClock.Now() (call 2) = %v, want %v", date2.Time(), fixedTime)
	}
	if !date3.Time().Equal(fixedTime) {
		t.Errorf("FixedClock.Now() (call 3) = %v, want %v", date3.Time(), fixedTime)
	}

	// Verify all three are equal to each other
	if !date1.Time().Equal(date2.Time()) || !date2.Time().Equal(date3.Time()) {
		t.Error("FixedClock.Now() returned different times on successive calls")
	}
}

func TestFixedClock_From(t *testing.T) {
	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(fixedTime)

	// Test with different time
	testTime := time.Date(2025, 5, 15, 8, 30, 0, 0, time.UTC)
	date := clock.From(testTime)

	// From() should use the provided time, not the fixed time
	if !date.Time().Equal(testTime) {
		t.Errorf("FixedClock.From() = %v, want %v", date.Time(), testTime)
	}
}

// TestFixedClock_DeterministicTesting demonstrates how FixedClock enables deterministic tests
func TestFixedClock_DeterministicTesting(t *testing.T) {
	// This test demonstrates a deterministic test pattern using FixedClock

	// Create a fixed clock for a specific test scenario
	testTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
	clock := NewFixedClock(testTime)

	// Use the clock in test code
	now := clock.Now()

	// We can now make deterministic assertions
	expected := "2026-02-09 12:00:00"
	if now.String() != expected {
		t.Errorf("String() = %v, want %v", now.String(), expected)
	}

	expectedUnix := int64(1770638400)
	if now.Unix() != expectedUnix {
		t.Errorf("Unix() = %d, want %d", now.Unix(), expectedUnix)
	}
}

// TestClock_Interface verifies that both implementations satisfy the Clock interface
func TestClock_Interface(t *testing.T) {
	var _ Clock = &DefaultClock{}
	var _ Clock = &FixedClock{}

	// This test will fail at compile time if either type doesn't implement Clock
}

// TestClock_EdgeCases tests edge cases for clock implementations
func TestClock_EdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		fixedTime time.Time
	}{
		{
			name:      "epoch",
			fixedTime: time.Unix(0, 0),
		},
		{
			name:      "year 0001",
			fixedTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "year 9999",
			fixedTime: time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			name:      "with nanoseconds",
			fixedTime: time.Date(2026, 2, 9, 12, 30, 45, 123456789, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clock := NewFixedClock(tt.fixedTime)
			date := clock.Now()

			if !date.Time().Equal(tt.fixedTime) {
				t.Errorf("FixedClock.Now() = %v, want %v", date.Time(), tt.fixedTime)
			}
		})
	}
}

// TestClock_Timezones verifies that clocks preserve timezone information
func TestClock_Timezones(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("Skipping timezone test: %v", err)
	}

	berlinTime := time.Date(2026, 2, 9, 12, 0, 0, 0, loc)

	// Test DefaultClock
	defaultClock := NewClock()
	date1 := defaultClock.From(berlinTime)
	if date1.Time().Location() != loc {
		t.Errorf("DefaultClock.From() location = %v, want %v", date1.Time().Location(), loc)
	}

	// Test FixedClock
	fixedClock := NewFixedClock(berlinTime)
	date2 := fixedClock.Now()
	if date2.Time().Location() != loc {
		t.Errorf("FixedClock.Now() location = %v, want %v", date2.Time().Location(), loc)
	}
}

