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
