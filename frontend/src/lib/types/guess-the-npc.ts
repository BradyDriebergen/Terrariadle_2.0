export interface GuessTheNpcGuessData {
  name: string;
  path: string;
}

export interface GuessTheNpcSession {
  quote: string;
  finished: boolean;
  guesses: GuessTheNpcGuessData[];
}

export interface GuessTheNpcCheckResult {
  finished: boolean;
  guess: GuessTheNpcGuessData;
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