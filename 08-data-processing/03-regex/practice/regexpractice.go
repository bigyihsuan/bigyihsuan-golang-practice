package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "abc123in"
	threeNs := "abcinnnnnn"
	fmt.Printf("isAlNum(\"abc123in\"):       %v\n", isAlNum(str))
	fmt.Printf("isIThenNs(\"abc123in\"):     %v\n", isIThenNs(str))
	fmt.Printf("isIThen1or2Ns(\"abc123in\"): %v\n", isIThen1or2Ns(str))
	fmt.Printf("isIThen1or2Ns(\"abc123in\"): %v\n", isIThen1or2Ns(str))
	fmt.Printf("isIThen3Ns(\"abcinnnnnn\"):  %v\n", isIThen3Ns(threeNs))
	fmt.Printf("isIThen3Ns(\"abc123in\"):    %v\n", isIThen3Ns(str))
}

func isAlNum(s string) bool {
	ok, _ := regexp.MatchString("[A-Za-z0-9]+", s)
	return ok
}

func isIThenNs(s string) bool {
	ok, _ := regexp.MatchString("in*", s)
	return ok
}
func isIThen1or2Ns(s string) bool {
	ok, _ := regexp.MatchString("in{1,2}", s)
	return ok
}
func isIThen3Ns(s string) bool {
	ok, _ := regexp.MatchString("in{3}", s)
	return ok
}
