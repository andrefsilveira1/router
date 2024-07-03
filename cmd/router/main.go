package main

import (
	"andrefsilveira1/router/internal/application/service"
	"andrefsilveira1/router/internal/domain"
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

	var path1, path2 []*domain.Node

	go func() {
		defer wg.Done()
		var hops int
		path1, hops = mesh.StarAlgorithm(startNode1, finalNode1, path2, startNode2, finalNode2, 'A')
		if path1 == nil {
			fmt.Println("No path found or missing path for payload 1")
		} else {
			fmt.Printf("Path found with %d hops for payload 1:\n", hops)
			mesh.PrintMesh(startNode1, finalNode1, startNode2, finalNode2, path1, path2)
			fmt.Println("\nPayload at goal node 1:")
			mesh.PrintPayload(finalNode1.Payload)
		}
	}()

	go func() {
		defer wg.Done()
		var hops int
		path2, hops = mesh.StarAlgorithm(startNode2, finalNode2, path1, startNode1, finalNode1, 'B')
		if path2 == nil {
			fmt.Println("No path found or missing path for payload 2")
		} else {
			fmt.Printf("Path found with %d hops for payload 2:\n", hops)
			mesh.PrintMesh(startNode1, finalNode1, startNode2, finalNode2, path1, path2)
			fmt.Println("\nPayload at goal node 2:")
			mesh.PrintPayload(finalNode2.Payload)
		}
	}()

	wg.Wait()

	fmt.Println("\nFinal payloads at goal nodes:")
	mesh.PrintPayload(finalNode1.Payload)
	mesh.PrintPayload(finalNode2.Payload)
}
