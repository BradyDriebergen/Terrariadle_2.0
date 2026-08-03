package store

import (
	"fmt"
	"strings"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

func toAnswerDomain(ad repo.AnswerData, cs CatalogStore) (domain.DailyAnswers, error) {
	currentWeapon, ok := cs.GetWeapon(ad.DailySlash.CurrentWeaponID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find current weapon with id: %v", ad.DailySlash.CurrentWeaponID)
	}

	prevWeapon, ok := cs.GetWeapon(ad.DailySlash.PrevWeaponID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find previous weapon with id: %v", ad.DailySlash.PrevWeaponID)
	}

	npc, ok := cs.GetNpc(ad.GuessTheNpc.NpcID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find npc with id: %v", ad.GuessTheNpc.NpcID)
	}

	enemy, ok := cs.GetEnemy(ad.Hangman.EnemyID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find enemy with id: %v", ad.Hangman.EnemyID)
	}
	enemy.Name = strings.ToUpper(enemy.Name)

	triviaQuestions := make([]domain.TriviaQuestion, 7)
	for i, q := range ad.TerraTrivia.QuestionIDs {
		triviaQuestions[i], ok = cs.GetTriviaQuestion(q)
		if !ok {
			return domain.DailyAnswers{}, fmt.Errorf("failed to find enemy with id: %v", ad.Hangman.EnemyID)
		}
	}

	return domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: currentWeapon,
			PrevWeapon:    prevWeapon,
		},
		Connections: domain.ConnectionAnswer{
			CategoryIDs: ad.Connections.CategoryIDs,
			Options:     optionsToAnswers(ad.Connections.Options),
		},
		GuessTheNpc: domain.NpcAnswer{
			NpcID:       npc.ID,
			Npc:         npc.NPC,
			Quote:       ad.GuessTheNpc.Quote,
			Name:        ad.GuessTheNpc.Name,
			NameOptions: ad.GuessTheNpc.NameOptions,
		},
		Hangman: domain.HangmanAnswer{
			Enemy: enemy,
		},
		TerraTrivia: domain.TerraTriviaAnswer{
			Questions: triviaQuestions,
		},
		ResetTime:     ad.ResetTime,
		NextResetTime: ad.NextResetTime,
	}, nil
}

func fromAnswerDomain(da domain.DailyAnswers) repo.AnswerData {
	triviaQuestionIDs := make([]int, 7)
	for i, q := range da.TerraTrivia.Questions {
		triviaQuestionIDs[i] = q.ID
	}

	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: da.DailySlash.CurrentWeapon.ID,
			PrevWeaponID:    da.DailySlash.PrevWeapon.ID,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: da.Connections.CategoryIDs,
			Options:     answersToOptions(da.Connections.Options),
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       da.GuessTheNpc.NpcID,
			Quote:       da.GuessTheNpc.Quote,
			Name:        da.GuessTheNpc.Name,
			NameOptions: da.GuessTheNpc.NameOptions,
		},
		Hangman: repo.HangmanData{
			EnemyID: da.Hangman.Enemy.ID,
		},
		TerraTrivia: repo.TerraTriviaData{
			QuestionIDs: triviaQuestionIDs,
		},
		ResetTime:     da.ResetTime,
		NextResetTime: da.NextResetTime,
	}
}

func answersToOptions(answers []domain.ConnectionOption) []repo.ConnectionOption {
	options := make([]repo.ConnectionOption, len(answers))
	for i, a := range answers {
		options[i] = repo.ConnectionOption{
			Option:     a.Option,
			CategoryID: a.CategoryID,
		}
	}
	return options
}

func optionsToAnswers(options []repo.ConnectionOption) []domain.ConnectionOption {
	answers := make([]domain.ConnectionOption, len(options))
	for i, o := range options {
		answers[i] = domain.ConnectionOption{
			Option:     o.Option,
			CategoryID: o.CategoryID,
		}
	}
	return answers
}

func toGuessCountDomain(guessCounts repo.PlayerGuessCounts) domain.PlayerGuessCounts {
	return domain.PlayerGuessCounts{
		DailySlashCount:  guessCounts.DailySlashCount,
		ConnectionsCount: guessCounts.ConnectionsCount,
		GuessTheNpcCount: guessCounts.GuessTheNpcCount,
		HangmanCount:     guessCounts.HangmanCount,
		TerraTriviaCount: guessCounts.TerraTriviaCount,
	}
}

func fromGuessCountDomain(guessCounts domain.PlayerGuessCounts) repo.PlayerGuessCounts {
	return repo.PlayerGuessCounts{
		DailySlashCount:  guessCounts.DailySlashCount,
		ConnectionsCount: guessCounts.ConnectionsCount,
		GuessTheNpcCount: guessCounts.GuessTheNpcCount,
		HangmanCount:     guessCounts.HangmanCount,
		TerraTriviaCount: guessCounts.TerraTriviaCount,
	}
}
