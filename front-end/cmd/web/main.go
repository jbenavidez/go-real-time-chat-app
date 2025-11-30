package main

import (
	"fmt"
	"frontend/internal/config"
	"frontend/internal/handlers"
	"frontend/internal/render"
	"log"
	"net/http"
)

const port = ":3000"

var app config.AppConfig

func main() {

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))
	//
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err := srv.ListenAndServe()
	log.Fatal(err)

}
