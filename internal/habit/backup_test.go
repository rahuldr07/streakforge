package habit

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestBackupCopiesStore(t *testing.T) {
	path := filepath.Join(t.TempDir(), "store.json")
	if err := os.WriteFile(path, []byte(`{"habits":[]}`), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}

	backupPath, err := Backup(path, time.Date(2026, 5, 26, 12, 30, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("Backup returned error: %v", err)
	}
	data, err := os.ReadFile(backupPath)
	if err != nil {
		t.Fatalf("read backup: %v", err)
	}
	if string(data) != `{"habits":[]}` {
		t.Fatalf("backup data = %s", string(data))
	}
}

func TestRestoreCopiesBackupAndKeepsSafetyBackup(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "store.json")
	backupPath := filepath.Join(dir, "store.backup.json")
	if err := os.WriteFile(path, []byte(`{"habits":[{"id":"old"}]}`), 0o600); err != nil {
		t.Fatalf("write store: %v", err)
	}
	if err := os.WriteFile(backupPath, []byte(`{"habits":[{"id":"new"}]}`), 0o600); err != nil {
		t.Fatalf("write backup: %v", err)
	}

	safetyPath, err := Restore(path, backupPath, time.Date(2026, 5, 27, 12, 30, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("Restore returned error: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read restored store: %v", err)
	}
	if string(data) != `{"habits":[{"id":"new"}]}` {
		t.Fatalf("restored data = %s", string(data))
	}
	if _, err := os.Stat(safetyPath); err != nil {
		t.Fatalf("missing safety backup: %v", err)
	}
}
