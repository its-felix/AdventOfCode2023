package day11

import (
	"github.com/its-felix/AdventOfCode2023/inputs"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	println(SolvePart1(inputs.GetInputLines("day11.txt")))
}

func TestSolvePart2(t *testing.T) {
	println(SolvePart2(inputs.GetInputLines("day11.txt")))
}

func BenchmarkSolvePart1Full(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart1(inputs.GetInputLines("day11.txt")) != 9543156 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2Full(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if SolvePart2(inputs.GetInputLines("day11.txt")) != 625243292686 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart1Prepared(b *testing.B) {
	galaxies, expandedRows, expandedCols := parse(inputs.GetInputLines("day11.txt"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sumDistance(galaxies, expandedRows, expandedCols, 2) != 9543156 {
			b.FailNow()
		}
	}
}

func BenchmarkSolvePart2Prepared(b *testing.B) {
	galaxies, expandedRows, expandedCols := parse(inputs.GetInputLines("day11.txt"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if sumDistance(galaxies, expandedRows, expandedCols, 1000000) != 625243292686 {
			b.FailNow()
		}
	}
}
