package utils

import (
	"customer/internal/logger"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertStringToInt(input string) (*int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		logger.NewLogger().Error("Input is not number")
		return nil, err
	}
	return &i, nil
}

func StringToTime(timeStr string) (time.Time, error) {
	layout := "15:04" // specify the "hh:mm" format
	return time.Parse(layout, timeStr)
}

func StringToDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02" // specify the "yyyy-mm-dd" format
	return time.Parse(layout, dateStr)
}
func ConvertTimestampToDate(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
}

func ConvertTimestampToTime(timestamp *timestamppb.Timestamp) time.Time {
	seconds := timestamp.GetSeconds()
	nanos := int64(timestamp.GetNanos())
	return time.Unix(seconds, nanos).UTC()
}
