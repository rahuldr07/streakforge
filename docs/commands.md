# Commands

## Habit lifecycle

- `streakforge path` prints the active store path.
- `streakforge add <name> [description]`
- `streakforge list [--all]` prints active habits, or includes archived habits with status labels.
- `streakforge today [done|todo]` prints today's status, optionally filtered to completed or pending habits.
- `streakforge next` prints unfinished active habits for today, prioritized by the streak each habit can preserve.
- `streakforge missed` prints active habits that were not completed yesterday, including their last completion date.
- `streakforge lagging [days]` prints active habits that have not been completed for the given number of days.
- `streakforge done <habit-id>`
- `streakforge log <habit-id> <yyyy-mm-dd>` records a completion for a specific date.
- `streakforge undo <habit-id>` removes today's completion from an active habit.
- `streakforge show <habit-id>` prints one habit's details, status, streak, and recent completions.
- `streakforge search [--all] <query>` finds habits by name or description, optionally including archived habits.
- `streakforge archive <habit-id>`
- `streakforge unarchive <habit-id>` restores an archived habit.
- `streakforge rename <habit-id> <name> [description]`

## Reporting

- `streakforge streaks` prints current streaks sorted strongest first, including the latest completion date.
- `streakforge pulse` prints the last seven days of active habit completions.
- `streakforge calendar [days]` prints recent daily completion rates across active habits.
- `streakforge stats [json]` prints active and archived habit counts, total completions, best current streak, and the largest completion count on a single habit.
- `streakforge doctor` checks the local store for duplicate IDs, empty names, invalid dates, duplicate completions, and future completions.
- `streakforge repair` deduplicates completion dates and sorts each habit's completion history.
- `streakforge export` writes CSV with habit IDs, names, descriptions, creation timestamps, archive status, and completion counts.
- `streakforge export json` writes the full habit store as formatted JSON.
- `streakforge export completions` writes one CSV row per completion event.
- `streakforge backup` writes a timestamped copy of the JSON store next to the original file.
- `streakforge restore <backup-path>` restores a backup and saves the previous current store first.
