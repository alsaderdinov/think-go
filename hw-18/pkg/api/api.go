package api

import (
	"encoding/json"
	"net/http"
	"think-go/hw-18/pkg/storage"

	"github.com/gorilla/mux"
)

// Define a struct to represent the JSON request body for adding a link.
var body struct {
	Link string `json:"link"`
}

// API represents the API server with endpoints for adding and retrieving links.
type API struct {
	Router  *mux.Router
	storage *storage.Service
}

// New creates a new instance of the API with the provided storage service.
func New(s *storage.Service) *API {
	api := API{
		Router:  mux.NewRouter(),
		storage: s,
	}

	api.endpoints()

	return &api
}

// endpoints sets up the API HTTP endpoints and middleware.
func (api *API) endpoints() {
	api.Router.Use(headersMiddleware)

	api.Router.HandleFunc("/api/v1/lynks/add", api.addLinkHandler).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/v1/lynks/get/{shortLink}", api.getLinkHandler).Methods(http.MethodGet)
}

// addLinkHandler handles the HTTP POST request for adding a new link.
func (api *API) addLinkHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortLink := api.storage.AddLink(body.Link)
	resp := map[string]string{"shortLink": shortLink}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// getLinkHandler handles the HTTP GET request for retrieving the original link associated with a short link.
func (api *API) getLinkHandler(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	shortLink := v["shortLink"]

	link := api.storage.GetLink(shortLink)
	resp := map[string]string{"origLink": link}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
