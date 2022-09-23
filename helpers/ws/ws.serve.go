package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

// ServeWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request, roomId string, userId int64) {
	logrus.Info("ROOM ID Connection", roomId)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId, userId}
	H.register <- s
	go s.writePump()
	go s.readPump()
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		fmt.Println("Defer close unregister")
		H.unregister <- s
		err := c.ws.Close()
		if err != nil {
			fmt.Sprintf("Error close %v", err.Error())
			return
		}
	}()
	c.ws.SetReadLimit(maxMessageSize)
	err := c.ws.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		logrus.Error("Websocket set read deadline error")
	}
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		logrus.Info("Broadcast message", msg)
		m := message{msg, s.room}
		H.broadcast <- m
	}
}

// write a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	err := c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	if err != nil {
		return err
	}
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			logrus.Info("Write pump message ", string(message))
			if !ok {
				logrus.Error("Write pump message not OK ", string(message))
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
