package main

import (
	"log"
	"net/http"
	"strings"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if id == "" || strings.Contains(id, "/") {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	switch r.Method {
	case http.MethodPut:
		w.Write([]byte("update customer " + id))
	case http.MethodDelete:
		w.Write([]byte("delete customer " + id))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func customersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("list customers"))
	case http.MethodPost:
		w.Write([]byte("create customer"))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Working website placeholder"))
}

func main() {
	// create a new request multiplexer (router)
	// this allocates a ServeMux that maps URL paths to -> handlers
	mux := http.NewServeMux()

	// register the function for a specific path
	// HandleFunc converts the function into an http.Handler internally
	// mux.Handle("/", http.HandlerFunc(rootHandler)) explicit form
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/customers", customersHandler)
	mux.HandleFunc("/customers/", queryHandler)

	addr := ":8000"
	log.Printf("Starting GO API service on %s", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Printf("server stopped: %v", err)
	}
}
