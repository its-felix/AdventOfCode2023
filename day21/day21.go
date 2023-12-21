package day21

import "github.com/its-felix/AdventOfCode2023/util"

const (
	north = iota
	east
	south
	west

	plot  = rune('.')
	rock  = rune('#')
	start = rune('S')
)

type node struct {
	kind      rune
	connected [4]*node
}

type state struct {
	n              *node
	remainingSteps int
}

func SolvePart1(input <-chan string) int {
	return numPossiblePlots(parse(input), 64)
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func numPossiblePlots(s *node, steps int) int {
	queue := []state{{
		n:              s,
		remainingSteps: steps,
	}}

	seen := make(util.Set[state])
	final := make(util.Set[*node])

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.remainingSteps == 0 {
			final.Add(curr.n)
		}

		if curr.remainingSteps > 0 && seen.AddIfAbsent(curr) {
			for _, conn := range curr.n.connected {
				if conn == nil || conn.kind != plot {
					continue
				}

				queue = append(queue, state{
					n:              conn,
					remainingSteps: curr.remainingSteps - 1,
				})
			}
		}
	}

	return len(final)
}

func parse(input <-chan string) *node {
	grid := make([][]*node, 0)
	var s *node

	row := 0
	for line := range input {
		for col, r := range line {
			var n *node
			grid, n = getOrCreateNode(grid, row, col)
			n.kind = r

			for direction, idx := range getConnected(row, col) {
				var conn *node
				grid, conn = getOrCreateNode(grid, idx[0], idx[1])
				n.connected[direction] = conn
			}

			if r == start {
				s = n
				s.kind = plot
			}
		}
		row++
	}

	return s
}

func getConnected(row, col int) [4][2]int {
	var conn [4][2]int
	conn[north] = [2]int{row - 1, col}
	conn[east] = [2]int{row, col + 1}
	conn[south] = [2]int{row + 1, col}
	conn[west] = [2]int{row, col - 1}

	return conn
}

func getOrCreateNode(grid [][]*node, row, col int) ([][]*node, *node) {
	if row < 0 || col < 0 {
		return grid, nil
	}

	missing := row - len(grid) + 1
	for missing > 0 {
		grid = append(grid, make([]*node, 0))
		missing--
	}

	missing = col - len(grid[row]) + 1
	for missing > 0 {
		grid[row] = append(grid[row], &node{})
		missing--
	}

	return grid, grid[row][col]
}
