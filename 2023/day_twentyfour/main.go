package main

import "fmt"

const (
	xyMin = 200000000000000
	xyMax = 400000000000000
)

func main() {
	inputCh := make(chan string)

	go loadData(inputCh)

	hails := []*hail{}

	for c := range inputCh {
		h := newHail(c)
		hails = append(hails, h)
	}
	count := 0
	for i := range hails {
		for j := i + 1; j < len(hails); j++ {
			if p, ok := hails[i].intersect(hails[j]); ok {
				fmt.Println("found intersection")
				if p.within(xyMin, xyMax) {
					count++
				}
			}
		}
	}

	fmt.Printf("There are %d intersections within the test area\n", count)
}
