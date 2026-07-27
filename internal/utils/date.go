package utils

import (
	"attendance-api/pkg/constants"
	"time"
)

func ParseDate(date string) (time.Time, error) {
	return time.Parse(constants.DateFormat, date)
}

func FormatDate(data time.Time) string {
	return data.Format(constants.DateFormat)
}
