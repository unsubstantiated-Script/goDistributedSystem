package main

import (
	"goDistributedSystem/internal/master"
	"goDistributedSystem/pkg/pb"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	nodeServer := master.NewNodeServer()

	pb.RegisterNodeServiceServer(grpcServer, nodeServer)

	// Start HTTP API server on port 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", nodeServer.AddTaskHandler)

	go func() {
		log.Println("HTTP API listening on port 8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	log.Println("master gRPC listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
