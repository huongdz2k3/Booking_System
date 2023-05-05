package flight

import (
	"customer/ent"
	"customer/internal/logger"
	"customer/internal/utils"

	"context"

	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func createGRPCClient() (FlightServiceClient, error) {
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	// Open a connection to the server.
	conn, err := grpc.Dial(":"+get("FLIGHT_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.NewLogger().Fatal("failed connecting to server", zap.Error(err))
	}
	return NewFlightServiceClient(conn), nil
}

func CreateFlight(input *ent.CreateFlightInput) (*Flight, error) {
	// create a gRPC client instance
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	createFlightInput, err := ConvertCreateFlightInput(input)
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	resp, err := client.CreateFlight(context.Background(), createFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateFlight(input *ent.UpdateFlightInput, id int) (*Flight, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	updateFlightInput, _ := ConvertUpdateFlightInput(input, id)
	// connect gRPC server
	resp, err := client.UpdateFlight(context.Background(), updateFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SearchFlights(input *ent.SearchFlightInput) (*SearchFlightResponse, error) {
	client, err := createGRPCClient()
	if err != nil {
		return nil, err
	}
	searchFlightInput, _ := ConvertSearchFlightInput(input)
	// connect gRPC server
	resp, err := client.SearchFlights(context.Background(), searchFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
