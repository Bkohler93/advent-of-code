package processor

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var types = []string{
	"soil",
	"fertilizer",
	"water",
	"light",
	"temperature",
	"humidity",
	"location",
}

type Processor struct {
	CurrentVal int //current val being processed, starts as a seed
	seeds      []int
	processes  []process
}

func New(input []string) (p *Processor, e error) {
	p = &Processor{
		seeds:     []int{},
		processes: []process{},
	}

	seedString, _ := strings.CutPrefix(input[0], "seeds: ")
	seedStrs := strings.Split(seedString, " ")

	for _, seedS := range seedStrs {
		newSeed, err := strconv.Atoi(seedS)
		if err != nil {
			e = err
			return
		}
		p.seeds = append(p.seeds, newSeed)
	}

	var newProcess process
	for i := 2; i < len(input); i++ {
		if strings.Contains(input[i], "map") {
			newProcess = process{}
			continue
		}
		if input[i] == "\n" || input[i] == "" {
			p.processes = append(p.processes, newProcess)
		} else {
			instructionVals := strings.Split(input[i], " ")
			destinationVal, _ := strconv.Atoi(instructionVals[0])
			sourceVal, _ := strconv.Atoi(instructionVals[1])
			rangeVal, _ := strconv.Atoi(instructionVals[2])
			ins := instruction{
				minIncVal:  sourceVal,
				maxIncVal:  sourceVal + rangeVal - 1,
				outputDiff: destinationVal - sourceVal,
			}
			newProcess.instructions = append(newProcess.instructions, ins)
		}
	}
	return
}

func (p *Processor) Run(w http.ResponseWriter) {
	minLocation := math.MaxInt
	for _, seed := range p.seeds {
		p.CurrentVal = seed
		for i, pcess := range p.processes {
			newVal := pcess.Run(p.CurrentVal)
			w.Write([]byte(fmt.Sprintf("%s = %d\n", types[i], newVal)))
			p.CurrentVal = newVal
		}
		w.Write([]byte(fmt.Sprintf("Location = %d\n\n", p.CurrentVal)))
		minLocation = min(minLocation, p.CurrentVal)
	}

	w.Write([]byte(fmt.Sprintf("min location = %d\n", minLocation)))

}

type process struct {
	instructions []instruction
}

type instruction struct {
	minIncVal  int
	maxIncVal  int
	outputDiff int //from input this = destination - source
}

func (p process) Run(currVal int) int {
	out := currVal

	for _, i := range p.instructions {
		if out >= i.minIncVal && out <= i.maxIncVal {
			out = out + i.outputDiff
			return out
		}
	}

	return out
}
