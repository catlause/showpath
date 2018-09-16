package message

import (
	_ "showpath/core/graph"
	"showpath/core/ptlayer"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handles messages
func MessageHandler(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
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

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	X     int    `json:"xpos"`
	Y     int    `json:"ypos"`
	Color string `json:"color"`
}

type Edge struct {
	X0    int    `json:"x0"`
	Y0    int    `json:"y0"`
	X1    int    `json:"x1"`
	Y1    int    `json:"y1"`
	Color string `json:"color"`
}

func getgraph() (g Graph, err error) {
	err = nil
	pts := ptlayer.RandPts(20, 700, 670)
	var nodes []Node = nil
	for _, pt := range pts {
		nodes = append(nodes, Node{
			X:     pt.X,
			Y:     pt.Y,
			Color: "yellow",
		})
	}
	g = Graph{
		Nodes: nodes,
	}
	return
}
