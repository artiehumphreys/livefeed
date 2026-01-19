"use client";

import { useScoreboard } from "@/lib/polling";
import { GameCard } from "./GameCard";

export function GameGrid() {
  const { games, loading } = useScoreboard(60_000);

  if (loading) return <p>Loading games…</p>;

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
      {games.map((g) => (
        <GameCard key={g.GameID} game={g} />
      ))}
    </div>
  );
}
