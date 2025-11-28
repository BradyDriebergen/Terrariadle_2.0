export interface DailySlashInitializer {
	previousWeaponData: PreviousWeapon;
	hint1: string;
	hint2: string;
	hint3: string;
}

export interface PreviousWeapon {
	name: string;
	path: string;
	rarity: string;
}

export interface SimpleWeapon {
	id: number;
	name: string;
	path: string;
}

export interface Weapon {
  id: number;
  name: string;
  "weapon-type": string;
  "mode-obtained": string;
  info: WeaponInfo;
}

export interface WeaponInfo {
  "image-path": string;
  "damage-type": string;
  damage: number;
  "use-time": string;
  rarity: string;
  operation: string;
  material: string;
  obtained: string[];
}