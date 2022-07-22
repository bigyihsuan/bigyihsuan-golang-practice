package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(unixToHuman(now))
}

func unixToHuman(unixTimestamp int64) string {
	return time.Unix(unixTimestamp, 0).Format(time.UnixDate)
}
