import { useEffect, useState } from "react";

export function useGame(gameId: string) {
  const [snapshot, setSnapshot] = useState<any>(null);

  useEffect(() => {
    const fetchData = async () => {
      await fetch(`http://localhost:8080/set-game?gameId=${gameId}`);

      const res = await fetch(`http://localhost:8080/snapshot`);
      if (!res.ok) return;
      setSnapshot(await res.json());
    };

    fetchData();
    const id = setInterval(fetchData, 5000);
    return () => clearInterval(id);
  }, [gameId]);

  return snapshot;
}
