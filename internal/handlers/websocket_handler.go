package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Agregar logica para permitir solo ciertos origenes
		return true // Para test cualquier origen es valido
	},
}

var clients = make(map[string]*websocket.Conn)

func Websocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("error al actualizar websocket: ", err.Error())
		return
	}
	defer conn.Close()

	clientID := c.Query("machine")
	clients[clientID] = conn

	log.Println(clientID)

	defer delete(clients, clientID)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error al leer mensaje:", err)
			break // Salir del bucle si hay un error
		}

		if messageType == websocket.TextMessage {
			fmt.Println("Mensaje recibido:", string(p))
			// Aquí puedes procesar el mensaje del cliente

			receiver := c.Query("receiver")

			if receiverConn, ok := clients[receiver]; ok {
				err := receiverConn.WriteMessage(websocket.TextMessage, p)
				if err != nil {
					log.Println("error al enviar mensaje: ", err.Error())
				}
			}
		}
	}
}
