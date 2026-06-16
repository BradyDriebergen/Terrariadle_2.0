export interface ConnectionsSession {
  attempts: number;
  finished: boolean;
  options: string[];
  solved_categories: SolvedCategory[];
}

export interface CategoryOption {
    id: number;
    option: string;
}

export interface SolvedCategory {
  name: string;
  options: string[];
}

export interface ConnectionsCheckResult {
  attempts: number;
  is_correct: boolean;
  one_away: boolean;
  correct_guess: SolvedCategory;
}

export interface ConnectionsWinningData {
	position: number;
}