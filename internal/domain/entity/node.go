package entity

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
