package handlers

import {
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/maxrmax/go-api-projects/project1/internal/middleware"
}

// starts with capital H, tells the compiler the function can be imported in other packages
// lower case is a private function -> main package wouldn't be able to import it
func Handler(r *chi.Mux) {
	// Global middleware

	// strips the trailing / from urls -> http://localhost:8080/account/coins/ < strips the last /
	r.Use(chimiddle.StripSlashes)

	// 
	r.Route("/account", func(router chi.Router) {
		// Middleware for /account route
		// our "bouncer" to check if the user is authorized to acces the data
		// if the auth fails, an error is returned and the rest is the function is not executed
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
