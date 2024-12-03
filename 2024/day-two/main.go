package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/bkohler93/aoc-helper/loader"
)

const (
	minDiff = 1
	maxDiff = 3
)

func main() {
	input := loader.GetInput("input.txt", "2024", "2", "53616c7465645f5fb169fa6c83641e8c3c98857c5d60a9e9b0ae64ebeeb1ba35d5af9d61d1ad955673a0d1d4291bdcd4a216a15bcb464fbbb7cc41b8c050ea3c")

	reports := parseReports(input)

	// reports = analyzeReports(reports)
	reports = analyzeReportsTolerate(reports)

	sum := sumSafeReports(reports)

	fmt.Printf("There are %d safe reports\n", sum)
}

func analyzeReportsTolerate(rs []report) []report {
	var wg sync.WaitGroup
	m := sync.Mutex{}

	for i, r := range rs {
		wg.Add(1)
		go func(r *report) {
			r.determineSafetyTolerate()
			m.Lock()
			rs[i] = *r
			m.Unlock()
			wg.Done()
		}(&r)
	}

	wg.Wait()
	return rs
}

func analyzeReports(rs []report) []report {
	var wg sync.WaitGroup
	m := sync.Mutex{}

	for i, r := range rs {
		wg.Add(1)
		go func(r *report) {
			r.determineSafety()
			m.Lock()
			rs[i] = *r
			m.Unlock()
			wg.Done()
		}(&r)
	}

	wg.Wait()
	return rs
}

func sumSafeReports(rs []report) int {
	count := 0

	for _, r := range rs {
		if r.isSafe {
			count++
		}
	}

	return count
}

func removeIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}

func (r *report) determineSafetyTolerate() {
	if determineSafety(r.levels) {
		r.isSafe = true
		return
	}

	for i := 0; i < len(r.levels); i++ {
		ls := removeIndex(r.levels, i)
		if determineSafety(ls) {
			r.isSafe = true
			return
		}
	}
	r.isSafe = false
}

func determineSafety(ls []int) bool {
	if len(ls) < 2 {
		return false
	}

	for i := 0; i < len(ls)-1; i++ {
		diffF := math.Abs(float64(ls[i] - ls[i+1]))
		d := int(diffF)

		if d < minDiff || d > maxDiff {
			return false
		}
	}

	//check for always increasing
	isInc := true
	for i := 0; i < len(ls)-1; i++ {
		if ls[i] > ls[i+1] {
			isInc = false
			break
		}
	}
	if isInc {
		return true
	}

	//check for always decreasing
	for i := 0; i < len(ls)-1; i++ {
		if ls[i] < ls[i+1] {
			return false
		}
	}
	return true
}

func (r *report) determineSafety() {
	//check for safe differences
	ls := r.levels

	if len(ls) < 2 {
		return
	}

	for i := 0; i < len(ls)-1; i++ {
		diffF := math.Abs(float64(ls[i] - ls[i+1]))
		d := int(diffF)

		if d < minDiff || d > maxDiff {
			r.isSafe = false
			return
		}
	}

	//check for always increasing
	isInc := true
	for i := 0; i < len(ls)-1; i++ {
		if ls[i] > ls[i+1] {
			isInc = false
			break
		}
	}
	if isInc {
		r.isSafe = true
		return
	}

	//check for always decreasing
	for i := 0; i < len(ls)-1; i++ {
		if ls[i] < ls[i+1] {
			r.isSafe = false
			return
		}
	}
	r.isSafe = true
}

func parseReports(input string) []report {
	reports := []report{}

	lines := strings.Split(input, "\n")

	for _, l := range lines {
		report := newReport(l)

		reports = append(reports, report)
	}

	return reports
}

type report struct {
	levels []int
	isSafe bool
}

func newReport(lStr string) report {
	ls := strings.Split(lStr, " ")

	levels := []int{}

	for _, l := range ls {
		level, _ := strconv.Atoi(l)
		levels = append(levels, level)
	}

	return report{
		levels: levels,
		isSafe: false,
	}
}
