package entities

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
