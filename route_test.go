package golam_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/inquizarus/golam"
	"github.com/stretchr/testify/assert"
)

func withTestHeader() golam.Middleware {
	return func(f http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("test", "yes")
			f.ServeHTTP(w, r)
		})
	}
}

var testRoute = golam.Route{
	Get: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from GET"))
	},
	Post: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from POST"))
	},
	Put: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from PUT"))
	},
	Delete: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from DELETE"))
	},
	Head: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from HEAD"))
	},
	Options: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from OPTIONS"))
	},
	NotSupported: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from NotSupported"))
	},
	Middlewares: []golam.Middleware{
		withTestHeader(),
	},
}

func TestThatRouteHandlersWorks(t *testing.T) {
	cases := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		"NotSupported",
	}
	for _, c := range cases {
		t.Logf("running test for Route handler %s", c)
		expectedBody := []byte(fmt.Sprintf("Hello from %s", c))
		res := httptest.NewRecorder()
		req := httptest.NewRequest(c, "/", strings.NewReader(""))
		testRoute.ServeHTTP(res, req)
		assert.Equal(t, expectedBody, res.Body.Bytes())
	}
}

func TestThatMiddlewaresAreApplied(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	testRoute.ServeHTTP(res, req)
	assert.Equal(t, "yes", res.Header().Get("test"))
}
