package main

import (
	"container/heap"

	"github.com/hannahapuan/best-wordle-word/ref"
)

// best wordle word finds the best wordle word by comparing
// the cardinality of each letter in each position in the wordle dictionary list

func main() {
	// create cardinality map of each occurrence of a letter in each position
	cardMap := initalizeCardMap(ref.WordleList)

	// create map of each letter's total cardinality utilizing the cardinality map
	wordScores := make(map[string]int)
	for _, word := range ref.WordleList {
		if !checkValid(word) {
			continue
		}
		wordScores[word] = determineWordScore(word, cardMap)
	}

	// sort into a priority queue where score is the priority
	ws := sortWordScores(wordScores)
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

// initalizeCardMap creates a map of each letter's cardinality in each position
// e.g.
//
//	a: {0: 1, 1: 1, 2: 1, 3: 1, 4: 1}
//	b: {0: 0, 1: 0, 2: 0, 3: 0, 4: 0}
func initalizeCardMap(wordList []string) map[string]map[int]int {
	cardMap := make(map[string]map[int]int)
	for _, word := range wordList {
		for i, letter := range word {
			l := string(letter)
			if _, ok := cardMap[l]; !ok {
				cardMap[l] = make(map[int]int)
			}
			cardMap[l][i] = cardMap[l][i] + 1
		}
	}

	return cardMap
}

// determineWordScore determines the score by adding the cardinality of each letter in each position
func determineWordScore(word string, cardMap map[string]map[int]int) int {
	var score int
	for j, letter := range word {
		l := string(letter)
		score += cardMap[l][j]
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
