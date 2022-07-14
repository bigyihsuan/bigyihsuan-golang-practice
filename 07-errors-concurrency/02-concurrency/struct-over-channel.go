package main

import (
	"fmt"
	"time"
)

type Packet struct {
	Value       string
	TimesPassed int
}

func main() {
	sendChannel := make(chan Packet, 1)
	receiveChannel := make(chan Packet, 1)

	// go send(sendChannel, receiveChannel)
	// go get(receiveChannel, sendChannel)
	go mix(sendChannel, receiveChannel)
	sendChannel <- Packet{"hello", 0}
	fmt.Printf("sender: %v\n", sendChannel)
	fmt.Printf("receiver: %v\n", receiveChannel)

	// for {
	// 	select {
	// 	case data := <-sendChannel:
	// 		fmt.Printf("main sendChannel data: %v\n", data)
	// 	case data := <-receiveChannel:
	// 		fmt.Printf("main receiveChannel data: %v\n", data)
	// 	}
	// }
	select {}
}

func mix(left, right chan Packet) {
	for {
		select {
		case data := <-left: // when this can read from left, do stuff
			fmt.Println("got left", data)
			time.Sleep(time.Second)
			right <- Packet{"from left", data.TimesPassed + 1}
		case data := <-right: // when this can read from right, do stuff
			fmt.Println("got right", data)
			time.Sleep(time.Second)
			left <- Packet{"from right", data.TimesPassed + 1}
		}
	}
}

// send() reads from input, and writes to output
func send(input <-chan Packet, output chan<- Packet) {
	for {
		// read from the send channel
		data := <-input
		fmt.Printf("data from send channel: %v\n", data)
		// manipulate data
		data.TimesPassed += 1
		// push the data out
		time.Sleep(time.Second / 4)
		fmt.Println("    send sending")
		output <- data
	}
}

// reads from input, resends it to output
func get(input <-chan Packet, output chan<- Packet) {
	for {
		// read from the send channel
		data := <-input
		fmt.Printf("data from receive channel: %v\n", data)
		// manipulate data
		data.TimesPassed += 1
		// push the data out
		time.Sleep(time.Second / 4)
		fmt.Println("    get sending")
		output <- data
	}
}
