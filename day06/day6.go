package day06

import (
	"github.com/its-felix/AdventOfCode2023/util"
	"strconv"
)

func SolvePart1(input <-chan string) int {
	res := 0
	for _, race := range parse(input) {
		ways := calculatePossibleWays(race[0], race[1])
		if res == 0 {
			res = len(ways)
		} else {
			res *= len(ways)
		}
	}

	return res
}

func SolvePart2(input <-chan string) int {
	race := fixKerning(parse(input))
	return len(calculatePossibleWays(race[0], race[1]))
}

func calculatePossibleWays(t, d int) []int {
	r := make([]int, 0)
	for hold := t - 1; hold >= 1; hold-- {
		travelT := t - hold
		travelD := travelT * hold

		if travelD > d {
			r = append(r, hold)
		}
	}

	return r
}

func fixKerning(input [][2]int) [2]int {
	r := [2]int{0, 0}
	for _, race := range input {
		r[0], _ = strconv.Atoi(strconv.Itoa(r[0]) + strconv.Itoa(race[0]))
		r[1], _ = strconv.Atoi(strconv.Itoa(r[1]) + strconv.Itoa(race[1]))
	}

	return r
}

func parse(input <-chan string) [][2]int {
	t, d := <-input, <-input
	t, d = util.TrimNonDigit(t), util.TrimNonDigit(d)

	_, times := util.ReadInts(t)
	_, distances := util.ReadInts(d)

	if len(times) != len(distances) {
		panic("lengths dont match")
	}

	r := make([][2]int, len(times))
	for i := 0; i < len(times); i++ {
		r[i] = [2]int{times[i], distances[i]}
	}

	return r
}
