package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"think-go/hw-13/pkg/crawler"

	"github.com/gorilla/mux"
)

// search handles the GET request to find specific documents by query.
func (api *API) search(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ids := api.index.Find(params["query"])

	var docs []crawler.Document

	for _, id := range ids {
		docs = append(docs, api.docs[id])
	}

	if err := json.NewEncoder(w).Encode(docs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getDocuments handles the GET request for all documents.
func (api *API) getDocuments(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(api.docs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// createDocument handles the POST request to create a new document
func (api *API) createDocument(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	if err := json.NewDecoder(r.Body).Decode(&doc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api.docs = append(api.docs, doc)
	w.WriteHeader(http.StatusCreated)
}

// getDocument handles the GET request for a specific document.
func (api *API) getDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, doc := range api.docs {
		if doc.ID == id {
			if err := json.NewEncoder(w).Encode(doc); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(w, r)
}

// deleteDocument handles the DELETE request for a specific document.
func (api *API) deleteDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for idx, doc := range api.docs {
		if doc.ID == id {
			api.docs = append(api.docs[:idx], api.docs[idx+1:]...)
			return
		}
	}

	http.NotFound(w, r)
}

// UpdateDocument handles the PUT request to update an existing document.
func (api *API) updateDocument(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedDoc crawler.Document
	if err := json.NewDecoder(r.Body).Decode(&updatedDoc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, doc := range api.docs {
		if doc.ID == id {
			api.docs[i] = updatedDoc
			return
		}
	}

	http.NotFound(w, r)
}
