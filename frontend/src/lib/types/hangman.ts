export interface HangmanSession {
  attempts: number;
  finished: boolean;
  phrase: string[];
  guesses: string[];
}

export interface HangmanCheckResult {
  phrase: string[];
  finished: boolean;
  is_correct: boolean;
  attempts: number;
}

export interface HangmanWinningData {
  position: number;
  enemy_name: string;
  enemy_path: string;
}