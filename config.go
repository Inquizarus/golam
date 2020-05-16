package golam

import "net/http"

// Config for determining how golam runs
type Config struct {
	Middlewares []func(http.Handler) http.Handler
	Routes      []Route
	UseTLSPort  int
	UsePort     int
	TLSCertPath string
	TLSKeyPath  string
	TLSEnabled  bool
}
