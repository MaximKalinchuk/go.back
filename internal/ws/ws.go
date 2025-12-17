package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	wsdto "go.back/internal/ws/dto"
	wshandlers "go.back/internal/ws/handlers"
	"go.back/pkg/customerror"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clientsConnections = make(map[string][]*websocket.Conn)
var clientsConnectionsMu sync.RWMutex

func HandleConnections(c *gin.Context) {

	var userId = c.GetString("userId")

	if userId == "" {
		customerror.HandleHTTPError(c, customerror.Unauthorized)
		return
	}

	logrus.Printf("[WebSocket] Попытка подключения пользователя %s", userId)

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		logrus.Fatal(err)
	}

	clientsConnectionsMu.Lock()
	clientsConnections[userId] = append(clientsConnections[userId], conn)
	clientsConnectionsMu.Unlock()

	logrus.Printf("[WebSocket] Пользователь %s успешно подключился", userId)
	logrus.Printf("[WebSocket] Количество соединений пользователя %s = %d", userId, len(clientsConnections[userId]))

	defer func() {
		conn.Close()

		clientsConnectionsMu.Lock()
		removeUserConnection(userId, conn)
		clientsConnectionsMu.Unlock()

		log.Printf("[WebSocket] Пользователь %s отключился", userId)
		logrus.Printf("[WebSocket] Количество соединений пользователя %s = %d", userId, len(clientsConnections[userId]))
	}()

	for {
		var parsedMessage wsdto.WsMessage

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Println(err)
			break
		}
		err = json.Unmarshal(message, &parsedMessage)

		if err != nil {
			logrus.Println(err)
			break
		}

		switch parsedMessage.Type {
		case "coordinates":
			wshandlers.HandleCoordinates(conn, parsedMessage.Data)
		default:
			logrus.Printf("[WebSocket] Неизветное событие: %s", parsedMessage.Type)
		}

		if err := conn.WriteMessage(messageType, message); err != nil {
			logrus.Println(err)
			break
		}
	}
}

func removeUserConnection(userId string, conn *websocket.Conn) {
	conns := clientsConnections[userId]

	for i, c := range conns {
		if c == conn {
			clientsConnections[userId] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	if len(clientsConnections[userId]) == 0 {
		delete(clientsConnections, userId)
	}
}
