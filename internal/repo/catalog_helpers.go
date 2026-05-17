package repo

import "terrariadle-backend/internal/domain"

func toDomainWeapons(src []weapon) []domain.Weapon {
	result := make([]domain.Weapon, len(src))
	for i, w := range src {
		result[i] = domain.Weapon{
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
	return result
}

func toDomainCategories(src []category) []domain.Category {
	result := make([]domain.Category, len(src))
	for i, c := range src {
		result[i] = domain.Category{
			ID:       c.ID,
			Category: c.Category,
			Options:  c.Options,
		}
	}

	return result
}

func toDomainNpcs(src []npc) []domain.Npc {
	result := make([]domain.Npc, len(src))
	for i, n := range src {
		result[i] = domain.Npc{
			ID:      n.ID,
			NPC:     n.NPC,
			NPCPath: n.NPCPath,
			Quotes:  n.Quotes,
			Names:   n.Names,
		}
	}

	return result
}

func toDomainEnemies(src []enemy) []domain.Enemy {
	result := make([]domain.Enemy, len(src))
	for i, e := range src {
		result[i] = domain.Enemy{
			ID:        e.ID,
			Name:      e.Name,
			ImagePath: e.ImagePath,
		}
	}

	return result
}
