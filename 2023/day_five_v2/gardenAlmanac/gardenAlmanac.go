package gardenalmanac

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type GardenAlmanac struct {
	seeds                 []int
	seedToSoilMap         []GardenMap
	soilToFertilizerMap   []GardenMap
	fertilizerToWaterMap  []GardenMap
	waterToLightMap       []GardenMap
	lightToTempMap        []GardenMap
	tempToHumidityMap     []GardenMap
	humidityToLocationMap []GardenMap
}

type GardenMap struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func NewGardenAlmanacPartTwo(input string) *GardenAlmanac {
	lines := strings.Split(input, "\n")
	g := GardenAlmanac{}

	for i, l := range lines {
		if strings.HasPrefix(l, "seeds: ") {
			s, _ := strings.CutPrefix(l, "seeds: ")
			seedStrs := strings.Split(s, " ")
			seeds := []int{}
			seedCh := make(chan int)

			var wg sync.WaitGroup

			for i := 0; i < len(seedStrs); i += 2 {
				in, _ := strconv.Atoi(seedStrs[i])
				r, _ := strconv.Atoi(seedStrs[i+1])

				wg.Add(1)

				go func(in, r int) {
					defer wg.Done()
					fmt.Println("adding seeds")
					for j := in; j < in+r; j++ {
						seed := j
						seedCh <- seed
					}
					fmt.Println("go func added seeds")
				}(in, r)
			}

			go func() {
				wg.Wait()
				close(seedCh)
			}()

			fmt.Println("listening for seeds")
			for s := range seedCh {
				seeds = append(seeds, s)
			}
			fmt.Println("finished adding seeds")

			// for i, ss := range seedStrs {
			// 	var in int
			// 	var r int

			// 	in, _ = strconv.Atoi(ss)
			// 	i++
			// 	r, _ = strconv.Atoi(seedStrs[i])

			// 	for j := in; j < in+r; j++ {
			// 		seed := in
			// 		seeds = append(seeds, seed)
			// 	}
			// 	fmt.Println("added seeds")
			// }
			g.seeds = seeds
		}

		if strings.HasPrefix(l, "seed-to-soil") {
			seedToSoilMap := extractGardenMap(&i, lines)
			g.seedToSoilMap = seedToSoilMap
		}

		if strings.HasPrefix(l, "soil-to-fertilizer") {
			soilToFertilizerMap := extractGardenMap(&i, lines)
			g.soilToFertilizerMap = soilToFertilizerMap
		}

		if strings.HasPrefix(l, "fertilizer-to-water") {
			fertilizerToWaterMap := extractGardenMap(&i, lines)
			g.fertilizerToWaterMap = fertilizerToWaterMap
		}

		if strings.HasPrefix(l, "water-to-light") {
			waterToLightMap := extractGardenMap(&i, lines)
			g.waterToLightMap = waterToLightMap
		}

		if strings.HasPrefix(l, "light-to-temperature") {
			lightToTempMap := extractGardenMap(&i, lines)
			g.lightToTempMap = lightToTempMap
		}

		if strings.HasPrefix(l, "temperature-to-humidity") {
			tempToHumidityMap := extractGardenMap(&i, lines)
			g.tempToHumidityMap = tempToHumidityMap
		}

		if strings.HasPrefix(l, "humidity-to-location") {
			humidityToLocationMap := extractGardenMap(&i, lines)
			g.humidityToLocationMap = humidityToLocationMap
		}
	}

	return &g
}

func NewGardenAlmanac(input string) *GardenAlmanac {
	lines := strings.Split(input, "\n")
	g := GardenAlmanac{}

	for i, l := range lines {
		if strings.HasPrefix(l, "seeds: ") {
			s, _ := strings.CutPrefix(l, "seeds: ")
			seedStrs := strings.Split(s, " ")
			seeds := []int{}

			for _, ss := range seedStrs {
				seed, _ := strconv.Atoi(ss)
				seeds = append(seeds, seed)
				g.seeds = seeds
			}
		}

		if strings.HasPrefix(l, "seed-to-soil") {
			seedToSoilMap := extractGardenMap(&i, lines)
			g.seedToSoilMap = seedToSoilMap
		}

		if strings.HasPrefix(l, "soil-to-fertilizer") {
			soilToFertilizerMap := extractGardenMap(&i, lines)
			g.soilToFertilizerMap = soilToFertilizerMap
		}

		if strings.HasPrefix(l, "fertilizer-to-water") {
			fertilizerToWaterMap := extractGardenMap(&i, lines)
			g.fertilizerToWaterMap = fertilizerToWaterMap
		}

		if strings.HasPrefix(l, "water-to-light") {
			waterToLightMap := extractGardenMap(&i, lines)
			g.waterToLightMap = waterToLightMap
		}

		if strings.HasPrefix(l, "light-to-temperature") {
			lightToTempMap := extractGardenMap(&i, lines)
			g.lightToTempMap = lightToTempMap
		}

		if strings.HasPrefix(l, "temperature-to-humidity") {
			tempToHumidityMap := extractGardenMap(&i, lines)
			g.tempToHumidityMap = tempToHumidityMap
		}

		if strings.HasPrefix(l, "humidity-to-location") {
			humidityToLocationMap := extractGardenMap(&i, lines)
			g.humidityToLocationMap = humidityToLocationMap
		}
	}

	return &g
}

func extractGardenMap(index *int, lines []string) []GardenMap {
	*index++
	gMap := []GardenMap{}

	for lines[*index] != "" {
		valStrs := strings.Split(lines[*index], " ")

		destinationStart, _ := strconv.Atoi(valStrs[0])
		sourceStart, _ := strconv.Atoi(valStrs[1])
		rangeLength, _ := strconv.Atoi(valStrs[2])

		gMap = append(gMap, GardenMap{
			destinationStart: destinationStart,
			sourceStart:      sourceStart,
			rangeLength:      rangeLength,
		})
		*index++
	}
	return gMap
}

func (g *GardenAlmanac) FindLowestLocationNumber() int {
	minLocation := math.MaxInt
	for _, s := range g.seeds {
		soilNumber := mapToNumber(s, g.seedToSoilMap)
		fertilizerNumber := mapToNumber(soilNumber, g.soilToFertilizerMap)
		waterNumber := mapToNumber(fertilizerNumber, g.fertilizerToWaterMap)
		lightNumber := mapToNumber(waterNumber, g.waterToLightMap)
		tempNumber := mapToNumber(lightNumber, g.lightToTempMap)
		humidityNumber := mapToNumber(tempNumber, g.tempToHumidityMap)
		locationNumber := mapToNumber(humidityNumber, g.humidityToLocationMap)
		minLocation = min(minLocation, locationNumber)
	}
	return minLocation
}

func mapToNumber(source int, m []GardenMap) int {
	for _, mm := range m {
		if source >= mm.sourceStart && source < mm.sourceStart+mm.rangeLength {
			diff := source - mm.sourceStart

			destination := mm.destinationStart + diff
			return destination
		}
	}
	return source
}
