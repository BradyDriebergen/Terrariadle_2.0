package services

import (
	"fmt"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

type WeaponOutput struct {
	PreviousWeaponData previousWeapon `json:"previousWeaponData"`
	Hint1              string         `json:"hint1"`
	Hint2              string         `json:"hint2"`
	Hint3              string         `json:"hint3"`
}

type previousWeapon struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Rarity string `json:"rarity"`
}

var rarities = map[string]int{
	"White":        0,
	"Blue":         1,
	"Green":        2,
	"Orange":       3,
	"Light Red":    4,
	"Pink":         5,
	"Light Purple": 6,
	"Lime":         7,
	"Yellow":       8,
	"Cyan":         9,
	"Red":          10,
}

var useTimes = map[string]int{
	"Insanely Fast":  7,
	"Very Fast":      6,
	"Fast":           5,
	"Average":        4,
	"Slow":           3,
	"Very Slow":      2,
	"Extremely Slow": 1,
	"Snail":          0,
}

func GetDailySlashSearchItems() ([]jsonreader.SearchWeaponResult, error) {
	result, ok := store.SearchWeaponsCache.Get()
	if !ok {
		return []jsonreader.SearchWeaponResult{}, fmt.Errorf("failed to get search data for weapons")
	}

	return result, nil
}

// Initializes the daily slash game for a user, returning the puzzle data and guessed weapons
func InitializeDailySlashGame(userId string) (WeaponOutput, []jsonreader.Weapon, error) {
	// Gets daily game data and weapons from memstore
	gameData, ok := store.GameData.Get()
	if !ok {
		return WeaponOutput{}, []jsonreader.Weapon{}, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return WeaponOutput{}, []jsonreader.Weapon{}, fmt.Errorf("failed to get weapons from memstore")
	}

	puzzleData := gameData.DailySlash
	weaponData := WeaponOutput{
		PreviousWeaponData: previousWeapon{
			Name:   puzzleData.PreviousWeapon.Name,
			Path:   puzzleData.PreviousWeapon.Info.ImagePath,
			Rarity: puzzleData.PreviousWeapon.Info.Rarity,
		},
		Hint1: puzzleData.CurrentWeapon.ModeObtained,
		Hint2: puzzleData.CurrentWeapon.WeaponType,
		Hint3: puzzleData.CurrentWeapon.Info.ImagePath,
	}

	// Gets the user from the collection
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return WeaponOutput{}, []jsonreader.Weapon{}, fmt.Errorf("failed to get user in user guesses API")
	}

	// Maps guessed weapons and returns puzzle data
	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			weaponSet := make(map[int]bool)
			for _, id := range user.Games[i].Guesses {
				weaponSet[id] = true
			}

			var filtered []jsonreader.Weapon
			for _, w := range weapons {
				if weaponSet[w.ID] {
					filtered = append(filtered, w)
				}
			}

			return weaponData, filtered, nil
		}
	}

	// Errors if can't find user guess data
	return WeaponOutput{}, []jsonreader.Weapon{}, fmt.Errorf("error getting player position")
}

// Checks a user's guess
func CheckDailySlashGuess(userId string, weaponId int) (bool, jsonreader.Weapon, error) {
	// Initial checks
	if weaponId > 450 {
		return false, jsonreader.Weapon{}, fmt.Errorf("invalid weapon ID")
	}
	if !isValidUUID(userId) {
		return false, jsonreader.Weapon{}, fmt.Errorf("invalid user ID")
	}

	// Gets user from database
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return false, jsonreader.Weapon{}, fmt.Errorf("failed to get user in check API")
	}

	// Gets memstore data
	gameData, ok := store.GameData.Get()
	if !ok {
		return false, jsonreader.Weapon{}, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return false, jsonreader.Weapon{}, fmt.Errorf("failed to get weapons from memstore")
	}

	// Checks guess
	won := false
	if weaponId == gameData.DailySlash.CurrentWeapon.ID {
		won = true
	}

	// Updates user based on guess
	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			// Sets Guesses
			user.Games[i].Guesses = append(user.Games[i].Guesses, weaponId)

			// Update weapon checks for this guess
			user.Games[i].Extra["WeaponChecks"] = checkGuess(weapons[weaponId-1], gameData.DailySlash.CurrentWeapon)

			// Sets winning
			if won {
				user.Games[i].HasWon = true
				gameData.GuessCounts.DailySlashCount += 1
				user.Games[i].Position = gameData.GuessCounts.DailySlashCount
				store.GameData.Set(gameData)

			}
			break
		}
	}
	err = models.UpdateUserData(user)
	if err != nil {
		return false, jsonreader.Weapon{}, err
	}

	guessedWeapon := weapons[weaponId-1]
	return won, guessedWeapon, nil
}

// Gets the winning position and total players for daily slash
func GetDailySlashWinningData(userId string) (int, int, error) {
	gameData, ok := store.GameData.Get()
	if !ok {
		return 0, 0, fmt.Errorf("failed to get data from memstore")
	}

	user, err := models.GetOrCreateUser(userId)
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

func checkGuess(guess, answer jsonreader.Weapon) models.WeaponChecks {
	damage := 1
	if answer.Info.Damage > guess.Info.Damage {
		damage = 2
	} else if answer.Info.Damage < guess.Info.Damage {
		damage = 0
	}

	useTime := 1
	if useTimes[answer.Info.UseTime] > useTimes[guess.Info.UseTime] {
		useTime = 2
	} else if useTimes[answer.Info.UseTime] < useTimes[guess.Info.UseTime] {
		useTime = 0
	}

	rarity := 1
	if rarities[answer.Info.Rarity] < rarities[guess.Info.Rarity] {
		rarity = 0
	} else if rarities[answer.Info.Rarity] > rarities[guess.Info.Rarity] {
		rarity = 2
	}

	obtained := guessObtained(guess.Info.Obtained, answer.Info.Obtained)

	return models.WeaponChecks{
		WeaponType: answer.WeaponType == guess.WeaponType,
		Damage:     damage,
		UseTime:    useTime,
		Rarity:     rarity,
		Operation:  answer.Info.Operation == guess.Info.Operation,
		Material:   answer.Info.Material == guess.Info.Material,
		Obtained:   obtained,
	}
}

func guessObtained(g, w []string) int {
	// Convert to maps for quick lookup
	gMap := make(map[string]bool)
	for _, val := range g {
		gMap[val] = true
	}
	wMap := make(map[string]bool)
	for _, val := range w {
		wMap[val] = true
	}

	// Check if they are identical sets
	if len(gMap) == len(wMap) {
		same := true
		for key := range gMap {
			if !wMap[key] {
				same = false
				break
			}
		}
		if same {
			return 2
		}
	}

	// Check for partial overlap
	for key := range gMap {
		if wMap[key] {
			return 1
		}
	}

	return 0
}
