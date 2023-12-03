package main

import (
	"github.com/its-felix/AdventOfCode2023/day1"
	"github.com/its-felix/AdventOfCode2023/inputs"
)

func main() {
	day1.Solve(inputs.GetInputLines(1), day1.Lookup1)
	day1.Solve(inputs.GetInputLines(1), day1.Lookup2)
}
