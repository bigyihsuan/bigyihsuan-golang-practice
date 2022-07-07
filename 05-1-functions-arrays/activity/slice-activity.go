package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter number of words: ")
	var wordCount int
	fmt.Scan(&wordCount)
	words := make([]string, wordCount)

	for i := range words {
		var word string
		fmt.Printf("Enter word %v: ", i+1)
		fmt.Scan(&word)
		words[i] = strings.ToLower(word)
	}

	avgLen := avgWordLen(words)
	longWords := getWordsLongerThan(words, avgLen)
	shortWords := getWordsShorterThan(words, avgLen)

	fmt.Printf("words: %v\n", words)
	fmt.Printf("avgLen: %v\n", avgLen)
	fmt.Printf("longWords: %v\n", longWords)
	fmt.Printf("shortWords: %v\n", shortWords)
}

func avgWordLen(words []string) int {
	var lens int
	var count = len(words)

	for _, word := range words {
		lens += len(word)
	}
	return lens / count
}

func getWordsLongerThan(words []string, minLen int) []string {
	var longWords []string
	for _, word := range words {
		if len(word) > minLen {
			longWords = append(longWords, word)
		}
	}
	return longWords
}

func getWordsShorterThan(words []string, maxLen int) []string {
	var shortWords []string
	for _, word := range words {
		if len(word) < maxLen {
			shortWords = append(shortWords, word)
		}
	}
	return shortWords
}
