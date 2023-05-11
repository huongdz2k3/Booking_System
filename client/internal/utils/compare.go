package utils

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func CompareTimeWithNow(flyTime *timestamppb.Timestamp, target int) bool {
	time1 := ConvertTimestampToTime(flyTime)
	time2 := time.Now()

	duration := time2.Sub(time1)

	return duration.Hours() > float64(target)
}

func CompareDateWithNow(flyDate *timestamppb.Timestamp, target int) bool {
	layout := "2006-01-02"
	time1 := ConvertTimestampToTime(flyDate)
	time2 := time.Now()

	date1, err1 := time.Parse(layout, time1.Format(layout))
	if err1 != nil {
		return false
	}

	date2, err2 := time.Parse(layout, time2.Format(layout))
	if err2 != nil {
		return false
	}

	duration := date1.Sub(date2)

	return duration.Hours()/24 > float64(target)
}

func CompareDate(date1 *timestamppb.Timestamp, date2 *timestamppb.Timestamp) bool {
	return date1.AsTime().Before(date2.AsTime())
}
