package quando

// Unit represents a time unit for arithmetic operations like Add and Sub,
// and for snap operations like StartOf and EndOf.
//
// Use the predefined constants (Seconds, Minutes, Hours, Days, Weeks, Months,
// Quarters, Years) rather than creating Unit values directly.
type Unit int

// Time unit constants for use with arithmetic and snap operations.
// These constants are ordered from smallest to largest unit, except for Quarters
// which is a special alias for 3 months.
const (
	// Seconds represents the seconds time unit.
	Seconds Unit = iota
	// Minutes represents the minutes time unit.
	Minutes
	// Hours represents the hours time unit.
	Hours
	// Days represents the days time unit.
	// When used with Add, this means calendar days (not 24-hour periods).
	Days
	// Weeks represents the weeks time unit (7 days).
	Weeks
	// Months represents the months time unit.
	// Adding months handles month-end overflow (e.g., Jan 31 + 1 month = Feb 28).
	Months
	// Quarters represents the quarters time unit (3 months).
	Quarters
	// Years represents the years time unit.
	Years
)

// String returns the string representation of the Unit.
// This is primarily useful for debugging and error messages.
func (u Unit) String() string {
	switch u {
	case Seconds:
		return "seconds"
	case Minutes:
		return "minutes"
	case Hours:
		return "hours"
	case Days:
		return "days"
	case Weeks:
		return "weeks"
	case Months:
		return "months"
	case Quarters:
		return "quarters"
	case Years:
		return "years"
	default:
		return "unknown"
	}
}
