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

type Queue struct {
	node     *Node
	priority int
	index    int
}
type PriorityQueue []*Queue

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

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool {
	return q[i].priority < q[j].priority
}

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *PriorityQueue) Push(x interface{}) {
	n := len(*q)
	queue := x.(*Queue)
	queue.index = n
	*q = append(*q, queue)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	queue := old[n-1]
	queue.index = -1
	*q = old[0 : n-1]
	return queue
}
