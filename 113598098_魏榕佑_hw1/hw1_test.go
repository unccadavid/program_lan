package main

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	path := "text_file.txt"
	expected := "This is a test. This test is only a test.\n"
	result := ReadFile(path)

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFilterCharsAndNormalize(t *testing.T) {
	input := "This is a test. This test is only a test."
	expected := "this is a test this test is only a test "
	result := FilterCharsAndNormalize(input)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestScan(t *testing.T) {
	input := "this is a test this test is only a test "
	expected := []string{"this", "is", "a", "test", "this", "test", "is", "only", "a", "test"}
	result := Scan(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestRemovestopwords(t *testing.T) {
	input := []string{"this", "is", "a", "test", "this", "test", "is", "only", "a", "test"}
	path := "stop_words.txt"
	expected := []string{"test", "test", "test"}
	stop_word_filter := RemoveStopWords(path)
	result := stop_word_filter(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFrequencies(t *testing.T) {
	input := []string{"test", "test", "test", "program", "program", "program", "program"}
	expected := map[string]int{"test": 3, "program": 4}
	result := Frequencies(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// func TestSort(t *testing.T) {
// 	input := map[string]int{"test": 3, "program": 4}
// 	expected := [("program",4),("test",3)]
// 	result := sort(input)
// 	if result != expected {
// 		t.Errorf("Expected %q but got %q", expected, result)
// 	}
// }
