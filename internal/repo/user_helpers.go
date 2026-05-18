package repo

import "terrariadle-backend/internal/domain"

func (u *userData) toDomain() domain.User {
	checks := make([]domain.WeaponChecks, len(u.dailySlash.checks))
	for i, c := range u.dailySlash.checks {
		checks[i] = domain.WeaponChecks{
			DamageType: c.damageType,
			Damage:     c.damage,
			UseTime:    c.useTime,
			Rarity:     c.rarity,
			Operation:  c.operation,
			Material:   c.material,
			Obtained:   c.obtained,
		}
	}

	return domain.User{
		UserID: u.userID,
		DailySlash: domain.DailySlashGame{
			Game:   toGameDomain(u.dailySlash.game),
			Checks: checks,
		},
		Connections: domain.ConnectionGame{
			Game:     toGameDomain(u.connections.game),
			Attempts: u.connections.attempts,
		},
		GuessTheNPC: domain.GuessTheNpcGame{
			Game:        toGameDomain(u.guessTheNPC.game),
			GuessedName: u.guessTheNPC.guessedName,
		},
		Hangman: domain.HangmanGame{
			Game:     toGameDomain(u.hangman.game),
			Attempts: u.hangman.attempts,
		},
	}
}

func toGameDomain(g game) domain.Game {
	return domain.Game{
		Guesses:  g.guesses,
		HasWon:   g.hasWon,
		Position: g.position,
	}
}

func fromDomain(u domain.User) userData {
	checks := make([]weaponChecks, len(u.DailySlash.Checks))
	for i, c := range u.DailySlash.Checks {
		checks[i] = weaponChecks{
			damageType: c.DamageType,
			damage:     c.Damage,
			useTime:    c.UseTime,
			rarity:     c.Rarity,
			operation:  c.Operation,
			material:   c.Material,
			obtained:   c.Obtained,
		}
	}

	return userData{
		dailySlash: dailySlashGame{
			game:   fromGameDomain(u.DailySlash.Game),
			checks: checks,
		},
		connections: connectionGame{
			game:     fromGameDomain(u.Connections.Game),
			attempts: u.Connections.Attempts,
		},
		guessTheNPC: guessTheNpcGame{
			game:        fromGameDomain(u.GuessTheNPC.Game),
			guessedName: u.GuessTheNPC.GuessedName,
		},
		hangman: hangmanGame{
			game:     fromGameDomain(u.Hangman.Game),
			attempts: u.Hangman.Attempts,
		},
	}
}

func fromGameDomain(g domain.Game) game {
	return game{
		guesses:  g.Guesses,
		hasWon:   g.HasWon,
		position: g.Position,
	}
}
