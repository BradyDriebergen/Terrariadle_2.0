package services

import (
	"terrariadle-backend/internal/utils/cache"
)

func GetDailySlashPlayersGuessed() (int, error) {
	gameData := cache.GetGameData()

	return gameData.GuessCounts.DailySlashCount, nil
}
