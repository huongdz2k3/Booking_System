package grpc

import (
	"customer/internal/utils"
	"customer/pb"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientName string
type GrpcHost string

const (
	BookingServiceClient  GrpcClientName = "BookingServiceClient"
	FlightServiceClient   GrpcClientName = "FlightServiceClient"
	CustomerServiceClient GrpcClientName = "CustomerServiceClient"
)

const (
	BookingServiceHost  GrpcHost = "BOOKING_SERVICE_HOST"
	FlightServiceHost   GrpcHost = "FLIGHT_SERVICE_HOST"
	CustomerServiceHost GrpcHost = "CUSTOMER_SERVICE_HOST"
)

var GrpcClientHostMap = map[GrpcHost]string{
	BookingServiceHost:  "BOOKING_SERVICE_HOST",
	FlightServiceHost:   "FLIGHT_SERVICE_HOST",
	CustomerServiceHost: "CUSTOMER_SERVICE_HOST",
}

var GrpcClientMap = map[GrpcClientName]string{
	BookingServiceClient:  "BOOKING_SERVICE_PORT",
	FlightServiceClient:   "FLIGHT_SERVICE_PORT",
	CustomerServiceClient: "CUSTOMER_SERVICE_PORT",
}

func NewGrpcClient(name GrpcClientName, host GrpcHost) (interface{}, error) {
	utils.LoadEnv()
	get := utils.GetEnvWithKey
	cc, err := grpc.Dial(get(GrpcClientHostMap[host])+":"+get(GrpcClientMap[name]), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	switch name {
	case BookingServiceClient:
		{
			return pb.NewBookingServiceClient(cc), nil
		}
	case FlightServiceClient:
		{
			return pb.NewFlightServiceClient(cc), nil
		}
	case CustomerServiceClient:
		{
			return pb.NewCustomerServiceClient(cc), nil
		}
	}

	return nil, errors.New("Invalid grpc client name")
}
