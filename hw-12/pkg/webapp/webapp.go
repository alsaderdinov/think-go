package webapp

import (
	"encoding/json"
	"net/http"
	"think-go/hw-12/pkg/crawler"
	"think-go/hw-12/pkg/index"

	"github.com/gorilla/mux"
)

const addr = "localhost:8080"

type Service struct {
	idx  *index.Service
	docs []crawler.Document
}

func New(idx *index.Service, docs []crawler.Document) *Service {
	return &Service{
		idx:  idx,
		docs: docs,
	}
}

func (s *Service) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/docs", s.docsHandler).Methods(http.MethodGet)
	r.HandleFunc("/index", s.indexHandler).Methods(http.MethodGet)

	return http.ListenAndServe(addr, r)
}

func (s *Service) docsHandler(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(s.docs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Service) indexHandler(w http.ResponseWriter, _ *http.Request) {
	if err := json.NewEncoder(w).Encode(s.idx.Data()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
