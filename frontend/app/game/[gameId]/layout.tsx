"use client";

import { useGame } from "@/hooks/useGame";
import { ScoreHeader } from "@/components/game/ScoreHeader";

export default function GameLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { gameId: string };
}) {
  const snapshot = useGame(params.gameId);
  if (!snapshot) return <p>Loading…</p>;

  return (
    <div className="max-w-5xl mx-auto p-4">
      <ScoreHeader snapshot={snapshot} />
      {children}
    </div>
  );
}
