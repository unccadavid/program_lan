package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	path := "text_file.txt"
	expected := "This is a test. This test is only a test."
	result := ReadFile(path)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFilterCharsAndNormalize(t *testing.T) {
	input := "This is a test. This test is only a test."
    expected := "this is a test this test is only a test "
	result := Filter_chars_and_normalize(input)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func test_Scan(t *testing.T) {
	input := "this is a test this test is only a test "
	expected := ["this","is","a","test","this","test","is","only","a","test"]
	result := scan(input)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func test_remove_stop_words(t *testing.T) {
	input := ["this","is","a","test","this","test","is","only","a","test"]
	path := "stop_words.txt"
	expected := ["test","test","test"]
	stop_word_filter := remove_stop_words(input)
	result := stop_word_filter(path)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func test_frequencies(t *testing.T) {
	input := ["test","test","test","program","program","program","program"]
	expected := {"test":3,"program":4}
	result := frequencies(input)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func test_sort(t *testing.T) {
	input := {"test":3,"program":4}
	expected := [("program",4),("test",3)]
	result := sort(input)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}