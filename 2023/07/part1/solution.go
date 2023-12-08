package main

import (
	"fmt"
	"sort"
	"strings"
)

type cardStrength int64

const (
	unknownCard cardStrength = iota
	twoCard
	threeCard
	fourCard
	fiveCard
	sixCard
	sevenCard
	eightCard
	nineCard
	tenCard
	jackCard
	queenCard
	kindCard
	aceCard
)

var cardEnum = map[string]cardStrength{
	"A": aceCard, "K": kindCard, "Q": queenCard, "J": jackCard, "T": tenCard, "9": nineCard, "8": eightCard, "7": sevenCard, "6": sixCard, "5": fiveCard, "4": fourCard, "3": threeCard, "2": twoCard,
}

type handStrength int64

const (
	unknownHand handStrength = iota
	highCard
	onePair
	twoPair
	threeKind
	fullHouse
	fourKind
	fiveKind
)

func solution(file string) (int64, error) {
	hands, err := ForEachLine(file, processLine)
	if err != nil {
		return -1, err
	}

	reverseSortHands(hands)
	return sumWinnings(hands), nil
}

type card struct {
	label    string
	strength cardStrength
}

func (c card) String() string {
	return c.label
}

type hand struct {
	cards    []*card
	cardsStr string
	strength handStrength
	bid      int64
}

func newHand(handCards string, bid int64) *hand {
	cards := []*card{}
	cardStr := strings.Split(handCards, "")
	for _, c := range cardStr {
		newCard := &card{
			label:    c,
			strength: cardEnum[c],
		}
		cards = append(cards, newCard)
	}

	cardCount := map[string]int64{}
	for _, c := range cards {
		cardCount[c.label]++
	}

	var five, four, three, two, one int64
	for _, count := range cardCount {
		switch {
		case count == 5:
			five++
		case count == 4:
			four++
		case count == 3:
			three++
		case count == 2:
			two++
		case count == 1:
			one++
		default:
			fmt.Println("hit default count=", count)
		}
	}

	var strength handStrength
	switch {
	case five == 1:
		strength = fiveKind
	case four == 1 && one == 1:
		strength = fourKind
	case three == 1 && two == 1:
		strength = fullHouse
	case three == 1 && one == 2:
		strength = threeKind
	case two == 2 && one == 1:
		strength = twoPair
	case two == 1 && one == 3:
		strength = onePair
	case one == 5:
		strength = highCard
	default:
		fmt.Println("unsupported card combination")
	}

	return &hand{
		cards:    cards,
		cardsStr: handCards,
		strength: strength,
		bid:      bid,
	}
}

func reverseSortHands(hands []*hand) {
	sort.Slice(hands, func(i, j int) bool {
		c, n := hands[i], hands[j]
		if c.strength == n.strength {
			// if both hands are the same strength, then sort based on the cards
			return lessThan(c.cards, n.cards)
		}
		return c.strength < n.strength
	})
}

func lessThan(l, g []*card) bool {
	for i := range l {
		l1, g1 := l[i].strength, g[i].strength
		if l1 == g1 {
			i++
			continue
		}
		if l1 < g1 {
			return true
		}
		if l1 > g1 {
			return false
		}
	}
	fmt.Println("cards are all equal")
	return false
}

// sumWinnings gets the total winnings of the set of hands by adding up the
// result of multiplying each hand's bid with its rank.
func sumWinnings(handsByRank []*hand) int64 {
	var totalWinnings int64
	for i, hand := range handsByRank {
		rank := i + 1 // the rank is off by 1 so increment once
		totalWinnings += int64(rank) * hand.bid
	}
	return totalWinnings
}
