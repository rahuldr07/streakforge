package app

import (
	"errors"
	"fmt"
	"io"
	"os"
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
}

var ErrUsage = errors.New("invalid command usage")

