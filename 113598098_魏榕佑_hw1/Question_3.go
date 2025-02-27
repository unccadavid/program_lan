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
	defer file.Close()
	var builder strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if builder.Len() > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	content := builder.String()
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
		for i, stopword := range stopWords {
			stopWords[i] = strings.TrimSpace(stopword)
		}
		var temp bool
		for _, word := range wordList {
			temp = false
			for _, stopWord := range stopWords {
				if word == stopWord {
					temp = true
					break
				}
			}
			if temp == false {
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

func PrintAll(word_freqs []struct {
	Word  string
	Count int
}) {
	for _, entry := range word_freqs {
		fmt.Printf("%s - %d\n", entry.Word, entry.Count)
	}
}

func main() {
	PrintAll(Sort(Frequencies(RemoveStopWords(Scan(FilterCharsAndNormalize(ReadFile(os.Args[1]))))(os.Args[2])))[0:25])
}
