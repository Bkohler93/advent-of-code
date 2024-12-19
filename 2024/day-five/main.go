package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("error opening file - %v", err)
	}

	u := NewUpdates(f)

	// sum := u.partOne()
	sum := u.partTwo()

	fmt.Printf("sum is %d\n", sum)
}

type updates struct {
	rules    []rule
	pageSets []pages
}

func NewUpdates(r io.Reader) updates {
	s := bufio.NewScanner(r)
	rules := make([]rule, 0)
	pageSets := make([]pages, 0)

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		strs := strings.Split(s.Text(), "|")
		r, err := newRule(strs)
		if err != nil {
			log.Fatalf("error creating rule - %v", err)
		}
		rules = append(rules, r)
	}

	for s.Scan() {
		strs := strings.Split(s.Text(), ",")
		p, err := newPages(strs)
		if err != nil {
			log.Fatalf("error creating pages - %v", err)
		}

		pageSets = append(pageSets, p)
	}

	return updates{
		rules:    rules,
		pageSets: pageSets,
	}
}

func (u updates) partOne() int {
	sum := 0
	for _, ps := range u.pageSets {
		isCorrect := true
		for _, r := range u.rules {
			if ps.includesRule(r) && !ps.followsRule(r) {
				isCorrect = false
			}
		}
		if isCorrect {
			sum += ps.middleNum()
		}
	}
	return sum
}

func (u updates) partTwo() int {
	sum := 0

	for _, ps := range u.pageSets {

		//find rule break
		isCorrect := true
		for _, r := range u.rules {
			if ps.includesRule(r) && !ps.followsRule(r) {
				isCorrect = false
			}
		}

		if !isCorrect {
			//fix page set
			fixedPs := fixPageOrder(ps, u.rules)
			sum += fixedPs.middleNum()
		}
	}

	return sum
}

func fixPageOrder(ps pages, rs []rule) pages {
	fps := make(pages, len(ps))
	copy(fps, ps)
	for {
		isCorrect := true
		for _, r := range rs {
			if fps.includesRule(r) && !fps.followsRule(r) {
				isCorrect = false
				nI := slices.Index(fps, r.num)
				cbI := slices.Index(fps, r.comesBefore)
				fps[nI], fps[cbI] = fps[cbI], fps[nI]
			}
		}
		if isCorrect {
			break
		}
	}

	return fps
}

type rule struct {
	num         int
	comesBefore int
}

func newRule(strs []string) (rule, error) {
	if len(strs) != 2 {
		return rule{}, fmt.Errorf("incorrect length of strs - %d", len(strs))
	}
	n, err := strconv.Atoi(strs[0])
	if err != nil {
		return rule{}, err
	}

	cb, err := strconv.Atoi(strs[1])
	if err != nil {
		return rule{}, err
	}

	return rule{
		num:         n,
		comesBefore: cb,
	}, nil
}

type pages []int

func newPages(strs []string) (pages, error) {
	ps := pages{}

	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			return ps, err
		}
		ps = append(ps, n)
	}
	return ps, nil
}

func (p pages) followsRule(r rule) bool {
	numI, cbI := 0, 0
	for i, n := range p {
		if n == r.num {
			numI = i
		}
		if n == r.comesBefore {
			cbI = i
		}
	}
	return numI < cbI
}

func (p pages) includesRule(r rule) bool {

	//check for num
	found := false
	for _, n := range p {
		if n == r.num {
			found = true
		}
	}
	if !found {
		return false
	}

	found = false
	//check for comesBefore num
	for _, n := range p {
		if n == r.comesBefore {
			found = true
		}
	}
	return found
}

func (p pages) middleNum() int {
	i := len(p) / 2
	return p[i]
}
