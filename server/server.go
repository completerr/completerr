package main

import (
	_ "completerr/config"
	"completerr/db"
	"completerr/services"
	"completerr/tasks"
	"completerr/web"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	tasks.StartScheduler()
	go services.ImportRadarrMovies()
	go services.ImportSonarrEpisodes()
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(cors)
	err := app.Serve()
	log.Println("Error", err)
}
