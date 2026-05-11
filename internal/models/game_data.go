package models

import (
	"context"
	"fmt"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/jsonreader"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GameData represents the the saved daily game data structure

type GameData struct {
	DailySlash    WeaponData            `bson:"dailySlash" json:"dailySlash"`
	Connections   []jsonreader.Category `bson:"connections" json:"connections"`
	GuessTheNpc   NPCdata               `bson:"guessTheNpc" json:"guessTheNpc"`
	Hangman       jsonreader.Enemy      `bson:"hangman" json:"hangman"`
	GuessCounts   PlayerGuessCounts     `bson:"guessCounts" json:"guessCounts"`
	ResetTime     time.Time             `bson:"resetTime" json:"resetTime"`
	NextResetTime time.Time             `bson:"nextResetTime" json:"nextResetTime"`
}

type PlayerGuessCounts struct {
	DailySlashCount  int `bson:"dailySlashCount" json:"dailySlashCount"`
	ConnectionsCount int `bson:"connectionsCount" json:"connectionsCount"`
	GuessTheNpcCount int `bson:"guessTheNpcCount" json:"guessTheNpcCount"`
	HangmanCount     int `bson:"hangmanCount" json:"hangmanCount"`
}

type WeaponData struct {
	CurrentWeapon  jsonreader.Weapon `json:"currentWeapon"`
	PreviousWeapon jsonreader.Weapon `json:"previousWeapon"`
}

type NPCdata struct {
	ID      int      `json:"id"`
	NPC     string   `json:"npc"`
	NPCPath string   `json:"npc-path"`
	Quote   string   `json:"quote"`
	Names   []string `json:"names"`
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
