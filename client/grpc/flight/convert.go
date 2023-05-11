package flight

import (
	"context"
	"customer/ent"
	"customer/ent/flight"
	"customer/grpc/shared"
	"customer/internal/utils"
	flight2 "customer/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCreateFlightInput(input *ent.CreateFlightInput, ctx context.Context) (*flight2.CreateFlightInput, error) {
	var returnDateTS *timestamp.Timestamp
	departDate, _ := utils.StringToDate(input.DepartDate)
	departTime, _ := utils.StringToTime(input.DepartTime)
	if input.Type.String() == "ONE_WAY" {
		returnDateTS = nil
	} else if input.Type.String() == "RETURN_TICKET" {
		if input.ReturnDate == nil {
			return nil, utils.WrapGQLBadRequestError(ctx, "Return Date is not blank when type is RETURN_TICKET")
		}

		returnDate, _ := utils.StringToDate(*input.ReturnDate)
		if returnDate.Before(departDate) {
			return nil, utils.WrapGQLBadRequestError(ctx, "Return Date must be after Depart Date")
		}
		returnDateTS = timestamppb.New(returnDate)

	}

	return &flight2.CreateFlightInput{
		Name:           input.Name,
		From:           input.From,
		To:             input.To,
		DepartDate:     timestamppb.New(departDate),
		DepartTime:     timestamppb.New(departTime),
		AvailableSlots: int32(input.AvailableSlots),
		FlightPlane:    input.FlightPlane,
		Status:         input.Status.String(),
		Type:           input.Type.String(),
		ReturnDate:     returnDateTS,
	}, nil
}

func ConvertUpdateFlightInput(input *ent.UpdateFlightInput, id int, ctx context.Context) (*flight2.UpdateFlightInput, error) {
	var returnDateTS *timestamp.Timestamp
	departDate, _ := utils.StringToDate(*input.DepartDate)
	departTime, _ := utils.StringToTime(*input.DepartTime)
	if input.Type.String() == "ONE_WAY" {
		returnDateTS = nil
	} else if input.Type.String() == "RETURN_TICKET" {
		if input.ReturnDate == nil {
			return nil, utils.WrapGQLBadRequestError(ctx, "Return Date is not blank when type is RETURN_TICKET")
		}

		returnDate, _ := utils.StringToDate(*input.ReturnDate)
		if returnDate.Before(departDate) {
			return nil, utils.WrapGQLBadRequestError(ctx, "Return Date must be after Depart Date")
		}
		returnDateTS = timestamppb.New(returnDate)

	}

	return &flight2.UpdateFlightInput{
		Id:             int32(id),
		Name:           *input.Name,
		From:           *input.From,
		To:             *input.To,
		DepartDate:     timestamppb.New(departDate),
		DepartTime:     timestamppb.New(departTime),
		AvailableSlots: int32(*input.AvailableSlots),
		FlightPlane:    *input.FlightPlane,
		Status:         input.Status.String(),
		Type:           input.Type.String(),
		ReturnDate:     returnDateTS,
	}, nil
}

func ConvertSearchFlightInput(input *ent.SearchFlightInput) (*flight2.QueryFlightInput, error) {
	returnConverted, _ := utils.StringToDate(input.ReturnDate)
	dateConverted, _ := utils.StringToDate(input.DepartDate)
	queryFlightInput := flight2.QueryFlightInput{
		From:       input.From,
		To:         input.To,
		DepartDate: timestamppb.New(dateConverted),
		ReturnDate: timestamppb.New(returnConverted),
		Size:       int32(input.Size),
		Page:       int32(input.Page),
		Type:       input.Type.String(),
	}
	return &queryFlightInput, nil
}

func FromProtoFlight(pbFlight *flight2.Flight) (*ent.Flight, error) {
	returnDate := ""
	if pbFlight.ReturnDate != nil {
		returnDate = utils.ConvertTimestampToDate(pbFlight.ReturnDate).Format("2006-01-02")
	}
	successBookingCount, err := shared.CountBookingsByStatus(&flight2.CountBookingsByStatusInput{Status: "SUCCESS", FlightId: pbFlight.Id})
	if err != nil {
		return nil, err
	}
	availableSlots := int(pbFlight.AvailableSlots) - int(successBookingCount.TotalRecords)
	return &ent.Flight{
		ID:             int(pbFlight.Id),
		Name:           pbFlight.Name,
		From:           pbFlight.From,
		To:             pbFlight.To,
		DepartDate:     utils.ConvertTimestampToDate(pbFlight.DepartDate).Format("2006-01-02"),
		DepartTime:     utils.ConvertTimestampToTime(pbFlight.DepartTime).Format("15:04:05"),
		Status:         GetStatusFromProto(pbFlight.Status),
		AvailableSlots: availableSlots,
		FlightPlane:    pbFlight.FlightPlane,
		Type:           GetTypeFromProto(pbFlight.Type),
		ReturnDate:     returnDate,
	}, nil
}

func ConvertListProtoToFlights(response *flight2.SearchFlightResponse) (*ent.SearchFlightResponse, error) {
	flightList := make([]*ent.Flight, len(response.Flights))
	for i, f := range response.Flights {
		flightList[i], _ = FromProtoFlight(f)
	}
	pageInfo, _ := convertProtoPaginationInfo(response)
	return &ent.SearchFlightResponse{
		Flights:        flightList,
		PaginationInfo: pageInfo,
	}, nil
}

func convertProtoPaginationInfo(response *flight2.SearchFlightResponse) (*ent.PaginationInfo, error) {

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

func GetStatusFromProto(status string) flight.Status {
	switch status {
	case "ON_TIME":
		return flight.StatusON_TIME
	case "DELAYED":
		return flight.StatusDELAYED
	case "CANCELLED":
		return flight.StatusDELAYED
	case "SCHEDULED":
		return flight.StatusSCHEDULED
	default:
		return ""
	}
}

func GetTypeFromProto(flightType string) flight.Type {
	switch flightType {
	case "ONE_WAY":
		return flight.TypeONE_WAY
	case "RETURN_TICKET":
		return flight.TypeRETURN_TICKET
	default:
		return ""
	}
}
