package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	rand.Seed(time.Now().Unix())
	var numbers []int
	n := 20
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			mutex.Lock()
			appendRandom(&numbers)
			fmt.Printf("%v", numbers)
			sortNums(&numbers)
			fmt.Printf(" => %v\n", numbers)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

func appendRandom(numbers *[]int) {
	*numbers = append(*numbers, rand.Intn(1000))
}

func sortNums(numbers *[]int) {
	sort.Ints(*numbers)
}
