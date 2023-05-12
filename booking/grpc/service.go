package grpc

import (
	"booking/ent"
	"booking/ent/booking"
	"booking/internal/utils"
	"booking/pb"
	"context"
	"crypto/rand"
	"encoding/base64"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"time"
)

type BookingService struct {
	client *ent.Client
	pb.UnimplementedBookingServiceServer
}

func BookingToProto(booking *ent.Booking) *pb.Booking {
	var cancelDate *timestamppb.Timestamp
	if booking.CancelDate != nil {
		cancelDate = timestamppb.New(*booking.CancelDate)
	}

	return &pb.Booking{
		Id:           int32(booking.ID),
		CreatedAt:    timestamppb.New(booking.CreatedAt),
		UpdatedAt:    timestamppb.New(booking.UpdatedAt),
		BookingCode:  booking.BookingCode,
		Address:      *booking.Address,
		Email:        *booking.Email,
		PhoneNumber:  *booking.PhoneNumber,
		Dob:          *booking.Dob,
		CustomerName: *booking.CustomerName,
		FlightId:     int32(booking.FlightID),
		CustomerId:   int32(*booking.CustomerID),
		LicenseId:    *booking.LicenseID,
		BookingDate:  timestamppb.New(booking.BookingDate),
		CancelDate:   cancelDate,
		Status:       booking.Status.String(),
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, input *pb.CreateBookingInput) (*pb.Booking, error) {
	book, err := s.client.Booking.Create().
		SetAddress(input.Address).
		SetEmail(input.Email).
		SetPhoneNumber(input.PhoneNumber).
		SetDob(input.Dob).
		SetCustomerName(input.CustomerName).
		SetLicenseID(input.LicenseId).
		SetFlightID(int(input.FlightId)).
		SetCustomerID(int(input.CustomerId)).
		SetBookingCode(generateRandomString(6)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	b, _ := book.Update().SetStatus(booking.StatusSUCCESS).Save(ctx)
	return BookingToProto(b), nil
}

func (s *BookingService) ViewBooking(ctx context.Context, input *pb.ViewBookingInput) (*pb.Booking, error) {
	booking, err := s.client.Booking.Query().Where(booking.BookingCode(input.GetBookingCode()), booking.Email(input.GetEmail()), booking.LicenseID(input.GetLicenseId())).First(ctx)
	if err != nil {
		return nil, err
	}
	return BookingToProto(booking), nil

}

func (s *BookingService) CancelBooking(ctx context.Context, input *pb.CancelBookingInput) (*pb.Booking, error) {
	booking, err := s.client.Booking.Query().Where(booking.BookingCode(input.GetBookingCode()), booking.Email(input.GetEmail()), booking.LicenseID(input.GetLicenseId())).First(ctx)
	if err != nil {
		return nil, err
	}
	booking, err = booking.Update().SetStatus("CANCEL").SetCancelDate(time.Now()).Save(ctx)
	if err != nil {
		return nil, utils.WrapBadRequestError(ctx, "You can't cancel this booking")
	}
	return BookingToProto(booking), nil
}

func (s *BookingService) GetBookingsHistory(ctx context.Context, input *pb.GetBookingsHistoryInput) (*pb.GetBookingsHistoryResponse, error) {
	offset := (input.Page - 1) * input.Size
	bookings, err := s.client.Booking.Query().Where(booking.CustomerID(int(input.GetCustomerId()))).Offset(int(offset)).Limit(int(input.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	total, err := s.client.Booking.Query().Where(booking.CustomerID(int(input.GetCustomerId()))).Offset(int(offset)).Limit(int(input.Size)).Count(ctx)
	// Calculate the total number of pages based on the total count and page size.
	totalPages := int(math.Ceil(float64(total) / float64(input.Size)))

	bookingsListProto := make([]*pb.Booking, len(bookings))
	for i, booking := range bookings {
		bookingsListProto[i] = BookingToProto(booking)
	}
	return &pb.GetBookingsHistoryResponse{
		Bookings:     bookingsListProto,
		TotalPages:   int32(totalPages),
		TotalRecords: int32(total),
		Page:         input.Page,
		Size:         input.Size,
	}, nil
}
func (s *BookingService) CountBookingsByStatus(ctx context.Context, input *pb.CountBookingsByStatusInput) (*pb.CountBookingsByStatusResponse, error) {
	total, err := s.client.Booking.Query().Where(booking.StatusIn(GetStatusFromProto(input.GetStatus())), booking.FlightID(int(input.GetFlightId()))).Count(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CountBookingsByStatusResponse{
		TotalRecords: int32(total),
		Status:       input.Status,
	}, nil
}

func NewBookingService(client *ent.Client) *BookingService {
	return &BookingService{
		client: client,
	}
}

func generateRandomString(length int) string {
	randomBytes := make([]byte, length)
	rand.Read(randomBytes)
	return base64.StdEncoding.EncodeToString(randomBytes)[:length]
}
