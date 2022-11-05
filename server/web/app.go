package web

import (
	"completerr/controllers"
	"completerr/utils"
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:webapp
var webapp embed.FS
var logger = utils.GetLogger()

type App struct {
	handlers map[string]http.HandlerFunc
}

func NewApp(cors bool) App {
	webAppFS, err := fs.Sub(webapp, "webapp")
	if err != nil {
		logger.Fatal(err)
	}
	app := App{
		handlers: make(map[string]http.HandlerFunc),
	}
	app.handlers["/api/sonarr/import"] = controllers.SonarrLibraryImport
	app.handlers["/api/sonarr/search"] = controllers.SonarrMissingSearch
	app.handlers["/api/sonarr/history"] = controllers.SonarrSearchHistory
	app.handlers["/api/radarr/import"] = controllers.RadarrLibraryImport
	app.handlers["/api/radarr/search"] = controllers.RadarrMissingSearch
	app.handlers["/api/radarr/history"] = controllers.RadarrSearchHistory
	app.handlers["/api/tasks/info"] = controllers.TaskInfo
	app.handlers["/api/tasks/history"] = controllers.TaskHistory
	app.handlers["/"] = loggingHandler(http.FileServer(http.FS(webAppFS))).ServeHTTP
	return app
}
func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	logger.Info("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
