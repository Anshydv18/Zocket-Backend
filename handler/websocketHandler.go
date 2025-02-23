package handler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[string]*websocket.Conn)
var mu sync.Mutex

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("webscoket upgrade failed")
		return
	}

	userID := c.Query("user_id")
	if userID == "" {
		fmt.Println("UserId is missing")
		return
	}

	mu.Lock()
	clients[userID] = conn
	mu.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("websocket  disconnect for user:", userID)
			mu.Lock()
			delete(clients, userID)
			mu.Unlock()
			break
		}
	}
}

func NotifyUser(userID, message string) {
	mu.Lock()
	defer mu.Unlock()

	if conn, ok := clients[userID]; ok {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
