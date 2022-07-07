package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 12345678901
	bytes := []byte(strconv.FormatInt(num, 10))
	fmt.Printf("bytes: %v %s\n", bytes, string(bytes))
	bytes = []byte(strconv.FormatInt(num, 2))
	fmt.Printf("bytes: %v %s\n", bytes, string(bytes))
	bytes = []byte(strconv.FormatInt(num, 16))
	fmt.Printf("bytes: %v %s %c %d\n", bytes, string(bytes), string(bytes)[1], bytes[1])
}
