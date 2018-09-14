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
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	X  int `json:"xpos"`
	Y  int `json:"ypos"`
}

type Edge struct {
	X0 int `json:"x0"`
	Y0 int `json:"y0"`
	X1 int `json:"x1"`
	Y1 int `json:"y1"`
}

// explore explores a path.
// If path is empty, it explores the user's home directory
func getgraph() (g Graph, err error) {
	err = nil
	g = Graph{
		Nodes: []Node{
			Node{X: 122, Y: 238},
			Node{X: 33, Y: 117},
		},
		Edges: []Edge{
			Edge{X0: 122, Y0: 238, X1: 33, Y1: 117},
		},
	}
	return
}
