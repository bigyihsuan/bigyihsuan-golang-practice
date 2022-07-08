package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Start with a paragraph of your choice. You may use any source for the paragraph.
// Convert all text to lowercase characters.
// Split the lowercase string into individual words.
// Count the number of words in the lowercase string.
// Determine the number of distinct words in the lowercase string.
// Calculate the number of times each word appears in the lowercase string.
// Use the appropriate data structure to store the words and their frequency of occurrence.
// Remove the punctuation from the lowercase string.
// Perform a count analysis on the text without punctuation characters.

var lorem = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus neque nisl, vehicula ut ipsum ut, ultricies imperdiet eros. Pellentesque sed fermentum dolor. Pellentesque lobortis tincidunt nulla, a lacinia lectus ullamcorper vitae. Maecenas lobortis ante id neque venenatis dignissim et vel diam. Maecenas vitae mattis elit. Fusce dolor mi, fermentum non libero vitae, elementum tristique nisi. Interdum et malesuada fames ac ante ipsum primis in faucibus. Proin rhoncus semper nisi, nec consectetur eros tempor eget. Mauris ac pellentesque lectus, finibus mattis leo. Donec quis risus pretium, cursus magna sed, ornare nisi. Sed eu magna ut elit tempor facilisis sed id felis. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum consequat pharetra mauris, nec laoreet ipsum viverra vitae. Vivamus auctor id leo in efficitur. Duis porta leo quis laoreet pretium. Sed id malesuada tortor.
`

func stripPunc(s string) string {
	var punctless []rune
	for _, c := range s {
		if !unicode.IsPunct(c) {
			punctless = append(punctless, c)
		}
	}
	return string(punctless)
}

func countWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word] += 1
	}
	return counts
}

func keys(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	var noPunctuation = stripPunc(lorem)
	var lowercase = strings.ToLower(noPunctuation)
	var words = strings.Split(lowercase, " ")
	var wordCount = len(words)

	var wordFrequency = countWords(words)
	var uniqueWords = keys(wordFrequency)

	fmt.Printf("noPunctuation: %v\n", noPunctuation)
	fmt.Printf("lowercase: %v\n", lowercase)
	fmt.Printf("words: %v\n", words)
	fmt.Printf("wordCount: %v\n", wordCount)
	fmt.Printf("wordFrequency: %v\n", wordFrequency)
	fmt.Printf("uniqueWords: %v\n", uniqueWords)
}
