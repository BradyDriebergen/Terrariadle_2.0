export interface TriviaItem {
  id: number;
  clue: string;
  letter_count: number;
  answer: string;
}

export interface TerraTriviaSession {
  finished: boolean;
  trivia_items: TriviaItem[];
  chunks: string[];
}

export interface TerraTriviaCheckResult {
  finished: boolean;
  guess_result: TriviaItem;
  is_correct: boolean;
}

export interface TerraTriviaWinningData {
  position: number;
}