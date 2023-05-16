package flight

import (
	"context"
	"customer/ent"
	"customer/grpc/booking"
	grpc_client "customer/internal/grpc"
	"customer/internal/logger"
	"customer/internal/utils"
	"customer/pb"
	"sync"

	"go.uber.org/zap"
)

type grpcClient struct {
	client pb.FlightServiceClient
	err    error
}

var (
	once     sync.Once
	instance *grpcClient
)

func GetFlightGRPCClient() *grpcClient {
	once.Do(func() {
		client, err := grpc_client.NewGrpcClient(grpc_client.FlightServiceClient, grpc_client.FlightServiceHost)
		if err != nil {
			logger.NewLogger().Fatal("failed connecting to server", zap.Error(err))
		}
		instance = &grpcClient{
			client: client.(pb.FlightServiceClient),
			err:    nil,
		}

	})
	return instance
}

func CreateFlight(input *ent.CreateFlightInput, ctx context.Context) (*pb.Flight, error) {
	createFlightInput, err := ConvertCreateFlightInput(input, ctx)
	if createFlightInput == nil {
		return nil, err
	}
	if utils.CompareDateWithNow(createFlightInput.DepartDate, 0) == false {
		return nil, utils.WrapGQLBadRequestError(ctx, "Depart date must be at least 1 day from today")
	}
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	resp, err := GetFlightGRPCClient().client.CreateFlight(context.Background(), createFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func UpdateFlight(input *ent.UpdateFlightInput, id int, cxt context.Context) (*pb.Flight, error) {

	updateFlightInput, err := ConvertUpdateFlightInput(input, id, cxt)
	if updateFlightInput == nil {
		return nil, err
	}
	if updateFlightInput.Status == "CANCELLED" {
		booking.CancelBookingWithFlightId(&pb.CancelBookingWithFlightIdInput{
			FlightId: int32(id),
		})
	}
	// connect gRPC server
	resp, err := GetFlightGRPCClient().client.UpdateFlight(context.Background(), updateFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func SearchFlights(input *ent.SearchFlightInput) (*pb.SearchFlightResponse, error) {

	searchFlightInput, _ := ConvertSearchFlightInput(input)
	// connect gRPC server
	resp, err := GetFlightGRPCClient().client.SearchFlights(context.Background(), searchFlightInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CountFlightsByStatus(status string, flight_id int32) (*pb.CountFlightsByStatusResponse, error) {
	resp, err := GetFlightGRPCClient().client.CountFlightsByStatus(context.Background(), &pb.CountFlightsByStatusInput{Status: status, FlightId: flight_id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
