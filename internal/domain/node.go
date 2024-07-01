package entity

type Node struct {
	Id       int
	X        int
	Y        int
	Blocked  bool
	Adjacent map[int]*Node
}
