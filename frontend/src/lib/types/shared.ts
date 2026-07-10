export interface DropdownListItem {
	id: number;
	name: string;
	path: string;
}

export interface UserGameResults {
	daily_slash: boolean;
	connections: boolean;
	guess_the_npc: boolean;
	hangman: boolean;
	terratrivia: boolean;
}
