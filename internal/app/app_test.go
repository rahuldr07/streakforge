package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestHelpCommand(t *testing.T) {
	var out bytes.Buffer
	err := RunWithConfig([]string{"help"}, Config{Stdout: &out, Home: t.TempDir()})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "streakforge add") {
		t.Fatalf("help output missing add command: %s", out.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	var out bytes.Buffer
	err := RunWithConfig([]string{"nope"}, Config{Stdout: &out, Home: t.TempDir()})
	if err == nil {
		t.Fatal("expected unknown command error")
	}
}

func TestDefaultConfigUsesUserHomeDirectory(t *testing.T) {
	cfg := DefaultConfig(&bytes.Buffer{}, &bytes.Buffer{})

	if cfg.Home == "" {
		t.Fatal("Home is empty")
	}
	if cfg.Stdout == nil {
		t.Fatal("Stdout is nil")
	}
	if cfg.Stderr == nil {
		t.Fatal("Stderr is nil")
	}
}

func TestPathPrintsStorePath(t *testing.T) {
	home := t.TempDir()
	var out bytes.Buffer
	err := RunWithConfig([]string{"path"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	want := filepath.Join(home, ".streakforge.json")
	if strings.TrimSpace(out.String()) != want {
		t.Fatalf("path = %q, want %q", strings.TrimSpace(out.String()), want)
	}
}

func TestListAllIncludesArchivedHabits(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "read", "name": "Read"},
    {"id": "old", "name": "Old", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"list", "--all"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "read  active    Read") {
		t.Fatalf("list output missing active habit: %s", out.String())
	}
	if !strings.Contains(out.String(), "old  archived  Old") {
		t.Fatalf("list output missing archived habit: %s", out.String())
	}
}

func TestStatsPrintsLongestCompletionCount(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "a", "name": "Read", "completions": ["2026-05-17"]},
    {"id": "b", "name": "Run", "completions": ["2026-05-16", "2026-05-17"]}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"stats"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "longest completions: 2") {
		t.Fatalf("stats output missing longest completions: %s", out.String())
	}
}

func TestStatsPrintsBestCurrentStreak(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	today := now.Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "a", "name": "Read", "completions": ["` + yesterday + `", "` + today + `"]},
    {"id": "b", "name": "Run", "completions": ["` + today + `"]}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"stats"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "best current streak: 2") {
		t.Fatalf("stats output missing best current streak: %s", out.String())
	}
}

func TestStatsJSONPrintsSummary(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "a", "name": "Read", "completions": ["2026-05-17"]},
    {"id": "b", "name": "Run", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"stats", "json"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), `"active_habits": 1`) {
		t.Fatalf("json stats missing active count: %s", out.String())
	}
	if !strings.Contains(out.String(), `"archived_habits": 1`) {
		t.Fatalf("json stats missing archived count: %s", out.String())
	}
}

func TestStreaksPrintsSortedReports(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	today := now.Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "walk", "name": "Walk", "completions": ["` + today + `"]},
    {"id": "read", "name": "Read", "completions": ["` + yesterday + `", "` + today + `"]}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"streaks"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("lines = %d, want 2; output=%q", len(lines), out.String())
	}
	if !strings.Contains(lines[0], "read  2 days  last: "+today+"  Read") {
		t.Fatalf("first streak line = %q", lines[0])
	}
	if !strings.Contains(lines[1], "walk  1 day  last: "+today+"  Walk") {
		t.Fatalf("second streak line = %q", lines[1])
	}
}

