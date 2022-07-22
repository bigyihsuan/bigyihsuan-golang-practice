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
	var d1, d2 time.Duration
	var d1str, d2str string
	var err error
	for {
		for {
			fmt.Print("Enter duration 1 (DdHhMmSs): ")
			d1str, _ = reader.ReadString('\n')
			d1str = strings.TrimSpace(d1str)
			d1, err = toDuration(d1str)
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
				total += d2
				break
			case "-":
				total -= d2
				break
			default:
				fmt.Printf("Invalid operation `%v`, enter again: ", operation)
			}
		}
		fmt.Printf("%v %s %v = \n", d1str, operation, d2str)
		result := toDHMS(total)
		fmt.Println(result)
		fmt.Println()
	}
}

type DurationResult struct {
	days, hours, minutes, seconds int
}

func (dr DurationResult) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%vd%vh%vm%vs\n", dr.days, dr.hours, dr.minutes, dr.seconds))
	builder.WriteString(fmt.Sprintf("%f days\n", dr.AsDays()))
	builder.WriteString(fmt.Sprintf("%f hours\n", dr.AsHours()))
	builder.WriteString(fmt.Sprintf("%f minutes\n", dr.AsMinutes()))
	builder.WriteString(fmt.Sprintf("%f seconds", dr.AsSeconds()))
	return builder.String()
}

func (dr DurationResult) AsDays() float64 {
	return float64(dr.days) + float64(dr.hours)/24 + float64(dr.minutes)/(24*60) + float64(dr.seconds)/(24*60*60)
}

func (dr DurationResult) AsHours() float64 {
	return dr.AsDays() * 24
}

func (dr DurationResult) AsMinutes() float64 {
	return dr.AsHours() * 60
}

func (dr DurationResult) AsSeconds() float64 {
	return dr.AsMinutes() * 60
}

func toDHMS(t time.Duration) DurationResult {
	days := int(t.Truncate(time.Hour*24) / (time.Hour * 24))
	hours := int((t.Truncate(time.Hour) - t.Truncate(time.Hour*24)) / time.Hour)
	minutes := int((t.Truncate(time.Minute) - t.Truncate(time.Hour)) / time.Minute)
	seconds := int((t.Truncate(time.Second) - t.Truncate(time.Minute)) / time.Second)
	return DurationResult{
		days:    days,
		hours:   hours,
		minutes: minutes,
		seconds: seconds,
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
