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
	// Create an 8x8 mesh
	mesh := service.CreateMesh()

	// Disable nodes
	service.DisableNodes(mesh)

	startNode1, _ := mesh.GetNode(0, 0)
	finalNode1, _ := mesh.GetNode(7, 7)

	startNode2, _ := mesh.GetNode(0, 1)
	finalNode2, _ := mesh.GetNode(7, 6)

	mesh.SetPayload(startNode1.X, startNode1.Y, payload)
	mesh.SetPayload(startNode2.X, startNode2.Y, payload2)

	fmt.Println("Initial payloads:")
	mesh.PrintPayload(startNode1.Payload)
	mesh.PrintPayload(startNode2.Payload)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		path1, hops1 := mesh.StarAlgorithm(startNode1, finalNode1, '+')
		if path1 == nil {
			fmt.Println("No path found or missing path for payload 1")
		} else {
			fmt.Printf("Path found with %d hops for payload 1:\n", hops1)
			mesh.PrintPayload(finalNode1.Payload)
		}
	}()

	go func() {
		defer wg.Done()
		path2, hops2 := mesh.StarAlgorithm(startNode2, finalNode2, '-')
		if path2 == nil {
			fmt.Println("No path found or missing path for payload 2")
		} else {
			fmt.Printf("Path found with %d hops for payload 2:\n", hops2)
			mesh.PrintPayload(finalNode2.Payload)
		}
	}()

	wg.Wait()

	fmt.Println("\nFinal payloads at goal nodes:")
	mesh.PrintPayload(finalNode1.Payload)
	mesh.PrintPayload(finalNode2.Payload)
}
