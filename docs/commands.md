# Commands

## Habit lifecycle

- `streakforge add <name> [description]`
- `streakforge list`
- `streakforge today`
- `streakforge done <habit-id>`
- `streakforge archive <habit-id>`
- `streakforge rename <habit-id> <name> [description]`

## Reporting

- `streakforge streaks`
- `streakforge pulse` prints the last seven days of active habit completions.
- `streakforge stats` prints active and archived habit counts, total completions, and the largest completion count on a single habit.
- `streakforge export` writes CSV with habit IDs, names, descriptions, creation timestamps, archive status, and completion counts.
- `streakforge backup` writes a timestamped copy of the JSON store next to the original file.
- `streakforge restore <backup-path>` restores a backup and saves the previous current store first.
