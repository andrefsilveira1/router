package domain

type Node struct {
	Id       int
	X        int
	Y        int
	Blocked  bool
	Adjacent map[int]*Node
}

type Mesh struct {
	Nodes map[int]*Node
	Size  int
}

func (g *Mesh) Disable(x, y int) {
	id := y*g.Size + x
	if node, exists := g.Nodes[id]; exists {
		node.Blocked = true
	}
}

func (g *Mesh) Enable(x, y int) {
	id := y*g.Size + x
	if node, exists := g.Nodes[id]; exists {
		node.Blocked = false
	}
}
