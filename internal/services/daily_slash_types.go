package services

import "terrariadle-backend/internal/domain"

type DailySlashInitData struct {
	PreviousWeapon WeaponPreview `json:"previous_weapon"`
	GuessedIDs     []int         `json:"guessed_ids"`
	Guesses        []WeaponGuess `json:"guesses"`
	Finished       bool          `json:"finished"`
}

type WeaponGuess struct {
	Weapon WeaponData      `json:"weapon"`
	Checks WeaponCheckData `json:"checks"`
}

type WeaponData struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	ImagePath  string   `json:"image_path"`
	DamageType string   `json:"damage_type"`
	Damage     int      `json:"damage"`
	UseTime    string   `json:"use_time"`
	Rarity     string   `json:"rarity"`
	Operation  string   `json:"operation"`
	Material   string   `json:"material"`
	Obtained   []string `json:"obtained"`
}

type SearchWeaponData struct {
	WeaponID int    `json:"weapon_id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

type DailySlashCheckData struct {
	Finished    bool        `json:"finished"`
	GuessResult WeaponGuess `json:"guess_result"`
}

type DailySlashWinningData struct {
	Position int `json:"position"`
}

type WeaponPreview struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Rarity string `json:"rarity"`
}

type WeaponCheckData struct {
	WeaponID   int                  `json:"weapon_id"`
	DamageType bool                 `json:"damage_type"`
	Damage     domain.CompareResult `json:"damage"`
	UseTime    domain.CompareResult `json:"use_time"`
	Rarity     domain.CompareResult `json:"rarity"`
	Operation  bool                 `json:"operation"`
	Material   bool                 `json:"material"`
	Obtained   domain.CompareResult `json:"obtained"`
}
