package main

/*
3) Create a map, sort the map based on key length in asecending order
*/

import (
	"fmt"
	"sort"
)

type PersonInfo struct {
	Address string
	ZipCode int
}

type Contacts map[string]PersonInfo
type Keys []string

func (k Keys) Len() int {
	return len(k)
}
func (k Keys) Less(i,j int) bool {
	return len(k[i]) < len(k[j])
}
func (k Keys) Swap(i,j int) {
	k[i], k[j] = k[j], k[i]
}


func (c Contacts) keys() Keys {
	var keys Keys
	for k := range c {
		keys = append(keys, k)
	}
	return keys
}
func (c Contacts) SortedKeys() Keys {
	keys := c.keys()
	sort.Sort(keys)
	return keys
}

func main() {
	contactInfo := Contacts{
		"patricia":  {"129 somewhere road", 65677},
		"jack": {"1 court", 1010},
		"michael": {"500 industrial park", 89892},
		"bob":  {"123 main", 890},
		"claus": {"9999 broad st", 1234},
	}
	sorted := contactInfo.SortedKeys()
	for _, k := range sorted {
		fmt.Println(k, contactInfo[k])
	}
}
// for k, v := range contactInfo {
// 	fmt.Println(k, v)
// }
// fmt.Println()
// fmt.Println()