package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("15:04:05"))
	fmt.Println(now.Add(time.Second * 5).Format("15:04:05"))
}
