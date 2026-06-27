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
- Local-first storage: keep everything in `~/.streakforge.json`, with readable JSON and no account.
- Durable history: archive habits without deleting completion history.
- Useful reporting: view habit details, search habits, streaks, weekly pulse, summary stats, and CSV exports.
- Recovery tools: create timestamped backups and restore them when needed.

## Quick Start

Requires Go 1.22 or newer.

```bash
go run ./cmd/streakforge help
go run ./cmd/streakforge add read "Read ten pages"
go run ./cmd/streakforge list
go run ./cmd/streakforge today
go run ./cmd/streakforge next
go run ./cmd/streakforge done <habit-id>
go run ./cmd/streakforge log <habit-id> 2026-06-26
go run ./cmd/streakforge show <habit-id>
go run ./cmd/streakforge streaks
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
| `add <name> [description]` | Create a habit. |
| `list` | Show active habits. |
| `today` | Show today's done/todo status. |
| `next` | Show unfinished habits prioritized by streaks at risk. |
| `done <habit-id>` | Complete a habit for today. |
| `log <habit-id> <yyyy-mm-dd>` | Complete a habit for a specific date. |
| `undo <habit-id>` | Remove today's completion from a habit. |
| `show <habit-id>` | Show details, status, streak, and recent completions. |
| `search <query>` | Find active habits by name or description. |
| `streaks` | Show current streak lengths. |
| `pulse` | Show the last seven days of completion counts. |
| `stats` | Show active, archived, completion, and rate totals. |
| `archive <habit-id>` | Hide a habit without deleting its history. |
| `rename <habit-id> <name> [description]` | Update a habit's name or description. |
| `backup` | Save a timestamped copy of the store. |
| `restore <backup-path>` | Restore a backup after saving the current store. |
| `export` | Print habit data as CSV. |

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
