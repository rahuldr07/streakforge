package app

import (
	"errors"
	"fmt"
	"io"
	"os"
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
	case "add":
		return runAdd(args[1:], cfg)
	case "list":
		return runList(cfg)
	case "done":
		return runDone(args[1:], cfg)
	case "streaks":
		return runStreaks(cfg)
	case "stats":
		return runStats(cfg)
	case "archive":
		return runArchive(args[1:], cfg)
	case "export":
		return runExport(cfg)
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
	fmt.Fprintln(w, "  streakforge add <name> [description]")
	fmt.Fprintln(w, "  streakforge list")
	fmt.Fprintln(w, "  streakforge done <habit-id>")
	fmt.Fprintln(w, "  streakforge streaks")
	fmt.Fprintln(w, "  streakforge stats")
	fmt.Fprintln(w, "  streakforge archive <habit-id>")
	fmt.Fprintln(w, "  streakforge export")
}

func runExport(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	return habit.WriteCSV(cfg.Stdout, store)
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

func runStats(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	summary := habit.Summarize(store)
	fmt.Fprintf(cfg.Stdout, "active habits: %d\n", summary.ActiveHabits)
	fmt.Fprintf(cfg.Stdout, "archived habits: %d\n", summary.ArchivedHabits)
	fmt.Fprintf(cfg.Stdout, "total completions: %d\n", summary.TotalCompletions)
	fmt.Fprintf(cfg.Stdout, "longest completions: %d\n", summary.LongestCompletion)
	return nil
}

func runStreaks(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	for _, h := range habit.Active(store) {
		fmt.Fprintf(cfg.Stdout, "%s  %d days  %s\n", h.ID, habit.CurrentStreak(h, time.Now()), h.Name)
	}
	return nil
}

func runDone(args []string, cfg Config) error {
	if len(args) != 1 {
		return fmt.Errorf("%w: done requires a habit id", ErrUsage)
	}
	path := habit.DefaultPath(cfg.Home)
	store, err := habit.Load(path)
	if err != nil {
		return err
	}
	store, completed, err := habit.Complete(store, args[0], time.Now())
	if err != nil {
		return err
	}
	if err := habit.Save(path, store); err != nil {
		return err
	}
	fmt.Fprintf(cfg.Stdout, "completed %q for today\n", completed.Name)
	return nil
}

func runList(cfg Config) error {
	store, err := habit.Load(habit.DefaultPath(cfg.Home))
	if err != nil {
		return err
	}
	habits := habit.Active(store)
	if len(habits) == 0 {
		fmt.Fprintln(cfg.Stdout, "no active habits")
		return nil
	}
	for _, h := range habits {
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
