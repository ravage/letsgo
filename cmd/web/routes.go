package main

import (
	"letsgo/cmd/web/config"
	"letsgo/cmd/web/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes creates routes
func Routes(app *config.Application, handlers *handlers.HandlerContext) *mux.Router {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", fileServer))

	router.Handle("/", handlers.HomeIndex())

	return router
}
