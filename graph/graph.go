package graph

import (
	"fmt"

	"github.com/pawilonek/algo/queue"
)

type Graph struct {
	adjacencyList map[Node][]Node
}

func NewGraph() Graph {
	return Graph{
		adjacencyList: make(map[Node][]Node),
	}
}

func (g *Graph) AddNode(node Node) {
	g.adjacencyList[node] = make([]Node, 0)
}

func (g *Graph) AddEdge(nodeA Node, nodeB Node) {
	g.adjacencyList[nodeA] = append(g.adjacencyList[nodeA], nodeB)
	g.adjacencyList[nodeB] = append(g.adjacencyList[nodeB], nodeA)
}

func (g *Graph) Print() {
	for node, neighbors := range g.adjacencyList {
		fmt.Printf("Node %s: ", node.Value)
		for _, neighborNode := range neighbors {
			fmt.Printf("%s, ", neighborNode.Value)
		}
		fmt.Println()
	}
}

func (g *Graph) BFS(startNode Node, endNode Node) []Node {
	q := queue.NewQueue()
	q.Push(startNode)

	visited := make(map[Node]bool, len(g.adjacencyList))
	visited[startNode] = true

	prev := make(map[Node]*Node, len(g.adjacencyList))

	for !q.IsEmpty() {
		next, _ := q.Pop()
		currNode := next.(Node)

		for _, neighborNode := range g.adjacencyList[currNode] {
			if visited[neighborNode] {
				continue
			}

			q.Push(neighborNode)
			visited[neighborNode] = true
			prev[neighborNode] = &currNode
		}
	}

	// Reconstruct path
	path := make([]Node, 0)
	for at := &endNode; at != nil; at = prev[*at] {
		path = append(path, *at)
	}

	// Reverse the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if path[0] != startNode {
		return []Node{}
	}

	return path
}
