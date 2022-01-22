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

func main() {
	path := flag.String("dictionary", "", "Specify dictionary file path")
	flag.Parse()

	var initialWordList = getInitialWordList(*path)
	fmt.Printf("%v\n", initialWordList)

	var runeScoreMap = getRuneScoreMap(initialWordList)
	for k, v := range runeScoreMap {
		fmt.Printf("%c:%d\n", k, v)
	}

	var wordScoreMap = getWordScoreMap(initialWordList, runeScoreMap)
	var sortedwordScoreMap = rankByWordCount(wordScoreMap)
	fmt.Printf("There are %d words. Here are some of them: %v", sortedwordScoreMap.Len(), sortedwordScoreMap[0:10])

	// TODO: prompt for guess and result
	// TODO: narrow down word list from guess and result
	// TODO: recalculate runeScoreMap && wordScoreMap (e.g recur)
}
