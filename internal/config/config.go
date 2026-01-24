package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	HistoryConfig HistoryConfig
}

type HistoryConfig struct {
	File *Value[string]
	Size *Value[int]
}

func New() *Config {
	return &Config{
		HistoryConfig: HistoryConfig{
			File: NewValue("~/.hermit_history", ValidPath),
			Size: NewValue(1000, InRange(1, 10000)),
		},
	}
}

func GetRCPath() string {
	// TODO: make this configurable
	return expandPath("~/.hermitrc")
}

func CreateDefaultRC(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("creating rc file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(defaultRCContent)
	if err != nil {
		return fmt.Errorf("writing rc file: %w", err)
	}

	return nil
}

func expandPath(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	return filepath.Join(home, path[2:])
}

const defaultRCContent = `
echo "Welcome to Hermit Shell!"
`
