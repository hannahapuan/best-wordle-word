package main

import (
	"container/heap"
	"fmt"

	"github.com/hannahapuan/best-wordle-word/ref"
)

// best wordle word finds the best wordle word by comparing:
// 1) the cardinality of each letter in each position in the wordle dictionary list
// 2) the overall possibility of a letter being in any solution

func main() {
	// create cardinality map of each occurrence of a letter in each position
	cardLetterMap := initializeCardMapByPosition(ref.WordleList)

	// create map of each letter's total cardinality
	cardWordMap := initializeCardMapByWord(ref.WordleList)

	// create map of each letter's total cardinality utilizing the cardinality maps by position and in all positions
	wordScores := make(map[string]int)
	for _, word := range ref.WordleList {
		if !checkValid(word) {
			continue
		}
		wordScores[word] = determineWordScore(word, cardLetterMap, cardWordMap)
	}

	// sort into a priority queue where score is the priority
	ws := sortWordScores(wordScores)

	fmt.Println("Word\tScore")
	fmt.Println("----\t-----")
	ws.Print()
}

// checkValid uses the ruleset set in ref/rules.go to determine if a word is valid
// suffix rules contains strings that the word cannot end with, currently "s" and "ed"
func checkValid(word string) bool {
	// check if word is valid
	for _, rule := range ref.SuffixRules {
		if word[len(word)-len(rule):] == rule {
			return false
		}
	}
	return true
}

// initializeCardMap creates a map of each letter's cardinality in each position
// e.g.
//
//	a: {0: 1, 1: 1, 2: 1, 3: 1, 4: 1}
//	b: {0: 0, 1: 0, 2: 0, 3: 0, 4: 0}
func initializeCardMapByPosition(wordList []string) map[string]map[int]int {
	cardLetterMap := make(map[string]map[int]int)
	for _, word := range wordList {
		for i, letter := range word {
			l := string(letter)
			if _, ok := cardLetterMap[l]; !ok {
				cardLetterMap[l] = make(map[int]int)
			}
			cardLetterMap[l][i] = cardLetterMap[l][i] + 1
		}
	}

	return cardLetterMap
}

// initializeCardMap creates a map of each letter's cardinality in all words
// e.g.
//
//	a: 245, b: 123, c: 12
func initializeCardMapByWord(wordList []string) map[string]int {
	cardWordMap := make(map[string]int)
	for _, word := range wordList {
		for _, letter := range word {
			l := string(letter)
			cardWordMap[l] += 1
		}
	}

	return cardWordMap
}

// determineWordScore determines the score by
// 1) adding the cardinality of each letter in each position
// 2) adding the cardinality of each letter in all words
func determineWordScore(word string, cardLetterMap map[string]map[int]int, cardWordMap map[string]int) int {
	var score int
	for j, letter := range word {
		l := string(letter)
		score += cardLetterMap[l][j] + cardWordMap[l]
	}

	return score
}

// sortWordScores sorts the word scores into a priority queue
func sortWordScores(wordScores map[string]int) ref.PriorityQueue {
	sortedWords := make(ref.PriorityQueue, 0)

	for word, score := range wordScores {
		heap.Push(&sortedWords, &ref.Item{Value: word, Priority: score, Index: 0})
	}
	return sortedWords
}

// func findMax(wordScores map[string]int) string {
// 	var max int
// 	var maxWord string
// 	for word, score := range wordScores {
// 		if score > max {
// 			max = score
// 			maxWord = word
// 		}
// 	}
// 	return maxWord
// }
