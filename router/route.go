package router

import (
	"net/http"
)

// Route ...
type Route struct {
	Path           string
	MethodHandlers map[string]http.HandlerFunc
}

// NewRoute ...
func NewRoute(mappingPath string) *Route {
	rt := new(Route)
	rt.Path = mappingPath
	rt.MethodHandlers = make(map[string]http.HandlerFunc)
	return rt
}

// Handler ...
func (rt *Route) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handleOrNotFound(w, r, rt.MethodHandlers[r.Method])
	})
}

// Methods get available HTTP method list
func (rt *Route) Methods() []string {
	methods := make([]string, 0)
	for key := range rt.MethodHandlers {
		methods = append(methods, key)
	}
	return methods
}

func handleOrNotFound(w http.ResponseWriter, r *http.Request, hFunc http.HandlerFunc) {
	if hFunc != nil {
		hFunc.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}
