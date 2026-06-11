export const CompareResult = {
	Lower: -1,
	Match: 0,
	Higher: 1,
	NoMatch: 2,
	PartialMatch: 3
} as const;

// if (checks.damage === CompareResult.Higher) { ... }
export type CompareResult = (typeof CompareResult)[keyof typeof CompareResult];

export interface WeaponPreview {
	name: string;
	path: string;
	rarity: string;
}

export interface WeaponCheck {
	weapon_id: number;
	damage_type: boolean;
	damage: CompareResult;
	use_time: CompareResult;
	rarity: CompareResult;
	operation: boolean;
	material: boolean;
	obtained: CompareResult;
}

export interface Weapon {
	id: number;
	name: string;
	image_path: string;
	damage_type: string;
	damage: number;
	use_time: string;
	rarity: string;
	operation: string;
	material: string;
	obtained: string[];
}

export interface WeaponGuess {
	weapon: Weapon;
	checks: WeaponCheck;
}

export interface WeaponListItem {
	weapon_id: number;
	name: string;
	path: string;
}

export interface DailySlashSession {
	previous_weapon: WeaponPreview;
	guessed_ids: number[];
	guesses: WeaponGuess[];
	finished: boolean;
}

export interface DailySlashCheckResult {
	finished: boolean;
	guess_result: WeaponGuess;
}

export interface DailySlashWinningData {
	position: number;
}
