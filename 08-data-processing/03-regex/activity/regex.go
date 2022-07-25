package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Assignment 1) Find a valid mobile no 10 digit long and begin with your country code
	countrycode := regexp.MustCompile(`\+1`)
	usa := "+18005555555"
	notusa := "+671234567891"

	fmt.Printf("countrycode: %v\n", countrycode)
	okusa := countrycode.MatchString(usa)
	fmt.Println(usa, okusa)
	oknotusa := countrycode.MatchString(notusa)
	fmt.Println(notusa, oknotusa)

	// Assignment 2) Search for IP address of range 192.160.1 - 192.170.1
	// 192.160.1.1    - 192.170.1.255
	matches1To255 := `[1-9]|[1-9][0-9]|[1-2][0-4][0-9]|25[0-5]`
	matches160To170 := `16[0-9]|170`
	ip := regexp.MustCompile(fmt.Sprintf("192\\.(%s)\\.(%s)\\.(%s)", matches160To170, matches1To255, matches1To255))
	goodip := `192.160.222.5`
	badip := `192.180.232.23`

	fmt.Printf("ip: %v\n", ip)
	fmt.Println(goodip, ip.MatchString(goodip))
	fmt.Println(badip, ip.MatchString(badip))

	// Next Assignment Find any three-letter words that start with the same letter and end with the same letter,
	// but which might have a different letter in between, such as cat or cot.

	// // NOT POSSIBLE IN GOLANG WITH `regexp`, but below is if it *did* work:
	// cac := regexp.MustCompile(`(.).(\1)`) // any char, then any char, then the first char
	// fmt.Println("cat", cac.MatchString("cat"))
	// fmt.Println("man", cac.MatchString("man"))
	// fmt.Println("bob", cac.MatchString("bob"))
	// fmt.Println("pop", cac.MatchString("pop"))

	// any char, then any char, then the first char
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string: ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	reg := fmt.Sprintf("^%[1]s.*%[1]s$", string(s[0]))
	for {
		fmt.Print("Enter string: ")
		k, _ := reader.ReadString('\n')
		k = strings.TrimSpace(k)
		ok, err := regexp.MatchString(reg, k)
		fmt.Println(k, ok, err)
	}
}
