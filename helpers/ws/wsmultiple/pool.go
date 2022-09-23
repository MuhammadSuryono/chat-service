package wsmultiple

import (
	"encoding/json"
	"fmt"
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
			//userId, _ := strconv.Atoi(client.ID)
			//conversations.AddNewUserOnline(int64(userId))
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
				var resp map[string]interface{}
				json.Unmarshal([]byte(message.Body), &resp)
				//iUser := resp["user"].(float64)
				//conversations.RemoveNewUserOnline(int64(iUser))
			}

			if message.Type == 1001 {
				var resp map[string]interface{}
				json.Unmarshal([]byte(message.Body), &resp)
				//iUser := resp["user"].(float64)
				//conversations.AddUserOnlineHistory(int64(iUser))
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
