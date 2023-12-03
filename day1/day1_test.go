package day1

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	Solve(inputs.GetInputLines(1), Lookup1)
}

func TestSolvePart2(t *testing.T) {
	Solve(inputs.GetInputLines(1), Lookup2)
}
