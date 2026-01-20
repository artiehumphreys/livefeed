export function ScoreHeader({ snapshot: raw }: { snapshot: any }) {
  const snapshot = raw?.snapshot ?? raw; // unwrap if needed

  console.log(snapshot);

  if (!snapshot?.Teams || snapshot.Teams.length !== 2) {
    return (
      <div className="flex justify-center">
        <div className="text-2xl font-bold">Loading Game...</div>
      </div>
    );
  }

  const homeBox = snapshot.BoxScore?.TeamBoxes?.find((t: any) => t.IsHome);
  const awayBox = snapshot.BoxScore?.TeamBoxes?.find((t: any) => !t.IsHome);

  const home =
    snapshot.Teams.find((t: any) => t.TeamID === homeBox?.TeamID) ??
    snapshot.Teams[0];
  const away =
    snapshot.Teams.find((t: any) => t.TeamID === awayBox?.TeamID) ??
    snapshot.Teams[1];

  return (
    <div className="flex justify-between items-center border-b pb-4">
      <div>
        <div className="text-2xl font-bold">{away?.Name}</div>
      </div>

      <div className="text-3xl font-bold">
        {awayBox?.Stats?.PTS ?? "–"} – {homeBox?.Stats?.PTS ?? "–"}
      </div>

      <div>
        <div className="text-2xl font-bold">{home?.Name}</div>
      </div>
    </div>
  );
}
