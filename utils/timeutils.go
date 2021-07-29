package utils

import (
	"log"
	"time"
)

const (
	LayoutIOS8601 = "2006-01-02T15:04+08:00"
	LayoutYYYYMMDDHHMMSS = "2006-01-02 15:04:05"
	LayoutYYYYMMDD = "2006-01-02"
)

func ConvertTimeFromIOS8601(iso8601 string) string {
	t, err := time.Parse(LayoutIOS8601, iso8601)
	if err != nil {
		log.Printf("ConvertTimeFromIOS8601 error: %v, fromTime: %s", err, iso8601)
		return iso8601
	}
	return t.Format(LayoutYYYYMMDDHHMMSS)
}

func ParseTimeFromIOS8601(iso8601 string) (time.Time, error) {
	return time.Parse(LayoutIOS8601, iso8601)
}