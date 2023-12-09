package day9

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines("day9.txt")))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines("day9.txt")))
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart1(inputs.GetInputLines("day9.txt")) != 1584748274 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart2(inputs.GetInputLines("day9.txt")) != 1026 {
			b.FailNow()
		}
	}
}
