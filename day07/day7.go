package day07

import (
	"cmp"
	"math/bits"
	"slices"
	"strconv"
)

var strengthsPart1, maxStrengthPart1 = prepareStrengths([]rune{
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

var strengthsPart2, maxStrengthPart2 = prepareStrengths([]rune{
	'J',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	'T',
	'Q',
	'K',
	'A',
})

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
	hand [5]rune
	bid  int
}

type roundResult struct {
	hType  handType
	hValue uint64
	bid    int
}

func SolvePart1(input <-chan string) int {
	return calculateWinnings(scoreAll(strengthsPart1, maxStrengthPart1, parse(input), false))
}

func SolvePart2(input <-chan string) int {
	return calculateWinnings(scoreAll(strengthsPart2, maxStrengthPart2, parse(input), true))
}

func calculateWinnings(results []roundResult) int {
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

func scoreAll(strengths []uint64, maxStrength uint64, rounds []round, joker bool) []roundResult {
	results := make([]roundResult, 0)
	for _, r := range rounds {
		hType, hValue := score(strengths, maxStrength, r.hand, joker)
		results = append(results, roundResult{
			hType:  hType,
			hValue: hValue,
			bid:    r.bid,
		})
	}

	return results
}

func score(strengths []uint64, maxStrength uint64, hand [5]rune, joker bool) (handType, uint64) {
	counts := make([]int, int(maxStrength))
	handValue := uint64(0)

	for i, v := range hand {
		cardStrength := strengths[v]
		counts[cardStrength-1]++
		shiftBits := (len(hand) - i - 1) * (64 - bits.LeadingZeros64(maxStrength))
		cardValue := cardStrength << shiftBits
		handValue |= cardValue
	}

	countJoker := counts[strengths['J']-1]
	if joker {
		counts[strengths['J']-1] = 0
	}

	slices.SortFunc(counts, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	if joker {
		counts[0] += countJoker
	}

	return findHandType(counts), handValue
}

func findHandType(counts []int) handType {
	switch counts[0] {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if counts[1] == 2 {
			return fullHouse
		} else {
			return threeOfAKind
		}
	case 2:
		if counts[1] == 2 {
			return twoPair
		} else {
			return onePair
		}
	case 1:
		return highCard
	default:
		panic("no matching type")
	}
}

func parse(input <-chan string) []round {
	rounds := make([]round, 0)
	for line := range input {
		r := round{
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
