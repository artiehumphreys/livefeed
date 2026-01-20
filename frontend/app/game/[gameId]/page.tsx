import { redirect } from "next/navigation";

export default async function GamePage({
  params,
}: {
  params: { gameId: string };
}) {
  const p = await params;
  redirect(`/game/${p.gameId}/boxscore`);
}
