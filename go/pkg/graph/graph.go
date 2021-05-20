package graph

import (
	"fmt"

	"github.com/nij4t/cs-lab/go/pkg/queue"
)

type Graph struct {
	nodes map[Node][]Node
}

type Node interface{}

func New() *Graph {
	return &Graph{nodes: map[Node][]Node{}}
}

func (g *Graph) AddNode(n Node) {
	g.nodes[n] = []Node{}
}

func (g *Graph) AddEdge(s Node, d Node) {
	g.nodes[s] = append(g.nodes[s], d)
	g.nodes[d] = append(g.nodes[d], s)
}

func (g *Graph) GetEdges(n Node) []Node {
	return g.nodes[n]
}

func (g Graph) String() string {
	res := ""

	for node := range g.nodes {
		res += fmt.Sprintf("%v => %v\n", node, g.GetEdges(node))
	}

	return res
}

func (g Graph) BFS(start Node, cb func(n Node)) {
	q := make(queue.Queue, 0)
	visited := map[Node]bool{}

	q.Enqueue(start)

	for !q.Empty() {
		for _, i := range g.GetEdges(q.Dequeue()) {
			cb(i)

			if _, ok := visited[i]; !ok {
				visited[i] = true
				q.Enqueue(i)
			}
		}
	}
}
