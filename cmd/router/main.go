package main

import (
	"andrefsilveira1/router/internal/application/service"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// Criando malha 8x8
	mesh := service.CreateMesh()

	// Desaabilitando nós
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Insira os nós bloqueados no formato 'x,y x,y ...': e depois aperte Enter")
	blockedNodes, _ := reader.ReadString('\n')
	blockedNodes = strings.TrimSpace(blockedNodes)
	service.BlockNodes(mesh, blockedNodes)

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
