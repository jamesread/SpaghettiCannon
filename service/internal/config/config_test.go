package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestRequiredMigration(t *testing.T) {
	require.Equal(t, "1.cvars.sql", RequiredMigration)
}

func TestListenAddr(t *testing.T) {
	t.Setenv("PORT", "")
	cfg := &Config{ListenAddress: "0.0.0.0:9090"}
	require.Equal(t, "0.0.0.0:9090", ListenAddr(cfg))

	t.Setenv("PORT", "3000")
	require.Equal(t, "0.0.0.0:3000", ListenAddr(cfg))

	t.Setenv("PORT", "127.0.0.1:4000")
	require.Equal(t, "127.0.0.1:4000", ListenAddr(cfg))
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	require.Equal(t, "0.0.0.0:8080", cfg.ListenAddress)
	require.True(t, cfg.ShowFooter)
	require.NotNil(t, cfg.Database)
}
