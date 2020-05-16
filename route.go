package golam

import (
	"net/http"
)

// Route represents a single path that can be accessed
// over HTTP in the server
type Route struct {
	Pattern      string
	Middlewares  []Middleware
	Get          http.HandlerFunc
	Post         http.HandlerFunc
	Put          http.HandlerFunc
	Delete       http.HandlerFunc
	Head         http.HandlerFunc
	Options      http.HandlerFunc
	NotSupported http.HandlerFunc
}

// ServeHTTP implementation to support default library of Handler
func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		r.handleGet(res, req)
		return
	case http.MethodPut:
		r.handlePut(res, req)
		return
	case http.MethodPost:
		r.handlePost(res, req)
		return
	case http.MethodDelete:
		r.handleDelete(res, req)
		return
	case http.MethodHead:
		r.handleHead(res, req)
		return
	case http.MethodOptions:
		r.handleOptions(res, req)
		return
	}
	r.handleNotSupported(res, req)
}

func (r *Route) handleNotSupported(res http.ResponseWriter, req *http.Request) {
	if nil != r.NotSupported {
		r.handle(r.NotSupported, res, req)
		return
	}
}

func (r *Route) handleGet(res http.ResponseWriter, req *http.Request) {
	if nil != r.Get {
		r.handle(r.Get, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handlePost(res http.ResponseWriter, req *http.Request) {
	if nil != r.Post {
		r.handle(r.Post, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handlePut(res http.ResponseWriter, req *http.Request) {
	if nil != r.Put {
		r.handle(r.Put, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handleDelete(res http.ResponseWriter, req *http.Request) {
	if nil != r.Delete {
		r.handle(r.Delete, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handleHead(res http.ResponseWriter, req *http.Request) {
	if nil != r.Head {
		r.handle(r.Head, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handleOptions(res http.ResponseWriter, req *http.Request) {
	if nil != r.Options {
		r.handle(r.Options, res, req)
		return
	}
	r.NotSupported(res, req)
}

func (r *Route) handle(fn http.Handler, res http.ResponseWriter, req *http.Request) {
	for i := 0; i < len(r.Middlewares); i++ {
		fn = r.Middlewares[i](fn)
	}
	fn.ServeHTTP(res, req)
}
