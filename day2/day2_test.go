package day2

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	SolvePart1(inputs.GetInputLines(2), 12, 13, 14)
}

func TestSolvePart2(t *testing.T) {
	SolvePart2(inputs.GetInputLines(2))
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart1(inputs.GetInputLines(2), 12, 13, 14)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolvePart2(inputs.GetInputLines(2))
	}
}
