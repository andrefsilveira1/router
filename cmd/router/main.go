package main

import (
	"andrefsilveira1/router/internal/application/service"
	"fmt"
)

var payload = map[string]string{
	"message":   "This is the payload message",
	"timestamp": "2024-07-01T12:00:00Z",
}

func main() {
	// Criando malha 8x8
	mesh := service.CreateMesh()

	// Desabilitando nós
	service.DisableNodes(mesh)

	startNode, _ := mesh.GetNode(0, 0)
	mesh.SetPayload(startNode.X, startNode.Y, payload)
	mesh.PrintPayload(startNode.Payload)

	fmt.Printf("\n")

	finalNode, _ := mesh.GetNode(7, 7)

	path, hops := mesh.StarAlgorithm(startNode, finalNode)
	if path == nil {
		fmt.Println("No path found or missing path")
	} else {
		fmt.Printf("Path found with %d hops:\n", hops) // Buscar esse significado de "hops"
		mesh.Print(startNode, finalNode, path)
	}

	mesh.PrintPayload(finalNode.Payload)
}
