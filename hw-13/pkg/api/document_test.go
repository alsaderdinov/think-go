package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"think-go/hw-13/pkg/crawler"
	"think-go/hw-13/pkg/index"
)

var testAPI *API

func TestMain(m *testing.M) {
	docs := []crawler.Document{
		{
			ID:    0,
			Title: "go",
			URL:   "https://golang.org",
		},
		{
			ID:    1,
			Title: "elixir",
			URL:   "https://elixir-lang.org",
		},
	}

	index := index.New()

	for _, doc := range docs {
		index.Add(doc.Title, doc.ID)
	}

	testAPI = New(docs, index)
	os.Exit(m.Run())
}

func TestAPI_search(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/documents/search/go", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	want := `[{"ID":0,"URL":"https://golang.org","Title":"go","Body":""}]`
	got := rr.Body.String()

	if !strings.Contains(got, want) {
		t.Errorf("got %v want %v", got, string(want))
	}
}

func TestAPI_getDocuments(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/documents", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	want, _ := json.Marshal(testAPI.docs)
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got %v want %v", got, string(want))
	}
}

func TestAPI_createDocument(t *testing.T) {
	payload, _ := json.Marshal(crawler.Document{URL: "https://ruby-lang.org", Title: "ruby"})

	req, err := http.NewRequest("POST", "/api/v1/documents", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("got %v want %v", rr.Code, http.StatusCreated)
	}

}

func TestAPI_getDocument(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/documents/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", rr.Code, http.StatusOK)
	}

	got := strings.TrimSuffix(rr.Body.String(), "\n")
	want, _ := json.Marshal(testAPI.docs[0])

	if got != string(want) {
		t.Errorf("got %v want %v", got, string(want))
	}
}

func TestAPI_deleteDocument(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/v1/documents/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	req, _ = http.NewRequest("GET", "/api/v1/documents/1", nil)
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("got %v want %v", status, http.StatusNotFound)
	}
}

func TestAPI_updateDocument(t *testing.T) {
	updatedDoc := crawler.Document{
		ID:    0,
		URL:   "https://ruby-lang.org",
		Title: "Ruby",
		Body:  "ruby programming language",
	}

	payload, _ := json.Marshal(updatedDoc)
	req, err := http.NewRequest("PUT", "/api/v1/documents/0", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	testAPI.router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/api/v1/documents/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	testAPI.router.ServeHTTP(rr, req)

	want, _ := json.Marshal(updatedDoc)
	got := strings.TrimSuffix(rr.Body.String(), "\n")

	if got != string(want) {
		t.Errorf("got %v want %v", got, string(want))
	}
}
