package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	ConfigVersion int
	ListenAddress string
	WebUIDir      string
	LogLevel      string
	ShowFooter    bool

	Database *DatabaseConfig
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
		ListenAddress: "0.0.0.0:4337",
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
		"./",
		"../",
		"/config/",
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


func LoadConfig(cfg *Config) {
	cfgFile, _ := filepath.Abs(filepath.Join(findConfigDir(), "config.yaml"))

	log.WithFields(log.Fields{
		"cfgFile": cfgFile,
	}).Infof("Loading config file")

	f, err := os.ReadFile(cfgFile)

	if err != nil {
		log.Fatal(err)
		return
	}

	err = yaml.UnmarshalStrict(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
