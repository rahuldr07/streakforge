# Commands

## Habit lifecycle

- `streakforge add <name> [description]`
- `streakforge list`
- `streakforge today`
- `streakforge next` prints unfinished active habits for today, prioritized by the streak each habit can preserve.
- `streakforge missed` prints active habits that were not completed yesterday, including their last completion date.
- `streakforge done <habit-id>`
- `streakforge log <habit-id> <yyyy-mm-dd>` records a completion for a specific date.
- `streakforge undo <habit-id>` removes today's completion from an active habit.
- `streakforge show <habit-id>` prints one habit's details, status, streak, and recent completions.
- `streakforge search <query>` finds active habits by name or description.
- `streakforge archive <habit-id>`
- `streakforge rename <habit-id> <name> [description]`

## Reporting

- `streakforge streaks`
- `streakforge pulse` prints the last seven days of active habit completions.
- `streakforge calendar [days]` prints recent daily completion rates across active habits.
- `streakforge stats` prints active and archived habit counts, total completions, best current streak, and the largest completion count on a single habit.
- `streakforge export` writes CSV with habit IDs, names, descriptions, creation timestamps, archive status, and completion counts.
- `streakforge export json` writes the full habit store as formatted JSON.
- `streakforge backup` writes a timestamped copy of the JSON store next to the original file.
- `streakforge restore <backup-path>` restores a backup and saves the previous current store first.
