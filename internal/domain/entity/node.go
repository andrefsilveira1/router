package entity

import "fmt"

type Node struct {
	Id       int
	X        int
	Y        int
	Adjacent map[int]int
	Package  map[string]interface{}
}

func CreateNode(id, x, y int) *Node {
	return &Node{
		Id:       id,
		X:        x,
		Y:        y,
		Adjacent: make(map[int]int),
		Package:  map[string]interface{}{"node_id": id, "value": ""},
	}
}

func (n *Node) AppendAdjacent(adjacentId, weight int) {
	n.Adjacent[adjacentId] = weight
}

func (n *Node) SendPackage(pack map[string]interface{}) {
	n.Package = pack
}

func (n *Node) PrintPackageInfo() {
	fmt.Printf("Package info for Node %d:\n", n.Id)
	for key, value := range n.Package {
		fmt.Printf("  %s: %v\n", key, value)
	}
}
