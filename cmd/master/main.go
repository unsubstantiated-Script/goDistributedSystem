package main

import (
	"goDistributedSystem/internal/master"
	"goDistributedSystem/pkg/pb"
	"log"
	"net"

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

	log.Println("master listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
