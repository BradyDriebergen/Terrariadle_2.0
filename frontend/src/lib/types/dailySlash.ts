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

export type Rarity = 'White' | 'Blue' | 'Green' | 'Orange' | 'Light Red' | 'Pink' | 'Light Purple' | 'Lime' | 'Yellow' | 'Cyan' | 'Red';

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