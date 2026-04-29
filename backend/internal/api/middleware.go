package api

// provides Handler, HandlerFunc, Request, ResponseWriter
import "net/http"

// --- FUNCTION SIGNATURE ---
// AuthMiddleware is a higher-order function.
//
// Input:
//
//	next http.Handler
//	  - any value that implements:
//	      ServeHTTP(ResponseWriter, *Request)
//
// Output:
//
//	http.Handler
//	  - a NEW handler that wraps the input handler
//
// No mutation of `next` occurs.
// A new handler is constructed and returned.
func AuthMiddleware(next http.Handler) http.Handler {

	// --- RETURN VALUE ---
	// http.HandlerFunc is a function type:
	//
	//     type HandlerFunc func(ResponseWriter, *Request)
	//
	// It has a method:
	//
	//     func (f HandlerFunc) ServeHTTP(w, r) {
	//         f(w, r)
	//     }
	//
	// This allows a plain function to satisfy http.Handler.
	//
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// --- HEADER EXTRACTION ---
		// r.Header is a map-like structure:
		//     map[string][]string
		//
		// Get("Authorization"):
		//   - returns the FIRST value for that key
		//   - if key does not exist → returns empty string ""
		//
		token := r.Header.Get("Authorization")

		// --- VALIDATION CHECK ---
		// Exact string comparison.
		//
		// Required value:
		//   "Bearer demo-token"
		//
		// No parsing:
		//   - "Bearer" is not extracted separately
		//   - no trimming or normalization
		//
		// Any deviation results in failure.
		//
		if token != "Bearer demo-token" {

			// --- FAILURE RESPONSE ---
			// http.Error:
			//   - sets status code (401 here)
			//   - writes plain text body "unauthorized\n"
			//
			// After writing:
			//   response is considered complete
			//
			http.Error(w, "unauthorized", http.StatusUnauthorized)

			// Prevents execution of the next handler.
			return
		}

		// --- FORWARDING ---
		// Calls the wrapped handler.
		//
		// Control flow continues into:
		//   next.ServeHTTP(w, r)
		//
		// Same ResponseWriter and Request are passed through.
		//
		next.ServeHTTP(w, r)
	})
}

/* or
type authMiddleware struct {
	next http.Handler
}

func (m authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token != "Bearer demo-token" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	m.next.ServeHTTP(w, r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return authMiddleware{next: next}
}
*/

/* or
func AuthMiddleware(next http.Handler) http.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token != "Bearer demo-token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
*/

/* or
func isAuthorized(r *http.Request) bool {
	return r.Header.Get("Authorization") == "Bearer demo-token"
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isAuthorized(r) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
*/

/* or
type Middleware struct {
	Handler http.Handler
}

func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != "Bearer demo-token" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	m.Handler.ServeHTTP(w, r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return Middleware{Handler: next}
}
*/
