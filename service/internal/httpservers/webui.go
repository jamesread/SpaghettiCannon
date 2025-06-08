package httpservers

import (
	"github.com/jamesread/golure/pkg/dirs"
	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"net/http"
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
		"../frontend/dist/",
		"../frontend/",
		"/frontend/",
		dir,
		"/usr/share/SpaghettiCannon/frontend/",
		"/var/www/SpaghettiCannon/",
		"/etc/SpaghettiCannon/frontend/",
	}

	found, _ := dirs.GetFirstExistingDirectory("webui", directoriesToSearch)

	return found
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
