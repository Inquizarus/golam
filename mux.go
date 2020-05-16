package golam

import "net/http"

// Mux defines which functionality an injected mux library should
// have and aids in keeping Golam flexible in implementations
type Mux interface {
	HandleFunc(pattern string, handlerFn http.HandlerFunc)
	Handle(pattern string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Use(middleware ...http.Handler)
}
