package wsmultiple

import (
	"chat-service/Exception"
	msgRepo "chat-service/repository/message"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

var ConnectionPool *Pool

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user with ID " + client.ID + " joined"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 0, Body: "User " + client.ID + " Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			if message.Type == 1000 {
				defer Exception.GetError()
				var messageChat map[string]interface{}
				_ = json.Unmarshal([]byte(message.Body), &messageChat)
				s := msgRepo.Repository.SaveNewMessageChat(messageChat["message"].(string), int64(messageChat["group"].(float64)), int64(messageChat["sender"].(float64)))
				m, _ := json.Marshal(s)
				if Exception.ErrorMessage != nil {
					logrus.Error(fmt.Sprintf("%v", Exception.ErrorMessage))
				}

				message.Body = string(m)
			}

			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
