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

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(inputs.GetInputLines(1), Lookup1)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(inputs.GetInputLines(1), Lookup2)
	}
}
