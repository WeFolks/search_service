package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/WeFolks/search_service/api"
	"github.com/WeFolks/search_service/elasticsearch"
	"github.com/WeFolks/search_service/grpc"
	"github.com/WeFolks/search_service/middleware"
	g "google.golang.org/grpc"
)

func main() {

	go func() {
		client, err := elasticsearch.GetESClient()

		if err != nil {
			log.Fatal("Elastic Search client can't be setup", err)
			return
		}

		http.Handle("/search", middleware.LogReq(api.GetData(client)))
		err = http.ListenAndServe(":8080", nil)

		if err != nil {
			log.Fatal("Error starting server:", err)
		}
		fmt.Printf("HTTP server hosted on port 8080")
	}()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := grpc.Server{}

	grpcServer := g.NewServer()

	grpc.RegisterSearchServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
