package router

import (
	"github.com/gagansingh3785/typio-service/handlers"
	"github.com/gagansingh3785/typio-service/registry"
	"github.com/gorilla/mux"
)

type RouterOption interface {
	Apply(router *mux.Router)
}

type RouterOptionFunc func(router *mux.Router)

func (f RouterOptionFunc) Apply(router *mux.Router) {
	f(router)
}

func WithPingRoute() RouterOptionFunc {
	return func(r *mux.Router) {
		r.HandleFunc("/ping", handlers.PingHandler())
	}
}

func WithExternalRoutes(svcRegistry *registry.ServiceRegistry) RouterOptionFunc {
	return func(r *mux.Router) {
		r.HandleFunc("/v1/paragraphs", handlers.ParagraphsV1Handler(svcRegistry))
	}
}

func NewRouterWithOptions(options ...RouterOption) *mux.Router {
	router := mux.NewRouter()

	for _, option := range options {
		option.Apply(router)
	}

	return router
}
