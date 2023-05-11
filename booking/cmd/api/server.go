package api

import (
	"booking/ent"
	grpc2 "booking/grpc"
	"booking/internal/utils"
	"booking/pb"
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewServerCmd(logger *zap.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with grpc",
		Run: func(cmd *cobra.Command, args []string) {
			utils.LoadEnv()
			get := utils.GetEnvWithKey
			client, err := ent.Open("postgres", get("POSTGRES_CONNECTION_STRING"))

			if err != nil {
				logger.Fatal("failed opening connection to postgres", zap.Error(err))

			}

			// Run the migration tool (creating tables, etc).
			if err := client.Schema.Create(context.Background()); err != nil {
				logger.Fatal("failed creating schema resources", zap.Error(err))
			}
			// Initialize the generated Customer service.
			svc := grpc2.NewBookingService(client)
			// Create a new gRPC server (you can wire multiple services to a single server).
			server := grpc.NewServer()
			// Register the Customer service with the server.
			pb.RegisterBookingServiceServer(server, svc)

			// Open port 5000 for listening to traffic.
			lis, err := net.Listen("tcp", ":"+get("PORT"))
			logger.Info("server listening on port " + get("PORT"))
			if err != nil {
				log.Fatalf("failed listening: %s", err)
			}

			// Listen for traffic indefinitely.
			if err := server.Serve(lis); err != nil {
				logger.Fatal("failed serving traffic", zap.Error(err))
			}
		},
	}
}
