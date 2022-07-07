package main

import "fmt"

func main() {
	var amount int
	var price int = 0

	fmt.Print("Enter amount: ")
	fmt.Scanf("%d", &amount)

	if 100 >= amount && amount >= 0 {
		price = amount * 5
	} else if 200 >= amount && amount > 100 {
		price = (amount-100)*7 + 100*5
	} else if amount > 200 {
		price += (amount-200)*10 + 100*7 + 100*5
	}
	fmt.Printf("amount: %v hours\n", amount)
	fmt.Printf("price: $%v\n", price)

}
