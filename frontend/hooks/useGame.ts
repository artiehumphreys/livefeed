"use client";

import { useEffect, useState } from "react";

export function useGame(gameId: string | undefined) {
  const [snapshot, setSnapshot] = useState<any>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!gameId) return;

    let intervalId: NodeJS.Timeout;

    const start = async () => {
      const validate = await fetch(
        `http://localhost:8080/validate-game?gameId=${gameId}`,
      );

      if (!validate.ok) {
        setError(`Invalid game ID ${gameId}`);
        return;
      }

      await fetch(`http://localhost:8080/set-game?gameId=${gameId}`);

      const fetchSnapshot = async () => {
        const res = await fetch(`http://localhost:8080/snapshot`);
        if (res.ok) {
          setSnapshot(await res.json());
        }
      };

      await fetchSnapshot();

      intervalId = setInterval(fetchSnapshot, 5000);
    };

    start();

    return () => {
      if (intervalId) clearInterval(intervalId);
    };
  }, [gameId]);

  return { snapshot, error };
}
