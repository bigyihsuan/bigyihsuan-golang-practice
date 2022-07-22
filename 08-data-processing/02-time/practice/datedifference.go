package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date 1 (YYYY-MM-DD): ")
	d1str, _ := reader.ReadString('\n')
	d1str = strings.TrimSpace(d1str)
	d1, _ := time.Parse("2006-01-02", d1str)

	fmt.Print("Enter date 2 (YYYY-MM-DD): ")
	d2str, _ := reader.ReadString('\n')
	d2str = strings.TrimSpace(d2str)
	d2, _ := time.Parse("2006-01-02", d2str)

	if d1.Before(d2) {
		diff := d2.Sub(d1)
		fmt.Println("day difference:", durationToDays(diff))
	} else {
		diff := d1.Sub(d2)
		fmt.Println("day difference:", durationToDays(diff))
	}
}

func durationToDays(durr time.Duration) int {
	return int(durr.Round(time.Hour*24) / (time.Hour * 24))
}
