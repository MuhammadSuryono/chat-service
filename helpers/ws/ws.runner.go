package ws

import (
	"chat-service/Exception"
	msgRepo "chat-service/repository/message"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

func (h *hub) Run() {
	//defer Exception.GetError()
	for {
		select {
		case s := <-h.register:
			fmt.Println("register", s.room)
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*connection]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			fmt.Println("unregister", s, h.rooms, h.register, h.unregister, h.broadcast)
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					close(s.conn.send)
					fmt.Println("Close connection", fmt.Sprintf("%v", s.conn))
					if len(connections) == 0 {
						//delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			connections := h.rooms[m.room]
			if strings.Contains(string(m.data), "CH-1000") {
				defer Exception.GetError()
				logrus.Info("Message saved ", string(m.data))
				var messageChat map[string]interface{}
				_ = json.Unmarshal(m.data, &messageChat)
				s := msgRepo.Repository.SaveNewMessageChat(messageChat["message"].(string), int64(messageChat["group"].(float64)), int64(messageChat["sender"].(float64)))
				m.data, _ = json.Marshal(s)
				if Exception.ErrorMessage != nil {
					logrus.Error(fmt.Sprintf("%v", Exception.ErrorMessage))
				}
			}

			for c := range connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
