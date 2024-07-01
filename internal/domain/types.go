package domain

import (
	"container/heap"
	"fmt"
	"math"
)

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

func (m *Mesh) GetNode(x, y int) (*Node, bool) {
	pos := y*m.Size + x
	node, exists := m.Nodes[pos]
	return node, exists
}

func (g *Mesh) Disable(x, y int) {
	pos := y*g.Size + x
	if node, exists := g.Nodes[pos]; exists {
		node.Blocked = true
	}
}

func (m *Mesh) Enable(x, y int) {
	pos := y*m.Size + x
	if node, exists := m.Nodes[pos]; exists {
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

func Cost(a, b *Node) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

func (m *Mesh) StarAlgorithm(start, goal *Node) ([]*Node, int) {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &Queue{node: start, priority: 0})

	cameFrom := make(map[int]*Node)
	gScore := make(map[int]int)
	fScore := make(map[int]int)

	for id := range m.Nodes {
		gScore[id] = math.MaxInt32
		fScore[id] = math.MaxInt32
	}
	gScore[start.Id] = 0
	fScore[start.Id] = Cost(start, goal)

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Queue).node

		if current == goal {
			path := []*Node{}
			for current != nil {
				path = append([]*Node{current}, path...)
				current = cameFrom[current.Id]
			}
			return path, gScore[goal.Id]
		}

		for _, neighbor := range current.Adjacent {
			if neighbor.Blocked {
				continue
			}
			tentativeGScore := gScore[current.Id] + 1

			if tentativeGScore < gScore[neighbor.Id] {
				cameFrom[neighbor.Id] = current
				gScore[neighbor.Id] = tentativeGScore
				fScore[neighbor.Id] = tentativeGScore + Cost(neighbor, goal)
				heap.Push(&openSet, &Queue{node: neighbor, priority: fScore[neighbor.Id]})
			}
		}
	}

	return nil, -1
}

func (m *Mesh) Print(start, goal *Node, path []*Node) {
	size := m.Size
	mesh := make([][]rune, size)
	for y := 0; y < size; y++ {
		mesh[y] = make([]rune, size)
		for x := 0; x < size; x++ {
			node, _ := m.GetNode(x, y)
			if node.Blocked {
				mesh[y][x] = 'x'
			} else {
				mesh[y][x] = '.'
			}
		}
	}

	if start != nil {
		mesh[start.Y][start.X] = 'S'
	}
	if goal != nil {
		mesh[goal.Y][goal.X] = 'E'
	}
	for _, node := range path {
		if node != start && node != goal {
			mesh[node.Y][node.X] = '*'
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%c ", mesh[y][x])
		}
		fmt.Println()
	}
}
