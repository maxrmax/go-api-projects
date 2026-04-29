package api

// provides ServeMux and Handler
import "net/http"

// NewRouter constructs and returns a request multiplexer.
func NewRouter(authHandler http.Handler, customerHandler http.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/login", authHandler)
	mux.Handle("/customers", customerHandler)
	mux.Handle("/customers/", customerHandler)
	return mux
}
