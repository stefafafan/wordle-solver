package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

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

func getRuneScoreMap(wordList []string) map[rune]uint16 {
	runeScoreMap := map[rune]uint16{}
	for _, w := range wordList {
		for _, r := range w {
			runeScoreMap[r]++
		}
	}
	return runeScoreMap
}

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

func narrowDownWordList(wordList []string, guess string, result string) []string {
	var blankRunes []string
	yellowRunes := map[string]int{}
	greenRunes := map[string]int{}
	greenRunesCount := 0
	for index, r := range result {
		if string(r) == "b" {
			blankRunes = append(blankRunes, string(guess[index]))
		}
		if string(r) == "y" {
			yellowRunes[string(guess[index])] = index
		}
		if string(r) == "g" {
			greenRunes[string(guess[index])] = index
			greenRunesCount++
		}
	}

	var newWordList []string
	for _, w := range wordList {
		wordHasBlank := false
		wordHasWrongY := false
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
				if i == yindex && string(r) == y {
					wordHasWrongY = true
					break
				}
			}
			if wordHasWrongY {
				break
			}
			// choose words with green runes
			for g, gindex := range greenRunes {
				if i == gindex && string(r) == g {
					currentGCount++
				}
			}
		}
		if !wordHasBlank && !wordHasWrongY && greenRunesCount == currentGCount {
			newWordList = append(newWordList, w)
		}
	}
	return newWordList
}

func main() {
	path := flag.String("dictionary", "", "Specify dictionary file path")
	flag.Parse()

	var wordList = getInitialWordList(*path)
	try := 1

	for try < 7 {
		fmt.Printf("Try number %d\n", try)
		var runeScoreMap = getRuneScoreMap(wordList)
		var wordScoreMap = getWordScoreMap(wordList, runeScoreMap)
		var sortedwordScoreMap = rankByWordCount(wordScoreMap)
		if sortedwordScoreMap.Len() > 10 {
			fmt.Printf("There are %d candidates. Here are the first 10 candidates:\n%v\n", sortedwordScoreMap.Len(), sortedwordScoreMap[0:10])
		} else {
			fmt.Printf("There are %d candidates. Here are the candidates:\n%v\n", sortedwordScoreMap.Len(), sortedwordScoreMap)
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

		wordList = narrowDownWordList(wordList, guess, result)
		try++
	}
	fmt.Println("Game over.")
}
