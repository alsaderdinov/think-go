package storage

import "math/rand"

// Constants for generating random short links
const (
	letters = "abcdefghijklmnopqrstuwxyzABCDEFGHIJKLMNOPQRSTUWXYZ"
	strSize = 11
)

// Storage is an interface that defines methods for adding and retrieving links.
type Storage interface {
	AddLink(string) string
	GetLink(string) string
}

// Service is the implementation of the Storage
type Service struct {
	links map[string]string
}

// New creates a new Service instance and initializes its internal map.
func New() *Service {
	s := Service{
		links: make(map[string]string),
	}
	return &s
}

// AddLink generates a random short link and associates it with the provided original link.
// It returns the generated short link.
func (s *Service) AddLink(link string) string {
	genStr := randString(strSize)
	s.links[genStr] = link

	return genStr
}

// GetLink retrieves the original link associated with the provided short link.
// If the short link is not found, it returns an empty string.
func (s *Service) GetLink(short string) string {
	return s.links[short]
}

// randString generates a random string of the specified length using the given letters.
func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
