package main

import (
	"context"
	"log"
	"net"

	pb "distributed-task-manager/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTaskServiceServer
}

func (s *server) CreateTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	// Placeholder: Add logic to save task
	return &pb.TaskResponse{Id: "1", Status: "Created"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &server{})

	log.Println("Server is listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
