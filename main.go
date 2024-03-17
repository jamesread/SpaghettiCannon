package main

import (
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/httpservers"
	"github.com/jamesread/SpaghettiCannon/internal/api"
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

	go api.Start(cfg)

	httpservers.StartServers(cfg)
}
