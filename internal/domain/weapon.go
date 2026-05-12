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

type SearchWeaponResult struct {
	WeaponId int
	Name     string
	Path     string
}
