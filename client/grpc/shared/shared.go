package shared

import (
	"context"
	grpc_client "customer/internal/grpc"
	"customer/pb"
)

func CountBookingsByStatus(input *pb.CountBookingsByStatusInput) (*pb.CountBookingsByStatusResponse, error) {
	// connect gRPC server
	client, err := grpc_client.NewGrpcClient(grpc_client.BookingServiceClient, grpc_client.BookingServiceHost)
	if err != nil {
		return nil, err
	}
	service := client.(pb.BookingServiceClient)
	resp, err := service.CountBookingsByStatus(context.Background(), input)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetFlightById(id int) (*pb.Flight, error) {
	client, err := grpc_client.NewGrpcClient(grpc_client.FlightServiceClient, grpc_client.FlightServiceHost)
	if err != nil {
		return nil, err
	}
	service := client.(pb.FlightServiceClient)
	resp, err := service.GetFlightById(context.Background(), &pb.GetFlightByIdInput{Id: int32(id)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
