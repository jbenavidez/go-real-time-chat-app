package main

import (
	"fmt"
	"frontend/internal/config"
	"frontend/internal/handlers"
	"frontend/internal/render"
	pb "frontend/proto/generated"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":3000"

var app config.AppConfig

func main() {
	//set gRPC connection
	conn, err := grpc.Dial("broker-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	//set gRPC Client
	client := pb.NewChatMessagesServiceClient(conn)
	app.GRPCClient = client

	fmt.Println(fmt.Sprintf("Starting front-end on port %s", port))

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
