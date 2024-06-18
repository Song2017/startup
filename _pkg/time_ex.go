package pkg

import (
	"fmt"
	"strconv"
	"time"
)

var (
	OneHour = time.Duration(1) * time.Hour

	TimeFormatWithZone = "2006-01-02T15:04:05-07:00"
)

func GetTimeOutSeconds(durations ...int) time.Duration {
	if len(durations) == 0 {
		return time.Duration(5) * time.Second
	}
	return time.Duration(durations[0]) * time.Second
}

func GetTimestamp(t *time.Time) string {
	if t == nil {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return strconv.FormatInt(t.Unix(), 10)
}

func TimeFormat(t time.Time, location ...string) string {
	loc, err := time.LoadLocation("Asia/Shanghai")
	timeFormat := TimeFormatWithZone
	if len(location) == 1 {
		loc, err = time.LoadLocation(location[0])
	} else if len(location) == 2 {
		loc, err = time.LoadLocation(location[0])
		timeFormat = location[1]
	}
	if err != nil {
		fmt.Println("TimeFormat err----", err, err.Error())
	}

	return t.In(loc).Format(timeFormat)
}

func GetDay(t *time.Time) string {
	if t == nil {
		return time.Now().Format("20060102")
	}
	return t.Format("20060102")
}
func GetMin(t *time.Time) string {
	if t == nil {
		return time.Now().Format("200601021504")
	}
	return t.Format("200601021504")
}
func GetSec(t *time.Time) string {
	if t == nil {
		return time.Now().Format("2006-01-02T15:04:05")
	}
	return t.Format("2006-01-02T15:04:05")
}

func FromISOString(strTime string) time.Time {
	// 2024-04-10T12:59:55+08:00
	t, err := time.Parse(time.RFC3339, strTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Now().AddDate(-1, 0, 0)
	}
	return t
}
