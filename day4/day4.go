package day4

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"unicode"
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

		fmt.Printf("card=%d wins=%d copies=%d", cardNum, wins, copyCount)

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

		fmt.Printf(" sum=%d\n", sum)
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

	line = trimUntilDigit(line)
	line, cardNum, anyMatched = readInt(line)
	if !anyMatched {
		panic("could not find cardNum")
	}

	var winning, mine []int
	line, winning = readNumbers(trimUntilDigit(line))
	line, mine = readNumbers(trimUntilDigit(line))

	return cardNum, winning, mine
}

func readNumbers(s string) (string, []int) {
	res := make([]int, 0)
	for {
		var num int
		var anyMatched bool
		s, num, anyMatched = readInt(strings.TrimSpace(s))

		if anyMatched {
			res = append(res, num)
		} else {
			break
		}
	}

	return s, res
}

func readInt(s string) (string, int, bool) {
	maxI := -1
	num := 0
	anyMatched := false

	for i, r := range s {
		v := int(r) - '0'
		if v >= 0 && v < 10 {
			num *= 10
			num += v
			maxI = i
			anyMatched = true
		} else {
			break
		}
	}

	return s[maxI+1:], num, anyMatched
}

func trimUntilDigit(s string) string {
	return strings.TrimLeftFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
}
