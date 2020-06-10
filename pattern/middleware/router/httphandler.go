package router

import (
	"log"
	"net/http"
	"time"
)

type middleware func(handler http.Handler) http.Handler

type Router struct {
	middlewareChain []middleware
	mux             map[string]http.Handler
}

func (r *Router) Use(mw middleware) {
	r.middlewareChain = append(r.middlewareChain, mw)
}

func (r *Router) Add(path string, handler http.Handler) {
	var mergedHandler = handler

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[path] = mergedHandler
}

func timeMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		timeStart := time.Now()
		handler.ServeHTTP(res, req)
		timeElapsed := time.Since(timeStart)
		log.Print(timeElapsed)
	})
}
