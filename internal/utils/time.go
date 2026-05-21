package utils

import "time"

// Returns the start of the next day in t's location.
func NextMidnight(t time.Time) time.Time {
	// Midnight at the *start* of the next day in the same location.
	y, m, d := t.Date()
	loc := t.Location()
	return time.Date(y, m, d+1, 0, 0, 0, 0, loc)
}

// Returns the duration until the next midnight from t.
func TimeUntilNextMidnight(t time.Time) time.Duration {
	return NextMidnight(t).Sub(t)
}

// Returns the duration until the next midnight from now.
func TimeUntilNextMidnightFromNow() time.Duration {
	return TimeUntilNextMidnight(time.Now())
}

// Used for testing, get's the next ten seconds
func NextShortTime(t time.Time) time.Time {
	return t.Truncate(10 * time.Second).Add(10 * time.Second)
}
