package booking

import (
	"context"
	"customer/ent"
	customer2 "customer/grpc/customer"
	"customer/grpc/shared"
	grpc_client "customer/internal/grpc"
	"customer/internal/logger"
	"customer/internal/utils"
	"customer/pb"
	"go.uber.org/zap"
	"sync"
)

type grpcClient struct {
	client pb.BookingServiceClient
	err    error
}

var (
	once     sync.Once
	instance *grpcClient
)

func GetBookingGRPCClient() *grpcClient {
	once.Do(func() {
		client, err := grpc_client.NewGrpcClient(grpc_client.BookingServiceClient, grpc_client.BookingServiceHost)
		if err != nil {
			logger.NewLogger().Fatal("failed connecting to server", zap.Error(err))
		}
		instance = &grpcClient{
			client: client.(pb.BookingServiceClient),
			err:    nil,
		}
	})
	return instance
}

func CreateBooking(input *ent.CreateBookingInput, ctx context.Context) (*pb.Booking, error) {
	flight, err := shared.GetFlightById(input.FlightID)
	customer, _ := customer2.GetCustomerByEmail(&pb.GetCustomerByEmailInput{Email: input.Email})

	successBookingCount, err := shared.CountBookingsByStatus(&pb.CountBookingsByStatusInput{Status: "SUCCESS"})
	if err != nil {
		return nil, err
	}
	if flight.AvailableSlots-successBookingCount.TotalRecords <= 0 {
		return nil, utils.WrapGQLBadRequestError(ctx, "No available slot")
	}
	if err != nil {
		return nil, err
	}
	if flight.Status != "SCHEDULED" {
		return nil, utils.WrapGQLBadRequestError(ctx, "Flight is not available")
	}
	if utils.CompareDateWithNow(flight.DepartDate, 0) == true {
		if utils.CompareTimeWithNow(flight.DepartTime, 4) == false {
			return nil, utils.WrapGQLBadRequestError(ctx, "Booking time must be at least 4 hours before departure time")
		}
	}

	count, err := shared.CountBookingsByStatus(&pb.CountBookingsByStatusInput{Status: "SUCCESS"})
	if flight.AvailableSlots <= count.TotalRecords {
		return nil, utils.WrapGQLBadRequestError(ctx, "No available slot")
	}
	createBookingInput, err := ConvertCreateBookingInput(input, ctx)
	if customer != nil {
		createBookingInput.CustomerId = int32(customer.Id)
	}
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	resp, err := GetBookingGRPCClient().client.CreateBooking(context.Background(), createBookingInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ViewBooking(input *ent.ViewBookingInput) (*pb.Booking, error) {
	viewBookingInput, err := ConvertViewBookingInput(input)
	// connect gRPC server
	resp, err := GetBookingGRPCClient().client.ViewBooking(context.Background(), viewBookingInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CancelBooking(input *ent.CancelBookingInput) (*pb.Booking, error) {
	cancelBookingInput, err := ConvertCancelBookingInput(input)
	if err != nil {
		return nil, err
	}
	booking, err := ViewBooking(&ent.ViewBookingInput{BookingCode: input.BookingCode, LicenseID: input.LicenseID, Email: input.Email})
	if err != nil {
		return nil, err
	}
	flight, err := shared.GetFlightById(int(booking.FlightId))
	if utils.CompareTimeWithNow(flight.DepartDate, 2) == true {
		return nil, utils.WrapGQLBadRequestError(context.Background(), "You can only cancel booking 2 days before departure date")
	}
	// connect gRPC server
	resp, err := GetBookingGRPCClient().client.CancelBooking(context.Background(), cancelBookingInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetBookingsHistory(input *ent.PaginateInput, customerId int) (*pb.GetBookingsHistoryResponse, error) {
	getBookingsHistoryInput, err := ConvertGetBookingsHistoryInput(input, customerId)
	if err != nil {
		return nil, err
	}
	// connect gRPC server
	resp, err := GetBookingGRPCClient().client.GetBookingsHistory(context.Background(), getBookingsHistoryInput)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
