"use client";

import { redirect } from "next/navigation";
import { useParams } from "next/navigation";

export default function GamePage() {
  const p = useParams();
  redirect(`/game/${p.gameId}/boxscore`);
}
