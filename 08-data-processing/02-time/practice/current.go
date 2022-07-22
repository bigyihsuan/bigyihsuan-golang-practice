package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	_, week := now.ISOWeek()
	fmt.Println(week)
	fmt.Println(now.Weekday())
	fmt.Println(now.YearDay())
	fmt.Println(now.Day())
	fmt.Println(int(now.Weekday()))
}
