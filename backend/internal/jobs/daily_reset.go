package jobs

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
	"terrariadle-backend/internal/utils"
	"time"
)

// Starts the daily reset job, which randomizes the puzzle data at midnight everynight.
// Call it from a goroutine in main().
func StartResetJob() {
	// For Boise time: loc, _ := time.LoadLocation("America/Boise")
	loc := time.Now().Location()

	// Gets the game data to check if it's already updated
	gameData, err := models.GetGameData()
	if err != nil {
		fmt.Print(err)
		reset(loc)
	} else {
		// Initial reset for if the server crashes
		t := time.Now().In(loc)
		if t.After(gameData.NextResetTime) {
			reset(loc)
		} else {
			// Puts game data in memory if same day
			store.GameData.Set(*gameData)
			fmt.Println("Game data successfully restored")
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		now := time.Now().In(loc)
		// next := utils.NextMidnight(now)
		next := utils.NextShortTime(now)
		timer := time.NewTimer(time.Until(next))

		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			reset(loc)
			fmt.Println("Weapons reset on:", time.Now().In(loc))
			// then loop to compute the *next* midnight again
		}
	}
}

/* Dont store guess amounts on the backend. Instead, read the length of guesses
and assign them to each guess_counts (in case the server crashes)*/

func reset(loc *time.Location) {
	fmt.Println("Reseting daily game data at:", time.Now())

	err := models.DeleteUserData()
	if err != nil {
		log.Fatal(err)
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	game, err := models.GetGameData()
	if err != nil {
		log.Fatal("failed to get game data when resetting")
	}

	weapons, err := randomDailyWeapons(rnd, game.DailySlash.CurrentWeapon)
	if err != nil {
		log.Fatal(err)
	}

	categories, err := randomCategories(rnd)
	if err != nil {
		log.Fatal(err)
	}

	npc, err := randomNpcData(rnd)
	if err != nil {
		log.Fatal(err)
	}

	enemy, err := randomEnemy(rnd)
	if err != nil {
		log.Fatal(err)
	}

	guessCounts := resetGuessCounts()

	gameData := models.GameData{
		DailySlash:    weapons,
		Connections:   categories,
		GuessTheNpc:   npc,
		Hangman:       enemy,
		GuessCounts:   guessCounts,
		ResetTime:     time.Now().In(loc),
		NextResetTime: utils.NextMidnight(time.Now().In(loc)),
	}

	if err := models.UpsertGameData(gameData); err != nil {
		log.Fatal(err)
	}

	store.GameData.Set(gameData)
}
