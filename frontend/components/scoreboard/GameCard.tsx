import Link from "next/link";
import { GameSummary } from "@/lib/types";

export function GameCard({ game }: { game: GameSummary }) {
  return (
    <Link href={`/game/${game.GameID}`}>
      <div className="border p-4 rounded hover:bg-stone-50 cursor-pointer">
        <div className="font-semibold text-gray-200">
          {game.AwayTeam} @ {game.HomeTeam}
        </div>

        <div className="text-sm text-gray-500">
          {game.State === "pre"
            ? `${game.StartTime}`
            : `${game.AwayScore} - ${game.HomeScore} : ${game.State}`}
        </div>
      </div>
    </Link>
  );
}
