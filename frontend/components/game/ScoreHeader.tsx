export function ScoreHeader({ snapshot }: { snapshot: any }) {
  const [home, away] = snapshot.Teams;

  return (
    <div className="flex justify-between items-center border-b pb-4">
      <div className="text-xl font-bold">
        {away.Name} @ {home.Name}
      </div>
      <div className="text-sm">`${snapshot.BoxScore?.Period}`</div>
      <div className="text-lg">
        {snapshot.BoxScore?.Clock
          ? `${snapshot.BoxScore.Clock.Minutes}:${snapshot.BoxScore.Clock.Seconds}`
          : snapshot.BoxScore?.Period}
      </div>
    </div>
  );
}
