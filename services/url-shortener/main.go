package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/mralanlee/gh-workflow-demo/services/url-shortener/controllers"
	db "github.com/mralanlee/gh-workflow-demo/services/url-shortener/db"
)

var port = os.Getenv("PORT")

func main() {
	client := &db.Client{
		Sql: db.Init(),
	}

	http.Handle("/api/shorten", handlers.Shorten(client))
	http.Handle("/api/stats", handlers.Stats(client))
	http.Handle("/", handlers.Redirect(client))
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on:", port)
}
