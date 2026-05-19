package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/repo"
	"terrariadle-backend/internal/store"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Initial grabbing of the URI link
	err := godotenv.Load(".env")
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("dotenv load failed: %v (cwd=%s). Falling back to system env", err, wd)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MongoDB URI not defined")
	}

	// Connect to the database
	mongoDB, err := db.Connect(uri, "terrariadle")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background(), mongoDB)

	fmt.Print("Database connected")

	ur := repo.NewUserRepo(mongoDB)
	us := store.NewUserStore(ur)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := us.GetUser(ctx, "test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", user)

	// ar := repo.NewAnswerRepo(mongoDB)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// answers, err := ar.GetAnswerData(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("\n%+v\n", answers)

	// cr := repo.NewCatalogRepo(mongoDB)

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// result, err := cr.GetEnemies(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("\n%+v\n", result)

	// // Initialize atomicstore
	// store.InitializeStoreFromJson()

	// // Sets up endpoints
	// h := server.NewMux()

	// // Starts the reset job
	// go jobs.StartResetJob()

	// // Starts the server
	// if err := server.RunServer(":3000", h); err != nil {
	// 	log.Fatal(err)
	// }
}
