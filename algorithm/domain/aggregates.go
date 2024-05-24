package domain

type Graph struct {
	nodes map[int]*Node
}

func CreateGraph() *Graph {
	return &Graph{
		nodes: make(map[int]*Node),
	}
}

func (g *Graph) AddNode(node *Node) {
	if _, existent := g.nodes[node.Id]; !existent {
		g.nodes[node.Id] = node
	}
}

func (g *Graph) GetNode(id int) (*Node, bool) {
	node, exists := g.nodes[id]
	return node, exists
}
