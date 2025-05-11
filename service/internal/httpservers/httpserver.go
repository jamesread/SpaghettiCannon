package httpservers

/*
This file implements a very simple, lightweight reverse proxy so that REST and
the webui can be accessed from a single endpoint.

This makes external reverse proxies (treafik, haproxy, etc) easier, CORS goes
away, and several other issues.
*/

import (
	"encoding/json"
	config "github.com/jamesread/SpaghettiCannon/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start(cfg *config.Config) {
	log.WithFields(log.Fields{
		"address": cfg.ListenAddress,
	}).Info("Starting single HTTP frontend")

	mux := http.NewServeMux()

	apipath, apihandler := GetNewApiHandler(cfg)

	log.Infof("API path: %q", apipath)

	mux.HandleFunc("/api"+apipath, func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("api req: %q", r.URL)

		http.StripPrefix("/api", apihandler).ServeHTTP(w, r)
	})

	mux.HandleFunc("/webUiSettings.json", func(w http.ResponseWriter, r *http.Request) {
		jsonRet, _ := json.Marshal(generateWebUISettings(cfg))

		_, err := w.Write([]byte(jsonRet))

		if err != nil {
			log.Warnf("Could not write webui settings: %v", err)
		}

	})

	webuiHandler := GetNewWebUIHandler(cfg.WebUIDir)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("ui  req: %q", r.URL)

		webuiHandler.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
