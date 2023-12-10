package day10

import (
	"math"
)

const (
	north = 0
	east  = 1
	south = 2
	west  = 3
)

var connections = map[rune][]int{
	'|': {north, south},
	'-': {east, west},
	'L': {north, east},
	'J': {north, west},
	'7': {south, west},
	'F': {south, east},
	'.': {},
	'S': {},
}

type node struct {
	raw      rune
	valid    bool
	distance int
	conn     [4]*node
}

func SolvePart1(input <-chan string) int {
	start := parse(input)
	return distance(start).distance
	// return follow(start)
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func follow(n *node) int {
	for {
		for _, conn := range n.conn {
			if conn == nil || !conn.valid {
				continue
			}

			if n.distance > conn.distance {
				return n.distance
			}

			n = conn
			break
		}
	}
}

func distance(n *node) *node {
	var prev *node
	queue := [][2]*node{{n, nil}}
	dirty := map[*node]bool{n: true}

	farthest := n

	for len(queue) > 0 {
		idx := len(queue) - 1
		n = queue[idx][0]
		prev = queue[idx][1]
		queue = queue[:idx]

		dist := n.distance + 1

		for _, conn := range n.conn {
			if conn == nil || !conn.valid || conn == prev {
				continue
			}

			if conn.distance > dist {
				conn.distance = dist
				dirty[conn] = true
				queue = append(queue, [2]*node{conn, n})
			}
		}

		if v, _ := dirty[farthest]; v || n.distance > farthest.distance {
			farthest = n
		}

		delete(dirty, n)
	}

	return farthest
}

func parse(input <-chan string) *node {
	grid := make([][]*node, 0)
	var start *node

	row := 0
	for line := range input {
		for col, r := range line {
			var n *node
			grid, n = getOrCreateNode(grid, row, col)
			n.valid = true
			n.raw = r

			for _, direction := range connections[r] {
				connRow, connCol := directionToPos(row, col, direction)

				var connectedNode *node
				grid, connectedNode = getOrCreateNode(grid, connRow, connCol)

				if connectedNode != nil {
					n.conn[direction] = connectedNode
					connectedNode.conn[opposite(direction)] = n
				}
			}

			if r == 'S' {
				start = n
				start.distance = 0
			}
		}

		row++
	}

	return start
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
		grid[row] = append(grid[row], &node{distance: math.MaxInt})
		missing--
	}

	return grid, grid[row][col]
}

func directionToPos(row, col, direction int) (int, int) {
	switch direction {
	case north:
		return row - 1, col
	case east:
		return row, col + 1
	case south:
		return row + 1, col
	case west:
		return row, col - 1
	}

	panic("invalid direction")
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
