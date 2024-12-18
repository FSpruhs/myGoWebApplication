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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
