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

	// create map of each letter's total cardinality and fill it with each word's score
	wordScores := make(map[string]int)
	for _, word := range ref.WordleList {
		wordScores[word] = determineWordScore(word, cardMap)
	}

	// sort into a priority queue where score is the priority
	ws := sortWordScores(wordScores)
	ws.Print()
}

// initalizeCardMap creates a map of each letter's cardinality in each position
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
