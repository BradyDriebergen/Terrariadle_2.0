package services

import (
	"fmt"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/utils/memstore"
)

type WeaponOutput struct {
	PreviousWeaponName string `json:"previousWeaponName"`
	Hint1              string `json:"hint1"`
	Hint2              string `json:"hint2"`
	Hint3              string `json:"hint3"`
}

func InitializeDailySlashGame() (WeaponOutput, error) {
	gameData, ok := memstore.GameData.Get()
	if !ok {
		return WeaponOutput{}, fmt.Errorf("failed to get data from memstore")
	}

	weapons := gameData.DailySlash
	weaponData := WeaponOutput{
		PreviousWeaponName: weapons.PreviousWeapon.Name,
		Hint1:              weapons.CurrentWeapon.ModeObtained,
		Hint2:              weapons.CurrentWeapon.WeaponType,
		Hint3:              weapons.CurrentWeapon.Info.ImagePath,
	}

	return weaponData, nil
}

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

func GetDailySlashPlayersGuessed() (int, error) {
	gameData, ok := memstore.GameData.Get()
	if !ok {
		return 0, fmt.Errorf("failed to get data from memstore")
	}

	return gameData.GuessCounts.DailySlashCount, nil
}
