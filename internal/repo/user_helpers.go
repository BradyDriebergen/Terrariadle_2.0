package repo

import (
	"terrariadle/internal/domain"
	"time"
)

func toUser(u userData) domain.User {
	checks := make([]domain.WeaponChecks, len(u.DailySlash.Checks))
	for i, c := range u.DailySlash.Checks {
		checks[i] = domain.WeaponChecks{
			WeaponID:   c.WeaponID,
			DamageType: c.DamageType,
			Damage:     c.Damage,
			UseTime:    c.UseTime,
			Rarity:     c.Rarity,
			Operation:  c.Operation,
			Material:   c.Material,
			Obtained:   c.Obtained,
		}
	}

	return domain.User{
		UserID: u.UserID,
		DailySlash: domain.DailySlashGame{
			Game:   toGame(u.DailySlash.Game),
			Checks: checks,
		},
		Connections: domain.ConnectionGame{
			Game:     toGame(u.Connections.Game),
			Attempts: u.Connections.Attempts,
		},
		GuessTheNPC: domain.GuessTheNpcGame{
			Game:        toGame(u.GuessTheNPC.Game),
			GuessedName: u.GuessTheNPC.GuessedName,
		},
		Hangman: domain.HangmanGame{
			Game:     toGame(u.Hangman.Game),
			Attempts: u.Hangman.Attempts,
		},
		TerraTrivia: domain.TerraTriviaGame{
			Game: toGame(u.TerraTrivia.Game),
		},
		LastSeen: time.Now(),
		Dirty:    false,
	}
}

func toGame(g game) domain.Game {
	return domain.Game{
		Guesses:  g.Guesses,
		Finished: g.HasWon,
		Position: g.Position,
	}
}

func toUserData(u domain.User) userData {
	checks := make([]weaponChecks, len(u.DailySlash.Checks))
	for i, c := range u.DailySlash.Checks {
		checks[i] = weaponChecks{
			WeaponID:   c.WeaponID,
			DamageType: c.DamageType,
			Damage:     c.Damage,
			UseTime:    c.UseTime,
			Rarity:     c.Rarity,
			Operation:  c.Operation,
			Material:   c.Material,
			Obtained:   c.Obtained,
		}
	}

	return userData{
		UserID: u.UserID,
		DailySlash: dailySlashGame{
			Game:   toGameData(u.DailySlash.Game),
			Checks: checks,
		},
		Connections: connectionGame{
			Game:     toGameData(u.Connections.Game),
			Attempts: u.Connections.Attempts,
		},
		GuessTheNPC: guessTheNpcGame{
			Game:        toGameData(u.GuessTheNPC.Game),
			GuessedName: u.GuessTheNPC.GuessedName,
		},
		Hangman: hangmanGame{
			Game:     toGameData(u.Hangman.Game),
			Attempts: u.Hangman.Attempts,
		},
		TerraTrivia: terraTriviaGame{
			Game: toGameData(u.TerraTrivia.Game),
		},
	}
}

func toGameData(g domain.Game) game {
	return game{
		Guesses:  g.Guesses,
		HasWon:   g.Finished,
		Position: g.Position,
	}
}
