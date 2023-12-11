package day03

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines("day3.txt")))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines("day3.txt")))
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SolvePart1(inputs.GetInputLines("day3.txt"))
		if v != 556367 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := SolvePart2(inputs.GetInputLines("day3.txt"))
		if v != 89471771 {
			b.FailNow()
		}
	}
}
