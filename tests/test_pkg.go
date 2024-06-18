package main

import (
	"fmt"
	"time"

	pkg "startup/_pkg"
)

func main2() {
	fmt.Println("main")
	fmt.Println(FromISOString("2024-04-10T12:59:55+08:00"))
	fmt.Println(FromISOString("2006-01-02T15:04:05Z"))

	fmt.Println(pkg.ToFen(0 / float64(1)))
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
