package habit

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func Backup(path string, now time.Time) (string, error) {
	if _, err := os.Stat(path); err != nil {
		return "", err
	}
	backupPath := fmt.Sprintf("%s.%s.bak", path, now.Format("20060102-150405"))
	if err := copyFile(path, backupPath); err != nil {
		return "", err
	}
	return backupPath, nil
}

func Restore(path string, backupPath string, now time.Time) (string, error) {
	safetyPath := ""
	if _, err := os.Stat(path); err == nil {
		var backupErr error
		safetyPath, backupErr = Backup(path, now)
		if backupErr != nil {
			return "", backupErr
		}
	} else if !os.IsNotExist(err) {
		return "", err
	}
	if err := copyFile(backupPath, path); err != nil {
		return "", err
	}
	return safetyPath, nil
}

func copyFile(source string, destination string) error {
	if err := os.MkdirAll(filepath.Dir(destination), 0o700); err != nil {
		return err
	}
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o600)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
