package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var fileName string
	words := make(map[string]struct{})

	fmt.Scan(&fileName)

	makeSetWord(&words, fileName)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		if input == " " {
			continue
		}

		if input == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		wordsInput := strings.Split(input, " ")

		for index, value := range wordsInput {
			if _, ok := words[normalization(value)]; ok {
				wordsInput[index] = censor(len(value))
			}
		}

		sentence := strings.Join(wordsInput, " ")
		fmt.Println(sentence)
	}

}

func makeSetWord(words *map[string]struct{}, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		key := normalization(scanner.Text())
		(*words)[key] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func normalization(word string) string {
	if word == "" {
		return word
	}

	return strings.ToLower(strings.TrimSpace(word))
}

func censor(lenghtWord int) string {
	return strings.Repeat("*", lenghtWord)
}
