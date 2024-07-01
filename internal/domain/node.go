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
