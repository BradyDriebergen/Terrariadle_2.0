package services

import (
	"fmt"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

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

func GetDailySlashHint(hintNum int) (string, error) {
	gameData, ok := store.GameData.Get()
	if !ok {
		return "", fmt.Errorf("failed to get data from memstore")
	}

	puzzleData := gameData.DailySlash
	switch hintNum {
	case 1:
		return puzzleData.CurrentWeapon.ModeObtained, nil
	case 2:
		return puzzleData.CurrentWeapon.WeaponType, nil
	case 3:
		return puzzleData.CurrentWeapon.Info.ImagePath, nil
	default:
		return "", fmt.Errorf("requested hint doesn't exist")
	}
}

func GetDailySlashSearchItems() ([]jsonreader.SearchWeaponResult, error) {
	result, ok := store.SearchWeaponsCache.Get()
	if !ok {
		return []jsonreader.SearchWeaponResult{}, fmt.Errorf("failed to get search data for weapons")
	}

	return result, nil
}

// Initializes the daily slash game for a user, returning the puzzle data and guessed weapons
func InitializeDailySlashGame(userId string) (previousWeapon, []jsonreader.Weapon, []models.WeaponChecks, bool, error) {
	// Gets daily game data and weapons from memstore
	gameData, ok := store.GameData.Get()
	if !ok {
		return previousWeapon{}, []jsonreader.Weapon{}, []models.WeaponChecks{}, false, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return previousWeapon{}, []jsonreader.Weapon{}, []models.WeaponChecks{}, false, fmt.Errorf("failed to get weapons from memstore")
	}

	puzzleData := gameData.DailySlash
	previousWeaponData := previousWeapon{
		Name:   puzzleData.PreviousWeapon.Name,
		Path:   puzzleData.PreviousWeapon.Info.ImagePath,
		Rarity: puzzleData.PreviousWeapon.Info.Rarity,
	}

	// Gets the user from the collection
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return previousWeapon{}, []jsonreader.Weapon{}, []models.WeaponChecks{}, false, fmt.Errorf("failed to get user in user guesses API")
	}

	// Maps guessed weapons and returns puzzle data
	weaponByID := make(map[int]jsonreader.Weapon)
	for _, w := range weapons {
		weaponByID[w.ID] = w
	}

	var filtered []jsonreader.Weapon
	for _, id := range user.DailySlash.Game.Guesses {
		if w, ok := weaponByID[id]; ok {
			filtered = append(filtered, w)
		}
	}

	return previousWeaponData, filtered, user.DailySlash.Checks, user.DailySlash.Game.HasWon, nil
}

// Checks a user's guess
func CheckDailySlashGuess(userId string, weaponId int) (bool, jsonreader.Weapon, models.WeaponChecks, error) {
	// Initial checks
	if !isValidUUID(userId) {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, fmt.Errorf("invalid user ID")
	}

	// Gets user from database
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, fmt.Errorf("failed to get user in check API")
	}

	// Gets memstore data
	gameData, ok := store.GameData.Get()
	if !ok {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, fmt.Errorf("failed to get data from memstore")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, fmt.Errorf("failed to get weapons from memstore")
	}

	guessedWeapon := jsonreader.Weapon{}
	found := false
	for _, w := range weapons {
		if w.ID == weaponId {
			guessedWeapon = w
			found = true
			break
		}
	}
	if !found {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, fmt.Errorf("weapon with id %d not found", weaponId)
	}

	// Checks guess
	won := false
	if weaponId == gameData.DailySlash.CurrentWeapon.ID {
		won = true
	}

	// Updates user based on guess
	user.DailySlash.Game.Guesses = append([]int{weaponId}, user.DailySlash.Game.Guesses...)

	// Update weapon checks for this guess
	checkResults := checkGuess(guessedWeapon, gameData.DailySlash.CurrentWeapon)
	user.DailySlash.Checks = append([]models.WeaponChecks{checkResults}, user.DailySlash.Checks...)

	// Sets winning
	if won {
		user.DailySlash.Game.HasWon = true
		gameData.GuessCounts.DailySlashCount += 1
		user.DailySlash.Game.Position = gameData.GuessCounts.DailySlashCount
		store.GameData.Set(gameData)

	}

	err = models.UpdateUserData(user)
	if err != nil {
		return false, jsonreader.Weapon{}, models.WeaponChecks{}, err
	}

	return won, guessedWeapon, checkResults, nil
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
	if len(user.DailySlash.Game.Guesses) > 0 {
		return user.DailySlash.Game.Position, gameData.GuessCounts.DailySlashCount, nil
	}

	return 0, 0, fmt.Errorf("player doesn't exist")
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
		DamageType: answer.Info.DamageType == guess.Info.DamageType,
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
