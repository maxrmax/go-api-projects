package main

import (
	"fmt"
	"net/http"

	// package for webdev
	"github.com/go-chi/chi"
	// package from our own module from the /internal/handlers folder
	"github.com/maxrmax/go-api-projects/project1/internal/handlers"
	// logrus for error debugging aliased as log
	log "github.com/sirupsen/logrus"
	// install packages like this with "go mod tidy"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	err := http.ListenAndServe("localhost:8000", r)
	// ListenAndServe always returns a non-nil error.
	if err != nil {
		log.Error(err)
	}
}
