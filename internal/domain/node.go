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
	Payload  map[string]string
}

// Em alguns algoritmos isso é chamado de "heurística". Pra mim, faz mais sentido chamar de "custo" ou "score"
func Cost(a, b *Node) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
