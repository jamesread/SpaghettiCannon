package config

import (
	"io/ioutil"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ListenAddressSingleHTTPFrontend string
	ListenAddressRest               string
	ListenAddressWebUI              string
	ExternalRestAddress             string
	WebUIDir                        string
	LogLevel                        string
	ShowFooter                      bool

	Database *DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Pass 			string
	Name string
}

func DefaultConfig() *Config {
	return &Config{
		WebUIDir:                        "./webui/",
		LogLevel:                        "info",
		ExternalRestAddress:             ".",
		ListenAddressSingleHTTPFrontend: "0.0.0.0:4337",
		ListenAddressRest:               "localhost:4339",
		ListenAddressWebUI:              "localhost:4340",
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
	f, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
		return
	}

	err = yaml.UnmarshalStrict(f, cfg)

	if err != nil {
		log.Fatal(err)
	}
}
