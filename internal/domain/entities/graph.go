package entities

import (
	"fmt"
	"strings"
)

type Graph struct {
	Nodes map[int]*Node
}

func CreateGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
	}
}

func (g *Graph) AddNode(node *Node) {
	if _, existent := g.Nodes[node.Id]; !existent {
		g.Nodes[node.Id] = node
	}
}

func (g *Graph) GetNode(id int) (*Node, bool) {
	node, exists := g.Nodes[id]
	return node, exists
}

func (g *Graph) PrintGraphTree(startId int) {
	visited := make(map[int]bool)
	var builder strings.Builder

	var printTree func(node *Node, prefix string)
	printTree = func(node *Node, prefix string) {
		if node == nil || visited[node.Id] {
			return
		}
		visited[node.Id] = true
		builder.WriteString(fmt.Sprintf("%s%d\n", prefix, node.Id))
		newPrefix := prefix + "  "
		for adj, weight := range node.Adjacent {
			builder.WriteString(fmt.Sprintf("%s|-- %d (weight: %d)\n", newPrefix, adj, weight))
			printTree(g.Nodes[adj], newPrefix+"|  ")
		}
	}

	startNode, exists := g.GetNode(startId)
	if !exists {
		fmt.Println("Start node does not exist.")
		return
	}

	printTree(startNode, "")
	fmt.Println(builder.String())
}
