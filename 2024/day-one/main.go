package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/bkohler93/aoc-helper/loader"
)

func main() {
	input := loader.GetInput("input.txt", "2024", "1", "53616c7465645f5fb169fa6c83641e8c3c98857c5d60a9e9b0ae64ebeeb1ba35d5af9d61d1ad955673a0d1d4291bdcd4a216a15bcb464fbbb7cc41b8c050ea3c")

	leftIds, rightIds := splitLists(input)

	distances, err := calculateDistances(leftIds, rightIds)
	if err != nil {
		log.Fatal(err)
	}

	s := sum(distances)

	fmt.Printf("total distance = %d\n", s)

	fmt.Println("=== PART TWO ===")

	sScores := calculateSimilarityScores(leftIds, rightIds)

	sPartTwo := sum(sScores)

	fmt.Printf("total similarity score = %d\n", sPartTwo)
}

func calculateSimilarityScores(l, r []int) []int {
	s := []int{}

	for _, n := range l {

		count := 0
		for _, rN := range r {
			if n == rN {
				count++
			}
		}
		score := count * n
		s = append(s, score)
	}
	return s
}

func sum(s []int) int {
	sum := 0

	for _, n := range s {
		sum += n
	}

	return sum
}

func calculateDistances(l, r []int) ([]int, error) {
	ds := []int{}

	if len(l) != len(r) {
		return ds, fmt.Errorf("two lists are not same length: %d and %d", len(l), len(r))
	}

	for i := range l {
		d := math.Abs(float64(l[i] - r[i]))
		di := int(d)
		ds = append(ds, di)
	}

	return ds, nil
}

func splitLists(s string) ([]int, []int) {
	var l, r []int
	lines := strings.Split(s, "\n")

	lCh := make(chan int)
	rCh := make(chan int)
	var wg sync.WaitGroup

	for _, line := range lines {
		wg.Add(1)
		go func(line string) {
			parts := strings.Split(line, "   ")
			if len(parts) < 2 {
				wg.Done()
				return
			}

			left := parts[0]
			right := parts[1]
			lN, _ := strconv.Atoi(left)
			rN, _ := strconv.Atoi(right)

			lCh <- lN
			rCh <- rN
			wg.Done()
		}(line)
	}
	go func() {
		wg.Wait()
		close(lCh)
		close(rCh)
	}()

	for lCh != nil || rCh != nil {
		select {
		case n, ok := <-lCh:
			if !ok {
				lCh = nil
			} else {
				l = append(l, n)
			}
		case n, ok := <-rCh:
			if !ok {
				rCh = nil
			} else {
				r = append(r, n)
			}
		}
	}
	slices.Sort(l)
	slices.Sort(r)
	return l, r
}
