package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	yesterday := today.AddDate(0, 0, -1)

	fmt.Printf("yesterday: %v\n", yesterday.Format(time.RFC822))
	fmt.Printf("today:     %v\n", today.Format(time.RFC822))
	fmt.Printf("tomorrow:  %v\n", tomorrow.Format(time.RFC822))
}
