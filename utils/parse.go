package utils

import "time"

func ParseStringToTime(value string, timeFormat string, fieldName string) (time.Time, error) {
	parsedVal, err := time.Parse(timeFormat, value)
	if err != nil {
		return time.Time{}, err
	}

	return parsedVal, nil
}
