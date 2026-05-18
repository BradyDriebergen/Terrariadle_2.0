package repo

import "terrariadle-backend/internal/domain"

func toDomainWeapons(src []weapon) []domain.Weapon {
	result := make([]domain.Weapon, len(src))
	for i, w := range src {
		result[i] = domain.Weapon{
			ID:           w.id,
			Name:         w.name,
			WeaponType:   w.weaponType,
			ModeObtained: w.modeObtained,
			Info: domain.WeaponInfo{
				ImagePath:  w.info.imagePath,
				DamageType: w.info.damageType,
				Damage:     w.info.damage,
				UseTime:    w.info.useTime,
				Rarity:     w.info.rarity,
				Operation:  w.info.operation,
				Material:   w.info.material,
				Obtained:   w.info.obtained,
			},
		}
	}
	return result
}

func toDomainCategories(src []category) []domain.Category {
	result := make([]domain.Category, len(src))
	for i, c := range src {
		result[i] = domain.Category{
			ID:       c.id,
			Category: c.category,
			Options:  c.options,
		}
	}

	return result
}

func toDomainNpcs(src []npc) []domain.Npc {
	result := make([]domain.Npc, len(src))
	for i, n := range src {
		result[i] = domain.Npc{
			ID:      n.id,
			NPC:     n.npc,
			NPCPath: n.npcPath,
			Quotes:  n.quotes,
			Names:   n.names,
		}
	}

	return result
}

func toDomainEnemies(src []enemy) []domain.Enemy {
	result := make([]domain.Enemy, len(src))
	for i, e := range src {
		result[i] = domain.Enemy{
			ID:        e.id,
			Name:      e.name,
			ImagePath: e.imagePath,
		}
	}

	return result
}
