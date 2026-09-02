package testutils

import "time"

func TestingTime() time.Time {
	return time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC)
}
