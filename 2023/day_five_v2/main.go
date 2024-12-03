package main

import (
	"fmt"

	gardenalmanac "github.com/bkohler93/aoc/dayfive/gardenAlmanac"
	"github.com/bkohler93/aoc/dayfive/loader"
)

func main() {
	input := loader.GetInput()

	gardenAlmanac := gardenalmanac.NewGardenAlmanacPartTwo(input)

	lowestLocationNumber := gardenAlmanac.FindLowestLocationNumber()

	fmt.Println("Lowest location number is", lowestLocationNumber)
}
