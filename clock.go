package quando

import "time"

// Clock provides an abstraction for time operations to enable deterministic testing.
// Use DefaultClock in production and FixedClock in tests.
//
// Example production code:
//
//	clock := quando.NewClock()
//	date := clock.Now()
//
// Example test code:
//
//	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
//	clock := quando.NewFixedClock(fixedTime)
//	date := clock.Now() // Always returns Feb 9, 2026
type Clock interface {
	// Now returns the current date according to this clock.
	Now() Date

	// From converts a time.Time to a Date using this clock's configuration.
	From(t time.Time) Date
}

// DefaultClock is the standard clock implementation that uses the system time.
// It returns the actual current time when Now() is called.
type DefaultClock struct{}

// NewClock returns a new DefaultClock that uses the system time.
// This is the clock to use in production code.
//
// Example:
//
//	clock := quando.NewClock()
//	now := clock.Now()
func NewClock() Clock {
	return &DefaultClock{}
}

// Now returns the current date using the system time.
func (c *DefaultClock) Now() Date {
	return Now()
}

// From converts a time.Time to a Date.
func (c *DefaultClock) From(t time.Time) Date {
	return From(t)
}

// FixedClock is a clock implementation that always returns the same time.
// This is useful for deterministic testing.
type FixedClock struct {
	fixedTime time.Time
}

// NewFixedClock returns a new FixedClock that always returns the specified time.
// This is primarily intended for testing.
//
// Example:
//
//	// In tests
//	fixedTime := time.Date(2026, 2, 9, 12, 0, 0, 0, time.UTC)
//	clock := quando.NewFixedClock(fixedTime)
//	date := clock.Now() // Always returns Feb 9, 2026 12:00:00
func NewFixedClock(t time.Time) Clock {
	return &FixedClock{fixedTime: t}
}

// Now returns the fixed time configured for this clock.
func (c *FixedClock) Now() Date {
	return From(c.fixedTime)
}

// From converts a time.Time to a Date.
// For FixedClock, this behaves the same as the DefaultClock.
func (c *FixedClock) From(t time.Time) Date {
	return From(t)
}
