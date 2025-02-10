// package main

// import (
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"time"
// )

// func main() {
// 	url := "ws://localhost:8000"

// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		log.Fatal("Connection error:", err)
// 	}
// 	defer conn.Close()

// 	// Send messages
// 	go func() {

// 			msg := "Hello WebSocket " + time.Now().String()
// 			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
// 			if err != nil {
// 				log.Println("Write error:", err)
// 				return
// 			}

// 	}()

// 	// Receive messages
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("Read error:", err)
// 			return
// 		}
// 		log.Printf("Received: %s", msg)
// 	}
// }

// ///////////////////// V2
// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"sync/atomic"
// 	// "time"

// 	"github.com/gorilla/websocket"
// )

// func connectAndSend(id int, wg *sync.WaitGroup, successCount *uint64) {
// 	defer wg.Done()

// 	url := "ws://54.198.237.81:8000"
// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		log.Printf("Client %d connection error: %v", id, err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Increment successful connection count
// 	atomic.AddUint64(successCount, 1)

// 	msg := []byte(fmt.Sprintf("Client %d connected", id))
// 	err = conn.WriteMessage(websocket.TextMessage, msg)
// 	if err != nil {
// 		log.Printf("Client %d write error: %v", id, err)
// 	}
// }

// func main() {
// 	clientCount := 10000
// 	var wg sync.WaitGroup
// 	var successCount uint64

// 	for i := 0; i < clientCount; i++ {
// 		wg.Add(1)
// 		go connectAndSend(i, &wg, &successCount)

// 	}

// 	wg.Wait()
// 	fmt.Printf("Successful connections: %d out of %d\n", successCount, clientCount)
// }
///////////////////////// V3 

// package main

// import (
// 	"fmt"
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"sync" 
// 	"time"
// )

// func connectAndSend(wg *sync.WaitGroup, successCount *int) {
// 	defer wg.Done()

// 	url := "ws://54.198.237.81:8000"
// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		return
// 	}
// 	defer conn.Close()

// 	// Increment successful connection count
// 	*successCount++

// 	msg := []byte("Client connected")
// 	err = conn.WriteMessage(websocket.TextMessage, msg)
// 	if err != nil {
// 		log.Println("Write error")
// 	}
// }

// func main() {
// 	clientCount := 1000
// 	var wg sync.WaitGroup
// 	successCount := 0

// 	for i := 0; i < clientCount; i++ {
// 		wg.Add(1)
// 		go connectAndSend(&wg, &successCount)
// 		time.Sleep(100 * time.Millisecond) 
// 	}

// 	wg.Wait()
// 	fmt.Printf("Successful connections: %d out of %d\n", successCount, clientCount)
// }

///////////////////////////////// Data sending Version 
// package main

// import (
// 	"fmt"
// 	"github.com/gorilla/websocket"
// 	"io/ioutil"
// 	"log"
// 	"time"
// )

// func main() {
// 	url := "ws://54.198.237.81:8000"

// 	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
// 	if err != nil {
// 		log.Fatal("Connection error:", err)
// 	}
// 	defer conn.Close()

// 	// Read file
// 	fileContent, err := ioutil.ReadFile("example.txt")
// 	if err != nil {
// 		log.Fatal("File read error:", err)
// 	}

// 	// Send file
// 	err = conn.WriteMessage(websocket.BinaryMessage, fileContent)
// 	if err != nil {
// 		log.Fatal("Send error:", err)
// 	}

// 	// Wait for server response
// 	_, response, err := conn.ReadMessage()
// 	if err != nil {
// 		log.Fatal("Read error:", err)
// 	}
// 	fmt.Println("Server response:", string(response))

// 	// Keep connection open
// 	for {
// 		time.Sleep(10 * time.Second)
// 	}
// }

// ///////////////////////////// Data sending V2




//////////////////////// Data sending Version - looped
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://54.198.237.81:8000"

	// Read the file content once
	fileContent, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal("File read error:", err)
	}

	// Wait group to manage goroutines
	var wg sync.WaitGroup

	// Create 100 connections
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(connectionID int) {
			defer wg.Done()

			// Establish connection
			conn, _, err := websocket.DefaultDialer.Dial(url, nil)
			if err != nil {
				log.Printf("Connection error (Connection #%d): %v", connectionID, err)
				return
			}
			defer conn.Close()

			// Send file content
			if err := conn.WriteMessage(websocket.BinaryMessage, fileContent); err != nil {
				log.Printf("Send error (Connection #%d): %v", connectionID, err)
				return
			}

			// Wait for server response
			_, response, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error (Connection #%d): %v", connectionID, err)
				return
			}
			fmt.Printf("Server response (Connection #%d): %s\n", connectionID, string(response))

			// Keep connection open for demonstration purposes
			time.Sleep(10 * time.Second)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
