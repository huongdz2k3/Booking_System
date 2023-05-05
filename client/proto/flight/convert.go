package flight

import (
	"customer/ent"
	"customer/ent/flight"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertCreateFlightInput(input *ent.CreateFlightInput) (*CreateFlightInput, error) {
	return &CreateFlightInput{
		Name:           input.Name,
		From:           input.From,
		To:             input.To,
		DepartDate:     timestamppb.New(input.DepartDate),
		DepartTime:     timestamppb.New(input.DepartTime),
		AvailableSlots: int32(input.AvailableSlots),
		FlightPlane:    input.FlightPlane,
		Status:         input.Status.String(),
	}, nil
}

func ConvertUpdateFlightInput(input *ent.UpdateFlightInput, id int) (*UpdateFlightInput, error) {
	return &UpdateFlightInput{
		Id:             int32(id),
		Name:           *input.Name,
		From:           *input.From,
		To:             *input.To,
		DepartDate:     timestamppb.New(*input.DepartDate),
		DepartTime:     timestamppb.New(*input.DepartTime),
		AvailableSlots: int32(*input.AvailableSlots),
		FlightPlane:    *input.FlightPlane,
		Status:         input.Status.String(),
	}, nil
}

func ConvertSearchFlightInput(input *ent.SearchFlightInput) (*QueryFlightInput, error) {
	queryFlightInput := QueryFlightInput{
		From:       input.From,
		To:         input.To,
		DepartDate: timestamppb.New(input.DepartDate),
		ReturnDate: timestamppb.New(input.ReturnDate),
		Size:       int32(input.Size),
		Page:       int32(input.Page),
	}
	return &queryFlightInput, nil
}

func FromProtoFlight(pbFlight *Flight) (*ent.Flight, error) {

	return &ent.Flight{
		ID:             int(pbFlight.Id),
		Name:           pbFlight.Name,
		From:           pbFlight.From,
		To:             pbFlight.To,
		DepartDate:     pbFlight.DepartDate.AsTime(),
		DepartTime:     pbFlight.DepartTime.AsTime(),
		Status:         GetStatusFromProto(pbFlight.Status),
		AvailableSlots: int(pbFlight.AvailableSlots),
		FlightPlane:    pbFlight.FlightPlane,
	}, nil
}

func ConvertListProtoToFlights(response *SearchFlightResponse) (*ent.SearchFlightResponse, error) {
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

func convertProtoPaginationInfo(response *SearchFlightResponse) (*ent.PaginationInfo, error) {

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
