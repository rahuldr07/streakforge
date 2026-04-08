package app

import (
	"bytes"
	"strings"
	"testing"
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
