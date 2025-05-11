package httpservers

import (
	"net/http"

	"github.com/jamesread/SpaghettiCannon/internal/api"

	"github.com/rs/cors"

	connectcors "connectrpc.com/cors"

	"github.com/jamesread/SpaghettiCannon/gen/SpaghettiCannon/clientapi/v1/spagv1connect"
	config "github.com/jamesread/SpaghettiCannon/internal/config"
)

func withCors(h http.Handler) http.Handler {
	mw := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	return mw.Handler(h)
}

func GetNewApiHandler(cfg *config.Config) (string, http.Handler) {
	apiServer := api.NewServer(cfg)

	path, handler := spagv1connect.NewSpaghettiCannonApiServiceHandler(apiServer)

	return path, withCors(handler)
}
