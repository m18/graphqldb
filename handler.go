package main

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/m18/graphqldb/loader"
)

// GraphQLHandler is a graphql HTTP handler
type graphQLHandler struct {
	handler *relay.Handler
	loaders loader.Map
}

// NewHandler returns a new http.Handler
func newHandler(schema *graphql.Schema, loaders loader.Map) http.Handler {
	h := &graphQLHandler{
		handler: &relay.Handler{Schema: schema},
		loaders: loaders,
	}
	mux := http.NewServeMux()
	mux.Handle("/graphql", h)
	mux.Handle("/graphql/", h)
	return mux
}

// ServeHTTP serves http requests
func (h *graphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := h.loaders.Attach(r.Context())
	r = r.WithContext(ctx)
	h.handler.ServeHTTP(w, r)
}
