package grpc

import (
	"context"
	"flight-service/ent"
	"flight-service/ent/flight"
	"flight-service/internal/utils"
	"flight-service/pb"
	"math"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightService struct {
	client *ent.Client
	pb.UnimplementedFlightServiceServer
}

func FlightToProto(flight *ent.Flight) *pb.Flight {
	var returnDate *timestamppb.Timestamp
	if flight.ReturnDate != nil {
		returnDate = timestamppb.New(*flight.ReturnDate)
	}
	return &pb.Flight{
		Id:             int32(flight.ID),
		Name:           flight.Name,
		From:           flight.From,
		To:             flight.To,
		DepartDate:     timestamppb.New(flight.DepartDate),
		DepartTime:     timestamppb.New(flight.DepartTime),
		Status:         flight.Status.String(),
		AvailableSlots: int32(flight.AvailableSlots),
		FlightPlane:    flight.FlightPlane,
		CreatedAt:      timestamppb.New(flight.CreatedAt),
		UpdatedAt:      timestamppb.New(flight.UpdatedAt),
		Type:           flight.Type.String(),
		ReturnDate:     returnDate,
	}
}

func (s *FlightService) CreateFlight(ctx context.Context, input *pb.CreateFlightInput) (*pb.Flight, error) {

	flight := s.client.Flight.Create().
		SetName(input.Name).
		SetTo(input.To).
		SetFrom(input.From).
		SetFlightPlane(input.FlightPlane).
		SetAvailableSlots(int(input.AvailableSlots)).
		SetDepartTime(ConvertProtoTimestampToTime(input.DepartTime)).
		SetDepartDate(ConvertProtoTimestampToTime(input.DepartDate)).
		SetStatus(GetStatusFromProto(input.Status)).
		SetType(GetTypeFromProto(input.Type))

	if input.Type == "RETURN_TICKET" {
		returnDate := ConvertProtoTimestampToTime(input.ReturnDate)
		flight.SetReturnDate(returnDate)
	}

	fli, err := flight.Save(ctx)
	if err != nil {
		return nil, err
	}

	return FlightToProto(fli), nil
}

func (s *FlightService) UpdateFlight(ctx context.Context, input *pb.UpdateFlightInput) (*pb.Flight, error) {
	fli, err := s.client.Flight.UpdateOneID(int(input.Id)).SetName(input.Name).SetTo(input.To).SetFrom(input.From).SetFlightPlane(input.FlightPlane).SetAvailableSlots(int(input.AvailableSlots)).SetDepartDate(ConvertProtoTimestampToTime(input.DepartTime)).SetDepartDate(ConvertProtoTimestampToTime(input.DepartDate)).SetStatus(GetStatusFromProto(input.Status)).SetReturnDate(ConvertProtoTimestampToTime(input.ReturnDate)).SetUpdatedAt(time.Now()).SetType(GetTypeFromProto(input.Type)).Save(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}
	return FlightToProto(fli), nil
}

func (s *FlightService) SearchFlights(ctx context.Context, input *pb.QueryFlightInput) (*pb.SearchFlightResponse, error) {
	// Calculate the offset based on the page and size inputs.
	offset := (input.Page - 1) * input.Size
	// Determine the flight type query based on input.Type
	flightTypeQuery := flight.TypeONE_WAY
	if input.Type == "RETURN_TICKET" {
		flightTypeQuery = flight.TypeRETURN_TICKET
	}
	// Query flights with the specified conditions
	flightsList, err := s.client.Flight.Query().
		Where(
			flight.Or(
				flight.And(
					flight.From(input.From),
					flight.To(input.To),
					flight.DepartDateGTE(ConvertProtoTimestampToTime(input.DepartDate)),
					flight.StatusIn(flight.StatusSCHEDULED),
					flight.TypeEQ(flightTypeQuery),
				),
				flight.And(
					flight.From(input.To),
					flight.To(input.From),
					flight.DepartDateLTE(ConvertProtoTimestampToTime(input.ReturnDate)),
					flight.StatusIn(flight.StatusSCHEDULED),
					flight.TypeEQ(flightTypeQuery),
				),
			),
		).Offset(int(offset)).Limit(int(input.Size)).All(ctx)

	if err != nil {
		return nil, err
	}
	// Count the total number of records matching the criteria.
	totalCount, err := s.client.Flight.Query().
		Where(
			flight.Or(
				flight.And(
					flight.From(input.From),
					flight.To(input.To),
					flight.DepartDate(ConvertProtoTimestampToTime(input.DepartDate)),
					flight.StatusIn(flight.StatusSCHEDULED),
					flight.TypeEQ(flightTypeQuery),
				),
				flight.And(
					flight.From(input.To),
					flight.To(input.From),
					flight.DepartDate(ConvertProtoTimestampToTime(input.ReturnDate)),
					flight.StatusIn(flight.StatusSCHEDULED),
					flight.TypeEQ(flightTypeQuery),
				),
			),
		).Count(ctx)
	if err != nil {
		return nil, err
	}
	// Calculate the total number of pages based on the total count and page size.
	totalPages := int(math.Ceil(float64(totalCount) / float64(input.Size)))

	// Convert the list of flights to ProtoFlight objects.
	flightListProto := make([]*pb.Flight, len(flightsList))
	for i, f := range flightsList {
		flightListProto[i] = FlightToProto(f)
	}
	return &pb.SearchFlightResponse{
		Flights:      flightListProto,
		Size:         input.Size,
		Page:         input.Page,
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalCount),
	}, nil
}

func (s *FlightService) GetFlightById(ctx context.Context, input *pb.GetFlightByIdInput) (*pb.Flight, error) {
	fli, err := s.client.Flight.Query().Where(flight.IDEQ(int(input.Id))).First(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}
	return FlightToProto(fli), nil
}

func (s *FlightService) CountFlightsByStatus(ctx context.Context, input *pb.CountFlightsByStatusInput) (*pb.CountFlightsByStatusResponse, error) {
	fli, err := s.client.Flight.Query().Where(flight.StatusEQ(GetStatusFromProto(input.Status))).Count(ctx)
	if err != nil {
		return nil, utils.WrapNotFoundError(ctx)
	}
	return &pb.CountFlightsByStatusResponse{
		Status: input.Status,
		Total:  int32(fli),
	}, nil
}

func NewFlightService(client *ent.Client) *FlightService {
	return &FlightService{
		client: client,
	}
}

func ConvertProtoTimestampToTime(ts *timestamppb.Timestamp) time.Time {
	t := time.Unix(ts.GetSeconds(), int64(ts.GetNanos())).UTC()
	return t
}