func TestNextPrintsPendingHabitsByStreakPriority(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	today := now.Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	twoDaysAgo := now.AddDate(0, 0, -2).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + twoDaysAgo + `", "` + yesterday + `"]},
    {"id": "walk", "name": "Walk", "completions": ["` + yesterday + `"]},
    {"id": "water", "name": "Water", "completions": ["` + today + `"]},
    {"id": "old", "name": "Old", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"next"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("lines = %d, want 2; output=%q", len(lines), out.String())
	}
	if !strings.Contains(lines[0], "read  2 days  Read") {
		t.Fatalf("first next line = %q", lines[0])
	}
	if !strings.Contains(lines[1], "walk  1 day  Walk") {
		t.Fatalf("second next line = %q", lines[1])
	}
}

func TestTodayFiltersDoneAndTodo(t *testing.T) {
	home := t.TempDir()
	today := time.Now().Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + today + `"]},
    {"id": "walk", "name": "Walk"}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var done bytes.Buffer
	err := RunWithConfig([]string{"today", "done"}, Config{Stdout: &done, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(done.String(), "done  read  Read") || strings.Contains(done.String(), "walk") {
		t.Fatalf("done filter output = %s", done.String())
	}

	var todo bytes.Buffer
	err = RunWithConfig([]string{"today", "todo"}, Config{Stdout: &todo, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(todo.String(), "todo  walk  Walk") || strings.Contains(todo.String(), "read") {
		t.Fatalf("todo filter output = %s", todo.String())
	}
}

func TestMissedPrintsHabitsSkippedYesterday(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	twoDaysAgo := now.AddDate(0, 0, -2).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + twoDaysAgo + `"]},
    {"id": "walk", "name": "Walk", "completions": ["` + yesterday + `"]},
    {"id": "water", "name": "Water"},
    {"id": "old", "name": "Old", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"missed"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	output := out.String()
	if !strings.Contains(output, "water  last: never") {
		t.Fatalf("missed output missing never-completed habit: %s", output)
	}
	if !strings.Contains(output, "read  last: "+twoDaysAgo) {
		t.Fatalf("missed output missing older completion: %s", output)
	}
	if strings.Contains(output, "walk") || strings.Contains(output, "old") {
		t.Fatalf("missed output included completed or archived habits: %s", output)
	}
}

func TestLaggingPrintsQuietHabits(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	fourDaysAgo := now.AddDate(0, 0, -4).Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + fourDaysAgo + `"]},
    {"id": "walk", "name": "Walk", "completions": ["` + yesterday + `"]},
    {"id": "water", "name": "Water"}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"lagging", "3"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	output := out.String()
	if !strings.Contains(output, "water  quiet: 4 days  last: never") {
		t.Fatalf("lagging output missing never-completed habit: %s", output)
	}
	if !strings.Contains(output, "read  quiet: 4 days  last: "+fourDaysAgo) {
		t.Fatalf("lagging output missing quiet habit: %s", output)
	}
	if strings.Contains(output, "walk") {
		t.Fatalf("lagging output included recent habit: %s", output)
	}
}

func TestCalendarPrintsCompletionRates(t *testing.T) {
	home := t.TempDir()
	now := time.Now()
	today := now.Format(time.DateOnly)
	yesterday := now.AddDate(0, 0, -1).Format(time.DateOnly)
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "completions": ["` + today + `", "` + yesterday + `"]},
    {"id": "walk", "name": "Walk", "completions": ["` + today + `"]}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"calendar", "2"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	output := out.String()
	if !strings.Contains(output, yesterday+"  1/2   50%") {
		t.Fatalf("calendar output missing yesterday: %s", output)
	}
	if !strings.Contains(output, today+"  2/2  100%") {
		t.Fatalf("calendar output missing today: %s", output)
	}
}

func TestPulseAcceptsDayCount(t *testing.T) {
	home := t.TempDir()
	var out bytes.Buffer
	err := RunWithConfig([]string{"pulse", "3"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("lines = %d, want 3; output=%q", len(lines), out.String())
	}
}

func TestDoctorPrintsStoreHealthIssues(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "read", "name": "Read"},
    {"id": "read", "name": "Read Again"}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"doctor"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "error  duplicate habit id: read") {
		t.Fatalf("doctor output missing duplicate id: %s", out.String())
	}
}

func TestRepairCleansStore(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [{"id": "read", "name": "Read", "completions": ["2026-07-05", "2026-07-03", "2026-07-05"]}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"repair"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "removed duplicate completions: 1") {
		t.Fatalf("repair output missing duplicate count: %s", out.String())
	}
	updated, err := os.ReadFile(filepath.Join(home, ".streakforge.json"))
	if err != nil {
		t.Fatalf("read store: %v", err)
	}
	if strings.Count(string(updated), "2026-07-05") != 1 {
		t.Fatalf("store still has duplicate completion: %s", string(updated))
	}
}

func TestUnarchiveRestoresHabit(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [{"id": "read", "name": "Read", "archived": true}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"unarchive", "read"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), `unarchived "Read"`) {
		t.Fatalf("unexpected unarchive output: %s", out.String())
	}
	updated, err := os.ReadFile(filepath.Join(home, ".streakforge.json"))
	if err != nil {
		t.Fatalf("read store: %v", err)
	}
	if strings.Contains(string(updated), `"archived": true`) {
		t.Fatalf("store still marks habit archived: %s", string(updated))
	}
}

func TestExportJSONPrintsStore(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [{"id": "read", "name": "Read"}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"export", "json"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), `"id": "read"`) {
		t.Fatalf("json export missing habit id: %s", out.String())
	}
}

