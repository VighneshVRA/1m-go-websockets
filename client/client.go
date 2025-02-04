package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := "ws://localhost:8000/ws"

	// Read file content
	fileContent, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var wg sync.WaitGroup
	numClients := 100 // Number of WebSocket clients to simulate

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			// Connect to the server
			conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
			if err != nil {
				log.Printf("Client #%d: Connection error: %v", clientID, err)
				return
			}
			defer conn.Close()

			// Send file content
			err = conn.WriteMessage(websocket.TextMessage, fileContent)
			if err != nil {
				log.Printf("Client #%d: Send error: %v", clientID, err)
				return
			}

			// Receive response
			_, response, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Client #%d: Read error: %v", clientID, err)
				return
			}
			fmt.Printf("Client #%d received: %s\n", clientID, string(response))

			// Keep the connection alive for a while
			time.Sleep(5 * time.Second)
		}(i)
	}

	wg.Wait()
	log.Println("All clients finished")
}
