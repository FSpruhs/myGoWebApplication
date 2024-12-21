package main

import (
	"log"
	"net/http"

	"github.com/fspruhs/myGoWebApplication/pkg/config"
	"github.com/fspruhs/myGoWebApplication/pkg/handlers"
	"github.com/fspruhs/myGoWebApplication/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	log.Println("Starting application on port", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln("error starting server: ", err)
	}
}
