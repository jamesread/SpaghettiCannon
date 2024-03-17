package httpservers

import (
	//	log "github.com/sirupsen/logrus"
	"github.com/jamesread/SpaghettiCannon/internal/config"
)

func StartServers(cfg *config.Config) {
	go startWebUIServer(cfg)

	go StartSingleHTTPFrontend(cfg)

	startRestAPIServer(cfg)
}
