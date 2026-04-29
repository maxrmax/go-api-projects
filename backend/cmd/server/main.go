package main

import (
	// standard library logging (writes to stderr by default)
	"log"
	// HTTP server primitives: Handler interface, ServeMux, server loop
	"net/http"

	// HTTP layer: handlers, router, middleware
	"backend/internal/api"
	// business logic layer
	"backend/internal/service"
	// data layer (in-memory mock database)
	"backend/internal/store"
)

func main() {

	// Internally:
	// - two maps are created: customers and users
	// - both maps store pointers to structs (*Customer, *User)
	// - initially one customer and one user are inserted
	// Result:
	//	store ---> MemoryStore{ customers: map, users: map }
	store := store.NewMemoryStore()

	// --- SERVICE LAYER INITIALIZATION ---
	// Services are thin wrappers around the store.
	// copies pointer.
	//	authService.store		---> store
	//	customerService.store	---> store
	authService := service.NewAuthService(store)
	customerService := service.NewCustomerService(store)

	// --- HANDLER LAYER INITIALIZATION ---
	// Handlers are HTTP-facing components.
	// Chain now becomes:
	//	authHandler		---> authService		---> store
	//	customerHandler	---> customerService	---> store
	authHandler := api.NewAuthHandler(authService)
	customerHandler := api.NewCustomerHandler(customerService)

	// --- ROUTER CONSTRUCTION ---
	// api.NewRouter returns *http.ServeMux.
	// ServeMux is a request multiplexer:
	// - matches incoming HTTP request paths
	// - dispatches them to registered handlers
	// Second argument:
	//	api.AuthMiddleware(customerHandler)
	router := api.NewRouter(
		authHandler,
		api.AuthMiddleware(customerHandler),
	)

	log.Println("server running on 0.0.0.0:8000")

	// server runs indefinitely unless error occurs
	// return value (error) is ignored in this code
	http.ListenAndServe(":8000", router)
}
