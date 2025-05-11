package main

import (
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/httpservers"
	log "github.com/sirupsen/logrus"
)


func main() {
	log.WithFields(log.Fields {
		"version": buildinfo.Version,
	}).Infof("SpaghettiCannon")

	cfg := config.DefaultConfig()

	config.LoadConfig(cfg)

	httpservers.Start(cfg)
}
