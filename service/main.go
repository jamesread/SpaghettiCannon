package main

import (
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/SpaghettiCannon/internal/httpservers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"version": buildinfo.Version,
	}).Infof("SpaghettiCannon")

	cfg := config.DefaultConfig()

	if err := config.LoadConfig(cfg); err != nil {
		log.WithError(err).Warn("Starting without a valid config file; web UI will show the configuration error")
	}

	if level, err := log.ParseLevel(cfg.LogLevel); err == nil {
		log.SetLevel(level)
	}

	httpservers.Start(cfg)
}
