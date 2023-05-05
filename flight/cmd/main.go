package main

import (
	"context"
	"flight-service/ent"
	"flight-service/ent/proto/entpb"
	"flight-service/internal/logger"
	"flight-service/internal/utils"
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
	client, err := ent.Open("postgres", get("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		logger.NewLogger().Fatal("failed opening connection to postgres", zap.Error(err))

	}
	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.NewLogger().Fatal("failed creating schema resources", zap.Error(err))
	}
	svc := entpb.NewFlightService(client)
	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()
	//register the flight service with the server
	entpb.RegisterFlightServiceServer(server, svc)

	// Open port for listening traffic
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
