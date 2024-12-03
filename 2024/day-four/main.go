package main

import (
	"fmt"
	"log"

	"github.com/bkohler93/aoc-helper/loader"
)

func main() {
	input, err := loader.GetInputFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)
}
