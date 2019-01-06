package middleware

import (
	"net/http"
)

// MidHandlerFunc ...
type MidHandlerFunc func(next http.Handler) http.Handler

// SetMiddlewares ...
func SetMiddlewares(handler http.Handler) http.Handler {
	return middlewareChains(handler,
		Common,
		Auth,
	)
}

func middlewareChains(handler http.Handler, mds ...MidHandlerFunc) http.Handler {
	var final http.Handler
	lastIndex := len(mds) - 1
	for i := lastIndex; i >= 0; i-- {
		if i != lastIndex {
			final = mds[i](final)
		} else {
			final = mds[i](handler)
		}
	}
	return final
}
