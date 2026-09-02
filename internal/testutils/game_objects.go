package testutils

import (
	"strconv"
	"terrariadle/internal/domain"
)

func GenerateWeapon(id int) domain.Weapon {
	return domain.Weapon{
		ID:           id,
		Name:         "weapon" + strconv.Itoa(id),
		WeaponType:   "type",
		ModeObtained: "hardmode",
		Info: domain.WeaponInfo{
			ImagePath:  ".png",
			DamageType: "type",
			Damage:     15,
			UseTime:    "time",
			Rarity:     "rarity",
			Operation:  "operation",
			Material:   "no",
			Obtained:   []string{"obtained"},
		},
	}
}

func GenerateWeapons() []domain.Weapon {
	weapons := make([]domain.Weapon, 3)
	for i := range 3 {
		weapons[i] = GenerateWeapon(i)
	}
	return weapons
}

func GenerateCategory(id int) domain.Category {
	return domain.Category{
		ID:       id,
		Category: "cat" + strconv.Itoa(id),
		Options:  []string{"opt1", "opt2", "opt3", "opt4"},
	}
}

func GenerateCategories() []domain.Category {
	categories := make([]domain.Category, 3)
	for i := range 3 {
		categories[i] = GenerateCategory(i)
	}
	return categories
}

func GenerateNpc(id int) domain.Npc {
	return domain.Npc{
		ID:      id,
		NPC:     "npc" + strconv.Itoa(id),
		NpcPath: ".png",
		Quotes:  []string{"quote1", "quote2"},
		Names:   []string{"name1", "name2"},
	}
}

func GenerateNpcs() []domain.Npc {
	npcs := make([]domain.Npc, 3)
	for i := range 3 {
		npcs[i] = GenerateNpc(i)
	}
	return npcs
}

func GenerateEnemy(id int) domain.Enemy {
	return domain.Enemy{
		ID:        id,
		Name:      "enemy" + strconv.Itoa(id),
		ImagePath: ".png",
	}
}

func GenerateEnemies() []domain.Enemy {
	enemies := make([]domain.Enemy, 3)
	for i := range 3 {
		enemies[i] = GenerateEnemy(i)
	}
	return enemies
}

func GenerateTriviaQuestion(id int) domain.TriviaQuestion {
	return domain.TriviaQuestion{
		ID:         id,
		Answer:     "answer" + strconv.Itoa(id),
		Clue:       "clue",
		Chunks:     []string{"chunk1", "chunk2"},
		ChunkCount: 2,
	}
}

func GenerateTriviaQuestions() []domain.TriviaQuestion {
	questions := make([]domain.TriviaQuestion, 3)
	for i := range 3 {
		questions[i] = GenerateTriviaQuestion(i)
	}
	return questions
}
