package webapp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"think-go/hw-12/pkg/crawler"
	"think-go/hw-12/pkg/index"
)

func TestService_indexHandler(t *testing.T) {
	idx := index.New()
	docs := []crawler.Document{{ID: 1, Title: "Test Title", Body: "Test Body", URL: "https://example.com"}}
	service := New(idx, docs)

	for i, doc := range docs {
		doc.ID = i
		idx.Add(doc.Title, doc.ID)
	}

	req, err := http.NewRequest("GET", "/index", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	got := rr.Body.String()
	want := `{"test":[0],"title":[0]}`

	if !strings.Contains(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestService_docsHandler(t *testing.T) {
	idx := index.New()
	docs := []crawler.Document{{ID: 1, Title: "Test Title", Body: "Test Body", URL: "https://example.com"}}
	service := New(idx, docs)

	req, err := http.NewRequest("GET", "/docs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.docsHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	got := rr.Body.String()
	want := `[{"ID":1,"URL":"https://example.com","Title":"Test Title","Body":"Test Body"}]`

	if !strings.Contains(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
