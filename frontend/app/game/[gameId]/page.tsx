"use client";

import { useGameSnapshot } from "@/lib/polling";
import { ScoreHeader } from "@/components/game/ScoreHeader";
import { GameTabs } from "@/components/game/GameTabs";

export default function GamePage({ params }: { params: { gameId: string } }) {
  const { snapshot, loading, error } = useGameSnapshot(params.gameId);

  if (loading) return <p>Loading game…</p>;
  if (error || !snapshot) return <p>Error loading game</p>;

  return (
    <div className="p-6">
      <ScoreHeader snapshot={snapshot} />
      <GameTabs snapshot={snapshot} />
    </div>
  );
}
