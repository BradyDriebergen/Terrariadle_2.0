package services

import (
	"fmt"
	"terrariadle-backend/internal/db"
)

func GetDailySlashUserGuesses(userId string) ([]int, error) {
	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return []int{}, fmt.Errorf("failed to get user in user guesses API")
	}

	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			return user.Games[i].Guesses, nil
		}
	}

	return []int{}, fmt.Errorf("error getting player position")
}
