package loader

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetInput() string {
	fileName := "input.txt"
	url := "https://adventofcode.com/2023/day/5/input"
	sessionID := "53616c7465645f5fb169fa6c83641e8c3c98857c5d60a9e9b0ae64ebeeb1ba35d5af9d61d1ad955673a0d1d4291bdcd4a216a15bcb464fbbb7cc41b8c050ea3c"

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}
	info, err := f.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		f.Seek(0, 0)
		r, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}

		c := &http.Cookie{
			Name:  "session",
			Value: sessionID,
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
	buffer := bytes.Buffer{}
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		buffer.Write([]byte("\n"))
	}

	return buffer.String()
}
