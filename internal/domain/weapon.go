package domain

type Weapon struct {
	ID           int
	Name         string
	WeaponType   string
	ModeObtained string
	Info         WeaponInfo
}

type WeaponInfo struct {
	ImagePath  string
	DamageType string
	Damage     int
	UseTime    string
	Rarity     string
	Operation  string
	Material   string
	Obtained   []string
}

type WeaponChecks struct {
	WeaponID   int
	DamageType bool
	Damage     CompareResult
	UseTime    CompareResult
	Rarity     CompareResult
	Operation  bool
	Material   bool
	Obtained   CompareResult
}

type SearchWeaponResult struct {
	WeaponID int
	Name     string
	Path     string
}
