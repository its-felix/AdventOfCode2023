package day08

import "strings"

type node struct {
	lr   [2]*node
	exit bool
}

func SolvePart1(input <-chan string) int {
	instructions, starts := parse(
		input,
		func(s string) bool {
			return s == "AAA"
		},
		func(s string) bool {
			return s == "ZZZ"
		},
	)
	return steps(instructions, starts[0])
}

func SolvePart2(input <-chan string) int {
	instructions, starts := parse(
		input,
		func(s string) bool {
			return strings.HasSuffix(s, "A")
		},
		func(s string) bool {
			return strings.HasSuffix(s, "Z")
		},
	)
	res := 0

	for i, start := range starts {
		numSteps := steps(instructions, start)

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

func steps(instructions []int, curr *node) int {
	for i := 0; ; i++ {
		curr = curr.lr[instructions[i%len(instructions)]]
		if curr.exit {
			return i + 1
		}
	}
}

func parse(input <-chan string, isStart, isExit func(string) bool) ([]int, []*node) {
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
		self.exit = isExit(id)

		if isStart(id) {
			allStarts = append(allStarts, self)
		}
	}

	return instructions, allStarts
}

func getOrCreateNode(lookup map[string]*node, id string) *node {
	if n, ok := lookup[id]; ok {
		return n
	}

	n := &node{
		lr: [2]*node{nil, nil},
	}
	lookup[id] = n
	return n
}
