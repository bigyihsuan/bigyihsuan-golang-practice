package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	group   sync.WaitGroup
	channel chan string
)

func foo() {
	for i := 0; i <= 100; i++ {
		channel <- fmt.Sprintf("foo i=%v", i) // send message
		time.Sleep(time.Millisecond)
	}
	// channel <- fmt.Sprint("foo finish")
	fmt.Println("foo finish")
	group.Done()
}
func bar() {
	for i := 100; i >= 0; i-- {
		if s,ok := <- channel; ok { // receive message
			fmt.Println(s)
			// channel <- fmt.Sprintf("    bar i=%v", i)
			fmt.Printf("    bar i=%v\n", i)
			time.Sleep(time.Millisecond)
		}
	}
	// channel <- fmt.Sprint("    bar finish")
	fmt.Println("    bar finish")
	group.Done()
}

func main() {
	channel = make(chan string)
	group.Add(2)
	go foo()
	go bar()
	group.Wait()
	close(channel) // close it so nothing is waiting
	for s := range channel {
		fmt.Println(s)
	}
}

// package main
// 
// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
// 
// var (
// 	group   sync.WaitGroup
// 	channel chan string
// )
// 
// func Forword() {
// 	for i := 0; i <= 10; i++ {
// 		channel <- fmt.Sprintf("Forword i=%v", i) // send message
// 		time.Sleep(time.Millisecond)
// 	}
// 	
// 	fmt.Println("Forword finish its task ")
// 	group.Done()
// }
// //Printng both the values
// func Reverse () {
// 	for i := 10; i >= 0; i-- {
// 		if s,ok := <- channel; ok { // receive message
// 			fmt.Println(s) //Print the value of Forword
// 			
// 			fmt.Printf("   Reverse i=%v\n", i)// Print the value of Reverse 
// 			time.Sleep(time.Millisecond)
// 		}
// 	}
// 	
// 	fmt.Println(" Reverse finished task")
// 	group.Done()
// }
// 
// func main() {
// 	channel = make(chan string)
// 	group.Add(2)
// 	go Forword()
// 	go Reverse()
// 	group.Wait()
// 	close(channel) // close it so nothing is waiting
// 	for s := range channel {
// 		fmt.Println(s)
// 	}
// }
