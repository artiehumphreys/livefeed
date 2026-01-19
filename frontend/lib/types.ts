export interface GameSummary {
  GameID: number;
  HomeTeam: string;
  AwayTeam: string;
  HomeScore: number;
  AwayScore: number;
  State: "pre" | "live" | "final";
  StartTime: string;
  Clock: string;
}

export interface ScoreboardResponse {
  Games: GameSummary[];
  LastUpdated: number;
}
