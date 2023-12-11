package day01

import (
	"fmt"
)

var Lookup1 = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var Lookup2 = map[string]int{
	"1":     1,
	"one":   1,
	"2":     2,
	"two":   2,
	"3":     3,
	"three": 3,
	"4":     4,
	"four":  4,
	"5":     5,
	"five":  5,
	"6":     6,
	"six":   6,
	"7":     7,
	"seven": 7,
	"8":     8,
	"eight": 8,
	"9":     9,
	"nine":  9,
}

func Solve(input <-chan string, lookup map[string]int) {
	sum := 0

	for line := range input {
		sum += Combine(Parse(line, lookup))
	}

	println(sum)
}

func Combine(v1, v2 int) int {
	return v1*10 + v2
}

func Parse(line string, lookup map[string]int) (int, int) {
	return First(line, lookup), Last(line, lookup)
}

func First(line string, lookup map[string]int) int {
	for i := 0; i < len(line); i++ {
		for j := i + 1; j <= len(line); j++ {
			sub := line[i:j]
			if v, ok := lookup[sub]; ok {
				return v
			}
		}
	}

	panic(fmt.Sprintf("no first value found: %s", line))
}

func Last(line string, lookup map[string]int) int {
	for i := len(line); i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			sub := line[j:i]
			if v, ok := lookup[sub]; ok {
				return v
			}
		}
	}

	panic(fmt.Sprintf("no last value found: %s", line))
}
