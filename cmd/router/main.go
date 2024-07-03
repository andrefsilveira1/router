package main

import (
	"andrefsilveira1/router/internal/application/service"
	"fmt"
	"sync"
)

var payload = map[string]string{
	"message":   "This is the payload message",
	"timestamp": "2024-07-01T12:00:00Z",
}

var payload2 = map[string]string{
	"message":   "This is the second payload message",
	"timestamp": "2024-08-01T12:00:00Z",
}

func main() {
	// Criando malha 8x8
	mesh := service.CreateMesh()

	// Desabilitando nós
	service.DisableNodes(mesh)

	startNode, _ := mesh.GetNode(0, 0)
	finalNode, _ := mesh.GetNode(7, 7)

	startNode2, _ := mesh.GetNode(0, 1)
	finalNode2, _ := mesh.GetNode(7, 6)

	mesh.SetPayload(startNode.X, startNode.Y, payload)
	mesh.SetPayload(startNode2.X, startNode2.Y, payload2)

	fmt.Println("Initial payloads:")
	mesh.PrintPayload(startNode.Payload)
	mesh.PrintPayload(startNode.Payload)
	fmt.Println()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		path, hops := mesh.StarAlgorithm(startNode, finalNode)
		if path == nil {
			fmt.Println("No path found or missing path for payload 1")
		} else {
			fmt.Printf("Path found with %d hops for payload 1:\n", hops)
			mesh.Print(startNode, finalNode, path)
			fmt.Println("\nPayload at goal node 1:")
			mesh.PrintPayload(finalNode.Payload)
		}
	}()

	go func() {
		defer wg.Done()
		path, hops := mesh.StarAlgorithm(startNode2, finalNode2)
		if path == nil {
			fmt.Println("No path found or missing path for payload 2")
		} else {
			fmt.Printf("Path found with %d hops for payload 2:\n", hops)
			mesh.Print(startNode2, finalNode2, path)
			fmt.Println("\nPayload at goal node 2:")
			mesh.PrintPayload(finalNode2.Payload)
		}
	}()

	wg.Wait()

	mesh.PrintPayload(finalNode.Payload)
	mesh.PrintPayload(finalNode2.Payload)
}
