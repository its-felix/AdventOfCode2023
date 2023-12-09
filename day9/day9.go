package day9

import (
	"github.com/its-felix/AdventOfCode2023/util"
	"strings"
)

func SolvePart1(input <-chan string) int {
	sum := 0
	for _, nums := range parse(input) {
		sum += extrapolateFuture(nums)
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	sum := 0
	for _, nums := range parse(input) {
		sum += extrapolatePast(nums)
	}

	return sum
}

func extrapolateFuture(nums []int) int {
	ws := buildWs(nums)

	for i := len(ws) - 1; i >= 1; i-- {
		extrapolated := ws[i][len(ws[i])-1] + ws[i-1][len(ws[i-1])-1]
		ws[i-1] = append(ws[i-1], extrapolated)
	}

	return ws[0][len(ws[0])-1]
}

func extrapolatePast(nums []int) int {
	ws := buildWs(nums)

	for i := len(ws) - 1; i >= 1; i-- {
		extrapolated := ws[i-1][0] - ws[i][0]
		ws[i-1] = append([]int{extrapolated}, ws[i-1]...)
	}

	return ws[0][0]
}

func buildWs(nums []int) [][]int {
	ws := make([][]int, 1)
	ws[0] = nums

	for {
		nums = ws[len(ws)-1]
		diffs := make([]int, 0, len(nums)-1)
		allZero := true

		for i := 1; i < len(nums); i++ {
			diff := nums[i] - nums[i-1]
			diffs = append(diffs, diff)
			allZero = allZero && diff == 0
		}

		ws = append(ws, diffs)

		if allZero {
			break
		}
	}

	return ws
}

func parse(input <-chan string) [][]int {
	readings := make([][]int, 0)

	for line := range input {
		readings = append(readings, readInts(line))
	}

	return readings
}

func readInts(line string) []int {
	ints := make([]int, 0)
	for {
		line = strings.TrimSpace(line)
		if len(line) < 1 {
			break
		}

		neg := line[0] == '-'
		if neg {
			line = line[1:]
		}

		var num int
		var matched bool
		if line, num, matched = util.ReadInt(line); matched {
			if neg {
				num = -num
			}

			ints = append(ints, num)
		}
	}

	return ints
}
