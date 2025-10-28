package services

import (
	"fmt"
	"strings"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/utils/memstore"
)

type WeaponOutput struct {
	PreviousWeaponName string `json:"previousWeaponName"`
	Hint1              string `json:"hint1"`
	Hint2              string `json:"hint2"`
	Hint3              string `json:"hint3"`
}

type searchResult struct {
	WeaponId int    `json:"weaponId"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

func InitializeDailySlashGame(userId string) (WeaponOutput, []domain.Weapon, error) {
	// Gets daily game data and weapons from memstore
	gameData, ok := memstore.GameData.Get()
	if !ok {
		return WeaponOutput{}, []domain.Weapon{}, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := memstore.WeaponData.Get()
	if !ok {
		return WeaponOutput{}, []domain.Weapon{}, fmt.Errorf("failed to get weapons from memstore")
	}

	puzzleData := gameData.DailySlash
	weaponData := WeaponOutput{
		PreviousWeaponName: puzzleData.PreviousWeapon.Name,
		Hint1:              puzzleData.CurrentWeapon.ModeObtained,
		Hint2:              puzzleData.CurrentWeapon.WeaponType,
		Hint3:              puzzleData.CurrentWeapon.Info.ImagePath,
	}

	// Gets the user from the collection
	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return WeaponOutput{}, []domain.Weapon{}, fmt.Errorf("failed to get user in user guesses API")
	}

	// Maps guessed weapons and returns puzzle data
	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			weaponSet := make(map[int]bool)
			for _, id := range user.Games[i].Guesses {
				weaponSet[id] = true
			}

			var filtered []domain.Weapon
			for _, w := range weapons {
				if weaponSet[w.ID] {
					filtered = append(filtered, w)
				}
			}

			return weaponData, filtered, nil
		}
	}

	// Errors if can't find user guess data
	return WeaponOutput{}, []domain.Weapon{}, fmt.Errorf("error getting player position")
}

func DailySlashSearch(query string) ([]searchResult, error) {
	if query == "" {
		return []searchResult{}, fmt.Errorf("query cannot be empty")
	}

	weapons, ok := memstore.WeaponData.Get()
	if !ok {
		return []searchResult{}, fmt.Errorf("error getting weapon data from memstore")
	}

	// Gets the first 20 results of search
	query = strings.ToLower(query)
	results := []searchResult{}
	for _, w := range weapons {
		if strings.Contains(strings.ToLower(w.Name), query) {
			results = append(results, searchResult{
				WeaponId: w.ID,
				Name:     w.Name,
				Path:     w.Info.ImagePath,
			})
		}
		if len(results) >= 20 { // limit results
			break
		}
	}

	return results, nil
}

func CheckDailySlashGuess(userId string, weaponId int) (bool, domain.Weapon, error) {
	// Initial checks
	if weaponId > 450 {
		return false, domain.Weapon{}, fmt.Errorf("invalid weapon ID")
	}
	if !isValidUUID(userId) {
		return false, domain.Weapon{}, fmt.Errorf("invalid user ID")
	}

	// Gets user from database
	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return false, domain.Weapon{}, fmt.Errorf("failed to get user in check API")
	}

	// Gets memstore data
	gameData, ok := memstore.GameData.Get()
	if !ok {
		return false, domain.Weapon{}, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := memstore.WeaponData.Get()
	if !ok {
		return false, domain.Weapon{}, fmt.Errorf("failed to get weapons from memstore")
	}

	// Checks guess
	won := false
	if weaponId == gameData.DailySlash.CurrentWeapon.ID {
		won = true
	}

	// Updates user based on guess
	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			user.Games[i].Guesses = append(user.Games[i].Guesses, weaponId)
			if won {
				user.Games[i].HasWon = true
				gameData.GuessCounts.DailySlashCount += 1
				user.Games[i].Position = gameData.GuessCounts.DailySlashCount
				memstore.GameData.Set(gameData)
			}
			break
		}
	}
	err = updateUser(col, user)
	if err != nil {
		return false, domain.Weapon{}, err
	}

	guessedWeapon := weapons[weaponId-1]
	return won, guessedWeapon, nil
}

func GetDailySlashWinningData(userId string) (int, int, error) {
	gameData, ok := memstore.GameData.Get()
	if !ok {
		return 0, 0, fmt.Errorf("failed to get data from memstore")
	}

	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get user in player position API")
	}

	// Gets winning position and player count
	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			if len(user.Games[i].Guesses) > 0 {
				return user.Games[i].Position, gameData.GuessCounts.DailySlashCount, nil
			}
			return 0, 0, fmt.Errorf("player doesn't exist")
		}
	}

	return 0, 0, fmt.Errorf("error getting player position")
}
