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
	var d1, d2 time.Time
	var d1str, d2str string
	var err error
	for {
		for {
			fmt.Print("Enter birthday: ")
			d1str, _ = reader.ReadString('\n')
			d1str = strings.TrimSpace(d1str)
			d1, err = time.Parse("January 2, 2006", d1str)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		for {
			fmt.Print("Enter date: ")
			d2str, _ = reader.ReadString('\n')
			d2str = strings.TrimSpace(d2str)
			d2, err = time.Parse("January 2, 2006", d2str)
			if err != nil {
				fmt.Println(err)
			} else {
				break
			}
		}
		total := d2.Sub(d1)
		fmt.Printf("%v to %v = \n", d1str, d2str)
		result := toDHMS(total)
		fmt.Println(result)
		fmt.Println()
	}
}

type DurationResult struct {
	years, months, days float64
}

func (dr DurationResult) String() string {
	var builder strings.Builder
	builder.WriteString(dr.AsYearMonthDay())
	builder.WriteString(dr.AsMonthDay())
	builder.WriteString(dr.AsWeekDay())
	builder.WriteString(fmt.Sprintf("%d days\n", int(dr.AsDays())))
	builder.WriteString(fmt.Sprintf("%d hours\n", int(dr.AsHours())))
	builder.WriteString(fmt.Sprintf("%d minutes\n", int(dr.AsMinutes())))
	builder.WriteString(fmt.Sprintf("%d seconds\n", int(dr.AsSeconds())))
	return builder.String()
}

func (dr DurationResult) AsYearMonthDay() string {
	return fmt.Sprintf("%v years, %v months, %v days\n", dr.years, dr.months, dr.days)
}
func (dr DurationResult) AsMonthDay() string {
	return fmt.Sprintf("%v months, %v days\n", dr.years*12+dr.months, dr.days)
}
func (dr DurationResult) AsWeekDay() string {
	return fmt.Sprintf("%v weeks, %v days\n", (dr.years*12+dr.months)*4, dr.days)
}

func (dr DurationResult) AsDays() float64 {
	return (dr.years*12+dr.months)*4*7 + dr.days
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
	timeDay := time.Hour * 24
	timeMonth := timeDay * 30
	timeYear := timeMonth*12 + timeDay*5

	years := float64(t.Truncate(timeYear) / timeYear)
	months := float64((t.Truncate(timeMonth) - t.Truncate(timeYear)) / timeMonth)
	days := float64((t.Truncate(timeDay) - t.Truncate(timeMonth)) / (timeDay))
	return DurationResult{
		years:  years,
		months: months,
		days:   days,
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
