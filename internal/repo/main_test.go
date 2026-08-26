package repo_test

import (
	"context"
	"fmt"
	"os"
	"terrariadle-backend/internal/db"
	"testing"

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
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:7",
		testcontainers.WithLogger(log.NewNoopLogger()),
	)
	if err != nil {
		fmt.Printf("failed to start mongodb container: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate container: %v\n", err)
		}
	}()

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		fmt.Printf("failed to get connection string: %v\n", err)
		os.Exit(1)
	}

	testMongo, err = db.Connect(connStr, "terrariadle_test")
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(ctx, testMongo)

	os.Exit(m.Run())
}

// Helper method that creates a new connection name for each test
func freshCollection(t *testing.T) string {
	t.Helper()
	name := "puzzles_" + t.Name()
	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, name)
	})
	return name
}