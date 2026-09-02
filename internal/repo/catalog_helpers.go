package repo

import "terrariadle/internal/domain"

func toWeapon(w weapon) domain.Weapon {
	return domain.Weapon{
		ID:           w.ID,
		Name:         w.Name,
		WeaponType:   w.WeaponType,
		ModeObtained: w.ModeObtained,
		Info: domain.WeaponInfo{
			ImagePath:  w.Info.ImagePath,
			DamageType: w.Info.DamageType,
			Damage:     w.Info.Damage,
			UseTime:    w.Info.UseTime,
			Rarity:     w.Info.Rarity,
			Operation:  w.Info.Operation,
			Material:   w.Info.Material,
			Obtained:   w.Info.Obtained,
		},
	}
}

func toWeapons(src []weapon) []domain.Weapon {
	result := make([]domain.Weapon, len(src))
	for i, w := range src {
		result[i] = toWeapon(w)
	}
	return result
}

func toCategory(c category) domain.Category {
	return domain.Category{
		ID:       c.ID,
		Category: c.Category,
		Options:  c.Options,
	}
}

func toCategories(src []category) []domain.Category {
	result := make([]domain.Category, len(src))
	for i, c := range src {
		result[i] = toCategory(c)
	}

	return result
}

func toNpc(n npc) domain.Npc {
	return domain.Npc{
		ID:      n.ID,
		NPC:     n.NPC,
		NpcPath: n.NpcPath,
		Quotes:  n.Quotes,
		Names:   n.Names,
	}
}

func toNpcs(src []npc) []domain.Npc {
	result := make([]domain.Npc, len(src))
	for i, n := range src {
		result[i] = toNpc(n)
	}

	return result
}

func toEnemy(e enemy) domain.Enemy {
	return domain.Enemy{
		ID:        e.ID,
		Name:      e.Name,
		ImagePath: e.ImagePath,
	}
}

func toEnemies(src []enemy) []domain.Enemy {
	result := make([]domain.Enemy, len(src))
	for i, e := range src {
		result[i] = toEnemy(e)
	}

	return result
}

func toTriviaQuestion(t triviaQuestion) domain.TriviaQuestion {
	return domain.TriviaQuestion{
		ID:         t.ID,
		Answer:     t.Answer,
		Clue:       t.Clue,
		Chunks:     t.Chunks,
		ChunkCount: t.ChunkCount,
	}
}

func toTriviaQuestions(src []triviaQuestion) []domain.TriviaQuestion {
	result := make([]domain.TriviaQuestion, len(src))
	for i, e := range src {
		result[i] = toTriviaQuestion(e)
	}

	return result
}
