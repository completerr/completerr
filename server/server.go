package main

import (
	_ "completerr/config"
	"completerr/db"
	"completerr/scheduler"
	"completerr/services"
	"completerr/web"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// CORS is enabled only in prod profile
	db.InitDB()
	scheduler.StartScheduler()
	go services.ImportRadarrMovies()
	go services.ImportSonarrEpisodes()
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(cors)
	err := app.Serve()
	log.Println("Error", err)
}
