package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"terrariadle-backend/internal/db"

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
	defer mongoDB.Close(context.Background())

	fmt.Print("Database connected")

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
