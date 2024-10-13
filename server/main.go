package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "distributed-task-manager/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedTaskServiceServer
	mu       sync.Mutex
	tasks    map[string]*pb.Task // In-memory storage using a hashmap for task details
	taskList []*pb.Task          // Slice to maintain task order for streaming
}

func (s *server) CreateTask(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", len(s.tasks)+1)
	task := &pb.Task{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	s.tasks[id] = task
	s.taskList = append(s.taskList, task)

	return &pb.TaskResponse{Id: id, Status: "Created"}, nil
}

func (s *server) GetTaskStatus(ctx context.Context, req *pb.TaskStatusRequest) (*pb.TaskStatusResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[req.Id]
	if !exists {
		return nil, fmt.Errorf("task not found")
	}

	return &pb.TaskStatusResponse{Status: "In Progress"}, nil // Placeholder status
}

func (s *server) RequestTask(req *pb.Empty, stream pb.TaskService_RequestTaskServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.taskList {
		if err := stream.Send(task); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := &server{
		tasks: make(map[string]*pb.Task),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, s)

	log.Println("Server is listening on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
