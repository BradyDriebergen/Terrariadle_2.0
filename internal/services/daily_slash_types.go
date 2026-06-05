package services

import "terrariadle-backend/internal/domain"

type DailySlashInitData struct {
	PreviousWeapon domain.WeaponPreview `json:"previous_weapon"`
	GuessedIDs     []int                `json:"guessed_ids"`
	Guesses        []WeaponGuess        `json:"guesses"`
	Finished       bool                 `json:"finished"`
}

type WeaponGuess struct {
	Weapon WeaponData          `json:"weapon"`
	Checks domain.WeaponChecks `json:"checks"`
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
