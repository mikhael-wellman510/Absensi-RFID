package utils

import "time"

const dateFormat = "2006-01-02"

func ParseDate(date string) (time.Time, error) {
	return time.Parse(dateFormat, date)
}

func FormatDate(data time.Time) string {
	return data.Format(dateFormat)
}
