package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// Returns word list parsed from file.
func getInitialWordList(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var initialWordList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		initialWordList = append(initialWordList, word)
	}
	return initialWordList
}

// Returns the rune score map from the given word list.
func getRuneScoreMap(wordList []string) map[rune]uint16 {
	runeScoreMap := map[rune]uint16{}
	for _, w := range wordList {
		for _, r := range w {
			runeScoreMap[r]++
		}
	}
	return runeScoreMap
}

// Returns the word score map from the given word list and rune score map.
func getWordScoreMap(wordList []string, runeScoreMap map[rune]uint16) map[string]int {
	wordScoreMap := map[string]int{}
	for _, w := range wordList {
		runes := make(map[rune]uint16, len(w))
		for _, r := range w {
			runes[r]++
		}
		score := 0
		for k := range runes {
			score += int(runeScoreMap[k])
		}
		wordScoreMap[w] = score
	}
	return wordScoreMap
}

// Sorts the given word score map by value.
// borrowed from https://stackoverflow.com/a/18695740
func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Narrows down the word list from the guess and result.
// Receives yellowRunes and greenRunes to take into account previous results.
func narrowDownWordList(wordList []string, yellowRunes map[string]int, greenRunes map[string]int, guess string, result string) (map[string]int, map[string]int, []string) {
	// Parse/initialize values from the guess/result strings.
	var blankRunes []string
	for index, r := range result {
		if string(r) == "b" {
			blankRunes = append(blankRunes, string(guess[index]))
		}
		if string(r) == "y" {
			yellowRunes[string(guess[index])] = index
		}
		if string(r) == "g" {
			greenRunes[string(guess[index])] = index
		}
	}
	yellowRunesCount := len(yellowRunes)
	greenRunesCount := len(greenRunes)

	var newWordList []string
	for _, w := range wordList {
		wordHasBlank := false
		wordHasWrongY := false
		currentYCount := 0
		currentGCount := 0
		for i, r := range w {
			// remove words with blank runes
			for _, b := range blankRunes {
				if string(r) == b {
					wordHasBlank = true
					break
				}
			}
			if wordHasBlank {
				break
			}
			// choose words with wrong yellow runes
			for y, yindex := range yellowRunes {
				// if the word contains the yellow rune at the exact same spot, that word is wrong.
				if i == yindex && string(r) == y {
					wordHasWrongY = true
					break
				}
				// count the number of yellow runes.
				if i != yindex && string(r) == y {
					currentYCount++
				}
			}
			if wordHasWrongY {
				break
			}
			// choose words with green runes
			for g, gindex := range greenRunes {
				// count the number of green runes.
				if i == gindex && string(r) == g {
					currentGCount++
				}
			}
		}
		// Append to the new word list only if
		// - the word does not contain blank runes.
		// - the word does not have yellow runes at wrong spots.
		// - the word has a correct number of yellow runes.
		// - the word has a correct number of green runes.
		if !wordHasBlank && !wordHasWrongY && yellowRunesCount == currentYCount && greenRunesCount == currentGCount {
			newWordList = append(newWordList, w)
		}
	}
	// Return the yellow and green runes as well, for future use.
	return yellowRunes, greenRunes, newWordList
}

func main() {
	path := flag.String("dictionary", "", "Specify dictionary file path")
	flag.Parse()

	var wordList = getInitialWordList(*path)
	candidateList := wordList
	try := 1
	yellowRunes := map[string]int{}
	greenRunes := map[string]int{}

	// You only get 6 tries.
	for try < 7 {
		fmt.Printf("-----Try number %d-----\n", try)
		if len(candidateList) == 1 {
			fmt.Printf("There is 1 candidate. This is probably the answer: %s\n\n", candidateList[0])
		} else if len(candidateList) > 10 {
			fmt.Printf("There are %d candidates. Here are the first 10 candidates:\n%v\n\n", len(candidateList), candidateList[0:10])
		} else {
			fmt.Printf("There are %d candidates. Here are the candidates:\n%v\n\n", len(candidateList), candidateList)
		}

		// Note: the rune score map is calculated from the updated candidate list.
		// The word score map uses the initial word list and the updated rune score map for recommendation.
		var runeScoreMap = getRuneScoreMap(candidateList)
		var wordScoreMap = getWordScoreMap(wordList, runeScoreMap)
		var sortedwordScoreMap = rankByWordCount(wordScoreMap)
		if len(candidateList) != 1 {
			if sortedwordScoreMap.Len() > 10 {
				fmt.Printf("Here are the first 10 recommendations to help narrow down the word:\n%v\n\n", sortedwordScoreMap[0:10])
			} else {
				fmt.Printf("Here are the recommendations to help narrow down the word:\n%v\n\n", sortedwordScoreMap)
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		guess := scanner.Text()
		fmt.Print("Enter the result (b: blank, y: yellow, g: green): ")
		scanner.Scan()
		result := scanner.Text()
		fmt.Printf("Guess: %s, Result: %s\n\n", guess, result)
		if result == "ggggg" {
			fmt.Println("GJ!")
			break
		}

		yellowRunes, greenRunes, candidateList = narrowDownWordList(candidateList, yellowRunes, greenRunes, guess, result)
		try++
	}
	fmt.Println("Game over.")
}
