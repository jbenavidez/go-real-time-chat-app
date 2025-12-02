package main

import (
	"broker/repository"
	dbrepo "broker/repository/db_repo"
	"fmt"
	"log"
)

const (
	gRpcPort = "50001"
)

type Config struct {
	DSN string
	DB  repository.DatabaseRepo
}

func main() {
	fmt.Println("starting  broker service...")
	app := Config{}
	conn := app.connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}
	//set up db
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	//set up helper
	NewGrpcHelper(&app)
	// Set up gRPC
	app.gRPCListenAndServe()

}
