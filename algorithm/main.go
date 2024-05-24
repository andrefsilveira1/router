package main

import (
	"andrefsilveira1/router/algorithm/domain/entities"
	"fmt"
)

func printGraph(graph entities.Graph) {
	for id, node := range graph.Nodes {
		fmt.Printf("Node ID: %d, Adjacent: %+v\n", id, node.Adjacent)
	}
}

func main() {
	graph := entities.CreateGraph() // Create Graph

	node1 := entities.CreateNode(1) // Create node 1
	node2 := entities.CreateNode(2) // Create node 2

	node1.AppendAdjacent(2, 5) // Add adjacent nodes
	node2.AppendAdjacent(1, 5) // Add adjacent nodes

	// Add nodes to graph
	graph.AddNode(node1)
	graph.AddNode(node2)

	fmt.Println("Graph setup completed!")
	fmt.Println("")

	printGraph(*graph) // Print graph
}
