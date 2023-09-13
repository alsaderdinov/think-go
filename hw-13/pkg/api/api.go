package api

import (
	"net/http"
	"think-go/hw-13/pkg/crawler"
	"think-go/hw-13/pkg/index"

	"github.com/gorilla/mux"
)

// API is a struct that represents the API service.
type API struct {
	router *mux.Router
	index  *index.Service
	docs   []crawler.Document
}

// New create a new instance of the API service.
func New(docs []crawler.Document, idx *index.Service) *API {
	api := API{
		router: mux.NewRouter(),
		docs:   docs,
		index:  idx,
	}

	api.endpoints()
	return &api
}

// Router returns the router of the API service
func (api *API) Router() *mux.Router {
	return api.router
}

// endpoints of API service.
func (api *API) endpoints() {
	api.router.Use(requestIDMiddleware)
	api.router.Use(headersMiddleware)
	api.router.Use(logMiddleware)

	api.router.HandleFunc("/api/v1/documents/search/{query}", api.search).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/documents", api.getDocuments).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/documents", api.createDocument).Methods(http.MethodPost)
	api.router.HandleFunc("/api/v1/documents/{id}", api.getDocument).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/documents/{id}", api.updateDocument).Methods(http.MethodPut)
	api.router.HandleFunc("/api/v1/documents/{id}", api.deleteDocument).Methods(http.MethodDelete)
}

// Serve run the API service
func (api *API) Serve(addr string) error {
	return http.ListenAndServe(addr, api.router)
}
