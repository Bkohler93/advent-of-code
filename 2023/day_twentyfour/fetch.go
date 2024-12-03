package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func loadData(inputCh chan string) {
	defer close(inputCh)

	f, err := os.OpenFile("input.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		f.Seek(0, 0)
		r, err := http.NewRequest("GET", "https://adventofcode.com/2023/day/24/input", nil)
		if err != nil {
			panic(err)
		}

		c := &http.Cookie{
			Name:  "session",
			Value: "53616c7465645f5fb169fa6c83641e8c3c98857c5d60a9e9b0ae64ebeeb1ba35d5af9d61d1ad955673a0d1d4291bdcd4a216a15bcb464fbbb7cc41b8c050ea3c",
		}
		r.AddCookie(c)

		res, err := http.DefaultClient.Do(r)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		scanner := bufio.NewScanner(res.Body)

		for scanner.Scan() {
			input := scanner.Text()
			fmt.Println("writing to file")
			f.WriteString(input + "\n")
		}
	}

	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputCh <- scanner.Text()
	}
}
