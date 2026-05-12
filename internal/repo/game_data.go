package repo

import (
	"context"
	"fmt"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type GameRepo struct {
	database *db.MongoDB
}

func NewGameRepo(db *db.MongoDB) *UserRepo {
	return &UserRepo{
		database: db,
	}
}

// GetGameData retrieves the game data from the specified MongoDB collection
func GetGameData() (*GameData, error) {
	collection := db.GetCollection("terrariadle", "daily_data")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data GameData
	res := collection.FindOne(ctx, bson.M{"_id": 1})
	err := res.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to find record %v", err)
	}

	return &data, nil
}

func UpsertGameData(data GameData) error {
	collection := db.GetCollection("terrariadle", "daily_data")
	return db.UpsertRecord(collection, bson.M{"_id": 1}, bson.M{"$set": data})
}

func (r *GameRepo) GetGameData(ctx context.Context, userId string) (domain.User, error) {
	user, err := db.FindOne[gameData](ctx, r.database, "daily_game_data", db.Filter{"id": 1})
	if err != nil {
		return domain.User{}, err
	}

	return user.toDomain(), nil
}

func (r *GameRepo) UpsertGameData(ctx context.Context, user *domain.User) error {
	err := db.Upsert(ctx, r.database, "daily_game_data", db.Filter{"id": 1}, fromDomain(*user))
	return err
}
