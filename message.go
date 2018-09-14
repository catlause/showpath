package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "getgraph":
		// GetGraph
		if payload, err = getgraph(); err != nil {
			payload = err.Error()
			return
		}
	}
	return
}

// Exploration represents the results of an exploration
type Graph struct {
	Coords []Coord `json:"coords"`
}

type Coord struct {
	X int `json:"xpos"`
	Y int `json:"ypos"`
}

// explore explores a path.
// If path is empty, it explores the user's home directory
func getgraph() (g Graph, err error) {
	err = nil
	g = Graph{
		Coords: []Coord{
			Coord{X: 122, Y: 238},
			Coord{X: 33, Y: 117},
		},
	}
	return
}
