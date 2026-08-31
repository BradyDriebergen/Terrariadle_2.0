package repo

import (
	"terrariadle/internal/domain"
)

func toAnswerRef(ad answerData) domain.AnswerRefs {
	return domain.AnswerRefs{
		DailySlash: domain.WeaponRef{
			CurrentWeaponID: ad.DailySlash.CurrentWeaponID,
			PrevWeaponID:    ad.DailySlash.PrevWeaponID,
		},
		Connections: domain.ConnectionRef{
			CategoryIDs: ad.Connections.CategoryIDs,
			Options:     toCategoryOption(ad.Connections.Options),
		},
		GuessTheNpc: domain.NpcRef{
			NpcID:       ad.GuessTheNpc.NpcID,
			Quote:       ad.GuessTheNpc.Quote,
			Name:        ad.GuessTheNpc.Name,
			NameOptions: ad.GuessTheNpc.NameOptions,
		},
		Hangman: domain.HangmanRef{
			EnemyID: ad.Hangman.EnemyID,
		},
		TerraTrivia: domain.TerraTriviaRef{
			QuestionIDs: ad.TerraTrivia.QuestionIDs,
		},
		ResetTime:     ad.ResetTime,
		NextResetTime: ad.NextResetTime,
	}
}

func toAnswerData(da domain.AnswerRefs) answerData {
	return answerData{
		DailySlash: weaponData{
			CurrentWeaponID: da.DailySlash.CurrentWeaponID,
			PrevWeaponID:    da.DailySlash.PrevWeaponID,
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
			EnemyID: da.Hangman.EnemyID,
		},
		TerraTrivia: terraTriviaData{
			QuestionIDs: da.TerraTrivia.QuestionIDs,
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

func toPlayerGuessCounts(gc guessCounts) domain.PlayerGuessCounts {
	return domain.PlayerGuessCounts{
		DailySlashCount:  gc.DailySlashCount,
		ConnectionsCount: gc.ConnectionsCount,
		GuessTheNpcCount: gc.GuessTheNpcCount,
		HangmanCount:     gc.HangmanCount,
		TerraTriviaCount: gc.TerraTriviaCount,
	}
}

func toGuessCounts(gc domain.PlayerGuessCounts) guessCounts {
	return guessCounts{
		DailySlashCount:  gc.DailySlashCount,
		ConnectionsCount: gc.ConnectionsCount,
		GuessTheNpcCount: gc.GuessTheNpcCount,
		HangmanCount:     gc.HangmanCount,
		TerraTriviaCount: gc.TerraTriviaCount,
	}
}
