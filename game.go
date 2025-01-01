package main

import (
    "sync"

    "github.com/gorilla/websocket"
)

// player represents a connected player
type Player struct {
    ID   string
    Conn *websocket.Conn
    Score int
}

// GameRoom represents a single game instance or “room”
type GameRoom struct {
    RoomID      string
    Players     map[string]*Player
    CurrentQ    string // current question
    CurrentAns  string // correct answer
    mux         sync.Mutex
}

// NewGameRoom creates a new room
func NewGameRoom(roomID string) *GameRoom {
    return &GameRoom{
        RoomID:  roomID,
        Players: make(map[string]*Player),
    }
}

// AddPlayer adds a new player to the room
func (g *GameRoom) AddPlayer(player *Player) {
    g.mux.Lock()
    defer g.mux.Unlock()
    g.Players[player.ID] = player
}

// Broadcast sends a message to all players in the room
func (g *GameRoom) Broadcast(msg string) {
    g.mux.Lock()
    defer g.mux.Unlock()

    for _, p := range g.Players {
        p.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
    }
}
