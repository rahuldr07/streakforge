# StreakForge

StreakForge is a small Go CLI for tracking habits, tasks, and daily streaks in a local JSON file.
It is designed to be simple, portable, and easy to inspect.

## Features

- Add and list active habits
- Mark habits complete for today
- Track current streaks
- Archive habits without deleting history
- Print summary statistics
- Export habit data as CSV
- Store everything in `~/.streakforge.json`

## Quick Start

```bash
go run ./cmd/streakforge help
go run ./cmd/streakforge add read "Read ten pages"
go run ./cmd/streakforge list
go run ./cmd/streakforge done <habit-id>
go run ./cmd/streakforge streaks
```

## Development

```bash
make test
make build
```

## Docs

- [Commands](docs/commands.md)
- [Storage](docs/storage.md)
- [Example data](examples/streakforge.json)
