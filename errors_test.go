package quando

import (
	"errors"
	"testing"
)

// TestSentinelErrors verifies that all sentinel errors are defined
func TestSentinelErrors(t *testing.T) {
	tests := []struct {
		name  string
		err   error
		check string
	}{
		{"ErrInvalidFormat", ErrInvalidFormat, "invalid date format"},
		{"ErrInvalidTimezone", ErrInvalidTimezone, "invalid timezone"},
		{"ErrOverflow", ErrOverflow, "date overflow"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Errorf("%s is nil, expected error", tt.name)
			}
			if tt.err.Error() != tt.check {
				t.Errorf("%s.Error() = %q, want %q", tt.name, tt.err.Error(), tt.check)
			}
		})
	}
}

// TestErrorsIs verifies that errors.Is works with sentinel errors
func TestErrorsIs(t *testing.T) {
	// Test direct comparison
	if !errors.Is(ErrInvalidFormat, ErrInvalidFormat) {
		t.Error("errors.Is(ErrInvalidFormat, ErrInvalidFormat) should be true")
	}

	if !errors.Is(ErrInvalidTimezone, ErrInvalidTimezone) {
		t.Error("errors.Is(ErrInvalidTimezone, ErrInvalidTimezone) should be true")
	}

	if !errors.Is(ErrOverflow, ErrOverflow) {
		t.Error("errors.Is(ErrOverflow, ErrOverflow) should be true")
	}

	// Test that different errors are not equal
	if errors.Is(ErrInvalidFormat, ErrInvalidTimezone) {
		t.Error("ErrInvalidFormat and ErrInvalidTimezone should not be equal")
	}

	if errors.Is(ErrInvalidTimezone, ErrOverflow) {
		t.Error("ErrInvalidTimezone and ErrOverflow should not be equal")
	}
}

// TestErrorUniqueness verifies that each error is unique
func TestErrorUniqueness(t *testing.T) {
	allErrors := []error{
		ErrInvalidFormat,
		ErrInvalidTimezone,
		ErrOverflow,
	}

	// Check that no two errors are the same
	for i, err1 := range allErrors {
		for j, err2 := range allErrors {
			if i != j && err1 == err2 {
				t.Errorf("Errors at index %d and %d are the same", i, j)
			}
		}
	}
}
