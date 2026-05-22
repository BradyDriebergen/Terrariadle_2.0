package services

import "terrariadle-backend/internal/domain"

func toPreview(w domain.Weapon) domain.WeaponPreview {
	return domain.WeaponPreview{
		Name:   w.Name,
		Path:   w.Info.ImagePath,
		Rarity: w.Info.Rarity,
	}
}

func toWeaponData(w domain.Weapon) WeaponData {
	return WeaponData{
		ID:         w.ID,
		Name:       w.Name,
		ImagePath:  w.Info.ImagePath,
		DamageType: w.Info.DamageType,
		Damage:     w.Info.Damage,
		UseTime:    w.Info.UseTime,
		Rarity:     w.Info.Rarity,
		Operation:  w.Info.Operation,
		Material:   w.Info.Material,
		Obtained:   w.Info.Obtained,
	}
}
