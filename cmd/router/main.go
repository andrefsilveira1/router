package main

import (
	"andrefsilveira1/router/internal/domain"
	"andrefsilveira1/router/internal/application/service"
	"fmt"
)

var payload = map[string]string{
	"message":   "This is the payload message",
	"timestamp": "2024-07-01T12:00:00Z",
}

var payload2 = map[string]string{
	"message":   "This is the second payload message",
	"timestamp": "2024-08-01T12:00:00Z",
}

func main() {
	// Criando malha 8x8
	mesh := service.CreateMesh()

	// Desabilitando nós
	service.DisableNodes(mesh)

	startNode, _ := mesh.GetNode(0, 0)
	finalNode, _ := mesh.GetNode(7, 7)

	startNode2, _ := mesh.GetNode(0, 1)
	finalNode2, _ := mesh.GetNode(7, 6)

	mesh.SetPayload(startNode.X, startNode.Y, payload)
	mesh.SetPayload(startNode2.X, startNode2.Y, payload2)

	mesh.PrintPayload(startNode.Payload)
	mesh.PrintPayload(startNode.Payload)

	fmt.Printf("\n")

	done := make(chan bool, 2)

	go func() {
		path, hops := mesh.StarAlgorithm(startNode, finalNode)
		if path == nil {
			fmt.Println("No path found or missing path for payload 1")
		} else {
			fmt.Printf("Path found with %d hops for payload 1:\n", hops)
			mesh.Print(startNode, finalNode, path)
			fmt.Println("\nPayload at goal node 1:")
			mesh.PrintPayload(finalNode.Payload)
		}
		done <- true
	}()

	go func() {
		path, hops := mesh.StarAlgorithm(startNode2, finalNode2)
		if path == nil {
			fmt.Println("No path found or missing path for payload 2")
		} else {
			fmt.Printf("Path found with %d hops for payload 2:\n", hops)
			mesh.Print(startNode2, finalNode2, path)
			fmt.Println("\nPayload at goal node 2:")
			mesh.PrintPayload(finalNode2.Payload)
		}
		done <- true
	}()

	<-done
	<-done

	flit1 := &domain.Flit{
		Source:      startNode,
		Destination: finalNode,
		Payload:     payload,
	}

	flit2 := &domain.Flit{
		Source:      startNode2,
		Destination: finalNode2,
		Payload:     payload2,
	}

	startNode.SendFlit(flit1)
	startNode2.SendFlit(flit2)

	go func() {
		for {
			flit := finalNode.ReceiveFlit()
			if flit == nil {
				break
			}
			fmt.Printf("Flit received at final node 1: %v\n", flit.Payload)
		}
		done <- true
	}()

	go func() {
		for {
			flit := finalNode2.ReceiveFlit()
			if flit == nil {
				break
			}
			fmt.Printf("Flit received at final node 2: %v\n", flit.Payload)
		}
		done <- true
	}()

	<-done
	<-done

	mesh.PrintPayload(finalNode.Payload)
	mesh.PrintPayload(finalNode2.Payload)
}
