package grpc

import "flight-service/ent/flight"

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
