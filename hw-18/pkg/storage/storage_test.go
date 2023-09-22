package storage

import (
	"testing"
)

func TestService_AddLinkAndGetLink(t *testing.T) {
	s := New()
	origLink := "https://example.com"
	shortLink := s.AddLink(origLink)

	if retrievedLink := s.GetLink(shortLink); retrievedLink != origLink {
		t.Errorf("got %v want %v", retrievedLink, origLink)
	}
}
