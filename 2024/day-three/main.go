package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/bkohler93/aoc-helper/loader"
)

func main() {
	input, err := loader.GetInputFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// matches := findAllMultOps(input) //part one
	matches := findEnabledMultOps(input) //part two

	ops := parseOps(matches)

	sum := processOps(ops)

	fmt.Printf("result is %d\n", sum)
}

func findEnabledMultOps(s string) []string {
	ops := []string{}

	rxs := map[string]*regexp.Regexp{
		"do":   regexp.MustCompile(`do\(\)`),
		"dont": regexp.MustCompile(`don't\(\)`),
		"mul":  regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`),
	}

	mulOn := true
	for {
		var firstRxName, firstMatch string
		minIndex := math.MaxInt

		for name, rx := range rxs {
			loc := rx.FindStringIndex(s)
			if loc != nil && loc[0] < minIndex {
				minIndex = loc[0]
				firstRxName = name
				firstMatch = s[loc[0]:loc[1]]
			}
		}

		if firstRxName == "" {
			break
		}

		switch firstRxName {
		case "do":
			mulOn = true
		case "dont":
			mulOn = false
		case "mul":
			if mulOn {
				ops = append(ops, firstMatch)
			}
		}

		s = s[minIndex+len(firstMatch):]
	}

	return ops
}

func findAllMultOps(s string) []string {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(s, -1)
	return matches
}

func parseOps(ss []string) []operation {
	ops := make([]operation, len(ss))

	for _, s := range ss {
		parts := strings.Split(s, ",")
		l, r := parts[0], parts[1]

		xs := strings.TrimPrefix(l, "mul(")
		ys := strings.TrimSuffix(r, ")")

		x, _ := strconv.Atoi(xs)
		y, _ := strconv.Atoi(ys)

		op := operation{x, y}
		ops = append(ops, op)
	}

	return ops
}

func processOps(ops []operation) int {
	sum := 0

	for _, op := range ops {
		r := op.mult()
		sum += r
	}

	return sum
}

type operation struct {
	x int
	y int
}

func (o operation) mult() int {
	return o.x * o.y
}
