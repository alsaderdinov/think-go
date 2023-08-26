package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	network = "tcp"
	address = "0.0.0.0:8000"
)

func main() {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	srvRd := bufio.NewReader(conn)
	cmdRd := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your query for search: ")
		query, _, err := cmdRd.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(conn, "%s\n", string(query))

		for {
			reply, _, err := srvRd.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			if len(reply) == 0 {
				break
			}
			fmt.Println(string(reply))
		}
	}
}
