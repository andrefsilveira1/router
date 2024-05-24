package domain

type Node struct {
	Id       int
	Adjacent map[int]int
}

func Append_node(id int) *Node {
	return &Node{
		Id:       id,
		Adjacent: make(map[int]int),
	}
}

func (n *Node) Append_adjacent(adjancentId, weight int) {
	n.Adjacent[adjancentId] = weight
}
