package main

import (
	"context"
	"customer-service/ent"
	"customer-service/ent/proto/entpb"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Initialize an ent client.
	client, err := ent.Open("postgres", "host=localhost port=5432 user=huongbui dbname=booking_system password=huongbui sslmode=disable")
	if err != nil {
		log.Fatalf("Getting error connect to postgresql database", zap.Error(err))

	}

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Initialize the generated Todo service.
	svc := entpb.NewCustomerService(client)
	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()
	// Register the Todo service with the server.
	entpb.RegisterCustomerServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
