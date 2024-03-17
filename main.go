package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/jamesread/SpaghettiCannon/internal/httpservers"
	"github.com/jamesread/SpaghettiCannon/internal/config"
)

var (
	version string
	commit string
	date string
)

func main() {
	log.Infof("SpaghettiCannon")

	cfg := config.DefaultConfig()

	httpservers.StartServers(cfg)
}
