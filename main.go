package main

import (
	"chat-service/Exception"
	"chat-service/config/database"
	"chat-service/controller"
	"chat-service/helpers/ws"
	"chat-service/models/tables"
	"chat-service/server"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		panic(fmt.Sprintf("Env file notfound on directory %v", errLoadEnv))
	}
	database.Init()
	tables.Migrate()

	defer Exception.GetError()
	go ws.H.Run()

	router := server.ConfigServer()
	controller.Routes(router)
	if Exception.ErrorMessage != nil {
		log.Fatalf("Can't start server on error: %v", Exception.ErrorMessage)
	}

	server.RunServer()

}
