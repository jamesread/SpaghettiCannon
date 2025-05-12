package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
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

func LoadConfig(cfg *Config) {
	f, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
		return
	}

	err = yaml.UnmarshalStrict(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
