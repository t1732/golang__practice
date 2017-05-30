package main

import (
	"fmt"
	"time"
)

func main() {
	// get current time
	fmt.Println(time.Now())


	// parse time
	var timeformat = "2006-01-02"
	t, err := time.Parse(timeformat, "2013-01-01")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)

	// change time
	tDayAgo := t.Add(24 * time.Hour)
	fmt.Println(tDayAgo)
}
