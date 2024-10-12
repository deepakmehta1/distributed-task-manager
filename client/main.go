package main

import (
	"context"
	"log"
	"time"

	pb "distributed-task-manager/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.DialContext(context.Background(), "localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateTask(ctx, &pb.TaskRequest{Name: "Sample Task", Description: "This is a sample task"})
	if err != nil {
		log.Fatalf("Could not create task: %v", err)
	}

	log.Printf("Task created: %s (status: %s)", res.Id, res.Status)
}
