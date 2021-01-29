package api

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//DateStringToDate use to transform string into time.Time
//the date should be in YYYY-MM-DD format
func DateStringToDate(dateString string) (time.Time, error) {
	var (
		err      error
		dateInfo = make([]int, 3)
	)

	for i, info := range strings.Split(dateString, "-") {
		dateInfo[i], err = strconv.Atoi(info)
		if err != nil {
			return time.Time{}, err
		}
	}
	return time.Date(
		dateInfo[0],
		time.Month(dateInfo[1]),
		dateInfo[2],
		0,
		0,
		0,
		0,
		time.UTC,
	), nil

}

//TimeToDate transform time.Time in YYYY-MM-DD string
func TimeToDate(date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())
}
