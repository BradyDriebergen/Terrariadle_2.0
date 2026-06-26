export interface HangmanSession {
  attempts: number;
  finished: boolean;
  phrase: string[];
  guesses: HangmanGuess[];
}

export interface HangmanGuess {
    letter: string;
    correct: boolean;
}

export interface HangmanCheckResult {
  phrase: string[];
  finished: boolean;
  guess: HangmanGuess;
  attempts: number;
}

export interface HangmanWinningData {
  position: number;
  enemy_name: string;
  enemy_path: string;
}