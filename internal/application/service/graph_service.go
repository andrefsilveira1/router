package services

import (
	entities "andrefsilveira1/router/internal/domain/entity"
	"fmt"
)

func XYRoute(graph *entities.Graph, startId, destId int) []int {
	startNode, startExists := graph.GetNode(startId)
	destNode, destExists := graph.GetNode(destId)

	if !startExists || !destExists {
		fmt.Println("Start or destination node does not exist.")
		return nil
	}

	path := []int{startNode.Id}
	currentNode := startNode

	// First route along X dimension
	for currentNode.X != destNode.X {
		nextNodeId := -1
		for adjId := range currentNode.Adjacent {
			adjNode := graph.Nodes[adjId]
			if (currentNode.X < destNode.X && adjNode.X > currentNode.X) ||
				(currentNode.X > destNode.X && adjNode.X < currentNode.X) {
				nextNodeId = adjId
				break
			}
		}
		if nextNodeId == -1 {
			fmt.Println("No valid path found in X direction.")
			return nil
		}
		currentNode = graph.Nodes[nextNodeId]
		path = append(path, currentNode.Id)
	}

	// Then route along Y dimension
	for currentNode.Y != destNode.Y {
		nextNodeId := -1
		for adjId := range currentNode.Adjacent {
			adjNode := graph.Nodes[adjId]
			if (currentNode.Y < destNode.Y && adjNode.Y > currentNode.Y) ||
				(currentNode.Y > destNode.Y && adjNode.Y < currentNode.Y) {
				nextNodeId = adjId
				break
			}
		}
		if nextNodeId == -1 {
			fmt.Println("No valid path found in Y direction.")
			return nil
		}
		currentNode = graph.Nodes[nextNodeId]
		path = append(path, currentNode.Id)
	}

	return path
}
