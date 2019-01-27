package basic

import (
	"fmt"
	"time"
)

const timeformat = "2006-01-02"

func TimeParse(value string) (time.Time, error) {
	// parse time
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, err := time.ParseInLocation(timeformat, value, loc)
	if err != nil {
		fmt.Println(err)
		return t, err
	}
	return t, err
}

func AdvanceOneHourNow() (time.Time) {
	t := time.Now()
	// change time
	tDayAgo := t.Add(24 * time.Hour)
	return tDayAgo
}
