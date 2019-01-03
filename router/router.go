package router

import (
	"fmt"
	"log"
	"net/http"

	midd "github.com/TruongTrongThanh/ImageServer/middleware"
)

// Router ...
type Router struct {
	Root   string
	Routes []*Route
}

// NewRouter ...
func NewRouter(root string) *Router {
	rtr := new(Router)
	rtr.Root = root
	rtr.Routes = make([]*Route, 0)
	return rtr
}

// RegisterRoute ...
func (rtr *Router) RegisterRoute(mappingPath string, method string, handler http.HandlerFunc) {
	if frt, hasMethod := rtr.IsDuplicateRoute(mappingPath, method); frt != nil {
		if hasMethod {
			log.Printf("IGNORE duplicate route: %s - %s\n", mappingPath, method)
		} else {
			frt.MethodHandlers[method] = handler
		}
	} else {
		nrt := NewRoute(mappingPath)
		nrt.MethodHandlers[method] = handler
		rtr.Routes = append(rtr.Routes, nrt)
	}
}

// IsDuplicateRoute ...
func (rtr *Router) IsDuplicateRoute(path string, method string) (*Route, bool) {
	var foundRt *Route
	var hasMethod bool

	for _, r := range rtr.Routes {
		if path == r.Path {
			foundRt = r
			break
		}
	}
	if foundRt != nil && foundRt.MethodHandlers[method] != nil {
		hasMethod = true
	}
	return foundRt, hasMethod
}

// MappingTo ...
func (rtr *Router) MappingTo(mux *http.ServeMux) {
	for _, rt := range rtr.Routes {
		fmt.Printf("Mapping route: %s%s %v\n", rtr.Root, rt.Path, rt.Methods())
		handler := rt.Handler()
		mux.Handle(rtr.Root+rt.Path, midd.SetMiddlewares(handler))
	}
}
