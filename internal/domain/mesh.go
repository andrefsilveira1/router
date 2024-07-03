package domain

import (
	"container/heap"
	"fmt"
	"math"
	"time"
)

type Mesh struct {
	Nodes map[int]*Node
	Size  int
}

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

func (m *Mesh) SetPayload(x, y int, payload map[string]string) {
	pos := y*m.Size + x
	if node, exists := m.Nodes[pos]; exists {
		node.Payload = payload
	}
}

func (m *Mesh) PrintPayload(payload map[string]string) {
	for key, value := range payload {
		fmt.Printf("Payload => %s: %s\n", key, value)
	}
}

func (m *Mesh) StarAlgorithm(start, goal *Node, pathSymbol rune) ([]*Node, int) {
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

	var path []*Node

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Queue).node

		if current == goal {
			path = []*Node{}
			for current != nil {
				path = append([]*Node{current}, path...)
				current = cameFrom[current.Id]
			}
			return path, gScore[goal.Id]
		}

		for _, neighbor := range current.Adjacent {
			neighbor.mu.Lock()
			if neighbor.Blocked {
				neighbor.mu.Unlock()
				continue
			}
			tentativeGScore := gScore[current.Id] + 1

			if tentativeGScore < gScore[neighbor.Id] {
				cameFrom[neighbor.Id] = current
				neighbor.Payload = current.Payload
				gScore[neighbor.Id] = tentativeGScore
				fScore[neighbor.Id] = tentativeGScore + Cost(neighbor, goal)
				heap.Push(&openSet, &Queue{node: neighbor, priority: fScore[neighbor.Id]})
			}
			neighbor.mu.Unlock()
		}

		// Add the current node to the path to visualize the progress
		path = append(path, current)

		// Print the current state of the mesh
		m.PrintMesh(start, goal, pathSymbol, path)
		time.Sleep(100 * time.Millisecond) // Sleep to visualize the steps
		fmt.Print("\033[H\033[2J")
	}

	return nil, -1
}

func (m *Mesh) PrintMesh(start, goal *Node, pathSymbol rune, path []*Node) {
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

	// Mark the path
	for _, node := range path {
		if node != start && node != goal {
			mesh[node.Y][node.X] = pathSymbol
		}
	}

	// Mark the start and goal nodes
	if start != nil {
		mesh[start.Y][start.X] = 'S'
	}
	if goal != nil {
		mesh[goal.Y][goal.X] = 'E'
	}

	// Print the updated matrix
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Printf("%c ", mesh[y][x])
		}
		fmt.Println()
	}
}
