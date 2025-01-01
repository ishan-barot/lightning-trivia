package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
    // simple route for WebSocket
    http.HandleFunc("/ws", handleConnections)

    // start server on port 8080
    log.Println("Server started on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

// handleConnections upgrades the HTTP request to a WebSocket and
// adds the player to a game room, etc.
func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Error upgrading: %v", err)
        return
    }
    defer conn.Close()

    // for demonstration, just echo messages:
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Printf("Read error: %v", err)
            break
        }
        log.Printf("Received: %s", msg)

        // echo back to client
        err = conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            log.Printf("Write error: %v", err)
            break
        }
    }
}
