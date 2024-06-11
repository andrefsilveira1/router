package services

import (
	entities "andrefsilveira1/router/internal/domain/entity"
	"container/heap"
	"math"
)

type PriorityQueue []*Item

type Item struct {
	node     *entities.Node
	priority int
	index    int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, node *entities.Node, priority int) {
	item.node = node
	item.priority = priority
	heap.Fix(pq, item.index)
}

func Dijkstra(graph *entities.Graph, startId int) map[int]int {
	dist := make(map[int]int)
	for id := range graph.Nodes {
		dist[id] = math.MaxInt32
	}
	dist[startId] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{
		node:     graph.Nodes[startId],
		priority: 0,
	})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		currentNode := item.node
		currentDistance := dist[currentNode.Id]

		for adjacentId, weight := range currentNode.Adjacent {
			alt := currentDistance + weight
			if alt < dist[adjacentId] {
				dist[adjacentId] = alt
				heap.Push(pq, &Item{
					node:     graph.Nodes[adjacentId],
					priority: alt,
				})
			}
		}
	}

	return dist
}
