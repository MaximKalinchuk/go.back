package ws

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"go.back/pkg/customerror"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var usersConnections = make(map[string][]*websocket.Conn)
var usersConnectionsMu sync.RWMutex

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

	usersConnectionsMu.Lock()
	usersConnections[userId] = append(usersConnections[userId], conn)
	usersConnectionsMu.Unlock()

	logrus.Printf("[WebSocket] Пользователь %s успешно подключился", userId)
	logrus.Printf("[WebSocket] Количество соединений пользователя %s = %d", userId, len(usersConnections[userId]))

	defer func() {
		conn.Close()

		usersConnectionsMu.Lock()
		removeUserConnection(userId, conn)
		usersConnectionsMu.Unlock()

		log.Printf("[WebSocket] Пользователь %s отключился", userId)
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Println(err)
			break
		}
		logrus.Printf("Сообщение от %s: %s", userId, message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			logrus.Println(err)
			break
		}
	}
}

func removeUserConnection(userId string, conn *websocket.Conn) {
	conns := usersConnections[userId]

	for i, c := range conns {
		if c == conn {
			usersConnections[userId] = append(conns[:i], conns[i+1:]...)
			break
		}
	}

	if len(usersConnections[userId]) == 0 {
		delete(usersConnections, userId)
	}
}
