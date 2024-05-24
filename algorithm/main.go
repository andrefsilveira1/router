package main

import (
	"andrefsilveira1/router/algorithm/domain"
	"fmt"
)

func printGraph(graph domain.Graph) {
	for id, node := range graph.Nodes {
		fmt.Printf("Node ID: %d, Adjacent: %+v\n", id, node.Adjacent)
	}
}

func main() {
	graph := domain.CreateGraph() // Create Graph

	node1 := domain.CreateNode(1) // Create node 1
	node2 := domain.CreateNode(2) // Create node 2

	node1.AppendAdjacent(2, 5) // Add adjacent nodes
	node2.AppendAdjacent(1, 5) // Add adjacent nodes

	// Add nodes to graph
	graph.AddNode(node1)
	graph.AddNode(node2)

	fmt.Println("Graph setup completed!")
	fmt.Println("")

	printGraph(*graph) // Print graph
}
