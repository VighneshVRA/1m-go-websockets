// package main

// import (
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"net/http"
// )

// func ws(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade connection
// 	upgrader := websocket.Upgrader{}
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}
// 	// Read messages from socket
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			conn.Close()
// 			return
// 		}
// 		log.Printf("msg: %s", string(msg))
// 	}
// }

// func main() {
// 	http.HandleFunc("/", ws)
// 	if err := http.ListenAndServe(":8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

//////////////////////////////v2

// package main

// import (
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof"

// 	"syscall"
// )

// var count int64

// func ws(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade connection
// 	upgrader := websocket.Upgrader{}
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}

	 
// 	defer func() {
// 		conn.Close()
// 	}()

// 	// Read messages from socket
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			return
// 		}
// 		log.Printf("msg: %s", string(msg))
// 	}
// }

// func main() {
// 	// Increase resources limitations
// 	var rLimit syscall.Rlimit
// 	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}
// 	rLimit.Cur = rLimit.Max
// 	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}

// 	// Enable pprof hooks
// 	go func() {
// 		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
// 			log.Fatalf("Pprof failed: %v", err)
// 		}
// 	}()

// 	http.HandleFunc("/", ws)
// 	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// //////////////////////////////v3
// package main

// import (
// 	"github.com/gorilla/websocket"
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof"
// 	"syscall"
// )

// var count int64

// func ws(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade connection 
// 	upgrader := websocket.Upgrader{}
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade failed")
// 		return
// 	}

// 	log.Println("New WebSocket connection established")
	
// 	defer func() { 
// 		conn.Close() 
// 	}()

// 	// Read messages from socket 
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println("WebSocket read error:", err)
// 			return
// 		}
// 		log.Printf("msg: %s", string(msg)) 
// 	}
// }

// func main() {
// 	// Increase resources limitations 
// 	var rLimit syscall.Rlimit
// 	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}
// 	rLimit.Cur = rLimit.Max
// 	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}

// 	// Enable pprof hooks 
// 	go func() {
// 		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
// 			log.Fatalf("Pprof failed: %v", err)
// 		}
// 	}()

// 	log.Println("WebSocket server starting on port 8000")
// 	http.HandleFunc("/", ws)
// 	if err := http.ListenAndServe(":8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
////////////////////////////v4
// package main

// import (
// 	"github.com/gobwas/ws"
// 	"github.com/gobwas/ws/wsutil"
// 	"log"
// 	"net/http"
// 	_ "net/http/pprof"
// 	"syscall"
// )

// var epoller *epoll

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade connection
// 	conn, _, _, err := ws.UpgradeHTTP(r, w)
// 	if err != nil {
// 		return
// 	}
// 	if err := epoller.Add(conn); err != nil {
// 		log.Printf("Failed to add connection %v", err)
// 		conn.Close()
// 	}
// }

// func main() {
// 	// Increase resources limitations
// 	var rLimit syscall.Rlimit
// 	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}
// 	rLimit.Cur = rLimit.Max
// 	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
// 		panic(err)
// 	}

// 	// Enable pprof hooks
// 	go func() {
// 		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
// 			log.Fatalf("pprof failed: %v", err)
// 		}
// 	}()

// 	// Start epoll
// 	var err error
// 	epoller, err = MkEpoll()
// 	if err != nil {
// 		panic(err)
// 	}

// 	go Start()

// 	http.HandleFunc("/", wsHandler)
// 	if err := http.ListenAndServe("0.0.0.0:8000", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func Start() {
// 	for {
// 		connections, err := epoller.Wait()
// 		if err != nil {
// 			log.Printf("Failed to epoll wait %v", err)
// 			continue
// 		}
// 		for _, conn := range connections {
// 			if conn == nil {
// 				break
// 			}
// 			if _, _, err := wsutil.ReadClientData(conn); err != nil {
// 				if err := epoller.Remove(conn); err != nil {
// 					log.Printf("Failed to remove %v", err)
// 				}
// 				conn.Close()
// 			} else {
// 				// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
// 				//log.Printf("msg: %s", string(msg))
// 			}
// 		}
// 	}
// }

 