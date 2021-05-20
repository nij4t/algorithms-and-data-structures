package graph

import (
	"reflect"
	"testing"
)

type friend struct {
	Node
	name        string
	sellsCatnip bool
}

type Friend interface {
	Name() string
	SellsCatnip() bool
}

func (d friend) Name() string {
	return d.name
}

func (d friend) SellsCatnip() bool {
	return d.sellsCatnip
}

func TestEdges(t *testing.T) {
	dealers := getFriends()
	graph1 := generateGraph(dealers)

	graph1.AddEdge(dealers[0], dealers[1])

	want := dealers[1]
	got := graph1.GetEdges(dealers[0])[0]

	if got != want {
		t.Errorf("%v != %v", got, want)
	}

}

func TestBFS(t *testing.T) {
	friends := getFriends()
	graph1 := generateGraph(friends)

	graph1.AddEdge(friends[0], friends[1])
	graph1.AddEdge(friends[0], friends[2])
	graph1.AddEdge(friends[0], friends[4])
	graph1.AddEdge(friends[4], friends[0])
	graph1.AddEdge(friends[4], friends[5])

	want := []Friend{friends[5]}
	got := []Friend{}

	graph1.BFS(friends[0], func(dealer Node) {
		if dealer.(friend).SellsCatnip() {
			got = append(got, dealer.(Friend))
		}
	})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func generateGraph(nodes []Friend) *Graph {
	graph1 := New()
	for _, node := range nodes {
		graph1.AddNode(node)
	}
	return graph1
}

func getFriends() []Friend {
	nodes := []Friend{
		friend{name: "Alice", sellsCatnip: false},
		friend{name: "Bob", sellsCatnip: false},
		friend{name: "Charlie", sellsCatnip: false},
		friend{name: "Drake", sellsCatnip: false},
		friend{name: "Elen", sellsCatnip: false},
		friend{name: "Freddie", sellsCatnip: true},
		friend{name: "George", sellsCatnip: false},
		friend{name: "Henry", sellsCatnip: true},
	}
	return nodes
}
