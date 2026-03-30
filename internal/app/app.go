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
	return RunWithConfig(args, Config{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Home:   os.Getenv("HOME"),
	})
}

func RunWithConfig(args []string, cfg Config) error {
	if len(args) == 0 {
		printHelp(cfg.Stdout)
		return nil
	}

	switch args[0] {
	case "add":
		return runAdd(args[1:], cfg)
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
