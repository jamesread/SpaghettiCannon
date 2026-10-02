package config

import (
	"path/filepath"
	"testing"
)

func TestLoadConfigFromFile_missingFile(t *testing.T) {
	cfg := DefaultConfig()
	cfgFile := filepath.Join(t.TempDir(), "config.yaml")

	err := loadConfigFromFile(cfg, cfgFile)
	if err == nil {
		t.Fatal("expected error for missing config file")
	}

	if cfg.LoadError == "" {
		t.Fatal("expected LoadError to be set")
	}

	if cfg.ListenAddress != "0.0.0.0:8080" {
		t.Fatalf("expected default listen address, got %q", cfg.ListenAddress)
	}
}
