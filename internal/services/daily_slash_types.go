package services

import "terrariadle-backend/internal/domain"

type DailySlashInitData struct {
	PreviousWeapon domain.WeaponPreview
	GuessedIDs     []int
	Guesses        []WeaponGuess
	HasWon         bool
}

type WeaponGuess struct {
	Weapon WeaponData
	Checks domain.WeaponChecks
}

type WeaponData struct {
	ID         int
	Name       string
	ImagePath  string
	DamageType string
	Damage     int
	UseTime    string
	Rarity     string
	Operation  string
	Material   string
	Obtained   []string
}

type DailySlashCheckData struct {
	Finished      bool
	GuessedWeapon domain.Weapon
	GuessChecks   domain.WeaponChecks
}

type DailySlashWinningData struct {
	Position    int
	PlayerCount int
}
