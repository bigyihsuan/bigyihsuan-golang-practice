package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter year: ")
	yearstr, _ := reader.ReadString('\n')
	yearstr = strings.TrimSpace(yearstr)
	year, _ := strconv.Atoi(yearstr)
	fmt.Println(year, isLeapYear(year))
}

func isLeapYear(year int) bool {
	switch {
	case year%4 != 0:
		return false
	case year%100 != 0:
		return true
	case year%400 != 0:
		return false
	default:
		return true
	}
}
