package api

import (
	// JSON encoding of responses
	"encoding/json"
	// HTTP interfaces and helpers
	"net/http"
	// string prefix check
	"strings"

	// provides CustomerService
	"backend/internal/service"
	// data layer (in-memory mock database)
	"backend/internal/store"
)

// This is a pointer reference used to call service methods.
// No customer data is stored here.
type CustomerHandler struct {
	service *service.CustomerService
}

func NewCustomerHandler(s *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: s}
}

// Called for every request routed to this handler.
func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// --- DEFAULT HEADER SET ---
	// All responses from this handler will declare JSON content,
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/customers" {
		h.handleCollection(w, r)
		return
	}

	// --- ROUTE: PREFIX MATCH ---
	// Matches:
	// "/customers/1"
	// "/customers/abc"
	// "/customers/anything"
	// Does NOT extract ID. Only routes based on prefix.
	if strings.HasPrefix(r.URL.Path, "/customers/") {
		h.handleItem(w, r)
		return
	}

	// --- FALLBACK ---
	// If neither condition matched:
	// write 404 response
	http.NotFound(w, r)
}

// Handles requests where path == "/customers"
func (h *CustomerHandler) handleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		customers := h.service.GetAllCustomers()
		json.NewEncoder(w).Encode(customers)
	case http.MethodPost:
		var c store.Customer
		json.NewDecoder(r.Body).Decode(&c)
		err := h.service.CreateCustomer(&c)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.WriteHeader(http.StatusCreated)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handles any path starting with "/customers/"
func (h *CustomerHandler) handleItem(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/customers/")
	if id == "" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case http.MethodGet:
		c, err := h.service.GetCustomer(id)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		json.NewEncoder(w).Encode(c)

	case http.MethodPut:
		var c store.Customer
		json.NewDecoder(r.Body).Decode(&c)

		err := h.service.UpdateCustomer(id, &c)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

	case http.MethodDelete:
		err := h.service.DeleteCustomer(id)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
