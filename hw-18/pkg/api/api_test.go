package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"think-go/hw-18/pkg/storage"
)

var api *API

func TestMain(m *testing.M) {
	s := storage.New()
	api = New(s)
}

func TestAPI_addLinkHandler(t *testing.T) {
	reqBody := `{"link": "https://example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/lynks/add", strings.NewReader(reqBody))
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	if _, ok := resp["shortLink"]; !ok {
		t.Error("Response does not contain 'shortLink' field")
	}
}

func TestAPI_getLinkHandler(t *testing.T) {
	origLink := "https://example.com"
	shortLink := api.storage.AddLink(origLink)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/lynks/get/"+shortLink, nil)
	rr := httptest.NewRecorder()

	api.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}

	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	got := resp["origLink"]
	want := origLink

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
