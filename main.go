package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"terrariadle/internal/api"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"terrariadle/internal/jobs"
	"terrariadle/internal/repo"
	"terrariadle/internal/services"
	"terrariadle/internal/store"
	"terrariadle/internal/web"
	"time"

	"github.com/joho/godotenv"
)

type stores struct {
	user       *store.CachedUserStore
	catalog    *store.CachedCatalogStore
	answer     *store.CachedAnswerStore
	guessCount *store.CachedGuessCountsStore
}

type gameServices struct {
	dailySlash  *services.DailySlash
	connections *services.Connections
	guessTheNpc *services.GuessTheNpc
	hangman     *services.Hangman
	terraTrivia *services.TerraTrivia
	sseService  *services.Common
}

// Variable assigned by build flag for version number provided by GitHub
var version = "dev"

func main() {
	uri := mustLoadURI()

	// Connects to the database
	mongoDB, err := db.Connect(uri, "terrariadle")
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer db.Close(context.Background(), mongoDB)
	log.Println("Database connected")

	jobCtx, cancelJobs := context.WithCancel(context.Background())
	defer cancelJobs()

	// Builds dependencies for jobs and server
	sseBroker := domain.NewBroker()
	s := mustCreateStores(jobCtx, mongoDB, sseBroker)
	svc := createServices(s)

	startBackgroundJobs(jobCtx, s)

	// Starts the server
	srv := api.NewServer(
		context.Background(),
		":8080",
		svc.dailySlash,
		svc.connections,
		svc.guessTheNpc,
		svc.hangman,
		svc.terraTrivia,
		svc.sseService,
		sseBroker,
		web.Assets(),
	)
	go func() {
		if err := srv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	log.Printf("Starting Terrariadle version %s, listening on port :8080\n", version)

	// On quit, shuts down the app cleanly
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancelJobs()

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
	log.Println("server stopped cleanly")
}

func mustLoadURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("dotenv load failed: %v (cwd=%s). Falling back to system env", err, wd)
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MongoDB URI not defined")
	}
	return uri
}

func mustCreateStores(ctx context.Context, mongoDB *db.MongoDB, broker *domain.Broker) stores {
	userRepo := repo.NewUserRepo(
		mongoDB,
		"user_data",
	)
	catalogRepo := repo.NewCatalogRepo(
		mongoDB,
		"daily_slash_weapons",
		"connections_categories",
		"guess_the_npc_npcs",
		"hangman_enemies",
		"terratrivia_questions",
	)
	answerRepo := repo.NewAnswerRepo(
		mongoDB,
		"daily_solutions",
		"player_guess_counts",
	)

	catalog, err := store.NewCatalogStore(ctx, catalogRepo)
	if err != nil {
		log.Fatalf("catalog store: %v", err)
	}
	answer, err := store.NewAnswerStore(ctx, answerRepo, catalog)
	if err != nil {
		log.Fatalf("answer store: %v", err)
	}
	guessCount, err := store.NewGuessCountStore(ctx, answerRepo, broker)
	if err != nil {
		log.Fatalf("guess count store: %v", err)
	}
	user := store.NewUserStore(userRepo)

	log.Println("Stores created")
	return stores{
		user:       user,
		catalog:    catalog,
		answer:     answer,
		guessCount: guessCount,
	}
}

func createServices(s stores) gameServices {
	log.Println("Services created")
	return gameServices{
		dailySlash:  services.NewDailySlashGame(s.answer, s.guessCount, s.catalog, s.user),
		connections: services.NewConnectionsGame(s.answer, s.guessCount, s.catalog, s.user),
		guessTheNpc: services.NewGuessTheNpcGame(s.answer, s.guessCount, s.catalog, s.user),
		hangman:     services.NewHangmanGame(s.answer, s.guessCount, s.catalog, s.user),
		terraTrivia: services.NewTerraTriviaGame(s.answer, s.guessCount, s.catalog, s.user),
		sseService:  services.NewSseStream(s.guessCount, s.user),
	}
}

func startBackgroundJobs(ctx context.Context, s stores) {
	puzzleRefresh := jobs.NewPuzzleRefresh(s.answer, s.guessCount, s.catalog, s.user)
	go puzzleRefresh.Start(ctx)
	go jobs.StartFlushJob(ctx, s.user)
	log.Println("Background jobs started")
}
