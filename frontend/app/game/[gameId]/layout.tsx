"use client";

import { useGame } from "@/hooks/useGame";
import { ScoreHeader } from "@/components/game/ScoreHeader";
import { useParams } from "next/navigation";

export default function GameLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const p = useParams();
  const snapshot = useGame(p.gameId);
  if (!snapshot) return <p>Loading...</p>;

  return (
    <div className="max-w-5xl mx-auto p-4">
      <ScoreHeader snapshot={snapshot} />
      {children}
    </div>
  );
}
