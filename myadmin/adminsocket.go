package myadmin

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

type client struct{} // Add more data to this type if needed

var clients = make(map[*websocket.Conn]client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)
func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = client{}
			fmt.Println("connection registered")

		case message := <-broadcast:
			fmt.Println("message received:", message)
			fmt.Println("message clients:", clients)

			// Send the message to all clients
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
					delete(clients, connection)
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func AdminSocketRoutes(app *fiber.App) {
	go runHub()
	adminSocket := app.Group("/ws/admin")
	adminSocket.Get("/shop/created", websocket.New(func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer func() {
			unregister <- c
			c.Close()
		}()
		// Register the client
		register <- c
		fmt.Println("client: ",clients)
		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return // Calls the deferred function, i.e. closes the connection on error
			}
			if messageType == websocket.TextMessage {
				broadcast <- string(message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}

	}))
}
