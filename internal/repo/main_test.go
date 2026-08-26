package repo_test

import (
	"context"
	"fmt"
	"os"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/repo"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

var testMongo *db.MongoDB

type Puzzle struct {
	ID   string `bson:"_id"`
	Mode string `bson:"mode"`
}

// Initial setup of test container for integration tests
func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:7", testcontainers.WithLogger(log.NewNoopLogger()))
	if err != nil {
		fmt.Printf("failed to start mongodb container: %v\n", err)
		return 1
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate container: %v\n", err)
		}
	}()

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		fmt.Printf("failed to get connection string: %v\n", err)
		return 1
	}

	testMongo, err = db.Connect(connStr, "terrariadle_test")
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		return 1
	}
	defer db.Close(ctx, testMongo)

	return m.Run()
}

// Helper method that creates a new connection name for each test
func freshCollection(t *testing.T) string {
	t.Helper()
	name := "repo_" + t.Name()
	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, name)
	})
	return name
}

func generateAnswerData1() repo.AnswerData {
	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: 42,
			PrevWeaponID:    17,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: []int{1, 2, 3, 4},
			Options: []repo.ConnectionOption{
				{Option: "Zenith", CategoryID: 1},
				{Option: "Terra Blade", CategoryID: 2},
			},
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       22,
			Quote:       "Nurses heal wounds. I heal broken bones.",
			Name:        "Nurse",
			NameOptions: []string{"Nurse", "Guide", "Merchant"},
		},
		Hangman: repo.HangmanData{
			EnemyID: 13,
		},
		TerraTrivia: repo.TerraTriviaData{
			QuestionIDs: []int{101, 102, 103},
		},
		ResetTime:     time.Date(2026, 8, 26, 0, 0, 0, 0, time.UTC),
		NextResetTime: time.Date(2026, 8, 27, 0, 0, 0, 0, time.UTC),
	}
}

func generateAnswerData2() repo.AnswerData {
	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: 30,
			PrevWeaponID:    10,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: []int{4, 5, 6, 7},
			Options: []repo.ConnectionOption{
				{Option: "Golden Delight", CategoryID: 3},
				{Option: "Skeleton", CategoryID: 4},
			},
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       22,
			Quote:       "Hunters shoot bows. I shoot guns.",
			Name:        "Arms Dealer",
			NameOptions: []string{"Pirate", "Angler", "Princess"},
		},
		Hangman: repo.HangmanData{
			EnemyID: 15,
		},
		TerraTrivia: repo.TerraTriviaData{
			QuestionIDs: []int{111, 112, 113},
		},
		ResetTime:     time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
		NextResetTime: time.Date(2026, 9, 27, 0, 0, 0, 0, time.UTC),
	}
}

func generateGuessCounts1() repo.PlayerGuessCounts {
	return repo.PlayerGuessCounts{
		DailySlashCount:  1,
		ConnectionsCount: 1,
		GuessTheNpcCount: 1,
		HangmanCount:     1,
		TerraTriviaCount: 1,
	}
}

func generateGuessCounts2() repo.PlayerGuessCounts {
	return repo.PlayerGuessCounts{
		DailySlashCount:  2,
		ConnectionsCount: 2,
		GuessTheNpcCount: 2,
		HangmanCount:     2,
		TerraTriviaCount: 2,
	}
}
