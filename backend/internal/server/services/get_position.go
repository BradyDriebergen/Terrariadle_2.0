package services

import (
	"fmt"
	"terrariadle-backend/internal/db"
)

func GetDailySlashPlayerPosition(userId string) (int, error) {
	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return 0, fmt.Errorf("failed to get user in player position API")
	}

	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			if len(user.Games[i].Guesses) > 0 {
				return user.Games[i].Position, nil
			}
			return 0, fmt.Errorf("player doesn't exist")
		}
	}

	return 0, fmt.Errorf("error getting player position")
}
