package domain

import (
	"math"
	"sync"
)

type Flit struct {
	Id 					int
	Payload map[string]string
	Source      *Node
	Destination *Node
}

type Node struct {
	Id       int
	X        int
	Y        int
	Blocked  bool
	Adjacent map[int]*Node
	Payload  map[string]string
	Buffer	 []*Flit 
	mu       sync.Mutex
}

// Em alguns algoritmos isso é chamado de "heurística". Pra mim, faz mais sentido chamar de "custo" ou "score"
func Cost(a, b *Node) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func (n *Node) SendFlit(flit *Flit) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Buffer = append(n.Buffer, flit)
}

func (n *Node) ReceiveFlit() *Flit {
	n.mu.Lock()
	defer n.mu.Unlock()
	if len(n.Buffer) == 0 {
		return nil
	}
	flit := n.Buffer[0]
	n.Buffer = n.Buffer[1:]
	return flit
}

func (n *Node) ClearBuffer() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Buffer = nil
}
