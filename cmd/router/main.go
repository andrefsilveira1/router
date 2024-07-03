package main

import (
	"andrefsilveira1/router/internal/application/service"
	"andrefsilveira1/router/internal/domain"
	"fmt"
	"sync"
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

		// Enviar flit para startNode
		flit1 := &domain.Flit{
			ID:        1,
			Source:    startNode,
			Destination: finalNode,
			Payload:   payload,
		}
		startNode.SendFlit(flit1)
	
		// Enviar flit para startNode2
		flit2 := &domain.Flit{
			ID:        2,
			Source:    startNode2,
			Destination: finalNode2,
			Payload:   payload2,
		}
		startNode2.SendFlit(flit2)

	// mesh.SetPayload(startNode.X, startNode.Y, payload)
	// mesh.SetPayload(startNode2.X, startNode2.Y, payload2)

	// mesh.PrintPayload(startNode.Payload)
	// mesh.PrintPayload(startNode.Payload)

	fmt.Printf("\n")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			flit := startNode.ReceiveFlit()
			if flit == nil {
				break
			}
			path, hops := mesh.StarAlgorithm(flit.Source, flit.Destination)
			if path == nil {
				fmt.Printf("No path found or missing path for payload %d\n", flit.ID)
			} else {
				fmt.Printf("Path found with %d hops for payload %d:\n", hops, flit.ID)
				mesh.Print(flit.Source, flit.Destination, path)
				fmt.Printf("\nPayload at goal node %d:\n", flit.Destination.Id)
				mesh.PrintPayload(flit.Destination.Payload)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			flit := startNode2.ReceiveFlit()
			if flit == nil {
				break
			}
			path, hops := mesh.StarAlgorithm(flit.Source, flit.Destination)
			if path == nil {
				fmt.Printf("No path found or missing path for payload %d\n", flit.ID)
			} else {
				fmt.Printf("Path found with %d hops for payload %d:\n", hops, flit.ID)
				mesh.Print(flit.Source, flit.Destination, path)
				fmt.Printf("\nPayload at goal node %d:\n", flit.Destination.Id)
				mesh.PrintPayload(flit.Destination.Payload)
			}
		}
	}()


	wg.Wait()

	// path, hops := mesh.StarAlgorithm(startNode, finalNode)
	// if path == nil {
	// 	fmt.Println("No path found or missing path")
	// } else {
	// 	fmt.Printf("Path found with %d hops:\n", hops) // Buscar esse significado de "hops"
	// 	mesh.Print(startNode, finalNode, path)
	// }

	mesh.PrintPayload(finalNode.Payload)
	mesh.PrintPayload(finalNode2.Payload)
}