func TestExportCompletionsPrintsCompletionRows(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [{"id": "read", "name": "Read", "completions": ["2026-07-05"]}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"export", "completions"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "read,Read,2026-07-05,false") {
		t.Fatalf("completion export missing row: %s", out.String())
	}
}

func TestExportActiveSkipsArchivedHabits(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [
  {"id": "read", "name": "Read"},
  {"id": "old", "name": "Old", "archived": true}
]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"export", "active"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "read,Read") {
		t.Fatalf("active export missing active habit: %s", out.String())
	}
	if strings.Contains(out.String(), "old,Old") {
		t.Fatalf("active export included archived habit: %s", out.String())
	}
}

func TestLogCompletesHabitForSpecificDate(t *testing.T) {
	home := t.TempDir()
	data := `{"habits": [{"id": "read", "name": "Read"}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"log", "read", "2026-05-17"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	updated, err := os.ReadFile(filepath.Join(home, ".streakforge.json"))
	if err != nil {
		t.Fatalf("read store: %v", err)
	}
	if !strings.Contains(string(updated), "2026-05-17") {
		t.Fatalf("store missing logged date: %s", string(updated))
	}
	if !strings.Contains(out.String(), "completed \"Read\" for 2026-05-17") {
		t.Fatalf("unexpected output: %s", out.String())
	}
}

func TestUndoRemovesTodayCompletion(t *testing.T) {
	home := t.TempDir()
	today := time.Now().Format(time.DateOnly)
	data := `{"habits": [{"id": "read", "name": "Read", "completions": ["` + today + `"]}]}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"undo", "read"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	updated, err := os.ReadFile(filepath.Join(home, ".streakforge.json"))
	if err != nil {
		t.Fatalf("read store: %v", err)
	}
	if strings.Contains(string(updated), today) {
		t.Fatalf("store still has today's completion: %s", string(updated))
	}
	if !strings.Contains(out.String(), "removed today's completion") {
		t.Fatalf("unexpected output: %s", out.String())
	}
}

func TestShowPrintsHabitDetails(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {
      "id": "read",
      "name": "Read",
      "description": "Read ten pages",
      "created_at": "2026-05-01T00:00:00Z",
      "completions": ["2026-05-17", "2026-05-18"]
    }
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"show", "read"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "description: Read ten pages") {
		t.Fatalf("show output missing description: %s", out.String())
	}
	if !strings.Contains(out.String(), "recent: 2026-05-18, 2026-05-17") {
		t.Fatalf("show output missing recent completions: %s", out.String())
	}
}

func TestSearchFindsMatchingHabits(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "read", "name": "Read", "description": "Books before bed"},
    {"id": "run", "name": "Run"},
    {"id": "old", "name": "Old Books", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"search", "books"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "read  Read  Books before bed") {
		t.Fatalf("search output missing match: %s", out.String())
	}
	if strings.Contains(out.String(), "old") {
		t.Fatalf("search output included archived habit: %s", out.String())
	}
}

func TestSearchAllIncludesArchivedHabits(t *testing.T) {
	home := t.TempDir()
	data := `{
  "habits": [
    {"id": "old", "name": "Old Books", "archived": true}
  ]
}`
	if err := os.WriteFile(filepath.Join(home, ".streakforge.json"), []byte(data), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	var out bytes.Buffer
	err := RunWithConfig([]string{"search", "--all", "books"}, Config{Stdout: &out, Home: home})
	if err != nil {
		t.Fatalf("RunWithConfig returned error: %v", err)
	}
	if !strings.Contains(out.String(), "old  archived  Old Books") {
		t.Fatalf("search --all output missing archived habit: %s", out.String())
	}
}
