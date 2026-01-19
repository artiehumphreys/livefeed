import Link from "next/link";
import { GameSummary } from "@/lib/types";

export function GameCard({ game }: { game: GameSummary }) {
  return (
    <Link href={`/game/${game.GameID}`}>
      <div className="border p-4 rounded hover:bg-gray-50 cursor-pointer">
        <div className="font-semibold">
          {game.AwayTeam} @ {game.HomeTeam}
        </div>
        <div className="text-sm text-gray-600">
          {game.StartTime} - {game.State}
        </div>
      </div>
    </Link>
  );
}
