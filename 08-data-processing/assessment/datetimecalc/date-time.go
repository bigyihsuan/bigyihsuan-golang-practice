// Yi-Hsuan Hsu 7/25/22
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var startingDays = regexp.MustCompile("([0-9]+)d")

func main() {
	fmt.Println("Press CTRL-C to quit any time")
	reader := bufio.NewReader(os.Stdin)
	var d1 time.Time
	var d2 time.Duration
	var d1str, d2str string
	var err error
	for {
		for {
			fmt.Print("Enter date: ")
			d1str, _ = reader.ReadString('\n')
			d1str = strings.TrimSpace(d1str)
			d1, err = time.Parse("January 2, 2006, 03:04:05 PM", d1str)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		for {
			fmt.Print("Enter duration 2 (DdHhMmSs): ")
			d2str, _ = reader.ReadString('\n')
			d2str = strings.TrimSpace(d2str)
			d2, err = toDuration(d2str)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		total := d1

		operation := ""
		fmt.Print("Operation (+/-): ")
		for operation != "+" && operation != "-" {
			operation, _ = reader.ReadString('\n')
			operation = strings.TrimSpace(operation)
			switch operation {
			case "+":
				total = total.Add(d2)
				break
			case "-":
				total = total.Add(-d2)
				break
			default:
				fmt.Printf("Invalid operation `%v`, enter again: ", operation)
			}
		}
		fmt.Printf("%v %s %v = \n", d1str, operation, d2str)
		result := total
		fmt.Println(result.Format("January 2, 2006, 03:04:05 PM"))
		fmt.Println()
	}
}

func toDuration(str string) (time.Duration, error) {
	// get the days
	daystr := startingDays.FindString(str)
	days, err := strconv.Atoi(daystr[:len(daystr)-1])
	if err != nil {
		return 0, err
	}
	// convert days to hours
	dayHour := time.Hour * 24 * time.Duration(days)
	// strip off the days
	hms := strings.TrimPrefix(str, daystr)
	d, err := time.ParseDuration(hms)
	if err != nil {
		return 0, err
	}
	// add it on
	d = d + dayHour
	return d, nil
}
