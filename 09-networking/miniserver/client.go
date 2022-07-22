package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

// var timezone string

// type TimeZones []string

// func (tz *TimeZones) String() string {
// 	out := ""
// 	for _, e := range *tz {
// 		out += e
// 		out += " "
// 	}
// 	out = strings.TrimSpace(out)
// 	return out
// }

// func (tz *TimeZones) Set(value string) error {
// 	*tz = append(*tz, value)
// 	return nil
// }

// var zones TimeZones

var wg sync.WaitGroup

func main() {
	// flag.Var(&zones, "zone", "list of many ")
	// flag.StringVar(&timezone, "zone", "localhost:8000", "Any number of servers to listen to")
	flag.Parse()
	wg.Add(1)
	// for _, timezone := range append(flag.Args(), zones...) {
	// for _, timezone := range append(flag.Args(), zones...)
	for _, place := range os.Args[1:] {
		pair := strings.Split(place, "=")
		location, timezone := pair[0], pair[1]
		fmt.Printf("Listening to %v for %v\n", timezone, location)
		go func() {
			conn, err := net.Dial("tcp", timezone)
			if err != nil {
				log.Fatal(err)
			}
			shouldCopy(os.Stdout, conn)
		}()
	}
	wg.Wait()
}

func shouldCopy(destination io.Writer, source net.Conn) {
	defer source.Close()
	if _, err := io.Copy(destination, source); err != nil {
		log.Fatal(err)
		return
	}
}
