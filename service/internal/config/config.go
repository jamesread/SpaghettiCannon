package config

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ConfigVersion int
	ListenAddress string
	WebUIDir      string
	LogLevel      string
	ShowFooter    bool

	Database *DatabaseConfig

	// Set when config.yaml could not be loaded; not read from YAML.
	ConfigFile string `yaml:"-"`
	LoadError  string `yaml:"-"`
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Pass string
	Name string
}

func DefaultConfig() *Config {
	return &Config{
		WebUIDir:      "./webui/",
		LogLevel:      "info",
		ListenAddress: "0.0.0.0:8080",
		Database: &DatabaseConfig{
			Host: "localhost",
			Port: 3306,
			User: "user",
			Pass: "password",
			Name: "spaghetticannon",
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

	return "./" // Should not exist
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
	f, err := os.ReadFile(cfgFile)
	if err != nil {
		if os.IsNotExist(err) {
			cfg.LoadError = fmt.Sprintf("config.yaml cannot be found at %s", cfgFile)
		} else {
			cfg.LoadError = fmt.Sprintf("config.yaml could not be read: %v", err)
		}

		return err
	}

	err = yaml.UnmarshalStrict(f, cfg)
	if err != nil {
		cfg.LoadError = fmt.Sprintf("config.yaml is invalid: %v", err)

		return err
	}

	cfg.LoadError = ""

	return nil
}
