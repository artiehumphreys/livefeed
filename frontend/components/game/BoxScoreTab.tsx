"use client";

import { useGameContext } from "./GameContext";

function safe(n: any) {
  return n == null ? "–" : n;
}

function pct(n: any, digits = 1) {
  if (n == null) return "–";
  return `${(Number(n) * 100).toFixed(digits)}%`;
}

function fixed(n: any, digits = 2) {
  if (n == null) return "–";
  const x = Number(n);
  if (Number.isNaN(x)) return "–";
  return x.toFixed(digits);
}

function StatRow({
  label,
  away,
  home,
}: {
  label: string;
  away: any;
  home: any;
}) {
  return (
    <div className="grid grid-cols-3 gap-2 py-2 text-sm border-b last:border-b-0">
      <div className="text-gray-600">{label}</div>
      <div className="text-right font-medium">{safe(away)}</div>
      <div className="text-right font-medium">{safe(home)}</div>
    </div>
  );
}

function Card({
  title,
  children,
}: {
  title: string;
  children: React.ReactNode;
}) {
  return (
    <div className="border rounded-lg overflow-hidden bg-white">
      <div className="px-4 py-3 border-b font-semibold">{title}</div>
      <div className="px-4 py-3 ">{children}</div>
    </div>
  );
}

function TeamBoxTable({ teamBox }: { teamBox: any }) {
  const players = (teamBox?.Players ?? []).slice();

  players.sort((a: any, b: any) => (b.MP ?? 0) - (a.MP ?? 0));

  const Row = ({ p }: { p: any }) => (
    <tr className="border-t text-gray-600">
      <td className="px-3 py-2 whitespace-nowrap">
        <div className="">
          {p.FirstName} {p.LastName}
        </div>
        <div className="text-xs text-gray-400">{p.Position}</div>
      </td>
      <td className="text-right px-3 py-2">{safe(p.MP)}</td>
      <td className="text-right px-3 py-2">{safe(p.PTS ?? 0)}</td>
      <td className="text-right px-3 py-2">{safe(p.REB ?? 0)}</td>
      <td className="text-right px-3 py-2">{safe(p.AST ?? 0)}</td>
      <td className="text-right px-3 py-2">{safe(p.TO ?? 0)}</td>
      <td className="text-right px-3 py-2">
        {safe(p.FGM)}-{safe(p.FGA)}
      </td>
      <td className="text-right px-3 py-2">
        {safe(p.TPM)}-{safe(p.TPA)}
      </td>
      <td className="text-right px-3 py-2">
        {safe(p.FTM)}-{safe(p.FTA)}
      </td>
    </tr>
  );

  return (
    <div className="border rounded-lg overflow-hidden bg-white">
      <div className="px-3 py-3 border-b flex items-center justify-between">
        <div className="font-semibold text-gray-700">{teamBox?.Name}</div>
        <div className="text-sm text-gray-600">
          Team PTS{" "}
          <span className="font-semibold text-gray-900">
            {teamBox?.Stats?.PTS ?? "–"}
          </span>
        </div>
      </div>

      <div className="overflow-auto">
        <table className="w-full text-sm text-gray-700">
          <thead className="bg-gray-50">
            <tr>
              <th className="text-left px-3 py-2">Player</th>
              <th className="text-right px-3 py-2">MIN</th>
              <th className="text-right px-3 py-2">PTS</th>
              <th className="text-right px-3 py-2">REB</th>
              <th className="text-right px-3 py-2">AST</th>
              <th className="text-right px-3 py-2">TO</th>
              <th className="text-right px-3 py-2">FG</th>
              <th className="text-right px-3 py-2">3PT</th>
              <th className="text-right px-3 py-2">FT</th>
            </tr>
          </thead>

          <tbody>
            {players.map((p: any) => (
              <Row key={p.ID} p={p} />
            ))}
          </tbody>
        </table>
      </div>

      <div className="px-4 py-3 border-t bg-gray-50 text-sm flex flex-wrap gap-x-4 gap-y-1 text-gray-700">
        <div>
          FG:{" "}
          <span className="font-semibold text-gray-900">
            {teamBox?.Stats?.FGM}-{teamBox?.Stats?.FGA}
          </span>{" "}
          <span className="text-gray-500">
            ({teamBox?.Stats?.FGpct ?? "–"}%)
          </span>
        </div>
        <div>
          3PT:{" "}
          <span className="font-semibold text-gray-900">
            {teamBox?.Stats?.TPM}-{teamBox?.Stats?.TPA}
          </span>{" "}
          <span className="text-gray-500">
            ({teamBox?.Stats?.TPpct ?? "–"}%)
          </span>
        </div>
        <div>
          FT:{" "}
          <span className="font-semibold text-gray-900">
            {teamBox?.Stats?.FTM}-{teamBox?.Stats?.FTA}
          </span>{" "}
          <span className="text-gray-500">
            ({teamBox?.Stats?.FTpct ?? "–"}%)
          </span>
        </div>
      </div>
    </div>
  );
}

