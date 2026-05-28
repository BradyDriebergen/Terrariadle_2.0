package services

import "terrariadle-backend/internal/domain"

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

func checkGuess(guess, answer domain.Weapon) domain.WeaponChecks {
	// Compares damage (1 is equal, 2 is greater, 0 is less)
	damage := domain.Match
	if answer.Info.Damage > guess.Info.Damage {
		damage = domain.Higher
	} else if answer.Info.Damage < guess.Info.Damage {
		damage = domain.Lower
	}

	// Compares use time (1 is equal, 2 is greater, 0 is less)
	useTime := domain.Match
	if useTimes[answer.Info.UseTime] > useTimes[guess.Info.UseTime] {
		useTime = domain.Higher
	} else if useTimes[answer.Info.UseTime] < useTimes[guess.Info.UseTime] {
		useTime = domain.Lower
	}

	// Compares rarities (1 is equal, 2 is greater, 0 is less)
	rarity := domain.Match
	if rarities[answer.Info.Rarity] > rarities[guess.Info.Rarity] {
		rarity = domain.Higher
	} else if rarities[answer.Info.Rarity] < rarities[guess.Info.Rarity] {
		rarity = domain.Lower
	}

	// Compares obtained values (2 is equal, 1 is partial, 0 is non-matching)
	obtained := sliceCompare(guess.Info.Obtained, answer.Info.Obtained)

	return domain.WeaponChecks{
		DamageType: answer.Info.DamageType == guess.Info.DamageType,
		Damage:     damage,
		UseTime:    useTime,
		Rarity:     rarity,
		Operation:  answer.Info.Operation == guess.Info.Operation,
		Material:   answer.Info.Material == guess.Info.Material,
		Obtained:   obtained,
	}
}

func sliceCompare(g, w []string) domain.CompareResult {
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
			return domain.Match
		}
	}

	// Check for partial overlap
	for key := range gMap {
		if wMap[key] {
			return domain.Higher
		}
	}

	return domain.Lower
}

func toPreview(w domain.Weapon) domain.WeaponPreview {
	return domain.WeaponPreview{
		Name:   w.Name,
		Path:   w.Info.ImagePath,
		Rarity: w.Info.Rarity,
	}
}

func toWeaponData(w domain.Weapon) WeaponData {
	return WeaponData{
		ID:         w.ID,
		Name:       w.Name,
		ImagePath:  w.Info.ImagePath,
		DamageType: w.Info.DamageType,
		Damage:     w.Info.Damage,
		UseTime:    w.Info.UseTime,
		Rarity:     w.Info.Rarity,
		Operation:  w.Info.Operation,
		Material:   w.Info.Material,
		Obtained:   w.Info.Obtained,
	}
}
