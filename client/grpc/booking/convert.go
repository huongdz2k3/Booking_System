package booking

import (
	"context"
	"customer/ent"
	"customer/ent/booking"
	"customer/internal/utils"
	"customer/pb"
)

func ConvertCreateBookingInput(input *ent.CreateBookingInput, ctx context.Context) (*pb.CreateBookingInput, error) {
	return &pb.CreateBookingInput{
		FlightId:     int32(input.FlightID),
		LicenseId:    input.LicenseID,
		Dob:          input.Dob,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		CustomerName: input.CustomerName,
		Address:      input.Address,
	}, nil
}

func ConvertViewBookingInput(input *ent.ViewBookingInput) (*pb.ViewBookingInput, error) {
	return &pb.ViewBookingInput{
		BookingCode: input.BookingCode,
		Email:       input.Email,
		LicenseId:   input.LicenseID,
	}, nil
}

func ConvertCancelBookingInput(input *ent.CancelBookingInput) (*pb.CancelBookingInput, error) {
	return &pb.CancelBookingInput{
		BookingCode: input.BookingCode,
		Email:       input.Email,
		LicenseId:   input.LicenseID,
	}, nil
}
func FromProtoToBooking(input *pb.Booking) *ent.Booking {
	cusId := int(input.CustomerId)

	var status booking.Status
	if input.Status == "CANCEL" {
		status = booking.StatusCANCEL
	} else if input.Status == "SUCCESS" {
		status = booking.StatusSUCCESS
	} else if input.Status == "PROCESS" {
		status = booking.StatusPROCESS
	}
	return &ent.Booking{
		CustomerID:   &cusId,
		FlightID:     int(input.FlightId),
		LicenseID:    &input.LicenseId,
		Dob:          &input.Dob,
		Email:        &input.Email,
		PhoneNumber:  &input.PhoneNumber,
		CustomerName: &input.CustomerName,
		Address:      &input.Address,
		BookingCode:  input.BookingCode,
		Status:       status,
		BookingDate:  utils.ConvertTimestampToDate(input.BookingDate).Format("2006-01-02 15:04:05"),
	}
}

func ConvertGetBookingsHistoryInput(input *ent.PaginateInput, customerId int) (*pb.GetBookingsHistoryInput, error) {
	return &pb.GetBookingsHistoryInput{
		Size:       int32(input.Size),
		Page:       int32(input.Page),
		CustomerId: int32(customerId),
	}, nil
}

func ConvertListProtoToBookingHistory(input *pb.GetBookingsHistoryResponse) (*ent.BookingHistoryResponse, error) {
	bookingsList := []*ent.Booking{}
	for _, booking := range input.Bookings {
		bookingsList = append(bookingsList, FromProtoToBooking(booking))
	}
	paginationInfo, _ := convertProtoPaginationInfo(input)
	return &ent.BookingHistoryResponse{
		Bookings:       bookingsList,
		PaginationInfo: paginationInfo,
	}, nil
}

func convertProtoPaginationInfo(response *pb.GetBookingsHistoryResponse) (*ent.PaginationInfo, error) {

	size := int(response.Size)
	totalRecords := int(response.TotalRecords)
	page := int(response.Page)
	totalPages := int(response.TotalPages)
	return &ent.PaginationInfo{
		Size:         &size,
		TotalRecords: &totalRecords,
		Page:         &page,
		TotalPages:   &totalPages,
	}, nil
}
