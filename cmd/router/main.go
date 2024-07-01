package main

import (
	"andrefsilveira1/router/internal/application/service"
	"fmt"
)

func main() {
	mesh := service.CreateMesh()
	mesh.Disable(3, 3)
	mesh.Disable(4, 4)
	mesh.Disable(2, 5)
	mesh.Disable(5, 2)

	startNode, _ := mesh.GetNode(0, 0)
	goalNode, _ := mesh.GetNode(7, 7)

	path, hops := mesh.StarAlgorithm(startNode, goalNode)
	if path == nil {
		fmt.Println("No path found.")
	} else {
		fmt.Printf("Path found with %d hops:\n", hops)
		mesh.Print(startNode, goalNode, path)
	}
}
