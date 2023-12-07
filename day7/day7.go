package day7

import (
	"cmp"
	"math/bits"
	"slices"
	"strconv"
)

var strengths, maxStrength = prepareStrengths([]rune{
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'T',
	'J',
	'Q',
	'K',
	'A',
})
var maxStrengthLeadingZeros = bits.LeadingZeros64(maxStrength)

type handType int

const (
	highCard = handType(iota)
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func prepareStrengths(values []rune) ([]uint64, uint64) {
	s := make([]uint64, 0)
	for i, v := range values {
		v := int(v)
		if missing := v - len(s) + 1; missing > 0 {
			s = append(s, make([]uint64, missing)...)
		}

		s[v] = uint64(i + 1)
	}

	return s, uint64(len(values))
}

type round struct {
	raw  string
	hand [5]rune
	bid  int
}

type roundResult struct {
	raw    string
	hType  handType
	hValue uint64
	bid    int
}

func SolvePart1(input <-chan string) int {
	results := make([]roundResult, 0)
	for _, r := range parse(input) {
		hType, hValue := score(r.hand)
		results = append(results, roundResult{
			raw:    r.raw,
			hType:  hType,
			hValue: hValue,
			bid:    r.bid,
		})
	}

	slices.SortFunc(results, func(a, b roundResult) int {
		if sType := cmp.Compare(a.hType, b.hType); sType != 0 {
			return sType
		}

		return cmp.Compare(a.hValue, b.hValue)
	})

	winnings := 0
	for i, r := range results {
		winnings += r.bid * (i + 1)
	}

	return winnings
}

func SolvePart2(input <-chan string) int {
	rounds := parse(input)
	println(rounds)
	return 0
}

func score(hand [5]rune) (handType, uint64) {
	counts := make([]int, int(maxStrength))
	handValue := uint64(0)

	for i, v := range hand {
		cardStrength := strengths[v]
		counts[cardStrength-1]++
		shiftBits := (len(hand) - i - 1) * (64 - maxStrengthLeadingZeros)
		cardValue := cardStrength << shiftBits
		handValue |= cardValue
	}

	slices.SortFunc(counts, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	switch counts[0] {
	case 5:
		return fiveOfAKind, handValue
	case 4:
		return fourOfAKind, handValue
	case 3:
		if counts[1] == 2 {
			return fullHouse, handValue
		} else {
			return threeOfAKind, handValue
		}
	case 2:
		if counts[1] == 2 {
			return twoPair, handValue
		} else {
			return onePair, handValue
		}
	case 1:
		return highCard, handValue
	default:
		panic("no matching type")
	}
}

func parse(input <-chan string) []round {
	rounds := make([]round, 0)
	for line := range input {
		r := round{
			raw:  line,
			hand: [5]rune{0, 0, 0, 0, 0},
		}

		for i, v := range line[:5] {
			r.hand[i] = v
		}

		r.bid, _ = strconv.Atoi(line[6:])
		rounds = append(rounds, r)
	}

	return rounds
}
