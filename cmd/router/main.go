package main

import (
	services "andrefsilveira1/router/internal/application/service"
	"andrefsilveira1/router/internal/domain/entities"
	"fmt"
)

func main() {
	// Instanciate graph and nodes
	graph := entities.CreateGraph()

	node1 := entities.CreateNode(1)
	node2 := entities.CreateNode(2)
	node3 := entities.CreateNode(3)
	node4 := entities.CreateNode(4)

	node1.AppendAdjacent(2, 1)
	node1.AppendAdjacent(3, 4)
	node2.AppendAdjacent(3, 2)
	node2.AppendAdjacent(4, 7)
	node3.AppendAdjacent(4, 3)

	graph.AddNode(node1)
	graph.AddNode(node2)
	graph.AddNode(node3)
	graph.AddNode(node4)

	// Applying dijkstra algorithm

	distances := services.Dijkstra(graph, 1)

	// graph.PrintGraphTree(1) This is terrible...

	// I really think that we should display this in a different way...
	// Maybe create a figma example using React ?

	fmt.Println("Shortest distances from node 1:")
	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}

}
