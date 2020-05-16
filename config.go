package golam

// Config for determining how golam runs
type Config struct {
	Middlewares []Middleware
	Routes      []Route
	UseTLSPort  int
	UsePort     int
	TLSCertPath string
	TLSKeyPath  string
	TLSEnabled  bool
}
