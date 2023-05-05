package entpb

import (
	"context"
	"flight-service/ent"
	"flight-service/ent/flight"
	"flight-service/internal/utils"
	"math"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightService struct {
	client *ent.Client
	UnimplementedFlightServiceServer
}

func FlightToProto(flight *ent.Flight) *Flight {
	return &Flight{
		Id:             int32(flight.ID),
		Name:           flight.Name,
		From:           flight.From,
		To:             flight.To,
		DepartDate:     timestamppb.New(flight.DepartDate),
		DepartTime:     timestamppb.New(flight.DepartTime),
		Status:         flight.Status,
		AvailableSlots: int32(flight.AvailableSlots),
		FlightPlane:    flight.FlightPlane,
		CreatedAt:      timestamppb.New(flight.CreatedAt),
		UpdatedAt:      timestamppb.New(flight.UpdatedAt),
	}
}

func (s *FlightService) CreateFlight(ctx context.Context, input *CreateFlightInput) (*Flight, error) {
	fli, err := s.client.Flight.Create().SetName(input.Name).SetTo(input.To).SetFrom(input.From).SetFlightPlane(input.FlightPlane).SetAvailableSlots(int(input.AvailableSlots)).SetDepartTime(ConvertProtoTimestampToTime(input.DepartTime)).SetDepartDate(ConvertProtoTimestampToTime(input.DepartDate)).SetStatus(input.Status).SetReturnDate(ConvertProtoTimestampToTime(input.ReturnDate)).Save(ctx)
	if err != nil {
		return nil, err
	}

	return FlightToProto(fli), nil

}

func (s *FlightService) UpdateFlight(ctx context.Context, input *UpdateFlightInput) (*Flight, error) {
	fli, err := s.client.Flight.UpdateOneID(int(input.Id)).SetName(input.Name).SetTo(input.To).SetFrom(input.From).SetFlightPlane(input.FlightPlane).SetAvailableSlots(int(input.AvailableSlots)).SetDepartDate(ConvertProtoTimestampToTime(input.DepartTime)).SetDepartDate(ConvertProtoTimestampToTime(input.DepartDate)).SetStatus(input.Status).SetReturnDate(ConvertProtoTimestampToTime(input.ReturnDate)).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}
	return FlightToProto(fli), nil
}

func (s *FlightService) SearchFlights(ctx context.Context, input *QueryFlightInput) (*SearchFlightResponse, error) {
	// Calculate the offset based on the page and size inputs.
	offset := (input.Page - 1) * input.Size
	flightsList, _ := s.client.Flight.Query().
		Where(
			flight.Or(
				flight.And(
					flight.From(input.From),
					flight.To(input.To),
					flight.DepartDate(ConvertProtoTimestampToTime(input.DepartDate)),
				),
				flight.And(
					flight.From(input.To),
					flight.To(input.From),
					flight.DepartDate(ConvertProtoTimestampToTime(input.ReturnDate)),
				),
			),
		).Offset(int(offset)).Limit(int(input.Size)).All(ctx)

	// Count the total number of records matching the criteria.
	totalCount, _ := s.client.Flight.Query().
		Where(
			flight.Or(
				flight.And(
					flight.From(input.From),
					flight.To(input.To),
					flight.DepartDate(ConvertProtoTimestampToTime(input.DepartDate)),
				),
				flight.And(
					flight.From(input.To),
					flight.To(input.From),
					flight.DepartDate(ConvertProtoTimestampToTime(input.ReturnDate)),
				),
			),
		).Count(ctx)

	// Calculate the total number of pages based on the total count and page size.
	totalPages := int(math.Ceil(float64(totalCount) / float64(input.Size)))

	// Convert the list of flights to ProtoFlight objects.
	flightListProto := make([]*Flight, len(flightsList))
	for i, f := range flightsList {
		flightListProto[i] = FlightToProto(f)
	}
	return &SearchFlightResponse{
		Flights:      flightListProto,
		Size:         input.Size,
		Page:         input.Page,
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalCount),
	}, nil
}

func NewFlightService(client *ent.Client) *FlightService {
	return &FlightService{
		client: client,
	}
}

func ConvertProtoTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos())).UTC()
}
