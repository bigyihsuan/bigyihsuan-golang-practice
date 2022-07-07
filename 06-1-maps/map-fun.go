package main

import "fmt"

func mapFromFn(item map[string]float64) map[string]float64 {
	db := make(map[string]float64)
	for k, v := range item {
		db[k] = v * 2
	}
	return db
}

func main() {
	phones := make(map[string]float64)
	phones["iphone"] = 12345.67
	phones["android"] = 500.30

	updated := mapFromFn(phones)
	delete(updated, "iphone")
	for k, v := range updated {
		fmt.Println(k, v)
	}
}