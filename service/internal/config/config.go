package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	log "github.com/sirupsen/logrus"
)

// RequiredMigration is the sql-migrate id this binary expects to be applied.
const RequiredMigration = "1.cvars.sql"

type Config struct {
	ConfigVersion     int             `koanf:"configVersion"`
	ListenAddress     string          `koanf:"listenAddress"`
	WebUIDir          string          `koanf:"webUIDir"`
	LogLevel          string          `koanf:"logLevel"`
	SiteTitle         string          `koanf:"siteTitle"`
	ShowFooter        bool            `koanf:"showFooter"`
	ShowNewVersions   bool            `koanf:"showNewVersions"`
	ShowVersionNumber bool            `koanf:"showVersionNumber"`
	Database          *DatabaseConfig `koanf:"database"`

	// Set when config.yaml could not be loaded; not read from YAML.
	ConfigFile string `koanf:"-" yaml:"-"`
	LoadError  string `koanf:"-" yaml:"-"`
}

type DatabaseConfig struct {
	Host string `koanf:"host"`
	Port int    `koanf:"port"`
	User string `koanf:"user"`
	Pass string `koanf:"pass"`
	Name string `koanf:"name"`
}

func DefaultConfig() *Config {
	return &Config{
		ConfigVersion:     1,
		WebUIDir:          "./webui/",
		LogLevel:          "info",
		ListenAddress:     "0.0.0.0:8080",
		SiteTitle:         "SpaghettiCannon",
		ShowFooter:        true,
		ShowNewVersions:   true,
		ShowVersionNumber: true,
		Database: &DatabaseConfig{
			Host: envOr("DB_HOST", "localhost"),
			Port: 3306,
			User: firstNonEmpty(os.Getenv("DB_USERNAME"), os.Getenv("DB_USER"), "user"),
			Pass: firstNonEmpty(os.Getenv("DB_PASSWORD"), os.Getenv("DB_PASS"), "password"),
			Name: firstNonEmpty(os.Getenv("DB_DATABASE"), os.Getenv("DB_NAME"), "spaghetticannon"),
		},
	}
}

func findConfigDir() string {
	directoriesToSearch := []string{
		"/config/",
		"./",
		"../",
	}

	for _, dir := range directoriesToSearch {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			log.WithFields(log.Fields{
				"dir": dir,
			}).Infof("Found the config directory")

			return dir
		}
	}

	log.Warnf("Did not find the config directory, you will probably crash loop.")

	return "./"
}

func LoadConfig(cfg *Config) error {
	cfgFile, err := filepath.Abs(filepath.Join(findConfigDir(), "config.yaml"))
	if err != nil {
		return err
	}

	cfg.ConfigFile = cfgFile

	log.WithFields(log.Fields{
		"cfgFile": cfgFile,
	}).Infof("Loading config file")

	return loadConfigFromFile(cfg, cfgFile)
}

func loadConfigFromFile(cfg *Config, cfgFile string) error {
	if _, err := os.Stat(cfgFile); err != nil {
		if os.IsNotExist(err) {
			cfg.LoadError = fmt.Sprintf("config.yaml cannot be found at %s", cfgFile)
		} else {
			cfg.LoadError = fmt.Sprintf("config.yaml could not be read: %v", err)
		}

		return err
	}

	k := koanf.New(".")
	if err := k.Load(file.Provider(cfgFile), yaml.Parser()); err != nil {
		cfg.LoadError = fmt.Sprintf("config.yaml could not be read: %v", err)
		return err
	}
	if err := k.Unmarshal("", cfg); err != nil {
		cfg.LoadError = fmt.Sprintf("config.yaml is invalid: %v", err)
		return err
	}

	// Older samples used lowercase loglevel.
	if cfg.LogLevel == "" {
		cfg.LogLevel = k.String("loglevel")
	}

	applyEnvOverrides(cfg)
	applyConfigFallbacks(cfg)

	cfg.LoadError = ""

	return nil
}

func applyEnvOverrides(cfg *Config) {
	if cfg.Database == nil {
		cfg.Database = &DatabaseConfig{}
	}
	setIfEnv(&cfg.Database.Host, "DB_HOST")
	setIfNonEmpty(&cfg.Database.Name, firstNonEmpty(os.Getenv("DB_DATABASE"), os.Getenv("DB_NAME")))
	setIfNonEmpty(&cfg.Database.User, firstNonEmpty(os.Getenv("DB_USERNAME"), os.Getenv("DB_USER")))
	setIfNonEmpty(&cfg.Database.Pass, firstNonEmpty(os.Getenv("DB_PASSWORD"), os.Getenv("DB_PASS")))
	setIfEnv(&cfg.SiteTitle, "SITE_TITLE")
}

func applyConfigFallbacks(cfg *Config) {
	applyStringFallbacks(cfg)
	if cfg.Database == nil {
		cfg.Database = DefaultConfig().Database
	}
}

func applyStringFallbacks(cfg *Config) {
	cfg.ListenAddress = defaultString(cfg.ListenAddress, "0.0.0.0:8080")
	cfg.SiteTitle = defaultString(cfg.SiteTitle, "SpaghettiCannon")
	cfg.LogLevel = defaultString(cfg.LogLevel, "info")
	cfg.WebUIDir = defaultString(cfg.WebUIDir, "./webui/")
}

func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}
	return value
}

// ListenAddr returns the bind address: $PORT if set, otherwise cfg.ListenAddress.
// PORT may be a bare port ("8080") or a full address (":8080" / "0.0.0.0:8080").
func ListenAddr(cfg *Config) string {
	if port := strings.TrimSpace(os.Getenv("PORT")); port != "" {
		return normalizePort(port)
	}
	if cfg != nil && cfg.ListenAddress != "" {
		return cfg.ListenAddress
	}
	return "0.0.0.0:8080"
}

func normalizePort(port string) string {
	if strings.Contains(port, ":") {
		return port
	}
	return "0.0.0.0:" + port
}

func setIfEnv(dst *string, key string) {
	if v := os.Getenv(key); v != "" {
		*dst = v
	}
}

func setIfNonEmpty(dst *string, v string) {
	if v != "" {
		*dst = v
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
