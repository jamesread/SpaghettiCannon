package httpservers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/jamesread/SpaghettiCannon/internal/buildinfo"
	"github.com/jamesread/SpaghettiCannon/internal/config"
	"github.com/jamesread/golure/pkg/dirs"
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
	ConfigError            string `json:"configError,omitempty"`
	ConfigFile             string `json:"configFile,omitempty"`
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
		ShowFooter:      cfg.ShowFooter,
		ShowNewVersions: cfg.ShowNewVersions,
		CurrentVersion:  buildinfo.Version,
		PageTitle:       cfg.SiteTitle,
		ConfigError:     cfg.LoadError,
		ConfigFile:      cfg.ConfigFile,
	}

	return ret
}

func GetNewWebUIHandler(dir string) http.Handler {
	root := findWebuiDir(dir)
	fs := http.FileServer(http.Dir(root))
	indexPath := filepath.Join(root, "index.html")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if serveIndexFallback(root, indexPath, r.URL.Path) {
			http.ServeFile(w, r, indexPath)
			return
		}
		fs.ServeHTTP(w, r)
	})
}

func serveIndexFallback(root, indexPath, urlPath string) bool {
	cleaned := filepath.Clean(urlPath)
	if skipIndexFallback(cleaned) {
		return false
	}
	if regularFileExists(filepath.Join(root, cleaned)) {
		return false
	}
	_, err := os.Stat(indexPath)
	return err == nil
}

func skipIndexFallback(cleaned string) bool {
	if cleaned == "/" || cleaned == "." {
		return true
	}
	return hasFileExtension(cleaned)
}

func regularFileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func hasFileExtension(path string) bool {
	base := filepath.Base(path)
	return strings.Contains(base, ".")
}
