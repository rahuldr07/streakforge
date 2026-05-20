# Storage

StreakForge stores data in a single JSON file:

```text
~/.streakforge.json
```

The file contains a list of habits, completion dates, archive status, and creation timestamps.
It is intentionally human-readable so it can be backed up, inspected, or moved between machines.

Writes are staged through a temporary file in the same directory before replacing the store file. The
store directory is created when needed, and files are written with owner-only permissions where the
platform supports them.

