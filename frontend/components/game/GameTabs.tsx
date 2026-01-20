"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

function Tab({
  href,
  label,
  active,
}: {
  href: string;
  label: string;
  active: boolean;
}) {
  return (
    <Link
      href={href}
      className={[
        "px-3 py-2 rounded-md text-sm font-medium",
        active ? "bg-black text-white" : "text-gray-600 hover:bg-gray-100",
      ].join(" ")}
    >
      {label}
    </Link>
  );
}

export function GameTabs({ gameId }: { gameId: string }) {
  const pathname = usePathname();

  const base = `/game/${gameId}`;
  const tabs = [
    { href: `${base}/boxscore`, label: "Boxscore" },
    { href: `${base}/runs`, label: "Runs" },
    { href: `${base}/pbp`, label: "Play-by-play" },
  ];

  return (
    <div className="flex gap-4 justify-center border-b pb-3 mt-4">
      {tabs.map((t) => (
        <Tab
          key={t.href}
          href={t.href}
          label={t.label}
          active={pathname === t.href}
        />
      ))}
    </div>
  );
}
