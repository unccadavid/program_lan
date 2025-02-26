package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Could not open file: %s", path)
	}
	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content
}

func FilterCharsAndNormalize(text string) string {
	re := regexp.MustCompile(`[\W_]+`)
	return strings.ToLower(re.ReplaceAllString(text, " "))
}

func Scan(text string) []string {
	return strings.Fields(text)
}

func RemoveStopWords(wordList []string) func(string) []string {
	stopwordfilter := func(path string) []string {
		stopWordsContent := ReadFile(path)
		var filteredWords []string
		stopWords := strings.Split(stopWordsContent, ",")
		for _, word := range wordList {
			var temp bool = true
			for _, stopWord := range stopWords {
				if word == stopWord {
					temp = false
					break
				}
			}
			if temp == true {
				filteredWords = append(filteredWords, word)
			}
		}
		return filteredWords
	}
	return stopwordfilter
}

func Frequencies(wordList []string) map[string]int {
	wordFreqs := make(map[string]int)
	for _, word := range wordList {
		wordFreqs[word]++
	}
	return wordFreqs
}

func Sort(wordFreqs map[string]int) []struct {
	Word  string
	Count int
} {
	var sortedWords []struct {
		Word  string
		Count int
	}

	for word, count := range wordFreqs {
		sortedWords = append(sortedWords, struct {
			Word  string
			Count int
		}{word, count})
	}

	sort.Slice(sortedWords, func(i, j int) bool {
		return sortedWords[i].Count > sortedWords[j].Count
	})

	return sortedWords
}

func main() {
	result := Sort(Frequencies(RemoveStopWords(Scan(FilterCharsAndNormalize(ReadFile(os.Args[1]))))(os.Args[2])))
	// //fmt.Printf(os.Args[1])
	// //fmt.Printf(ReadFile(os.Args[1]))
	// //fmt.Printf(ReadFile(os.Args[2]))
	// //fmt.Printf("%s", FilterCharsAndNormalize(ReadFile(os.Args[1])))
	// //fmt.Printf("%s", []string{"st", "st", "st", "prgrm", "prgrm", "prgrm", "prgrm"})
	// //fmt.Printf("%s", Scan(FilterCharsAndNormalize(ReadFile(os.Args[1]))))
	// //fmt.Printf("%s", RemoveStopWords(os.Args[2])(Scan(FilterCharsAndNormalize(ReadFile(os.Args[1])))))
	for _, entry := range result {
		fmt.Printf("%s - %d\n", entry.Word, entry.Count)
	}
}
