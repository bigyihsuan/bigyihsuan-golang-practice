package main

// 1)Create a map where values must be a structure type

import "fmt"

type PersonInfo struct {
	Address string
	ZipCode int
}

type Contacts map[string]PersonInfo

func main() {
	contactInfo := Contacts{
		"bob":  {"123 main", 890},
		"jack": {"1 court", 1010},
		"jill": {"500 industrial park", 89892},
		"rich": {"9999 broad st", 1234},
		"pat":  {"129 somewhere road", 65677},
	}

	for k, v := range contactInfo {
		fmt.Println(k, v)
	}
}