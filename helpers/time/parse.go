package time

import (
	"fmt"
	"time"
)

func JustDate(dateString string) (dateParse time.Time, err error) {
	var layoutFormat, value string

	layoutFormat = "2006-03-02"
	value = dateString
	dateParse, err = time.Parse(layoutFormat, value)

	if err != nil {
		return dateParse, err
	}

	return dateParse, nil
}

func JustTime(timeString string) (timeParse time.Time, err error) {
	var layoutFormat, value string

	layoutFormat = "08:00:00 Z0700"
	value = fmt.Sprintf("%s:00", timeString)
	timeParse, err = time.Parse(layoutFormat, value)

	if err != nil {
		return timeParse, err
	}

	return timeParse, nil
}
