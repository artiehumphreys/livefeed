"use client";

import { useParams } from "next/navigation";
import { useGame } from "@/hooks/useGame";
import { GameProvider } from "@/components/game/GameContext";
import { ScoreHeader } from "@/components/game/ScoreHeader";
import { GameTabs } from "@/components/game/GameTabs";

export default function GameLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const params = useParams<{ gameId: string }>();
  const { snapshot, error } = useGame(params?.gameId);

  return (
    <GameProvider value={{ snapshot, error }}>
      <div className="max-w-7xl mx-auto px-4 py-6">
        <ScoreHeader snapshot={snapshot} />
        <GameTabs gameId={params.gameId} />
        <div className="pt-4">{children}</div>
      </div>
    </GameProvider>
  );
}
