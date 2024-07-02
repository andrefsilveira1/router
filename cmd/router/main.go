package main

import (
	"andrefsilveira1/router/internal/application/service"
	"fmt"
)

func main() {
	// Criando malha 8x8
	mesh := service.CreateMesh()

	// Desaabilitando nós
	mesh.Disable(3, 3)
	mesh.Disable(4, 4)
	mesh.Disable(2, 5)
	mesh.Disable(5, 2)

	startNode, _ := mesh.GetNode(0, 0)
	goalNode, _ := mesh.GetNode(7, 7)

	path, hops := mesh.StarAlgorithm(startNode, goalNode)
	if path == nil {
		fmt.Println("No path found or missing path")
	} else {
		fmt.Printf("Path found with %d hops:\n", hops) // Buscar esse significado de "hops"
		mesh.Print(startNode, goalNode, path)
	}
}
