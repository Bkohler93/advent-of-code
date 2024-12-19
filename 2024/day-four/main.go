package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/bkohler93/aoc-helper/loader"
)

type wordSearch [][]rune

func main() {
	input, err := loader.GetInputFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ws := parseInput(input)

	// c := countXmas(ws)
	c := countXmasPartTwo(ws)

	fmt.Println("xmas count is ", c)
}

func parseInput(input string) wordSearch {
	ws := make(wordSearch, 0, len(input))

	s := bufio.NewScanner(strings.NewReader(input))

	for s.Scan() {
		ws = append(ws, []rune(s.Text()))
	}

	return ws
}

func countXmasPartTwo(ws wordSearch) int {
	c := 0
	s := len(ws)

	for col := 0; col < s-2; col++ {
		for row := 0; row < s-2; row++ {
			//select downslope diagonal
			dd := make([]rune, 0, 3)
			for i := 0; i < 3; i++ {
				dd = append(dd, ws[row+i][col+i])
			}

			//select upslope diagonal
			ud := make([]rune, 0, 3)
			for i := 0; i < 3; i++ {
				ud = append(ud, ws[row+2-i][col+i])
			}

			//if both both contains "MAS" or "SAM" increment c
			if hasMas(dd, ud) {
				c++
			}
		}
	}

	return c
}

func hasMas(l1, l2 []rune) bool {
	return (string(l1) == "SAM" || string(l1) == "MAS") && (string(l2) == "SAM" || string(l2) == "MAS")
}

func countXmas(ws wordSearch) int {
	c := 0

	c += countHorizontal(ws)

	c += countVertical(ws)

	c += countDownSlopeDiagonal(ws)

	c += countUpSlopeDiagonal(ws)
	return c
}

func countUpSlopeDiagonal(ws wordSearch) int {
	c := 0
	s := len(ws)

	// check diagonals starting at top left of wordsearch, moving down to bottom
	for row := 0; row < s; row++ {
		l := make([]rune, 0, row+1)
		cp := cap(l)
		for col := 0; col < cp; col++ {
			l = append(l, ws[row-col][col])
		}
		c += countLine(l)
	}

	// check diagonals starting at top right (row+1) of wordsearch, moving down to bottom
	for row := 1; row < s; row++ {
		l := make([]rune, 0, s-row)
		i := 0
		for col := s - 1; col >= 0+row; col-- {
			l = append(l, ws[row+i][col])
			i++
		}
		c += countLine(l)
	}

	return c
}

func countDownSlopeDiagonal(ws wordSearch) int {
	c := 0
	s := len(ws)

	// check diagonals starting at top left of wordsearch, moving down to bottom
	for row := 0; row < s; row++ {
		l := make([]rune, 0, s-row)
		cp := cap(l)
		for col := 0; col < cp; col++ {
			l = append(l, ws[row+col][col])
		}
		c += countLine(l)
	}

	// check diagonals starting at top left (+1 col) of wordsearch, moving across to end of row
	for col := 1; col < s; col++ {
		l := make([]rune, 0, s-col)
		cp := cap(l)
		for row := 0; row < cp; row++ {
			l = append(l, ws[row][col+row])
		}
		c += countLine(l)
	}

	return c
}

func countVertical(ws wordSearch) int {
	c := 0
	s := len(ws[0])

	for col := 0; col < s; col++ {
		l := make([]rune, 0, s)
		for row := 0; row < s; row++ {
			l = append(l, ws[row][col])
		}
		c += countLine(l)
	}

	return c
}

func countHorizontal(ws wordSearch) int {
	c := 0

	for _, l := range ws {
		c += countLine(l)
	}

	return c
}

func countLine(l []rune) int {
	line := string(l)

	return strings.Count(line, "XMAS") + strings.Count(line, "SAMX")
}
