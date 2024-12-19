package main

import (
	"testing"
)

func TestXmasCountPartTwo(t *testing.T) {
	ws := parseInput(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 9
	actual := countXmasPartTwo(ws)

	if expected != actual {
		t.Errorf("expected %d got %d", expected, actual)
	}
}

func TestXmasCount(t *testing.T) {
	ws := parseInput(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`)

	expected := 18

	actual := countXmas(ws)

	if expected != actual {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
