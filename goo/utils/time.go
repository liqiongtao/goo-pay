package gooUtils

import (
	"strings"
	"time"
)

func NextDateTime(num int) time.Time {
	t := time.Now()
	return t.AddDate(0, 0, num)
}

func NextDate(num int) time.Time {
	val := time.Now().AddDate(0, 0, num).Format("2006-01-02")
	ttime, _ := time.ParseInLocation("2006-01-02", val, time.Local)
	return ttime
}

func Today() time.Time {
	val := time.Now().Format("2006-01-02")
	ttime, _ := time.ParseInLocation("2006-01-02", val, time.Local)
	return ttime
}

func WeekNum() int {
	week := time.Now().Weekday().String()
	switch strings.ToLower(week) {
	case "monday":
		return 1
	case "tuesday":
		return 2
	case "wednesday":
		return 3
	case "thursday":
		return 4
	case "friday":
		return 5
	case "saturday":
		return 6
	case "sunday":
		return 7
	}
	return 7
}
