# Changelog

## Unreleased

- Add `path` to print the active store file.
- Add `today done` and `today todo` filters.
- Add `list --all` to include archived habits.
- Add `stats json` for machine-readable summaries.
- Add `repair` to deduplicate and sort completion history.
- Add `calendar` to show recent daily completion rates.
- Add `doctor` to report common local store issues.
- Add `lagging` to find habits quiet for several days.
- Add `export completions` for one CSV row per completion event.
- Sort `streaks` by strongest current streak and include latest completion dates.
- Add `log` to record completions for a specific date.
- Add `undo` to remove today's completion from a habit.
- Add `show` to inspect one habit's details and recent completions.
- Add `search` to find active habits by name or description.
- Show the best current active streak in `stats`.
- Add `next` to prioritize today's unfinished habits by streaks at risk.
- Reject completions for archived habits.
- Show the longest per-habit completion count in `stats`.
- Create store directories automatically and replace store files through a temporary file.
- Include descriptions and creation timestamps in CSV exports.
- Add today status reporting.
- Add backup restore support.
- Add completion rate to stats output.

## 0.1.0

- Added habit creation, listing, completion, and archiving.
- Added streak and summary reporting.
- Added CSV export.
- Added JSON storage and tests.
