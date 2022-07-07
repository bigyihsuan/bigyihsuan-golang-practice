/*
Create a map that includes at least ten key:value pairs.
The pairs should make sense, but you can choose whatever you wish.
Suggestions include
	produce categories (e.g., apple:fruit and onion:vegetable),
	animal categories (e.g., mammals, birds, fish),
	city populations,
	or word definitions.
Prompt the user to input a search term.
Display all key:value pairs that include the input search term,
	either in the key or in the value.
For example, if the topic is produce,
	the user can enter either apple or fruit and see all matching map entries.
If the value does not exist in the dictionary,
	display a user-friendly error message.
Prompt the user to start over again or exit the program.
*/

package main

import (
	"fmt"
	"strings"
)


var produce = map[string]string{
	"apple":  "fruit",
	"onion":  "vegetable",
	"wheat":  "grain",
	"orange": "fruit",
	"carrot": "vegetable",
	"rice":   "grain",
	"pear":   "fruit",
	"garlic": "vegetable",
	"rye":    "grain",
	"banana": "fruit",
	"tomato": "vegetable",
}

var foodTypes = reversedMap(produce)

func reversedMap(mapping map[string]string) map[string][]string {
	newMap := make(map[string][]string)
	for k, v := range mapping {
		// eles := newMap[v]
		fmt.Println(newMap[v])
		newMap[v] = append(newMap[v], k)
		fmt.Println(newMap[v])
	}
	return newMap
}

func prettify(elements []string) string {
	out := ""
	for i, ele := range elements {
		if i != len(elements)-1 {
			out += fmt.Sprintf("%s, ", ele)
		} else {
			out += fmt.Sprintf("and %s", ele)
		}
	}
	return out
}

func main() {
	var response string
	for {
		fmt.Print("Enter fruit name or produce type (fruit, vegetable, grain): ")
		fmt.Scan(&response)
		response = strings.ToLower(response)

		if category, inMap := produce[response]; inMap {
			// response is a food
			fmt.Printf("%v is a %v\n", response, category)
		} else if fruits, inMap := foodTypes[response]; inMap {
			// response is a category
			fmt.Printf("%v are %vs\n", prettify(fruits), response)
		} else {
			// restart
			fmt.Printf("%v is neither a fruit nor a category\n", response)
		}
		fmt.Print("exit? (y to exit) ")
		fmt.Scanln(&response)
		response = strings.ToLower(response)
		if response == "y" {
			fmt.Println("exiting...")
			break
		}
	}
}
