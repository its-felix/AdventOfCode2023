package day4

import (
	"github.com/its-felix/AdventOfCode2023/util"
	"math"
	"slices"
)

func SolvePart1(input <-chan string) int {
	sum := 0

	for line := range input {
		_, winning, mine := Parse(line)
		sum += Score(winning, mine)
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	sum := 0
	copies := make([]int, 0)

	for line := range input {
		cardNum, winning, mine := Parse(line)
		for len(copies) < cardNum {
			copies = append(copies, 0)
		}

		wins := CountWinning(winning, mine)
		copyCount := copies[cardNum-1] // number of copies of this card from previous cards

		for i := 0; i < wins; i++ {
			wonCardIndex := cardNum + i
			for len(copies) <= wonCardIndex {
				copies = append(copies, 0)
			}

			copies[wonCardIndex] += 1         // one for the original
			copies[wonCardIndex] += copyCount // one for each copy
		}

		sum += 1                      // one for the original
		sum += wins * (copyCount + 1) // one for each copy
	}

	return sum
}

func Score(winning, mine []int) int {
	wins := CountWinning(winning, mine)
	if wins == 0 {
		return 0
	}

	return int(math.Pow(2, float64(wins-1)))
}

func CountWinning(winning, mine []int) int {
	wins := 0
	for _, num := range mine {
		if slices.Contains(winning, num) {
			wins += 1
		}
	}

	return wins
}

func Parse(line string) (int, []int, []int) {
	var anyMatched bool
	var cardNum int

	line = util.TrimNonDigit(line)
	line, cardNum, anyMatched = util.ReadInt(line)
	if !anyMatched {
		panic("could not find cardNum")
	}

	var winning, mine []int
	line, winning = util.ReadInts(util.TrimNonDigit(line))
	line, mine = util.ReadInts(util.TrimNonDigit(line))

	return cardNum, winning, mine
}
