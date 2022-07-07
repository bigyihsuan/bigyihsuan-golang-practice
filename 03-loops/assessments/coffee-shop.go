//! Yi-Hsuan Hsu
//! 6/30/22
//! C277 Golang
//! Assessment: Coffee Shop

package main

import (
	"fmt"
	"strings"
)

func main() {
	const SIZE_PROMPT string = "Do you want small, medium, or large? "
	const KIND_PROMPT string = "Do you want brewed, expresso, or cold press? "
	const FLAVOR_PROMPT string = "Do you want flavored syrup? (yes or no) "
	const FLAVOR_SELECTION string = "Do you want hazelnut, vanilla, or caramel? "
	const NO_FLAVOR string = "No flavor selected."
	const ORDER_SUMMARY string = "You asked for a %s cup of %s coffee with %s syrup.\n"
	const CUP_PRICE string = "Your cup of coffee costs $%.02f\n"
	const CUP_AND_TIP_PRICE string = "The price with tip costs $%.02f\n"

	var size, kind, flavor string
	var temp string
	var price float64 = 0

	// ask size of cup: small, medium, large
	fmt.Print(SIZE_PROMPT)
	fmt.Scan(&size)
	size = strings.ToLower(size)

	// ask kind of coffee: brewed, expresso, cold press
	fmt.Print(KIND_PROMPT)
	fmt.Scanf("%s %s\n", &kind, &temp)
	kind = strings.ToLower(kind)
	if kind == "cold" {
		// handle Scan() and Scanln() stopping at the first space
		kind += temp
	}

	// ask flavoring (optional): hazelnut, vanilla, caramel (none)
	fmt.Print(FLAVOR_PROMPT)
	fmt.Scan(&flavor)
	flavor = strings.ToLower(flavor)
	if flavor == "yes" {
		fmt.Print(FLAVOR_SELECTION)
		fmt.Scan(&flavor)
		flavor = strings.ToLower(flavor)
		// handle "none" case for flavor
		switch flavor {
		case "hazelnut", "vanilla", "caramel":
			// ok
		default:
			flavor = "none"
		}
	} else {
		fmt.Println(NO_FLAVOR)
		flavor = "none"
	}

	// calculate price:
	// * small = 2, medium = 3, large = 4
	switch size {
	case "small":
		price += 2
	case "medium":
		price += 3
	case "large":
		price += 4
	}

	// * brewed = 0, expresso = 0.50, cold brew = 1
	switch kind {
	case "brewed":
		// price += 0
	case "expresso":
		price += 0.50
	case "cold press":
		price += 1
	}

	// * no flavor = 0, else = 0.50
	switch flavor {
	case "none":
		// price += 0
		flavor = "no"
	default:
		price += 0.50
	}

	// show staterment that summarizes order
	fmt.Printf(ORDER_SUMMARY, size, kind, flavor)

	// show total cost of cup of coffee, and cost with 15% tip (round to 2 decimals)
	fmt.Printf(CUP_PRICE, price)
	fmt.Printf(CUP_AND_TIP_PRICE, price+0.15*price)
}
