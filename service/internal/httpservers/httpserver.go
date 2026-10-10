package httpservers

/*
This file implements a very simple, lightweight reverse proxy so that REST and
the webui can be accessed from a single endpoint.

This makes external reverse proxies (treafik, haproxy, etc) easier, CORS goes
away, and several other issues.
*/

import (
	"encoding/json"
	"net/http"
	"path"

	config "github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

func Start(cfg *config.Config) {
	addr := config.ListenAddr(cfg)
	log.WithFields(log.Fields{
		"address": addr,
	}).Info("Starting single HTTP frontend")

	mux := http.NewServeMux()

	apipath, apihandler := GetNewApiHandler(cfg)

	log.Infof("API path: %q", apipath)

	mux.Handle("/api/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn := path.Base(r.URL.Path)
		r.URL.Path = apipath + fn

		log.Debugf("api req: %q", r.URL.Path)

		apihandler.ServeHTTP(w, r)
	}))

	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/webUiSettings.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		jsonRet, _ := json.Marshal(generateWebUISettings(cfg))

		_, err := w.Write([]byte(jsonRet))

		if err != nil {
			log.Warnf("Could not write webui settings: %v", err)
		}

	})

	webuiHandler := GetNewWebUIHandler(cfg.WebUIDir)

	mux.Handle("/", webuiHandler)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
