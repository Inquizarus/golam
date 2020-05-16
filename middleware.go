package golam

import "net/http"

// Middleware helps modify incoming request and outgoing responses
// in a modular way and is mostly an alias for Handler from http package
type Middleware func(http.Handler) http.Handler
