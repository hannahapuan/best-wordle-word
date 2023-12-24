package main

import (
	"container/heap"

	"github.com/hannahapuan/best-wordle-word/ref"
)

// best wordle word finds the best wordle word by comparing
// the cardinality of each letter in each position in the wordle dictionary list

func main() {
	// create cardinality map of each occurrence of a letter in each position
	cardMap := make(map[string]map[int]int)

	// initialize map of each letter's cardinality in each position
	for _, word := range ref.WordleList {
		for i, letter := range word {
			l := string(letter)
			if _, ok := cardMap[l]; !ok {
				cardMap[l] = make(map[int]int)
			}
			cardMap[l][i] = cardMap[l][i] + 1
		}
	}

	// create map of each letter's total cardinality
	wordScores := make(map[string]int)
	for _, word := range ref.WordleList {
		var score int
		for j, letter := range word {
			l := string(letter)
			score += cardMap[l][j]
		}
		wordScores[word] = score
	}

	ws := sortWordScores(wordScores)
	ws.Print()
}

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
// 	fmt.Println(maxWord, max)
// 	return maxWord
// }

// func printPretty(cardMap map[string]map[int]int) {
// 	for letter, pos := range cardMap {
// 		fmt.Println(letter, printPrettyInts(pos))
// 	}
// }

// func printPrettyInts(intsMap map[int]int) {
// 	for letter, pos := range cardMap {
// 		fmt.Println(letter, pos)
// 	}
// }
