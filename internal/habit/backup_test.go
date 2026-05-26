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
