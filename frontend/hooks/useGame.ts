import { useEffect, useState } from "react";

export function useGame(gameId: string) {
  const [snapshot, setSnapshot] = useState<any>(null);

  useEffect(() => {
    const fetchData = async () => {
      const res = await fetch(
        `http://localhost:8080/snapshot?gameId=${gameId}`,
      );
      if (!res.ok) return;
      setSnapshot(await res.json());
    };

    fetchData();
    const id = setInterval(fetchData, 5000);
    return () => clearInterval(id);
  }, [gameId]);

  return snapshot;
}
