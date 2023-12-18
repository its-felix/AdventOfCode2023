package day17

import (
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
	row       int
	col       int
	cost      uint64
	connected [4]*node
}

type traverseState struct {
	n              *node
	direction      int
	directionCount int
	cost           uint64
	past4          [4]*node
}

func SolvePart1(input <-chan string) uint64 {
	start, end := parse(input)
	return traverse(start, end, 3)
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func traverse(start, end *node, maxSameDirection int) uint64 {
	minFinalCost := uint64(math.MaxUint64)
	minEntryCost := make(map[[4]*node]uint64)
	queue := make([]traverseState, 0)

	for _, direction := range []int{east, south} {
		queue = append(queue, traverseState{
			n:              start,
			direction:      direction,
			directionCount: 1,
			cost:           0,
		})
	}

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		if state.n == end {
			if state.cost < minFinalCost {
				minFinalCost = state.cost
			}
			continue
		}

		for direction, conn := range state.n.connected {
			if conn == nil || !conn.valid || direction == opposite(state.direction) {
				continue
			}

			cost := state.cost + conn.cost
			if cost+distance(conn, end) >= minFinalCost {
				continue
			}

			past4 := [4]*node{conn, state.past4[0], state.past4[1], state.past4[2]}
			if mCost, ok := minEntryCost[past4]; ok && cost >= mCost {
				continue
			}

			directionCount := 1
			if direction == state.direction {
				directionCount = state.directionCount + 1
				if directionCount > maxSameDirection {
					continue
				}
			}

			minEntryCost[past4] = cost
			queue = append(queue, traverseState{
				n:              conn,
				direction:      direction,
				directionCount: directionCount,
				cost:           cost,
				past4:          past4,
			})
		}
	}

	return minFinalCost
}

func distance(start, end *node) uint64 {
	rowStart, colStart, rowEnd, colEnd := start.row, start.col, end.row, end.col
	if rowStart > rowEnd {
		temp := rowStart
		rowStart = rowEnd
		rowEnd = temp
	}

	if colStart > colEnd {
		temp := colStart
		colStart = colEnd
		colEnd = temp
	}

	return uint64((rowEnd - rowStart) + (colEnd - colStart))
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
	var end *node
	grid := make([][]*node, 0)

	row := 0
	for line := range input {
		for col, r := range line {
			cost := int(r) - '0'
			if cost < 0 || cost > 9 {
				panic("invalid input")
			}

			var n *node
			grid, n = getOrCreateNode(grid, row, col)
			n.valid = true
			n.cost = uint64(cost)

			for direction, idx := range getConnected(row, col) {
				var conn *node
				grid, conn = getOrCreateNode(grid, idx[0], idx[1])
				n.connected[direction] = conn
			}

			end = n
		}

		row++
	}

	return grid[0][0], end
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
		grid[row] = append(grid[row], &node{row: row, col: col})
		missing--
	}

	return grid, grid[row][col]
}
