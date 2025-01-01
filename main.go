package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/ws", handleConnections)

    log.Println("Server starting on :" + port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("Error upgrading: %v", err)
        return
    }
    defer conn.Close()

    // Simple echo loop
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Printf("Read error: %v", err)
            break
        }
        log.Printf("Received: %s", msg)

        err = conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            log.Printf("Write error: %v", err)
            break
        }
    }
}
