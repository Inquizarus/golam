package golam

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

var tlsConfig = tls.Config{
	MinVersion:               tls.VersionTLS12,
	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	PreferServerCipherSuites: true,
	CipherSuites: []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	},
}

// Run the server
func Run(handler Mux, cfg Config) {
	handler.Use(cfg.Middlewares...)
	for i := 0; i < len(cfg.Routes); i++ {
		handler.Handle(cfg.Routes[i].Pattern, &cfg.Routes[i])
	}
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.UsePort),
		Handler: handler,
	}
	if true == cfg.TLSEnabled {
		server.TLSConfig = &tlsConfig
		server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
		server.ListenAndServeTLS(cfg.TLSCertPath, cfg.TLSKeyPath)
	}
	if false == cfg.TLSEnabled {
		server.ListenAndServe()
	}
}
