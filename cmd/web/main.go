package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Berhtz/go-app/pkg/config"
	"github.com/Berhtz/go-app/pkg/handlers"
	"github.com/Berhtz/go-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache in main.go")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
