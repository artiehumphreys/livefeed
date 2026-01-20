import { GameGrid } from "@/components/scoreboard/GameGrid";

export default function Home() {
  return (
    <div className="p-6">
      <main className="text-3xl font-bold mb-4 ">
        <h1 className="mb-4 ">Today&apos;s Slate</h1>
        <GameGrid></GameGrid>
      </main>
    </div>
  );
}
