package entity

type Node struct {
	Id       int
	Adjacent map[int]int
}

func CreateNode(id int) *Node {
	return &Node{
		Id:       id,
		Adjacent: make(map[int]int),
	}
}

func (n *Node) AppendAdjacent(adjacentId, weight int) {
	n.Adjacent[adjacentId] = weight
}
