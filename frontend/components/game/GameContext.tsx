"use client";

import { createContext, useContext } from "react";

type GameContextType = {
  snapshot: any;
  error: string | null;
};

const GameContext = createContext<GameContextType | null>(null);

export function GameProvider({
  value,
  children,
}: {
  value: GameContextType;
  children: React.ReactNode;
}) {
  return <GameContext.Provider value={value}>{children}</GameContext.Provider>;
}

export function useGameContext() {
  const ctx = useContext(GameContext);
  if (!ctx) throw new Error("useGameContext must be used within GameProvider");
  return ctx;
}
