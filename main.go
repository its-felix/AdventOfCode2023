package main

import (
	"github.com/its-felix/AdventOfCode2023/day01"
	"github.com/its-felix/AdventOfCode2023/inputs"
)

func main() {
	day01.Solve(inputs.GetInputLines("day1.txt"), day01.Lookup1)
	day01.Solve(inputs.GetInputLines("day1.txt"), day01.Lookup2)
}
