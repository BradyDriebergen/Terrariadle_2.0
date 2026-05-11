package jobs

import (
	"fmt"
	"math/rand"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

// ---------------------------------------------
// Helper methods
// ---------------------------------------------

// Shuffles a slice (no reference needed)
func shuffle[T any](list []T, rnd *rand.Rand) {
	rnd.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

// ---------------------------------------------
// Get random guess data methods
// ---------------------------------------------

// Generates a pair of random weapons
func randomDailyWeapons(rnd *rand.Rand, prevWeapon jsonreader.Weapon) (models.WeaponData, error) {
	weapons, ok := store.WeaponsCache.Get()
	if !ok {
		return models.WeaponData{}, fmt.Errorf("failed to get weapons from memstore")
	}
	if len(weapons) < 2 {
		return models.WeaponData{}, fmt.Errorf("not enough weapons to choose from: got %d", len(weapons))
	}

	shuffle(weapons, rnd)

	return models.WeaponData{CurrentWeapon: weapons[0], PreviousWeapon: prevWeapon}, nil
}

// Generates a set of categories with unique options
func randomCategories(rnd *rand.Rand) ([]jsonreader.Category, error) {
	categories, ok := store.CategoriesCache.Get()
	if !ok {
		return []jsonreader.Category{}, fmt.Errorf("failed to get categories from memstore")
	}
	if len(categories) < 4 {
		return []jsonreader.Category{}, fmt.Errorf("need at least 4 categories, got %d", len(categories))
	}

	shuffle(categories, rnd)

	selectedCategories := categories[:4]
	allOptions := []string{}

	// Loops through the four categories and gets options
	for i := range selectedCategories {
		opts := selectedCategories[i].Options
		if len(opts) < 4 {
			return []jsonreader.Category{}, fmt.Errorf("categories don't have 4 options: %v", selectedCategories[i])
		}

		shuffle(opts, rnd)

		selectedCategories[i].Options = opts[:4]
		allOptions = append(allOptions, opts[:4]...)
	}

	return selectedCategories, nil
}

// Gets random NPC data with unique names
func randomNpcData(rnd *rand.Rand) (models.NPCdata, error) {
	npcs, ok := store.NpcsCache.Get()
	if !ok {
		return models.NPCdata{}, fmt.Errorf("failed to get npcs from memstore")
	}
	if len(npcs) < 4 {
		return models.NPCdata{}, fmt.Errorf("need at least 4 NPCs, got %d", len(npcs))
	}

	shuffle(npcs, rnd)
	// Helper method that picks a random string from a string slice
	pickStr := func(ss []string) string { return ss[rnd.Intn(len(ss))] }

	primaryNPC := npcs[0]
	if len(primaryNPC.Quotes) == 0 {
		return models.NPCdata{}, fmt.Errorf("npc %q has no quotes", primaryNPC.NPC)
	}
	if len(primaryNPC.Names) == 0 {
		return models.NPCdata{}, fmt.Errorf("npc %q has no names", primaryNPC.NPC)
	}

	// NPC data that will be sent, ensures that names is a slice of strings
	out := models.NPCdata{
		ID:      primaryNPC.ID,
		NPC:     primaryNPC.NPC,
		NPCPath: primaryNPC.NPCPath,
		Quote:   pickStr(primaryNPC.Quotes),
		Names:   make([]string, 0, 4),
	}
	out.Names = append(out.Names, pickStr(primaryNPC.Names))

	// Loops through three other random NPCs and picks a name
	secondaryNPCs := npcs[1:4]
	for i := range secondaryNPCs {
		if len(secondaryNPCs[i].Names) == 0 {
			return models.NPCdata{}, fmt.Errorf("npc %q has no names", secondaryNPCs[i].NPC)
		}
		out.Names = append(out.Names, pickStr(secondaryNPCs[i].Names))
	}

	return out, nil
}

// Gets a random enemy
func randomEnemy(rnd *rand.Rand) (jsonreader.Enemy, error) {
	enemies, ok := store.EnemiesCache.Get()
	if !ok {
		return jsonreader.Enemy{}, fmt.Errorf("failed to get weapons from memstore")
	}

	shuffle(enemies, rnd)

	return enemies[0], nil
}

// Resets the guess counts to 0
func resetGuessCounts() models.PlayerGuessCounts {
	return models.PlayerGuessCounts{
		DailySlashCount:  0,
		ConnectionsCount: 0,
		GuessTheNpcCount: 0,
		HangmanCount:     0,
	}
}
