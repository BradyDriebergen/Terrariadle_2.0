export interface NpcGuess {
	name: string;
	path: string;
}

export interface GuessTheNpcSession {
	quote: string;
	finished: boolean;
	guesses: NpcGuess[];
	guessed_ids: number[];
}

export interface GuessTheNpcCheckResult {
	finished: boolean;
	guess: NpcGuess;
}

export interface GuessTheNpcWinningData {
	position: number;
	names: string[];
	guessed_name: string;
	correct_name: string;
}

export interface GuessTheNpcMiniGameResult {
	guessed_name: string;
	correct_name: string;
}
