package main

import (
	"testing"
)

func TestFilterCharsAndNormalize(t *testing.T) {
	sentence := "This is a test. This test is only a test."
	expected := "this is a test  this test is only a test "

	// give me a slice that convert the string in golang to a slice of runes
	data = []rune(sentence)
	FilterCharsAndNormalize()

	result := string(data)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

func TestFilterCharsAndNormalize2(t *testing.T) {
	sentence := "This is a test. This test is only a test!@#&"
	expected := "this is a test  this test is only a test    "

	data = []rune(sentence)
	FilterCharsAndNormalize()

	result := string(data)
	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
