package api

import (
	"context"
	"customer-service/ent"
	grpc2 "customer-service/grpc"
	"customer-service/internal/utils"
	"customer-service/migration"
	"customer-service/pb"
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
			logger.Info("URL: " + get("POSTGRES_CONNECTION_STRING"))
			// Initialize an ent client.
			client, err := ent.Open("postgres", get("POSTGRES_CONNECTION_STRING"))

			if err != nil {
				logger.Fatal("failed opening connection to postgres", zap.Error(err))

			}

			// Run the migration tool (creating tables, etc).
			if err := client.Schema.Create(context.Background()); err != nil {
				logger.Fatal("failed creating schema resources", zap.Error(err))
			}
			migration.CreateAdmin(client)
			// Initialize the generated Customer service.
			svc := grpc2.NewCustomerService(client)
			// Create a new gRPC api (you can wire multiple services to a single api).
			server := grpc.NewServer()
			// Register the Customer service with the api.
			pb.RegisterCustomerServiceServer(server, svc)

			// Open port 5000 for listening to traffic.
			lis, err := net.Listen("tcp", ":"+get("PORT"))
			logger.Info("api listening on port " + get("PORT"))
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
