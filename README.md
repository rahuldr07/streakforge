# StreakForge

StreakForge is a small Go CLI for tracking habits, tasks, and daily streaks in a local JSON file.
It is designed to be simple, portable, and easy to inspect.

## Features

- Add and list active habits
- Review today's done/todo status
- See the next unfinished habits to protect streaks at risk
- Mark habits complete for today
- Track current streaks
- Archive habits without deleting history
- Print summary statistics
- Restore from timestamped backups
- Export habit data as CSV
- Store everything in `~/.streakforge.json`

## Quick Start

Requires Go 1.22 or newer.

```bash
go run ./cmd/streakforge help
go run ./cmd/streakforge add read "Read ten pages"
go run ./cmd/streakforge list
go run ./cmd/streakforge today
go run ./cmd/streakforge next
go run ./cmd/streakforge done <habit-id>
go run ./cmd/streakforge streaks
```

`next` prints unfinished active habits for today, sorted by the current streak you can preserve by
completing them.

## Development

```bash
make test
make build
```

`make build` writes a local `streakforge` binary in the repository root.

## Docs

- [Commands](docs/commands.md)
- [Storage](docs/storage.md)
- [Example data](examples/streakforge.json)
