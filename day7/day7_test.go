package day7

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines(7)))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines(7)))
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart1(inputs.GetInputLines(7)) != 251136060 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart2(inputs.GetInputLines(7)) != 249400220 {
			b.FailNow()
		}
	}
}
