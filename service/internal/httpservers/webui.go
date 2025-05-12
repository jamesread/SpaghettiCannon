package httpservers

import (
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

type webUISettings struct {
	Rest                   string
	ThemeName              string
	ShowFooter             bool
	ShowNavigation         bool
	ShowNewVersions        bool
	CurrentVersion         string
	PageTitle              string
	SectionNavigationStyle string
}

func findWebuiDir(dir string) string {
	directoriesToSearch := []string{
		dir,
		"../frontend/dist/",
		"../frontend/",
		"/frontend/",
		"/usr/share/SpaghettiCannon/frontend/",
		"/var/www/SpaghettiCannon/",
		"/etc/SpaghettiCannon/frontend/",
	}

	for _, dir := range directoriesToSearch {
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			absdir, _ := filepath.Abs(dir)

			log.WithFields(log.Fields{
				"dir": dir,
				"absdir": absdir,
			}).Infof("Found the webui directory")

			return absdir
		}
	}

	log.Warnf("Did not find the webui directory, you will probably get 404 errors.")

	return "./webui" // Should not exist
}

func serveWebuiSettings(w http.ResponseWriter, r *http.Request) {
}

func generateWebUISettings(cfg *config.Config) *webUISettings {
	ret := &webUISettings{
		ShowFooter: cfg.ShowFooter,
		CurrentVersion: buildinfo.Version,
	}

	return ret
}

func GetNewWebUIHandler(dir string) http.Handler {
	return http.FileServer(http.Dir(findWebuiDir(dir)))
}
