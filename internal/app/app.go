package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/rahuldr07/streakforge/internal/habit"
)

type Config struct {
	Stdout io.Writer
	Stderr io.Writer
	Home   string
}

func Run(args []string) error {
	return RunWithConfig(args, DefaultConfig(os.Stdout, os.Stderr))
}

func DefaultConfig(stdout io.Writer, stderr io.Writer) Config {
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.Getenv("HOME")
	}
	return Config{
		Stdout: stdout,
		Stderr: stderr,
		Home:   home,
	}
}

func RunWithConfig(args []string, cfg Config) error {
	if len(args) == 0 {
		printHelp(cfg.Stdout)
		return nil
	}

	switch args[0] {
	case "path":
		return runPath(cfg)
	case "add":
		return runAdd(args[1:], cfg)
	case "list":
		return runList(args[1:], cfg)
	case "today":
		return runToday(args[1:], cfg)
	case "next":
		return runNext(cfg)
	case "missed":
		return runMissed(cfg)
	case "lagging":
		return runLagging(args[1:], cfg)
	case "done":
		return runDone(args[1:], cfg)
	case "log":
		return runLog(args[1:], cfg)
	case "undo":
		return runUndo(args[1:], cfg)
	case "show":
		return runShow(args[1:], cfg)
	case "search":
		return runSearch(args[1:], cfg)
	case "streaks":
		return runStreaks(cfg)
	case "stats":
		return runStats(args[1:], cfg)
	case "pulse":
		return runPulse(cfg)
	case "calendar":
		return runCalendar(args[1:], cfg)
	case "doctor":
		return runDoctor(cfg)
	case "repair":
		return runRepair(cfg)
	case "archive":
		return runArchive(args[1:], cfg)
	case "unarchive":
		return runUnarchive(args[1:], cfg)
	case "rename":
		return runRename(args[1:], cfg)
	case "backup":
		return runBackup(cfg)
	case "restore":
		return runRestore(args[1:], cfg)
	case "export":
		return runExport(args[1:], cfg)
	case "help", "--help", "-h":
		printHelp(cfg.Stdout)
		return nil
	default:
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func printHelp(w io.Writer) {
	fmt.Fprintln(w, "StreakForge - local habit and streak tracker")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  streakforge help")
	fmt.Fprintln(w, "  streakforge path")
	fmt.Fprintln(w, "  streakforge add <name> [description]")
	fmt.Fprintln(w, "  streakforge list [--all]")
	fmt.Fprintln(w, "  streakforge today [done|todo]")
	fmt.Fprintln(w, "  streakforge next")
	fmt.Fprintln(w, "  streakforge missed")
	fmt.Fprintln(w, "  streakforge lagging [days]")
	fmt.Fprintln(w, "  streakforge done <habit-id>")
	fmt.Fprintln(w, "  streakforge log <habit-id> <yyyy-mm-dd>")
	fmt.Fprintln(w, "  streakforge undo <habit-id>")
	fmt.Fprintln(w, "  streakforge show <habit-id>")
	fmt.Fprintln(w, "  streakforge search [--all] <query>")
	fmt.Fprintln(w, "  streakforge streaks")
	fmt.Fprintln(w, "  streakforge stats [json]")
	fmt.Fprintln(w, "  streakforge pulse")
	fmt.Fprintln(w, "  streakforge calendar [days]")
	fmt.Fprintln(w, "  streakforge doctor")
	fmt.Fprintln(w, "  streakforge repair")
	fmt.Fprintln(w, "  streakforge archive <habit-id>")
	fmt.Fprintln(w, "  streakforge unarchive <habit-id>")
	fmt.Fprintln(w, "  streakforge rename <habit-id> <name> [description]")
	fmt.Fprintln(w, "  streakforge backup")
	fmt.Fprintln(w, "  streakforge restore <backup-path>")
	fmt.Fprintln(w, "  streakforge export [csv|json|completions]")
}

func runPath(cfg Config) error {
	fmt.Fprintln(cfg.Stdout, habit.DefaultPath(cfg.Home))
	return nil
}

func runNext(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	actions := habit.Next(store, time.Now())
	if len(actions) == 0 {
		fmt.Fprintln(cfg.Stdout, "no pending habits")
		return nil
	}
	for _, action := range actions {
		fmt.Fprintf(cfg.Stdout, "%s  %s  %s\n", action.ID, streakText(action.CurrentStreak), action.Name)
	}
	return nil
}

func runMissed(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	missed := habit.Missed(store, time.Now())
	if len(missed) == 0 {
		fmt.Fprintln(cfg.Stdout, "no missed habits yesterday")
		return nil
	}
	for _, h := range missed {
		last := h.LastCompleted
		if last == "" {
			last = "never"
		}
		fmt.Fprintf(cfg.Stdout, "%s  last: %s  streak: %s  %s\n", h.ID, last, streakText(h.CurrentStreak), h.Name)
	}
	return nil
}

func runLagging(args []string, cfg Config) error {
	days := 3
	if len(args) > 1 {
		return fmt.Errorf("%w: lagging accepts an optional day count", ErrUsage)
	}
	if len(args) == 1 {
		parsed, err := strconv.Atoi(args[0])
		if err != nil || parsed < 0 {
			return fmt.Errorf("%w: lagging days must be zero or greater", ErrUsage)
		}
		days = parsed
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	lagging := habit.Lagging(store, time.Now(), days)
	if len(lagging) == 0 {
		fmt.Fprintf(cfg.Stdout, "no habits quiet for %s or more\n", streakText(days))
		return nil
	}
	for _, h := range lagging {
		last := h.LastCompleted
		if last == "" {
			last = "never"
		}
		fmt.Fprintf(cfg.Stdout, "%s  quiet: %s  last: %s  %s\n", h.ID, streakText(h.DaysQuiet), last, h.Name)
	}
	return nil
}

func streakText(days int) string {
	if days == 1 {
		return "1 day"
	}
	return fmt.Sprintf("%d days", days)
}

func runRestore(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: restore requires a backup path", ErrUsage)
	}
	safetyPath, err := habit.Restore(habit.DefaultPath(cfg.Home), args[0], time.Now())
	if err != nil {
		return err
	}
	if safetyPath == "" {
		fmt.Fprintln(cfg.Stdout, "restored store")
		return nil
	}
	fmt.Fprintf(cfg.Stdout, "restored store; previous data backed up to %s\n", safetyPath)
	return nil
}

func runToday(args []string, cfg Config) error {
	filter := "all"
	if len(args) > 1 {
		return fmt.Errorf("%w: today accepts done or todo", ErrUsage)
	}
	if len(args) == 1 {
		filter = strings.ToLower(args[0])
		if filter != "done" && filter != "todo" {
			return fmt.Errorf("%w: today accepts done or todo", ErrUsage)
		}
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	statuses := habit.Today(store, time.Now())
	filtered := make([]habit.TodayStatus, 0, len(statuses))
	for _, status := range statuses {
		if filter == "done" && !status.Completed {
			continue
		}
		if filter == "todo" && status.Completed {
			continue
		}
		filtered = append(filtered, status)
	}
	if len(filtered) == 0 {
		fmt.Fprintln(cfg.Stdout, "no active habits")
		return nil
	}
	for _, status := range filtered {
		mark := "todo"
		if status.Completed {
			mark = "done"
		}
		fmt.Fprintf(cfg.Stdout, "%s  %s  %s\n", mark, status.ID, status.Name)
	}
	return nil
}

func runBackup(cfg Config) error {
	backupPath, err := habit.Backup(habit.DefaultPath(cfg.Home), time.Now())
	if err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "backup written to %s\n", backupPath)
	return nil
}

func runRename(args []string, cfg Config) error {
	if len(args) < 2 {
		return fmt.Errorf("%w: rename requires a habit id and name", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	var description string
	if len(args) > 2 {
		description = strings.Join(args[2:], " ")
	}
	store, renamed, err := habit.Rename(store, args[0], args[1], description)
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "renamed %s %q\n", renamed.ID, renamed.Name)
	return nil
}

func runPulse(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	for _, day := range habit.WeeklyPulse(store, time.Now()) {
		fmt.Fprintf(cfg.Stdout, "%s  %s %d\n", day.Date, pulseBar(day.Completion), day.Completion)
	}
	return nil
}

func runCalendar(args []string, cfg Config) error {
	days := 14
	if len(args) > 1 {
		return fmt.Errorf("%w: calendar accepts an optional day count", ErrUsage)
	}
	if len(args) == 1 {
		parsed, err := strconv.Atoi(args[0])
		if err != nil || parsed <= 0 {
			return fmt.Errorf("%w: calendar days must be a positive number", ErrUsage)
		}
		days = parsed
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	for _, day := range habit.CompletionCalendar(store, time.Now(), days) {
		percent := 0
		if day.Total > 0 {
			percent = int(day.Rate*100 + 0.5)
		}
		fmt.Fprintf(cfg.Stdout, "%s  %d/%d  %3d%%  %s\n", day.Date, day.Complete, day.Total, percent, calendarBar(day.Rate))
	}
	return nil
}

func runDoctor(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	issues := habit.Doctor(store, time.Now())
	if len(issues) == 0 {
		fmt.Fprintln(cfg.Stdout, "ok  store looks healthy")
		return nil
	}
	for _, issue := range issues {
		fmt.Fprintf(cfg.Stdout, "%s  %s\n", issue.Severity, issue.Message)
	}
	return nil
}

func runRepair(cfg Config) error {
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, result := habit.Repair(store)
	if !result.Changed {
		fmt.Fprintln(cfg.Stdout, "store already clean")
		return nil
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "removed duplicate completions: %d\n", result.DuplicateRemoved)
	fmt.Fprintf(cfg.Stdout, "sorted habit histories: %d\n", result.CompletionsSorted)
	return nil
}

func calendarBar(rate float64) string {
	switch {
	case rate >= 1:
		return "#####"
	case rate >= 0.75:
		return "####-"
	case rate >= 0.5:
		return "###--"
	case rate > 0:
		return "#----"
	default:
		return "-----"
	}
}

func pulseBar(count int) string {
	if count <= 0 {
		return "□"
	}
	if count == 1 {
		return "■"
	}
	if count == 2 {
		return "■■"
	}
	return "■■■"
}

func runExport(args []string, cfg Config) error {
	format := "csv"
	if len(args) > 1 {
		return fmt.Errorf("%w: export accepts csv, json, or completions", ErrUsage)
	}
	if len(args) == 1 {
		format = strings.ToLower(args[0])
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	switch format {
	case "csv":
		return habit.WriteCSV(cfg.Stdout, store)
	case "json":
		return habit.WriteJSON(cfg.Stdout, store)
	case "completions":
		return habit.WriteCompletionCSV(cfg.Stdout, store)
	default:
		return fmt.Errorf("%w: export accepts csv, json, or completions", ErrUsage)
	}
}

func runArchive(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: archive requires a habit id", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, archived, err := habit.Archive(store, args[0])
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "archived %q\n", archived.Name)
	return nil
}

func runUnarchive(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: unarchive requires a habit id", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, restored, err := habit.Unarchive(store, args[0])
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "unarchived %q\n", restored.Name)
	return nil
}

func runStats(args []string, cfg Config) error {
	format := "text"
	if len(args) > 1 {
		return fmt.Errorf("%w: stats accepts json", ErrUsage)
	}
	if len(args) == 1 {
		format = strings.ToLower(args[0])
		if format != "json" {
			return fmt.Errorf("%w: stats accepts json", ErrUsage)
		}
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	summary := habit.SummarizeAt(store, time.Now())
	if format == "json" {
		encoder := json.NewEncoder(cfg.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(summary)
	}
	fmt.Fprintf(cfg.Stdout, "active habits: %d\n", summary.ActiveHabits)
	fmt.Fprintf(cfg.Stdout, "archived habits: %d\n", summary.ArchivedHabits)
	fmt.Fprintf(cfg.Stdout, "total completions: %d\n", summary.TotalCompletions)
	fmt.Fprintf(cfg.Stdout, "longest completions: %d\n", summary.LongestCompletion)
	fmt.Fprintf(cfg.Stdout, "best current streak: %d\n", summary.BestCurrentStreak)
	fmt.Fprintf(cfg.Stdout, "completion rate: %.2f per habit\n", summary.CompletionRate)
	return nil
}

func runStreaks(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	reports := habit.StreakReports(store, time.Now())
	if len(reports) == 0 {
		fmt.Fprintln(cfg.Stdout, "no active habits")
		return nil
	}
	for _, report := range reports {
		last := report.LastCompleted
		if last == "" {
			last = "never"
		}
		fmt.Fprintf(cfg.Stdout, "%s  %s  last: %s  %s\n", report.ID, streakText(report.CurrentStreak), last, report.Name)
	}
	return nil
}

func runDone(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: done requires a habit id", ErrUsage)
	}
	return completeForDay(args[0], time.Now(), cfg)
}

func runLog(args []string, cfg Config) error {
	if len(args) != 2 {
		return fmt.Errorf("%w: log requires a habit id and date", ErrUsage)
	}
	day, err := time.Parse(time.DateOnly, args[1])
	if err != nil {
		return fmt.Errorf("%w: date must use yyyy-mm-dd", ErrUsage)
	}
	return completeForDay(args[0], day, cfg)
}

func completeForDay(id string, day time.Time, cfg Config) error {
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, completed, err := habit.Complete(store, id, day)
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	if day.Format(time.DateOnly) == time.Now().Format(time.DateOnly) {
		fmt.Fprintf(cfg.Stdout, "completed %q for today\n", completed.Name)
		return nil
	}
	fmt.Fprintf(cfg.Stdout, "completed %q for %s\n", completed.Name, day.Format(time.DateOnly))
	return nil
}

func runUndo(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: undo requires a habit id", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, h, removed, err := habit.Uncomplete(store, args[0], time.Now())
	if err != nil {
		return err
	}
	if !removed {
		fmt.Fprintf(cfg.Stdout, "%q was not completed today\n", h.Name)
		return nil
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "removed today's completion for %q\n", h.Name)
	return nil
}

func runShow(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: show requires a habit id", ErrUsage)
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	h, err := habit.Find(store, args[0])
	if err != nil {
		return err
	}
	status := "active"
	if h.Archived {
		status = "archived"
	}
	description := h.Description
	if description == "" {
		description = "-"
	}
	created := "unknown"
	if !h.CreatedAt.IsZero() {
		created = h.CreatedAt.Format(time.DateOnly)
	}
	fmt.Fprintf(cfg.Stdout, "id: %s\n", h.ID)
	fmt.Fprintf(cfg.Stdout, "name: %s\n", h.Name)
	fmt.Fprintf(cfg.Stdout, "description: %s\n", description)
	fmt.Fprintf(cfg.Stdout, "status: %s\n", status)
	fmt.Fprintf(cfg.Stdout, "created: %s\n", created)
	fmt.Fprintf(cfg.Stdout, "current streak: %d\n", habit.CurrentStreak(h, time.Now()))
	fmt.Fprintf(cfg.Stdout, "completions: %d\n", len(h.Completions))
	recent := recentCompletions(h.Completions, 5)
	if len(recent) > 0 {
		fmt.Fprintf(cfg.Stdout, "recent: %s\n", strings.Join(recent, ", "))
	}
	return nil
}

func recentCompletions(completions []string, limit int) []string {
	recent := append([]string(nil), completions...)
	sort.Sort(sort.Reverse(sort.StringSlice(recent)))
	if len(recent) > limit {
		return recent[:limit]
	}
	return recent
}

func runSearch(args []string, cfg Config) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: search requires a query", ErrUsage)
	}
	includeArchived := false
	if args[0] == "--all" {
		includeArchived = true
		args = args[1:]
	}
	if len(args) == 0 {
		return fmt.Errorf("%w: search requires a query", ErrUsage)
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	matches := habit.SearchWithOptions(store, strings.Join(args, " "), includeArchived)
	if len(matches) == 0 {
		fmt.Fprintln(cfg.Stdout, "no matching habits")
		return nil
	}
	for _, h := range matches {
		status := ""
		if includeArchived {
			status = "active"
			if h.Archived {
				status = "archived"
			}
		}
		if h.Description == "" {
			if includeArchived {
				fmt.Fprintf(cfg.Stdout, "%s  %s  %s\n", h.ID, status, h.Name)
				continue
			}
			fmt.Fprintf(cfg.Stdout, "%s  %s\n", h.ID, h.Name)
			continue
		}
		if includeArchived {
			fmt.Fprintf(cfg.Stdout, "%s  %s  %s  %s\n", h.ID, status, h.Name, h.Description)
			continue
		}
		fmt.Fprintf(cfg.Stdout, "%s  %s  %s\n", h.ID, h.Name, h.Description)
	}
	return nil
}

func runList(args []string, cfg Config) error {
	includeArchived := false
	if len(args) > 1 {
		return fmt.Errorf("%w: list accepts only --all", ErrUsage)
	}
	if len(args) == 1 {
		if args[0] != "--all" {
			return fmt.Errorf("%w: list accepts only --all", ErrUsage)
		}
		includeArchived = true
	}
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	habits := habit.Active(store)
	if includeArchived {
		habits = store.Habits
	}
	if len(habits) == 0 {
		if includeArchived {
			fmt.Fprintln(cfg.Stdout, "no habits")
		} else {
			fmt.Fprintln(cfg.Stdout, "no active habits")
		}
		return nil
	}
	for _, h := range habits {
		if includeArchived && h.Archived {
			fmt.Fprintf(cfg.Stdout, "%s  archived  %s\n", h.ID, h.Name)
			continue
		}
		if includeArchived {
			fmt.Fprintf(cfg.Stdout, "%s  active    %s\n", h.ID, h.Name)
			continue
		}
		fmt.Fprintf(cfg.Stdout, "%s  %s\n", h.ID, h.Name)
	}
	return nil
}

var ErrUsage = errors.New("invalid command usage")

func runAdd(args []string, cfg Config) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: add requires a habit name", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	var description string
	if len(args) > 1 {
		description = strings.Join(args[1:], " ")
	}
	store, created, err := habit.Add(store, args[0], description, time.Now())
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "added %s %q\n", created.ID, created.Name)
	return nil
}
