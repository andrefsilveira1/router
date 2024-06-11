package main

import (
	services "andrefsilveira1/router/internal/application/service"
	"andrefsilveira1/router/internal/domain/entity"
	"fmt"
)

func main() {
	graph := entity.CreateGraph()

	n1 := entity.CreateNode(1, 0, 0)
	n2 := entity.CreateNode(2, 1, 0)
	n3 := entity.CreateNode(3, 2, 0)
	n4 := entity.CreateNode(4, 2, 1)
	n5 := entity.CreateNode(5, 2, 2)

	// Append adjacent nodes (assume weight is 1 for simplicity)
	n1.AppendAdjacent(2, 1)
	n2.AppendAdjacent(1, 1)
	n2.AppendAdjacent(3, 1)
	n3.AppendAdjacent(2, 1)
	n3.AppendAdjacent(4, 1)
	n4.AppendAdjacent(3, 1)
	n4.AppendAdjacent(5, 1)
	n5.AppendAdjacent(4, 1)

	// Add nodes to the graph
	graph.AddNode(n1)
	graph.AddNode(n2)
	graph.AddNode(n3)
	graph.AddNode(n4)
	graph.AddNode(n5)

	// Test the XY routing algorithm
	path := services.XYRoute(graph, 1, 4)
	if path != nil {
		fmt.Println("Path found:", path)
	} else {
		fmt.Println("No path found.")
	}
}
