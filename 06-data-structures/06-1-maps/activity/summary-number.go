/*
2) Take 20 (any count) from console between 1 -100 )
then print the summary like
	no between
		1- 10 Count-5
		11 -20 count -7
		21-30 count -10
	etc
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

type Buckets map[int]int
type Keys []int

func (b Buckets) keys() Keys {
	var keys Keys
	for k := range b {
		keys = append(keys, k)
	}
	return keys
}
func (b Buckets) SortedKeys() Keys {
	keys := b.keys()
	sort.Ints(keys)
	return keys
}
func (b Buckets) summary() {
	fmt.Println("Numbers between:")
	for _, k := range b.SortedKeys() {
		base := k * 10
		fmt.Printf("    %2v-%2v: %v\n", base, base+9, b[k])
	}
}

func main() {
	buckets := make(Buckets)
	for i := 0; i < 10; i++ {
		buckets[i] = 0
	}

	for i := 0; i < 20; i++ {
		var num int
		fmt.Printf("Enter number 0<=n<100 #%v: ", i)
		fmt.Scan(&num)
		key := int(math.Abs(float64(num)))
		key %= 100
		key /= 10
		num %= 100
		buckets[key] += 1
	}
	buckets.summary()
}
