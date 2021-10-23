package graph_test

import (
	"fmt"
	"testing"

	"github.com/Pawilonek/algo/graph"
)

func TestBreadthFirstSearch(t *testing.T) {
	g := graph.NewGraph()

	// Preapre example graph: https://youtu.be/oDqjPvD54Ss?t=102
	nodes := make([]graph.Node, 0)
	for i := 0; i < 14; i++ {
		node := graph.Node{Value: fmt.Sprintf("%02d", i)}

		nodes = append(nodes, node)
		g.AddNode(node)
	}

	g.AddEdge(nodes[0], nodes[7])
	g.AddEdge(nodes[0], nodes[9])
	g.AddEdge(nodes[0], nodes[11])
	g.AddEdge(nodes[1], nodes[8])
	g.AddEdge(nodes[1], nodes[10])
	g.AddEdge(nodes[2], nodes[3])
	g.AddEdge(nodes[2], nodes[12])
	g.AddEdge(nodes[3], nodes[4])
	g.AddEdge(nodes[3], nodes[7])
	g.AddEdge(nodes[5], nodes[6])
	g.AddEdge(nodes[6], nodes[7])
	g.AddEdge(nodes[7], nodes[11])
	g.AddEdge(nodes[8], nodes[9])
	g.AddEdge(nodes[8], nodes[12])
	g.AddEdge(nodes[9], nodes[10])

	// Couple of scenarios for the graph
	var bfsScenarions = []struct {
		scenario     string
		startNode    graph.Node
		endNode      graph.Node
		expectedPath []graph.Node
	}{
		{
			"Non existing path",
			nodes[0],
			nodes[13],
			[]graph.Node{},
		},
		{
			"Scenario 1",
			nodes[0],
			nodes[4],
			[]graph.Node{
				nodes[0],
				nodes[7],
				nodes[3],
				nodes[4],
			},
		},
		{
			"Scenario 2",
			nodes[11],
			nodes[2],
			[]graph.Node{
				nodes[11],
				nodes[7],
				nodes[3],
				nodes[2],
			},
		},
		{
			"Scenario 3",
			nodes[10],
			nodes[5],
			[]graph.Node{
				nodes[10],
				nodes[9],
				nodes[0],
				nodes[7],
				nodes[6],
				nodes[5],
			},
		},
	}

	// Testing
	for _, tt := range bfsScenarions {
		t.Run(tt.scenario, func(t *testing.T) {
			result := g.BFS(tt.startNode, tt.endNode)

			if len(tt.expectedPath) != len(result) {
				t.Errorf("Expected patch length is invalid")

				return
			}

			for key, expectedNode := range tt.expectedPath {
				if expectedNode != result[key] {
					t.Errorf(
						"Paths do not on %d step. Expected: %s but got: %s",
						key,
						expectedNode.Value,
						result[key].Value,
					)

					return
				}
			}
		})
	}
}
