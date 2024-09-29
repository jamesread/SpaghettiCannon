package httpservers

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/jamesread/SpaghettiCannon/internal/api"

	//"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/jamesread/SpaghettiCannon/gen/proto/spagv1connect"
	config "github.com/jamesread/SpaghettiCannon/internal/config"
)

var (
	cfg *config.Config
)

func SetGlobalRestConfig(config *config.Config) {
	cfg = config
}

func startRestAPIServer(globalConfig *config.Config) error {
	cfg = globalConfig

	apiServer := api.NewServer(cfg)

	log.WithFields(log.Fields{
		"address": cfg.ListenAddressRest,
	}).Info("Starting REST API")

	mux := http.NewServeMux()
	path, handler := spagv1connect.NewSpaghettiCannonApiServiceHandler(
		apiServer,
	)
	mux.Handle(path, handler)

	return http.ListenAndServe(
		cfg.ListenAddressRest, 
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
