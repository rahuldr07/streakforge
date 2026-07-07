# StreakForge

StreakForge is a tiny Go CLI for building habits without giving your data to a service.
It tracks daily completions, streaks, backups, and exports in one local JSON file that is easy to
inspect, move, or restore.

```bash
streakforge next
read  2 days  Read ten pages
walk  1 day   Evening walk
```

## Features

- Daily habit tracking: add habits, mark them done, and review today's todo/done status.
- Flexible completion editing: backfill a missed date with `log` or undo today's completion.
- Streak protection: use `next` to see unfinished habits sorted by the streak you can preserve.
- Recovery view: use `missed` to find habits skipped yesterday and catch up quickly.
- Lagging view: find habits that have been quiet for several days.
- Local-first storage: keep everything in `~/.streakforge.json`, with readable JSON and no account.
- Store discovery: use `path` to print the active data file location.
- Durable history: archive habits without deleting completion history.
- Useful reporting: view habit details, search habits, streaks, weekly pulse, summary stats, and CSV exports.
- Calendar view: scan recent completion rates across all active habits.
- Recovery tools: create timestamped backups and restore them when needed.
- Store diagnostics: run `doctor` to spot duplicate IDs, bad dates, and future completions.
- Store repair: run `repair` to deduplicate and sort completion history.

## Quick Start

Requires Go 1.22 or newer.

```bash
go run ./cmd/streakforge help
go run ./cmd/streakforge path
go run ./cmd/streakforge add read "Read ten pages"
go run ./cmd/streakforge list
go run ./cmd/streakforge today
go run ./cmd/streakforge next
go run ./cmd/streakforge missed
go run ./cmd/streakforge lagging 3
go run ./cmd/streakforge done <habit-id>
go run ./cmd/streakforge log <habit-id> 2026-06-26
go run ./cmd/streakforge show <habit-id>
go run ./cmd/streakforge streaks
go run ./cmd/streakforge calendar 14
go run ./cmd/streakforge doctor
go run ./cmd/streakforge repair
```

## Everyday Flow

```bash
# See what still needs attention today.
go run ./cmd/streakforge next

# Mark one habit complete.
go run ./cmd/streakforge done <habit-id>

# Backfill yesterday if you forgot to log it.
go run ./cmd/streakforge log <habit-id> 2026-06-26

# Check your streaks.
go run ./cmd/streakforge streaks
```

`next` skips completed and archived habits, then puts the longest current streaks first so you can
protect the habits most at risk.

## Commands

| Command | What it does |
| --- | --- |
| `path` | Print the active store path. |
| `add <name> [description]` | Create a habit. |
| `list [--all]` | Show active habits, or include archived habits. |
| `today [done|todo]` | Show today's status, optionally filtered. |
| `next` | Show unfinished habits prioritized by streaks at risk. |
| `missed` | Show active habits that were not completed yesterday. |
| `lagging [days]` | Show active habits quiet for the given number of days. |
| `done <habit-id>` | Complete a habit for today. |
| `log <habit-id> <yyyy-mm-dd>` | Complete a habit for a specific date. |
| `undo <habit-id>` | Remove today's completion from a habit. |
| `show <habit-id>` | Show details, status, streak, and recent completions. |
| `search <query>` | Find active habits by name or description. |
| `streaks` | Show current streak lengths sorted strongest first. |
| `pulse` | Show the last seven days of completion counts. |
| `calendar [days]` | Show recent daily completion rates for active habits. |
| `stats [json]` | Show active, archived, completion, and rate totals. |
| `doctor` | Check the local store for common data issues. |
| `repair` | Deduplicate and sort completion history. |
| `archive <habit-id>` | Hide a habit without deleting its history. |
| `rename <habit-id> <name> [description]` | Update a habit's name or description. |
| `backup` | Save a timestamped copy of the store. |
| `restore <backup-path>` | Restore a backup after saving the current store. |
| `export [csv|json|completions]` | Print habit data as CSV, formatted JSON, or completion-event CSV. |

## Storage

StreakForge writes to:

```text
~/.streakforge.json
```

The store is intentionally plain JSON. You can inspect it, back it up, sync it with your own tools,
or move it to another machine.

## Development

```bash
make test
make build
```

`make build` writes a local `streakforge` binary in the repository root.

Run the binary directly after building:

```bash
./streakforge help
```

## Docs

- [Commands](docs/commands.md)
- [Storage](docs/storage.md)
- [Example data](examples/streakforge.json)
