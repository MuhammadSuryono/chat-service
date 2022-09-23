package wsmultiple

import (
	"fmt"
	"net/http"
)

func ServeWs(w http.ResponseWriter, r *http.Request, clientId string) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		Conn: conn,
		Pool: ConnectionPool,
		ID:   clientId,
	}

	ConnectionPool.Register <- client
	client.Read()
}

func InitWsPool() {
	ConnectionPool = NewPool()
	go ConnectionPool.Start()
}
