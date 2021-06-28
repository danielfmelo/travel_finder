package graph_test

import (
	"reflect"
	"testing"

	"github.com/danielfmelo/travel_finder/repositories/graph"
)

func TestDijkstra(t *testing.T) {
	t.Run("one", func(tt *testing.T) {
		graph := graph.NewGraph()
		graph.AddEdge("S", "P", 2)
		graph.AddEdge("S", "U", 3)
		graph.AddEdge("P", "Q", 5)
		graph.AddEdge("P", "X", 4)
		graph.AddEdge("U", "X", 1)
		graph.AddEdge("U", "V", 3)
		graph.AddEdge("X", "Q", 7)
		graph.AddEdge("X", "Y", 6)
		graph.AddEdge("X", "V", 8)
		graph.AddEdge("V", "W", 4)
		graph.AddEdge("Y", "R", 1)
		graph.AddEdge("Y", "W", 3)
		graph.AddEdge("Q", "R", 2)
		graph.AddEdge("R", "T", 6)
		graph.AddEdge("W", "T", 5)
		path, cost := graph.GetPath("S", "T")
		if cost != 15 {
			t.Errorf("expected 15 got %d", cost)
		}

		pathExpected := []string{"S", "P", "Q", "R", "T"}
		if !reflect.DeepEqual(path, pathExpected) {
			t.Errorf("expected %v got %v", pathExpected, path)
		}
	})

	t.Run("two", func(tt *testing.T) {
		graph := graph.NewGraph()
		graph.AddEdge("S", "B", 4)
		graph.AddEdge("S", "C", 2)
		graph.AddEdge("B", "C", 1)
		graph.AddEdge("B", "D", 5)
		graph.AddEdge("C", "D", 8)
		graph.AddEdge("C", "E", 10)
		graph.AddEdge("D", "E", 2)
		graph.AddEdge("D", "T", 6)
		graph.AddEdge("E", "T", 2)
		path, cost := graph.GetPath("S", "T")
		if cost != 12 {
			t.Errorf("expected 12 got %d", cost)
		}

		pathExpected := []string{"S", "C", "B", "D", "E", "T"}
		if !reflect.DeepEqual(path, pathExpected) {
			t.Errorf("expected %v got %v", pathExpected, path)
		}
	})
}
