import { useEffect, useState } from "react";
import { GameSummary, ScoreboardResponse } from "@/types";

export function useScoreboard(interval: number) {
  const [games, setGames] = useState<GameSummary[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await fetch("http://localhost:8080/scoreboard");
        if (!res.ok) throw new Error("failed");

        const data: ScoreboardResponse = await res.json();

        setGames(data.Games ?? []);
        setLoading(false);
      } catch {
        setGames([]);
        setLoading(false);
      }
    };

    fetchData();
    const id = setInterval(fetchData, interval);
    return () => clearInterval(id);
  }, [interval]);

  return { games, loading };
}
