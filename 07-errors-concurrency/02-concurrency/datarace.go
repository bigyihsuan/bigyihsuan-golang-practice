package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().Unix())
	// these 5 all share global counter variable, causing a data race
	n := 25
	l := 5
	wg.Add(l)
	format := "% " + strconv.Itoa(l) + "v"
	for i := 0; i < l; i++ {
		go incrementor(fmt.Sprintf(format, strings.Repeat("*", i+1)), n)
	}
	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}

func incrementor(s string, n int) {
	for i := 0; i <= n; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
		// counter++
		atomic.AddInt64(&counter, 1)
		current := atomic.LoadInt64(&counter)
		fmt.Println(s, i, "counter", current)
	}
	wg.Done()
}
