package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bkohler93/advent-of-code/2023/day-six/processor"
)

const (
	port = ":8080"
)

func main() {

	input, err := readInput("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	processor, err := processor.New(input)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		processor.Run(w)
	})

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func readInput(filePath string) ([]string, error) {
	input := []string{}

	f, err := os.Open(filePath)
	if err != nil {
		return input, fmt.Errorf("couldn't open file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}
