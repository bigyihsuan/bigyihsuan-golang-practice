package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var port = flag.String("port", "8000", "The port for the server")
var timeZoneOffset = flag.String("zone", "", "Timezone location")

func main() {
	flag.Parse()
	fmt.Printf("Listening on port %v\n", *port)
	fmt.Printf("Using timezone %v\n", *timeZoneOffset)
	os.Setenv("TZ", *timeZoneOffset)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("Got connection from %v\b", conn.RemoteAddr())
	for {
		_, err := io.WriteString(conn, fmt.Sprintf("%v %v", os.Getenv("TZ"), time.Now().Format("15:04:05\n")))
		if err != nil {
			log.Printf("Connection to %v closed\n", conn.RemoteAddr())
			return
		}
		time.Sleep(time.Second)
	}

}
