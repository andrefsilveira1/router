package service

import (
	"fmt"
	"strings"
	"andrefsilveira1/router/internal/domain"
)

var size = 8 // Tamanho da malha será fixo aqwui

func CreateMesh() *domain.Mesh {
	mesh := &domain.Mesh{
		Nodes: make(map[int]*domain.Node),
		Size:  size,
	}
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			id := y*size + x
			mesh.Nodes[id] = &domain.Node{
				Id:       id,
				X:        x,
				Y:        y,
				Blocked:  false,
				Adjacent: make(map[int]*domain.Node),
			}
		}
	}
	for _, node := range mesh.Nodes {
		x, y := node.X, node.Y
		if x > 0 {
			node.Adjacent[(y*size)+(x-1)] = mesh.Nodes[(y*size)+(x-1)]
		}
		if x < size-1 {
			node.Adjacent[(y*size)+(x+1)] = mesh.Nodes[(y*size)+(x+1)]
		}
		if y > 0 {
			node.Adjacent[((y-1)*size)+x] = mesh.Nodes[((y-1)*size)+x]
		}
		if y < size-1 {
			node.Adjacent[((y+1)*size)+x] = mesh.Nodes[((y+1)*size)+x]
		}
	}

	return mesh
}

func BlockNodes(mesh *domain.Mesh, blockedNodes string) {
	pairs := strings.Split(blockedNodes, " ")
	for _, pair := range pairs {
		coords := strings.Split(pair, ",")
		if len(coords) != 2 {
			fmt.Println("Formato inválido, use 'x,y'")
			continue
		}
		var x, y int
		fmt.Sscanf(coords[0], "%d", &x)
		fmt.Sscanf(coords[1], "%d", &y)
		mesh.Disable(x, y)
	}
}
