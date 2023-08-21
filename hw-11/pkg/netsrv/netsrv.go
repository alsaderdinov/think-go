package netsrv

import (
	"bufio"
	"fmt"
	"net"
	"think-go/hw-11/pkg/crawler"
	"think-go/hw-11/pkg/index"
	"time"
)

// Server represents a network server that handles incoming connections and searches for documents in an index.
type Server struct {
	host  string
	port  string
	index *index.Service
	docs  []crawler.Document
}

// New creates and initializes a new Server instance.
func New(host, port string, idx *index.Service, docs []crawler.Document) *Server {
	return &Server{
		host:  host,
		port:  port,
		index: idx,
		docs:  docs,
	}
}

// Run starts the server and listens for incoming connections.
func (s *Server) Run() error {
	listener, err := net.Listen("tcp", net.JoinHostPort(s.host, s.port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go s.handleConn(conn)
	}
}

// handleConn processes an incoming connection by searching for documents and sending their details.
func (s *Server) handleConn(conn net.Conn) error {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	for {
		msg, _, err := rd.ReadLine()
		if err != nil {
			return err
		}

		ids := s.index.Find(string(msg))

		for _, id := range ids {
			doc := s.docs[id]
			if _, err := fmt.Fprintf(conn, "%s: %s\n", doc.Title, doc.URL); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintf(conn, "\n"); err != nil {
			return nil
		}

		conn.SetDeadline(time.Now().Add(30 * time.Second))
	}
}
