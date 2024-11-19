package main

import (
	"os"
	"strings"
	"testing"
)

func TestNormalization(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{" Hello ", "hello"},
		{"WORLD", "world"},
		{"  GoLang  ", "golang"},
		{"", ""},
	}

	for _, c := range cases {
		result := normalization(c.input)
		if result != c.expected {
			t.Errorf("normalization(%q) = %q; want %q", c.input, result, c.expected)
		}
	}
}

func TestCensor(t *testing.T) {
	cases := []struct {
		length   int
		expected string
	}{
		{5, "*****"},
		{3, "***"},
		{0, ""},
		{1, "*"},
	}

	for _, c := range cases {
		result := censor(c.length)
		if result != c.expected {
			t.Errorf("censor(%d) = %q; want %q", c.length, result, c.expected)
		}
	}
}

func TestMakeSetWord(t *testing.T) {
	fileContent := "hello\nworld\nGolang\n"
	fileName := "test_words.txt"
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(fileName)

	words := make(map[string]struct{})
	makeSetWord(&words, fileName)

	expectedWords := map[string]struct{}{
		"hello":  {},
		"world":  {},
		"golang": {},
	}

	for word := range expectedWords {
		if _, exists := words[word]; !exists {
			t.Errorf("Expected word %q to be in the set", word)
		}
	}
}

func TestIntegration(t *testing.T) {
	fileContent := "hello\nworld\n"
	fileName := "test_words_integration.txt"
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(fileName)

	words := make(map[string]struct{})
	makeSetWord(&words, fileName)

	cases := []struct {
		input    string
		expected string
	}{
		{"hello world", "***** *****"},
		{"golang", "golang"},
		{"hello golang", "***** golang"},
		{"exit", "Bye!"},
	}

	for _, c := range cases {
		input := c.input
		if input == "exit" {
			if c.expected != "Bye!" {
				t.Errorf("Expected %q but got %q", c.expected, "Bye!")
			}
			break
		}

		wordsInput := strings.Split(input, " ")
		for index, value := range wordsInput {
			if _, ok := words[normalization(value)]; ok {
				wordsInput[index] = censor(len(value))
			}
		}

		sentence := strings.Join(wordsInput, " ")
		if sentence != c.expected {
			t.Errorf("Expected %q but got %q", c.expected, sentence)
		}
	}
}
