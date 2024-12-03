package hands

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Hand struct {
	cards string
	rank  int
	bet   int
	t     handType
}

func (h Hand) String() string {
	return fmt.Sprintf("Hand{cards: %s, rank: %d, bet: %d, type: %s}", h.cards, h.rank, h.bet, h.t)
}

type handType int

const (
	fiveKind handType = iota
	fourKind
	fullHouse
	threeKind
	twoPair
	onePair
	highCard
)

func (h handType) String() string {
	return []string{"five of a kind", "four of a kind", "full house", "three of a kind", "two pair", "one pair", "high card"}[h]
}

type Hands []Hand

func NewHands(input string) Hands {
	hands := []Hand{}

	lines := strings.Split(input, "\n")

	ch := make(chan Hand)
	var wg sync.WaitGroup

	for _, l := range lines {
		if len(l) == 0 {
			break
		}
		wg.Add(1)

		go func() {
			defer wg.Done()
			cards := strings.Split(l, " ")[0]
			t := getHandType(cards)
			betStr := strings.Split(l, " ")[1]
			bet, _ := strconv.Atoi(betStr)

			ch <- Hand{
				cards: cards,
				rank:  0,
				bet:   bet,
				t:     t,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for h := range ch {
		hands = append(hands, h)
	}

	return hands
}

func getHandType(cards string) handType {
	m := make(map[rune]int)

	for _, card := range cards {
		m[card]++
	}

	for v, c := range m {
		if c == 5 {
			return fiveKind
		} else if c == 4 {
			return fourKind
		} else if c == 3 {
			for _, cc := range m {
				if cc == 2 {
					return fullHouse
				}
			}
			return threeKind
		} else if c == 2 {
			for vv, cc := range m {
				if vv != v && cc == 2 {
					return twoPair
				}
			}
			return onePair
		}
	}
	return highCard
}

func (h Hands) SortOrder() {
	for i, hh := range h {
		for j := i; j < len(h); j++ {
			if hh.t == highCard {
				fmt.Println(hh)
			}
		}
	}
}

func (h Hands) CalculateScore() int {
	return 0
}
