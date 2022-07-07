// Yi-Hsuan Hsu
// C277 Golang
// 6/28/2022

package main

import (
	"fmt"
	"strings"
)

func main() {
	var numQuestions = 5
	var numCorrect = 0

	var r1, r2, r3, r4, r5 string
	fmt.Println("3 + 3 = 4?")
	fmt.Scan(&r1)
	if strings.ToLower(r1) == "false" {
		numCorrect += 1
	}
	fmt.Println("5 is a number?")
	fmt.Scan(&r2)
	if strings.ToLower(r2) == "true" {
		numCorrect += 1
	}
	fmt.Println("10 is divisible by 3?")
	fmt.Scan(&r3)
	if strings.ToLower(r3) == "false" {
		numCorrect += 1
	}
	fmt.Println("apples are fruit?")
	fmt.Scan(&r4)
	if strings.ToLower(r4) == "true" {
		numCorrect += 1
	}
	fmt.Println("1 - 1 != 0?")
	fmt.Scan(&r5)
	if strings.ToLower(r5) == "false" {
		numCorrect += 1
	}
	var percent = float32(numCorrect) / float32(numQuestions) * 100
	fmt.Println("Answers are: false, true, false, true, false")
	fmt.Printf("Answered:    %v, %v, %v, %v, %v\n", r1, r2, r3, r4, r5)
	fmt.Printf("Answer rate: %v/%v = %v%%\n", numCorrect, numQuestions, int(percent))

}
