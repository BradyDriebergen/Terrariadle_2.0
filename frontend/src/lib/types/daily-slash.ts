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

export type Rarity =
	| 'White'
	| 'Blue'
	| 'Green'
	| 'Orange'
	| 'Light Red'
	| 'Pink'
	| 'Light Purple'
	| 'Lime'
	| 'Yellow'
	| 'Cyan'
	| 'Red';

export const colors: Record<Rarity, string> = {
	White: '#ffffff',
	Blue: '#177cff',
	Green: '#35d400',
	Orange: '#ffa600',
	'Light Red': '#ff8080',
	Pink: '#ffbdf0',
	'Light Purple': '#be9bdc',
	Lime: '#69f75c',
	Yellow: '#fbff00',
	Cyan: '#00fff7',
	Red: '#ff0000'
};

export const backgrounds: Record<Rarity, string> = {
	White: "url('/daily-slash/backgrounds/WoodWall.png')",
	Blue: "url('/daily-slash/backgrounds/StoneWall.png')",
	Green: "url('/daily-slash/backgrounds/DungeonWall.png')",
	Orange: "url('/daily-slash/backgrounds/HellstoneWall.png')",
	'Light Red': "url('/daily-slash/backgrounds/PearlstoneWall.png')",
	Pink: "url('/daily-slash/backgrounds/CrystalBlockWall.png')",
	'Light Purple': "url('/daily-slash/backgrounds/ChlorophyteBrickWall.png')",
	Lime: "url('/daily-slash/backgrounds/LihzahrdBrickWall.png')",
	Yellow: "url('/daily-slash/backgrounds/MartianConduitWall.png')",
	Cyan: "url('/daily-slash/backgrounds/SmoothMarbleWall.png')",
	Red: "url('/daily-slash/backgrounds/LuminiteBrickWall.png')"
};

export const borders: Record<Rarity, string> = {
	White: "url('/daily-slash/borders/Wood.png')",
	Blue: "url('/daily-slash/borders/StoneBlock.png')",
	Green: "url('/daily-slash/borders/DungeonBrick.png')",
	Orange: "url('/daily-slash/borders/HellstoneBrick.png')",
	'Light Red': "url('/daily-slash/borders/PearlstoneBlock.png')",
	Pink: "url('/daily-slash/borders/CrystalBlock.png')",
	'Light Purple': "url('/daily-slash/borders/ChlorophyteBrick.png')",
	Lime: "url('/daily-slash/borders/LihzahrdBrick.png')",
	Yellow: "url('/daily-slash/borders/MartianConduitPlating.png')",
	Cyan: "url('/daily-slash/borders/SmoothMarbleBlock.png')",
	Red: "url('/daily-slash/borders/LuminiteBrick.png')"
};
