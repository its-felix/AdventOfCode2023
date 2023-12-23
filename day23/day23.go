package day23

import (
	"github.com/its-felix/AdventOfCode2023/util"
	"maps"
	"math"
)

const (
	north = iota
	east
	south
	west
)

type node struct {
	valid     bool
	slope     int
	connected [4]*node
}

type state struct {
	n         *node
	direction int
	steps     int
	seen      util.Set[*node]
}

func SolvePart1(input <-chan string) int {
	return solve(parse(input))
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func solve(start, end *node) int {
	queue := make([]state, 0)
	queue = append(queue, state{
		n:         start,
		direction: south,
		steps:     0,
		seen:      make(util.Set[*node]),
	})

	maxSteps := math.MinInt
	for len(queue) > 0 {
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if curr.n == end {
			maxSteps = max(maxSteps, curr.steps)
		}

		for direction, conn := range curr.n.connected {
			if conn == nil || !conn.valid || direction == opposite(curr.direction) {
				continue
			}

			if conn.slope != -1 && conn.slope != direction {
				continue
			}

			if seen := maps.Clone(curr.seen); seen.AddIfAbsent(curr.n) {
				queue = append(queue, state{
					n:         conn,
					direction: direction,
					steps:     curr.steps + 1,
					seen:      seen,
				})
			}
		}
	}

	return maxSteps
}

func opposite(direction int) int {
	switch direction {
	case north:
		return south

	case east:
		return west

	case south:
		return north

	case west:
		return east
	}

	panic("invalid direction")
}

func parse(input <-chan string) (*node, *node) {
	grid := make([][]*node, 0)
	var start, end *node

	row := 0
	for line := range input {
		for col, r := range line {
			if r == '#' {
				continue
			}

			var n *node
			grid, n = getOrCreateNode(grid, row, col)
			n.valid = true

			switch r {
			case '>':
				n.slope = east

			case '<':
				n.slope = west

			case 'v':
				n.slope = south

			case '^':
				n.slope = north

			default:
				n.slope = -1
			}

			for direction, idx := range getConnected(row, col) {
				var conn *node
				grid, conn = getOrCreateNode(grid, idx[0], idx[1])
				n.connected[direction] = conn
			}

			if start == nil {
				start = n
			}

			end = n
		}
		row++
	}

	return start, end
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
