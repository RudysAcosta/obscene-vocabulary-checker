package main

import (
	"os"
	"testing"
)

// TestNormalization verifica la función normalization
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

// TestCensor verifica la función censor
func TestCensor(t *testing.T) {
	cases := []struct {
		length   int
		expected string
	}{
		{5, "*****"},
		{3, "***"},
		{0, ""},
	}

	for _, c := range cases {
		result := censor(c.length)
		if result != c.expected {
			t.Errorf("censor(%d) = %q; want %q", c.length, result, c.expected)
		}
	}
}

// TestMakeSetWord verifica la función makeSetWord
func TestMakeSetWord(t *testing.T) {
	// Crea un archivo temporal
	fileContent := "hello\nworld\ngolang\n"
	fileName := "test_words.txt"
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(fileName) // Elimina el archivo después de la prueba

	words := make(map[string]struct{})
	makeSetWord(&words, fileName)

	// Verifica si las palabras se agregaron correctamente al conjunto
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

// TestIntegration prueba la funcionalidad completa del programa
func TestIntegration(t *testing.T) {
	// Crea un archivo temporal
	fileContent := "hello\nworld\n"
	fileName := "test_words_integration.txt"
	err := os.WriteFile(fileName, []byte(fileContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(fileName) // Elimina el archivo después de la prueba

	words := make(map[string]struct{})
	makeSetWord(&words, fileName)

	// Simula entradas y salidas
	inputs := []string{"hello", "world", "golang", "exit"}
	expectedOutputs := []string{"*****", "*****", "golang", "Bye!"}

	for i, input := range inputs {
		input = normalization(input)
		if input == "exit" {
			if expectedOutputs[i] != "Bye!" {
				t.Errorf("Expected %q but got %q", expectedOutputs[i], "Bye!")
			}
			break
		}

		if _, exists := words[input]; exists {
			output := censor(len(input))
			if output != expectedOutputs[i] {
				t.Errorf("Expected %q but got %q", expectedOutputs[i], output)
			}
		} else {
			if input != expectedOutputs[i] {
				t.Errorf("Expected %q but got %q", expectedOutputs[i], input)
			}
		}
	}
}
