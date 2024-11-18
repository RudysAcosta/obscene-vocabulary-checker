package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var fileName string

	fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
