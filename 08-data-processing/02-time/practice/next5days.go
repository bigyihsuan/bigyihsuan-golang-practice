package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	for i := 0; i <= 5; i++ {
		fmt.Printf("today + %v days = %s\n", i, today.AddDate(0, 0, i).Format("2006-01-02"))
	}
}
