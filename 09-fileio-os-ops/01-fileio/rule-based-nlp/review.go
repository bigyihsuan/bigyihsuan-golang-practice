package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// 1)Create Text file(Documants containing summary of context like movie review ,either in possitive or negetive )
// 2)Read the text file and depending upon the words available in the Document give a conclusion about review (Possitive/negetive)
// in terms of % of possitive/negetive
// 3)Give high priority to Awesome than Good , and then batter and so on

func main() {
	negativeReview, negerr := NewArticleFromFile("negativeReview.txt")
	if negerr != nil {
		fmt.Println(negerr)
		return
	}
	positiveReview, poserr := NewArticleFromFile("positiveReview.txt")
	if poserr != nil {
		fmt.Println(poserr)
		return
	}
	negativeWords, negwordserr := NewArticleFromFile("negative-words.txt")
	if negwordserr != nil {
		fmt.Println(negwordserr)
		return
	}
	positiveWords, poswordserr := NewArticleFromFile("positive-words.txt")
	if poswordserr != nil {
		fmt.Println(poswordserr)
		return
	}
	fmt.Println("negative review score:", negativeReview.Score(positiveWords, negativeWords))
	fmt.Println("positive review score:", positiveReview.Score(positiveWords, negativeWords))
}

type Article struct {
	Original   string
	WordCounts map[string]int
}

func NewArticleFromFile(filename string) (*Article, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lowered := strings.ToLower(string(bytes))
	notalphaspace := regexp.MustCompile("[^A-Za-z\\- \n]")
	nospace := notalphaspace.ReplaceAllLiteralString(lowered, "")
	nospace = strings.ReplaceAll(nospace, "-", " ")
	nonewline := strings.ReplaceAll(nospace, "\n", " ")
	words := strings.Split(nonewline, " ")
	article := new(Article)
	article.Original = lowered
	article.WordCounts = CountWords(words)
	return article, nil
}

func CountWords(words []string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range words {
		if len(word) > 0 {
			wordCount[word] += 1
		}
	}
	return wordCount
}

func (a Article) Score(positiveWords *Article, negativeWords *Article) float64 {
	posWords := positiveWords.WordCounts
	negWords := negativeWords.WordCounts
	var positiveScore, negativeScore, totalWordCount int
	for word, count := range a.WordCounts {
		if _, ok := posWords[word]; ok {
			positiveScore += count
		} else if _, ok := negWords[word]; ok {
			negativeScore += count
		}
		totalWordCount += count
	}
	return float64(positiveScore-negativeScore) / float64(totalWordCount)
}
