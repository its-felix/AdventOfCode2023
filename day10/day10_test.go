package day10

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines("day10_example1.txt")))
	println(SolvePart1(inputs.GetInputLines("day10_example2.txt")))
	println(SolvePart1(inputs.GetInputLines("day10_example3.txt")))
	println(SolvePart1(inputs.GetInputLines("day10.txt")))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines("day10.txt")))
}
