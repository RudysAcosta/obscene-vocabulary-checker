package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var fileName, word string
	words := make(map[string]struct{})

	fmt.Scan(&fileName)

	makeSetWord(&words, fileName)

	for {
		fmt.Scan(&word)

		if word == "" {
			continue
		}

		if word == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		if _, ok := words[normalization(word)]; ok {
			fmt.Println(censor(len(word)))
		} else {
			fmt.Println(word)
		}
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
