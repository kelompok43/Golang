package time

import (
	"fmt"
	"strconv"
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

func DateTime(dateString, timeString string) (dateTimeParse string, err error) {
	var layoutFormat, value string

	layoutFormat = "2006-01-02 15:04:05"
	value = fmt.Sprintf("%s %s:00", dateString, timeString)
	dateTime, err := time.ParseInLocation(layoutFormat, value, time.Local)

	if err != nil {
		return dateTimeParse, err
	}

	dateTimeParse = strconv.Itoa(int(dateTime.UnixNano() / 1000000))
	return dateTimeParse, nil
}
