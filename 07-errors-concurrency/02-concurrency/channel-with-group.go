package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	printNow chan bool
	i        int
	group sync.WaitGroup
)

func main() {
	printNow = make(chan bool)
	group.Add(2)
	go printer()
	go sender()
	group.Wait()
	fmt.Println("finished")
}
func printer() {
	// for {
	for j := 0; j < 10; j++ {
		// wait for something to come in on printnow
		if _, ok := <-printNow; ok {
			fmt.Println("Recieved !", i, j)
		}
		// else {
		// 	return
		// 	//   os.Exit(0)
		// }
	}
	group.Done()
	// }
}
func sender() {
	// for {
	for i = 0; i < 10; i++ {
		fmt.Println("Call", i)
		printNow <- true // send "signal" to print something
		time.Sleep(1 * time.Millisecond)
	}
	group.Done()
	// close(printNow)
	//return
	// }
}