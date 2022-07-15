package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vehicula quis nunc iaculis tempus. Nullam et neque a dui placerat malesuada. Sed sit amet sollicitudin neque. Mauris vel cursus augue. Praesent nibh felis, consequat et libero non, viverra cursus mauris. Phasellus sodales leo et venenatis lobortis. Sed eu placerat mauris, ut interdum est. Pellentesque urna odio, venenatis eu sem at, scelerisque rutrum erat. In at convallis enim. Interdum et malesuada fames ac ante ipsum primis in faucibus. Phasellus ut nulla a tortor bibendum ornare. Ut nec nisl condimentum, interdum tellus tincidunt, imperdiet metus. Praesent egestas dapibus mi, fermentum aliquam libero placerat id. Pellentesque velit nisl, placerat in dignissim non, commodo sed lacus. Praesent congue tortor nisl, eget sollicitudin ante rutrum et. Nullam gravida lectus nibh, a consectetur ex vehicula non. Praesent a dui ligula. Donec vel consectetur massa, vel elementum tellus. Etiam vel ligula eu turpis iaculis fringilla id vitae nisi. Maecenas volutpat, nulla ut tempus dictum, augue ligula feugiat urna, in sagittis massa lorem a leo. Integer eleifend vestibulum egestas. Etiam sed odio vel risus tincidunt cursus et id leo. Nulla facilisi. Morbi vehicula velit vitae felis condimentum laoreet. Nam ultricies massa ut nisl sollicitudin mollis. Quisque nec varius.`
	lorem = strings.ToLower(lorem)
	var sentences []string
	for _, sentence := range strings.SplitAfter(lorem, ".") {
		sentences = append(sentences, strings.TrimSpace(sentence))
	}
	// fmt.Println(strings.Join(sentences, "\n"))
	for i, sentence := range sentences {
		wg.Add(2)
		go reverse(sentence, i)
		count := countWords(sentence, i)
		fmt.Printf("COUNT %v = %v\n", i, <-count)
	}
	wg.Wait()
}

func splitWords(sentence string) []string {
	return strings.Split(sentence, " ")
}

func reverse(sentence string, id int) {
	words := splitWords(sentence)
	fmt.Printf("REVERSED %v: ", id)
	for i := len(words) - 1; i >= 0; i-- {
		fmt.Printf("%v ", words[i])
	}
	fmt.Println()
	wg.Done()
}

func countWords(sentence string, id int) chan int {
	numWords := make(chan int)
	go func() {
		words := splitWords(sentence)
		numWords <- len(words)
		close(numWords)
		wg.Done()
	}()
	return numWords
}
