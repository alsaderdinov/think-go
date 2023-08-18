package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	g := newGame(11)
	g.start()
}

type game struct {
	sBoard  map[string]int
	tableCh chan string
	wScore  int
	sync.WaitGroup
}

func newGame(wScore int) *game {
	return &game{
		sBoard:  make(map[string]int),
		tableCh: make(chan string),
		wScore:  wScore,
	}
}

func (g *game) start() {
	g.Add(2)
	go g.player("Narrator")
	go g.player("Tyler")

	g.tableCh <- "begin"

	g.Wait()

	for k, v := range g.sBoard {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func (g *game) player(name string) {
	defer g.Done()

	for v := range g.tableCh {
		if g.hasWinner() {
			close(g.tableCh)
			break
		}

		var state string
		switch v {
		case "begin", "stop", "pong":
			state = "ping"
		case "ping":
			state = "pong"
		}
		fmt.Printf("%s: %s\n", name, state)

		if hasLuck() {
			g.addPoint(name)
			fmt.Println(name, "strikes hard and scores a point")
			g.tableCh <- "stop"
		} else {
			g.tableCh <- state
		}
	}
}

func (g *game) hasWinner() bool {
	for _, score := range g.sBoard {
		if score == g.wScore {
			return true
		}
	}
	return false
}

func (g *game) addPoint(name string) {
	g.sBoard[name]++
}

func hasLuck() bool {
	return rand.Intn(5) == 0
}
