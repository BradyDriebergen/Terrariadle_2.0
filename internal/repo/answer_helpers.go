package repo

import (
	"terrariadle/internal/domain"
)

func toAnswerData(da domain.DailyAnswers) answerData {
	triviaQuestionIDs := make([]int, 7)
	for i, q := range da.TerraTrivia.Questions {
		triviaQuestionIDs[i] = q.ID
	}

	return answerData{
		DailySlash: weaponData{
			CurrentWeaponID: da.DailySlash.CurrentWeapon.ID,
			PrevWeaponID:    da.DailySlash.PrevWeapon.ID,
		},
		Connections: connectionData{
			CategoryIDs: da.Connections.CategoryIDs,
			Options:     toCategoryAnswer(da.Connections.Options),
		},
		GuessTheNpc: npcData{
			NpcID:       da.GuessTheNpc.NpcID,
			Quote:       da.GuessTheNpc.Quote,
			Name:        da.GuessTheNpc.Name,
			NameOptions: da.GuessTheNpc.NameOptions,
		},
		Hangman: hangmanData{
			EnemyID: da.Hangman.Enemy.ID,
		},
		TerraTrivia: terraTriviaData{
			QuestionIDs: triviaQuestionIDs,
		},
		ResetTime:     da.ResetTime,
		NextResetTime: da.NextResetTime,
	}
}

func toCategoryAnswer(answers []domain.ConnectionOption) []connectionOption {
	options := make([]connectionOption, len(answers))
	for i, a := range answers {
		options[i] = connectionOption{
			Option:     a.Option,
			CategoryID: a.CategoryID,
		}
	}
	return options
}

func toCategoryOption(options []connectionOption) []domain.ConnectionOption {
	answers := make([]domain.ConnectionOption, len(options))
	for i, o := range options {
		answers[i] = domain.ConnectionOption{
			Option:     o.Option,
			CategoryID: o.CategoryID,
		}
	}
	return answers
}
