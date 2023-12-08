package day8

import (
	"strings"
)

type node struct {
	lr   [2]*node
	exit bool
}

func SolvePart1(input <-chan string) int {
	instructions, lookup, _ := parse(input)
	start, end := lookup["AAA"], lookup["ZZZ"]

	return steps(instructions, start, func(n *node) bool {
		return n == end
	})
}

func SolvePart2(input <-chan string) int {
	instructions, _, starts := parse(input)
	res := 0

	for i, start := range starts {
		numSteps := steps(instructions, start, func(n *node) bool {
			return n.exit
		})

		if i == 0 {
			res = numSteps
		} else {
			res = lcm(res, numSteps)
		}
	}

	return res
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}

func lcm(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	return x * y / gcd(x, y)
}

func steps(instructions []int, curr *node, isExit func(*node) bool) int {
	for i := 0; ; i++ {
		curr = curr.lr[instructions[i%len(instructions)]]
		if isExit(curr) {
			return i + 1
		}
	}
}

func parse(input <-chan string) ([]int, map[string]*node, []*node) {
	line := <-input
	instructions := make([]int, 0, len(line))

	for _, r := range line {
		var idx int
		if r == 'L' {
			idx = 0
		} else {
			idx = 1
		}

		instructions = append(instructions, idx)
	}

	lookup := make(map[string]*node)
	allStarts := make([]*node, 0)

	for line = range input {
		if line == "" {
			continue
		}

		id := line[:3]
		self := getOrCreateNode(lookup, id)
		self.lr[0] = getOrCreateNode(lookup, line[7:10])
		self.lr[1] = getOrCreateNode(lookup, line[12:15])

		if strings.HasSuffix(id, "A") {
			allStarts = append(allStarts, self)
		}
	}

	return instructions, lookup, allStarts
}

func getOrCreateNode(lookup map[string]*node, id string) *node {
	if n, ok := lookup[id]; ok {
		return n
	}

	n := &node{
		lr:   [2]*node{nil, nil},
		exit: strings.HasSuffix(id, "Z"),
	}
	lookup[id] = n
	return n
}
