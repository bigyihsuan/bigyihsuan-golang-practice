package main

import (
	"fmt"
	"regexp"
)

func main() {
	// https://www.w3resource.com/python-exercises/re/
	regexes := map[int]string{
		1:  `[A-Za-z0-9]+`,
		2:  `ab*`,
		3:  `ab+`,
		4:  `ab?`,
		5:  `ab{3}`,
		6:  `ab{2,3}`,
		7:  `[a-z]+_[a-z]+`,
		8:  `[A-Z][a-z]*`,
		9:  `a.*b$`,
		10: `^[A-Za-z]+`,
		11: `[A-Za-z]+[[:punct:]]*$`,
		12: `z`,
		13: `\Bz\B`,
		14: `\w+`,
		15: `^[0-9]`,
	}
	fmt.Printf("regexes: %v\n", regexes)
}

func testRegex01To15(regexes map[int]string, s string) {
	for _, reg := range regexes {
		regexp.MatchString(reg, s)
	}
}

func regex16(s string) string {
	// remove leading 0s from an ip address
	return ""
}
