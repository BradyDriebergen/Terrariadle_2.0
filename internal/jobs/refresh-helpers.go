package jobs

import (
	"math/rand"
	"slices"
	"strings"
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

	newEnemy.Name = strings.ToUpper(newEnemy.Name)

	return domain.HangmanAnswer{
		Enemy: newEnemy,
	}
}

func (j *PuzzleRefreshJob) refreshTriviaQuestions() domain.TerraTriviaAnswer {
	allQuestions := j.catalogStore.GetTriviaQuestions()
	shuffle(allQuestions, j.rng)

	// Split questions up based on chunk size
	twoChunks := make([]*domain.TriviaQuestion, 0, 500)
	threeChunks := make([]*domain.TriviaQuestion, 0, 1300)
	fourChunks := make([]*domain.TriviaQuestion, 0, 200)
	for _, o := range allQuestions {
		switch o.ChunkCount {
		case 2:
			twoChunks = append(twoChunks, &o)
		case 3:
			threeChunks = append(threeChunks, &o)
		case 4:
			fourChunks = append(fourChunks, &o)
		}
	}

	oldQuestions := j.answerStore.GetAnswers().TerraTrivia.Questions
	oldIDs := make(map[int]struct{}, len(oldQuestions))
	for _, q := range oldQuestions {
		oldIDs[q.ID] = struct{}{}
	}

	// helper function for adding questions based on chunk size
	add := func(
		dst []domain.TriviaQuestion,
		pool []*domain.TriviaQuestion,
		count int,
	) []domain.TriviaQuestion {
		for _, q := range pool {
			if count == 0 {
				break
			}
			if _, seen := oldIDs[q.ID]; seen {
				continue
			}
			dst = append(dst, *q)
			count--
		}
		return dst
	}

	newQuestions := make([]domain.TriviaQuestion, 0, 7)
	if j.rng.Float64() < 0.70 {
		newQuestions = add(newQuestions, threeChunks, 6)
		newQuestions = add(newQuestions, twoChunks, 1)
	} else {
		newQuestions = add(newQuestions, fourChunks, 1)
		newQuestions = add(newQuestions, threeChunks, 4)
		newQuestions = add(newQuestions, twoChunks, 2)
	}

	if len(newQuestions) != 7 {
		// log error: reusing previous day's questions if fail
		return domain.TerraTriviaAnswer{
			Questions: oldQuestions,
		}
	}

	shuffle(newQuestions, j.rng)

	return domain.TerraTriviaAnswer{
		Questions: newQuestions,
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
