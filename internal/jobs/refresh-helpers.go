package jobs

import (
	"math/rand"
	"slices"
	"terrariadle-backend/internal/domain"
)

func (j *PuzzleRefreshJob) refreshWeapons() domain.WeaponAnswer {
	weapons := j.catalogStore.GetWeapons()
	oldWeapon := j.answerStore.GetAnswers().DailySlash.CurrentWeapon

	newWeapon := randomItem(weapons, j.rng)
	for newWeapon.ID == oldWeapon.ID {
		newWeapon = randomItem(weapons, j.rng)
	}

	return domain.WeaponAnswer{
		CurrentWeapon: newWeapon,
		PrevWeapon:    oldWeapon,
	}
}

func (j *PuzzleRefreshJob) refreshCategories() domain.ConnectionAnswer {
	categories := j.catalogStore.GetCategories()
	shuffle(categories, j.rng)

	oldCategories := j.answerStore.GetAnswers().Connections.CategoryIDs
	newCategoryOptions := make([]domain.ConnectionOption, 0, 16)
	newCategoryIDs := []int{}
	index := 0

	for len(newCategoryIDs) < 4 && index < len(categories) {
		cat := categories[index]

		if !slices.Contains(oldCategories, cat.ID) {
			newCategoryIDs = append(newCategoryIDs, cat.ID)

			shuffle(cat.Options, j.rng)

			for i := range 4 {
				newCategoryOptions = append(newCategoryOptions, domain.ConnectionOption{
					CategoryID: cat.ID,
					Option:     cat.Options[i],
				})
			}
		}

		index++
	}

	return domain.ConnectionAnswer{
		CategoryIDs: newCategoryIDs,
		Options:     newCategoryOptions,
	}
}

func (j *PuzzleRefreshJob) refreshNpc() domain.NpcAnswer {
	npcs := j.catalogStore.GetNpcs()
	shuffle(npcs, j.rng)

	oldNpc := j.answerStore.GetAnswers().GuessTheNpc
	index := 0

	randomNpc := npcs[index]
	for randomNpc.ID == oldNpc.NpcID {
		index++
		randomNpc = npcs[index]
	}

	quote := randomItem(randomNpc.Quotes, j.rng)
	name := randomItem(randomNpc.Names, j.rng)
	names := []string{name}

	for len(names) < 4 && index < len(npcs) {
		index++

		names = append(names, randomItem(npcs[index].Names, j.rng))
	}

	shuffle(names, j.rng)

	return domain.NpcAnswer{
		NpcID:       randomNpc.ID,
		Npc:         randomNpc.NPC,
		Quote:       quote,
		Name:        name,
		NameOptions: names,
	}
}

func (j *PuzzleRefreshJob) refreshEnemy() domain.HangmanAnswer {
	enemies := j.catalogStore.GetEnemies()
	oldEnemy := j.answerStore.GetAnswers().Hangman.Enemy

	newEnemy := randomItem(enemies, j.rng)
	for newEnemy.ID == oldEnemy.ID {
		newEnemy = randomItem(enemies, j.rng)
	}

	return domain.HangmanAnswer{
		Enemy: newEnemy,
	}
}

func shuffle[T any](list []T, rnd *rand.Rand) {
	rnd.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

func randomItem[T any](list []T, rng *rand.Rand) T {
	return list[rng.Intn(len(list))]
}
