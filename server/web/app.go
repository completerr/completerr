package web

import (
	"completerr/controllers"
	"completerr/utils"
	"net/http"
)

var logger = utils.GetLogger()

type App struct {
	handlers map[string]http.HandlerFunc
}

func NewApp(cors bool) App {
	app := App{
		handlers: make(map[string]http.HandlerFunc),
	}
	//if !cors {
	//	techHandler = disableCors(techHandler)
	//	helloHandler = disableCors(helloHandler)
	//}
	app.handlers["/api/sonarr/import"] = controllers.SonarrLibraryImport
	app.handlers["/api/sonarr/search"] = controllers.SonarrMissingSearch
	app.handlers["/api/sonarr/history"] = controllers.SonarrSearchHistory
	app.handlers["/api/radarr/import"] = controllers.RadarrLibraryImport
	app.handlers["/api/radarr/search"] = controllers.RadarrMissingSearch
	app.handlers["/api/radarr/history"] = controllers.RadarrSearchHistory
	app.handlers["/api/tasks/info"] = controllers.TaskInfo
	app.handlers["/api/tasks/history"] = controllers.TaskHistory
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
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
