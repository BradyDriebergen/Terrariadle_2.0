package repo

import "terrariadle-backend/internal/domain"

func (u *userData) toDomain() domain.User {
	checks := make([]domain.WeaponChecks, len(u.DailySlash.Checks))
	for i, c := range u.DailySlash.Checks {
		checks[i] = domain.WeaponChecks{
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
			Game:   toGameDomain(u.DailySlash.Game),
			Checks: checks,
		},
		Connections: domain.ConnectionGame{
			Game:     toGameDomain(u.Connections.Game),
			Attempts: u.Connections.Attempts,
		},
		GuessTheNPC: domain.GuessTheNpcGame{
			Game:        toGameDomain(u.GuessTheNPC.Game),
			GuessedName: u.GuessTheNPC.GuessedName,
		},
		Hangman: domain.HangmanGame{
			Game:     toGameDomain(u.Hangman.Game),
			Attempts: u.Hangman.Attempts,
		},
	}
}

func toGameDomain(g game) domain.Game {
	return domain.Game{
		Guesses:  g.Guesses,
		HasWon:   g.HasWon,
		Position: g.Position,
	}
}

func fromDomain(u domain.User) userData {
	checks := make([]weaponChecks, len(u.DailySlash.Checks))
	for i, c := range u.DailySlash.Checks {
		checks[i] = weaponChecks{
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
		DailySlash: dailySlashGame{
			Game:   fromGameDomain(u.DailySlash.Game),
			Checks: checks,
		},
		Connections: connectionGame{
			Game:     fromGameDomain(u.Connections.Game),
			Attempts: u.Connections.Attempts,
		},
		GuessTheNPC: guessTheNpcGame{
			Game:        fromGameDomain(u.GuessTheNPC.Game),
			GuessedName: u.GuessTheNPC.GuessedName,
		},
		Hangman: hangmanGame{
			Game:     fromGameDomain(u.Hangman.Game),
			Attempts: u.Hangman.Attempts,
		},
	}
}

func fromGameDomain(g domain.Game) game {
	return game{
		Guesses:  g.Guesses,
		HasWon:   g.HasWon,
		Position: g.Position,
	}
}
