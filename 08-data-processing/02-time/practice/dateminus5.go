package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	minus := dateMinusFive()
	fmt.Println(minus.Format("2006-01-02"))
}

func dateMinusFive() time.Time {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date (YYYY-MM-DD): ")
	datestr, _ := reader.ReadString('\n')
	datestr = strings.TrimSpace(datestr)
	date, err := time.Parse("2006-01-02", datestr)
	if err != nil {
		date = time.Now()
	}
	return date.Add(-time.Hour * 24 * 6)
}
