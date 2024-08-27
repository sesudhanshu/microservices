package main

import (
	"log"

	"github.com/sesudhanshu/Go_Microservice/internal/database"
	"github.com/sesudhanshu/Go_Microservice/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialise Database client %s", err)
	}

	server := server.NewEchoServer(db)
	if err := server.Start(); err != nil {
		log.Fatalf(err.Error())
	}
}
