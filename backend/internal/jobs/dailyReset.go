package jobs

import (
	"context"
	"fmt"
	"terrariadle-backend/internal/types"
	"time"
)

// StartMidnightReset blocks until ctx is canceled.
// Call it from a goroutine in main().
func StartResetJob(ctx context.Context, data *types.GameData) {
	// For Boise time: loc, _ := time.LoadLocation("America/Boise")
	loc := time.Now().Location()

	// Initial reset for if the server crashes
	reset(data)

	for {
		now := time.Now().In(loc)
		next := nextMidnight(now)
		timer := time.NewTimer(time.Until(next))

		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			reset(data)
			// then loop to compute the *next* midnight again
		}
	}
}

func nextMidnight(t time.Time) time.Time {
	// Midnight at the *start* of the next day in the same location.
	y, m, d := t.Date()
	loc := t.Location()
	return time.Date(y, m, d+1, 0, 0, 0, 0, loc)
}

/* Dont store guess amounts on the backend. Instead, read the length of guesses
and assign them to each guess_counts (in case the server crashes)*/

func reset(data *types.GameData) {
	fmt.Println("Reseting daily game data at:", time.Now())

	fmt.Printf("type=%T value=%+v\n", data, data) // shows pointer and fields
	if data != nil {
		fmt.Printf("deref=%+v\n", *data) // shows the struct values
	} else {
		fmt.Println("data is nil")
	}
}
