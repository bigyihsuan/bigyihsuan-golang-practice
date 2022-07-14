package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

var group sync.WaitGroup
func main() {
	group.Add(2)
	go increment("foo")
	go increment("    bar")
	fmt.Println("main")
	group.Wait()
	time.Sleep(time.Second)
}

func increment(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		mutex.Lock()
		counter++
		time.Sleep(time.Millisecond + time.Duration(i))
		fmt.Println(s, i, "counter =", counter)
		mutex.Unlock()
	}
	group.Done()	
}

func foo() {
	for i := 0; i < 50; i++ {
		fmt.Printf("foo i=%v\n", i)
		time.Sleep(time.Millisecond)
	}
}

func bar() {
	for i := 0; i < 50; i++ {
		fmt.Printf("    bar i=%v\n", i)
		time.Sleep(time.Millisecond)
	}
}