package grpc

import "booking/ent/booking"

func GetStatusFromProto(status string) booking.Status {
	switch status {
	case "PENDING":
		return booking.StatusPROCESS
	case "SUCCESS":
		return booking.StatusSUCCESS
	case "CANCEL":
		return booking.StatusCANCEL
	default:
		return ""
	}
}
