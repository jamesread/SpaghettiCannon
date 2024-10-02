package main

import (
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/httpservers"
	log "github.com/sirupsen/logrus"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	log.Infof("SpaghettiCannon")

	cfg := config.DefaultConfig()

	config.LoadConfig(cfg)

	httpservers.StartServers(cfg)
}
