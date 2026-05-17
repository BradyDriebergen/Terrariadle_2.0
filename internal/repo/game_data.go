package repo

import (
	"context"
	"terrariadle-backend/internal/db"
)

type GameRepo struct {
	database *db.MongoDB
}

func NewGameRepo(db *db.MongoDB) *UserRepo {
	return &UserRepo{
		database: db,
	}
}

func (r *GameRepo) GetGameData(ctx context.Context, userId string) (GameData, error) {
	gameData, err := db.FindOne[GameData](ctx, r.database, "daily_game_data", db.Filter{"id": 1})
	if err != nil {
		return GameData{}, err
	}

	return *gameData, nil
}

func (r *GameRepo) UpsertGameData(ctx context.Context, gameData *GameData) error {
	err := db.Upsert(ctx, r.database, "daily_game_data", db.Filter{"id": 1}, gameData)
	return err
}
