package main

import (
	"context"
	"customer-service/ent"
	"customer-service/ent/proto/entpb"
	"customer-service/internal/logger"
	"customer-service/internal/utils"
	"fmt"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

func main() {
	initLogger()
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	fmt.Println("URL: " + get("POSTGRES_CONNECTION_STRING"))
	// Initialize an ent client.
	client, err := ent.Open("postgres", get("POSTGRES_CONNECTION_STRING"))

	if err != nil {
		logger.NewLogger().Fatal("failed opening connection to postgres", zap.Error(err))

	}

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.NewLogger().Fatal("failed creating schema resources", zap.Error(err))
	}
	// Initialize the generated Customer service.
	svc := entpb.NewCustomerService(client)
	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()
	// Register the Customer service with the server.
	entpb.RegisterCustomerServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":"+get("PORT"))
	logger.NewLogger().Info("server listening on port " + get("PORT"))
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		logger.NewLogger().Fatal("failed serving traffic", zap.Error(err))
	}
}
