package domain

import (
	"math"
)

type Node struct {
	Id       int
	X        int
	Y        int
	Blocked  bool
	Adjacent map[int]*Node
}

func Cost(a, b *Node) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
