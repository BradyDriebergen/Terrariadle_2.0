package jobs

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/types"
	"terrariadle-backend/internal/utils"
	"terrariadle-backend/internal/utils/cache"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// StartMidnightReset blocks until ctx is canceled.
// Call it from a goroutine in main().
func StartResetJob() {
	// For Boise time: loc, _ := time.LoadLocation("America/Boise")
	loc := time.Now().Location()

	col := db.GetCollection("terrariadle", "daily_data")

	// Gets the game data to check if it's already updated
	gameData, err := db.GetGameData(col, bson.M{"_id": 1})
	if err != nil {
		fmt.Print(err)
		reset(col, loc)
	} else {
		// Initial reset for if the server crashes
		t := time.Now().In(loc)
		if t.After(gameData.NextResetTime) {
			reset(col, loc)
		} else {
			// Puts game data in memory if same day
			cache.SetGameData(*gameData)
			fmt.Println("Game data successfully restored")
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		now := time.Now().In(loc)
		next := utils.NextMidnight(now)
		timer := time.NewTimer(time.Until(next))

		select {
		case <-ctx.Done():
			timer.Stop()
			return
		case <-timer.C:
			reset(col, loc)
			// then loop to compute the *next* midnight again
		}
	}
}

/* Dont store guess amounts on the backend. Instead, read the length of guesses
and assign them to each guess_counts (in case the server crashes)*/

func reset(col *mongo.Collection, loc *time.Location) {
	fmt.Println("Reseting daily game data at:", time.Now())

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	game, err := db.GetGameData(col, bson.M{"_id": 1})
	if err != nil {
		log.Fatal("failed to get game data when resetting")
	}

	weapons, err := randomDailyWeapons(rnd, game.DailySlash.PreviousWeapon)
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

	gameData := types.GameData{
		DailySlash:    weapons,
		Connections:   categories,
		GuessTheNpc:   npc,
		Hangman:       enemy,
		GuessCounts:   guessCounts,
		ResetTime:     time.Now().In(loc),
		NextResetTime: utils.NextMidnight(time.Now().In(loc)),
	}

	if err := db.UpsertRecord(col, bson.M{"_id": 1}, bson.M{"$set": gameData}); err != nil {
		log.Fatal(err)
	}

	cache.SetGameData(gameData)
	fmt.Println("Game data stored in cache")
}