function TeamStatsCompare({
  awayBox,
  homeBox,
}: {
  awayBox: any;
  homeBox: any;
}) {
  const a = awayBox?.Stats ?? {};
  const h = homeBox?.Stats ?? {};

  return (
    <Card title="Team Stats">
      <div className="grid grid-cols-3 gap-2 pb-2 text-xs uppercase tracking-wide text-gray-500">
        <div />
        <div className="text-right">{awayBox?.Name ?? "Away"}</div>
        <div className="text-right">{homeBox?.Name ?? "Home"}</div>
      </div>

      <div className="divide-y text-gray-700">
        <StatRow label="Points" away={a.PTS} home={h.PTS} />
        <StatRow
          label="FG"
          away={`${a.FGM}-${a.FGA} (${safe(a.FGpct)}%)`}
          home={`${h.FGM}-${h.FGA} (${safe(h.FGpct)}%)`}
        />
        <StatRow
          label="3PT"
          away={`${a.TPM}-${a.TPA} (${safe(a.TPpct)}%)`}
          home={`${h.TPM}-${h.TPA} (${safe(h.TPpct)}%)`}
        />
        <StatRow
          label="FT"
          away={`${a.FTM}-${a.FTA} (${safe(a.FTpct)}%)`}
          home={`${h.FTM}-${h.FTA} (${safe(h.FTpct)}%)`}
        />
        <StatRow label="Rebounds" away={a.REB} home={h.REB} />
        <StatRow label="Off Reb" away={a.OREB} home={h.OREB} />
        <StatRow label="Assists" away={a.AST} home={h.AST} />
        <StatRow label="Turnovers" away={a.TO} home={h.TO} />
        <StatRow label="Steals" away={a.STL} home={h.STL} />
        <StatRow label="Blocks" away={a.BLK} home={h.BLK} />
        <StatRow label="Fouls" away={a.PF} home={h.PF} />
      </div>
    </Card>
  );
}

function AdvancedMetricsCompare({
  awayTeam,
  homeTeam,
}: {
  awayTeam: any;
  homeTeam: any;
}) {
  const a = awayTeam?.Metrics ?? {};
  const h = homeTeam?.Metrics ?? {};

  return (
    <Card title="Advanced">
      <div className="grid grid-cols-3 gap-2 pb-2 text-xs uppercase tracking-wide text-gray-500">
        <div />
        <div className="text-right">{awayTeam?.Name ?? "Away"}</div>
        <div className="text-right">{homeTeam?.Name ?? "Home"}</div>
      </div>

      <div className="divide-y text-gray-600">
        <StatRow label="PPP" away={fixed(a.PPP)} home={fixed(h.PPP)} />
        <StatRow
          label="TO%"
          away={pct(a.TurnoverPct)}
          home={pct(h.TurnoverPct)}
        />
        <StatRow
          label="OREB%"
          away={pct(a.OffensiveReboundPct)}
          home={pct(h.OffensiveReboundPct)}
        />
        <StatRow
          label="DREB%"
          away={pct(a.DefensiveReboundPct)}
          home={pct(h.DefensiveReboundPct)}
        />
        <StatRow
          label="3PA Rate"
          away={pct(a.ThreePointAttemptRate)}
          home={pct(h.ThreePointAttemptRate)}
        />
        <StatRow
          label="FT Rate"
          away={pct(a.FreeThrowRate)}
          home={pct(h.FreeThrowRate)}
        />
        <StatRow
          label="eFG%"
          away={fixed(a.EffectiveFGPct)}
          home={fixed(h.EffectiveFGPct)}
        />
      </div>
    </Card>
  );
}

export function BoxScoreTab() {
  const { snapshot, error } = useGameContext();
  const s = snapshot?.snapshot ?? snapshot;

  if (error) return <div className="text-red-600">{error}</div>;
  if (!s?.BoxScore?.TeamBoxes || !s?.Teams) return <div>Loading...</div>;

  const boxes = s.BoxScore.TeamBoxes as any[];
  const homeBox = boxes.find((t) => t.IsHome);
  const awayBox = boxes.find((t) => !t.IsHome);

  // Match computed metrics teams to home/away by TeamID
  const homeTeam = s.Teams.find((t: any) => t.TeamID === homeBox?.TeamID);
  const awayTeam = s.Teams.find((t: any) => t.TeamID === awayBox?.TeamID);

  return (
    <div className="grid grid-cols-1 lg:grid-cols-12 gap-6">
      {/* Left: box tables */}
      <div className="lg:col-span-8 space-y-6">
        <TeamBoxTable teamBox={awayBox} />
        <TeamBoxTable teamBox={homeBox} />
      </div>

      {/* Right: comparisons */}
      <div className="lg:col-span-4 space-y-6">
        <TeamStatsCompare awayBox={awayBox} homeBox={homeBox} />
        <AdvancedMetricsCompare awayTeam={awayTeam} homeTeam={homeTeam} />
      </div>
    </div>
  );
}
