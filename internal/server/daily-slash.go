package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

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

type previousWeapon struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Rarity string `json:"rarity"`
}

type dailySlashInit struct {
	PreviousWeapon previousWeapon
	GuessedWeapons []jsonreader.Weapon
	GuessChecks    []models.WeaponChecks
	HasWon         bool
}

// Gets the data required for loading the Daily Slash game
func getDailySlashInitData(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	initData, err := initDailySlash(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"previousWeapon": initData.PreviousWeapon,
		"guesses":        initData.GuessedWeapons,
		"checks":         initData.GuessChecks,
		"won":            initData.HasWon,
	})
}

func initDailySlash(userId string) (dailySlashInit, error) {
	returnData := dailySlashInit{
		PreviousWeapon: previousWeapon{},
		GuessedWeapons: []jsonreader.Weapon{},
		GuessChecks:    []models.WeaponChecks{},
		HasWon:         false,
	}

	// Initial pulls from stores and database
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get game data")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get weapons")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnData, fmt.Errorf("failed to get user information")
	}

	puzzleData := gameData.DailySlash
	returnData.PreviousWeapon = previousWeapon{
		Name:   puzzleData.PreviousWeapon.Name,
		Path:   puzzleData.PreviousWeapon.Info.ImagePath,
		Rarity: puzzleData.PreviousWeapon.Info.Rarity,
	}

	// Maps guessed weapons to their ids, this might need to be cached in the future
	// but only if this endpoint gets ~10,000 RPS
	weaponByID := make(map[int]jsonreader.Weapon)
	for _, w := range weapons {
		weaponByID[w.ID] = w
	}

	// Maps guessed ids to weapon objects
	for _, id := range user.DailySlash.Game.Guesses {
		if w, ok := weaponByID[id]; ok {
			returnData.GuessedWeapons = append(returnData.GuessedWeapons, w)
		}
	}

	returnData.GuessChecks = user.DailySlash.Checks
	returnData.HasWon = user.DailySlash.Game.HasWon

	return returnData, nil
}

// Gets the simplified weapon list from the database that the users can search through
func getDailySlashSearchItems(w http.ResponseWriter, r *http.Request) {
	searchItems, err := searchableWeapons()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, searchItems)
}

func searchableWeapons() ([]jsonreader.SearchWeaponResult, error) {
	result, ok := store.SearchWeaponsCache.Get()
	if !ok {
		return []jsonreader.SearchWeaponResult{}, fmt.Errorf("failed to get weapons")
	}

	return result, nil
}

// Gets a hint about the daily weapon
func getDailySlashHint(w http.ResponseWriter, r *http.Request) {
	// Decodes hint number and assigns it to an int
	hintNum, err := strconv.Atoi(r.PathValue("hintNum"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	searchItems, err := dailySlashHint(hintNum)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, searchItems)
}

func dailySlashHint(hintNum int) (string, error) {
	gameData, ok := store.GameData.Get()
	if !ok {
		return "", fmt.Errorf("failed to get game data")
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
		return "", fmt.Errorf("requested hint does not exist")
	}
}

type dailySlashGuessRequestBody struct {
	UserID string `json:"userId"`
	Guess  int    `json:"guess"`
}

type dailySlashCheck struct {
	HasWon        bool
	GuessedWeapon jsonreader.Weapon
	GuessChecks   models.WeaponChecks
}

// Checks Daily Slash guess and updates the database
func postDailySlashGuess(w http.ResponseWriter, r *http.Request) {
	var req dailySlashGuessRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidUUID(req.UserID) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user id"})
		return
	}

	checkData, err := checkDailySlashGuess(req.UserID, req.Guess)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"guess": checkData.GuessedWeapon,
		"check": checkData.GuessChecks,
		"won":   checkData.HasWon,
	})
}

func checkDailySlashGuess(userId string, weaponId int) (dailySlashCheck, error) {
	returnData := dailySlashCheck{
		GuessedWeapon: jsonreader.Weapon{},
		GuessChecks:   models.WeaponChecks{},
		HasWon:        false,
	}

	// Initial pulls from stores and database
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get game data")
	}
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get weapons")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnData, fmt.Errorf("failed to get user information")
	}

	// Finds guessed weapon, return if not found
	found := false
	for _, w := range weapons {
		if w.ID == weaponId {
			returnData.GuessedWeapon = w
			found = true
			break
		}
	}
	if !found {
		return returnData, fmt.Errorf("weapon with id %d not found", weaponId)
	}

	// Adds guessed weapon to guesses
	user.DailySlash.Game.Guesses = append([]int{weaponId}, user.DailySlash.Game.Guesses...)

	// Checks guessed weapon against daily weapon
	returnData.GuessChecks = checkGuess(returnData.GuessedWeapon, gameData.DailySlash.CurrentWeapon)
	user.DailySlash.Checks = append([]models.WeaponChecks{returnData.GuessChecks}, user.DailySlash.Checks...)

	// if correct, update user values
	if weaponId == gameData.DailySlash.CurrentWeapon.ID {
		user.DailySlash.Game.HasWon = true
		gameData.GuessCounts.DailySlashCount += 1
		user.DailySlash.Game.Position = gameData.GuessCounts.DailySlashCount
		store.GameData.Set(gameData)
	}

	err = models.UpdateUserData(user)
	if err != nil {
		return returnData, fmt.Errorf("failed to update user")
	}

	returnData.HasWon = user.DailySlash.Game.HasWon

	return returnData, nil
}

func checkGuess(guess, answer jsonreader.Weapon) models.WeaponChecks {
	// Compares damage (1 is equal, 2 is greater, 0 is less)
	damage := 1
	if answer.Info.Damage > guess.Info.Damage {
		damage = 2
	} else if answer.Info.Damage < guess.Info.Damage {
		damage = 0
	}

	// Compares use time (1 is equal, 2 is greater, 0 is less)
	useTime := 1
	if useTimes[answer.Info.UseTime] > useTimes[guess.Info.UseTime] {
		useTime = 2
	} else if useTimes[answer.Info.UseTime] < useTimes[guess.Info.UseTime] {
		useTime = 0
	}

	// Compares rarities (1 is equal, 2 is greater, 0 is less)
	rarity := 1
	if rarities[answer.Info.Rarity] < rarities[guess.Info.Rarity] {
		rarity = 0
	} else if rarities[answer.Info.Rarity] > rarities[guess.Info.Rarity] {
		rarity = 2
	}

	// Compares obtained values (2 is equal, 1 is partial, 0 is non-matching)
	obtained := sliceCompare(guess.Info.Obtained, answer.Info.Obtained)

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

func sliceCompare(g, w []string) int {
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

// Gets the position and players guessed numbers
func getDailySlashWinningData(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	pos, count, err := dailySlashWinningData(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"pos":   pos,
		"count": count,
	})
}

func dailySlashWinningData(userId string) (int, int, error) {
	// Initial pulls from stores and database
	gameData, ok := store.GameData.Get()
	if !ok {
		return 0, 0, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get user information")
	}

	// Checks if user guessed
	if len(user.DailySlash.Game.Guesses) > 0 {
		return user.DailySlash.Game.Position, gameData.GuessCounts.DailySlashCount, nil
	}

	return 0, 0, fmt.Errorf("user has not guessed")
}
