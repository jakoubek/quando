package quando

import (
	"testing"
)

func TestUnitConstants(t *testing.T) {
	// Verify that all constants are defined and have expected values
	tests := []struct {
		name     string
		unit     Unit
		expected int
		str      string
	}{
		{"Seconds", Seconds, 0, "seconds"},
		{"Minutes", Minutes, 1, "minutes"},
		{"Hours", Hours, 2, "hours"},
		{"Days", Days, 3, "days"},
		{"Weeks", Weeks, 4, "weeks"},
		{"Months", Months, 5, "months"},
		{"Quarters", Quarters, 6, "quarters"},
		{"Years", Years, 7, "years"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify the numeric value
			if int(tt.unit) != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, int(tt.unit), tt.expected)
			}

			// Verify the string representation
			if tt.unit.String() != tt.str {
				t.Errorf("%s.String() = %q, want %q", tt.name, tt.unit.String(), tt.str)
			}
		})
	}
}

func TestUnitOrdering(t *testing.T) {
	// Verify that units are ordered from smallest to largest (except Quarters)
	if Seconds >= Minutes {
		t.Error("Seconds should be < Minutes")
	}
	if Minutes >= Hours {
		t.Error("Minutes should be < Hours")
	}
	if Hours >= Days {
		t.Error("Hours should be < Days")
	}
	if Days >= Weeks {
		t.Error("Days should be < Weeks")
	}
	if Weeks >= Months {
		t.Error("Weeks should be < Months")
	}
	if Months >= Quarters {
		t.Error("Months should be < Quarters")
	}
	if Quarters >= Years {
		t.Error("Quarters should be < Years")
	}
}

func TestUnitString(t *testing.T) {
	tests := []struct {
		name     string
		unit     Unit
		expected string
	}{
		{"Seconds", Seconds, "seconds"},
		{"Minutes", Minutes, "minutes"},
		{"Hours", Hours, "hours"},
		{"Days", Days, "days"},
		{"Weeks", Weeks, "weeks"},
		{"Months", Months, "months"},
		{"Quarters", Quarters, "quarters"},
		{"Years", Years, "years"},
		{"Unknown", Unit(999), "unknown"},
		{"Negative", Unit(-1), "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.unit.String()
			if result != tt.expected {
				t.Errorf("Unit(%d).String() = %q, want %q", tt.unit, result, tt.expected)
			}
		})
	}
}

// TestUnitTypeSafety verifies that Unit provides compile-time type safety
func TestUnitTypeSafety(t *testing.T) {
	// This test primarily exists to demonstrate type safety at compile time
	// If this compiles, the type safety is working

	var u Unit
	u = Seconds
	u = Minutes
	u = Hours
	u = Days
	u = Weeks
	u = Months
	u = Quarters
	u = Years

	// Verify last assignment worked
	if u != Years {
		t.Error("Type safety test failed")
	}
}

// TestUnitComparability verifies that Units can be compared
func TestUnitComparability(t *testing.T) {
	if Seconds == Minutes {
		t.Error("Seconds should not equal Minutes")
	}
	if Seconds == Seconds {
		// This should be true
	} else {
		t.Error("Seconds should equal itself")
	}

	// Test in switch statement (common usage pattern)
	u := Days
	switch u {
	case Seconds, Minutes, Hours:
		t.Error("Days matched wrong case")
	case Days:
		// Expected
	case Weeks, Months, Quarters, Years:
		t.Error("Days matched wrong case")
	default:
		t.Error("Days didn't match any case")
	}
}

