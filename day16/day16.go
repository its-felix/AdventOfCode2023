package day16

import "github.com/its-felix/AdventOfCode2023/util"

const (
	north = iota
	east
	south
	west
)

func SolvePart1(input <-chan string) int {
	return countEnergized(parse(input), 0, 0, east)
}

func SolvePart2(input <-chan string) int {
	grid := parse(input)
	most := 0

	for row := 0; row < len(grid); row++ {
		most = max(most, countEnergized(grid, row, 0, east))
		most = max(most, countEnergized(grid, row, len(grid[row])-1, west))
	}

	for col := 0; col < len(grid[0]); col++ {
		most = max(most, countEnergized(grid, 0, col, south))
	}

	for col := 0; col < len(grid[len(grid)-1]); col++ {
		most = max(most, countEnergized(grid, len(grid)-1, col, north))
	}

	return most
}

func countEnergized(grid [][]rune, row, col, direction int) int {
	sum := 0
	queue := [][3]int{{row, col, direction}}
	energized := make(util.Set[[2]int])
	seen := make(util.Set[[3]int])

	for len(queue) > 0 {
		curr := queue[0]
		row, col, direction = curr[0], curr[1], curr[2]
		queue = queue[1:]

		if seen.AddIfAbsent(curr) {
			if energized.AddIfAbsent([2]int{row, col}) {
				sum++
			}

			for _, next := range nextPosAndDirection(grid[row][col], row, col, direction) {
				if !isValidIndex(grid, next[0], next[1]) {
					continue
				}

				queue = append(queue, next)
			}
		}
	}

	return sum
}

func isValidIndex(grid [][]rune, row, col int) bool {
	if row < 0 || col < 0 {
		return false
	}

	return row < len(grid) && col < len(grid[row])
}

func nextPosAndDirection(r rune, row, col, direction int) [][3]int {
	if r == '.' {
		return [][3]int{nextPos(row, col, direction)}
	}

	if r == '/' {
		switch direction {
		case north:
			return [][3]int{nextPos(row, col, east)}

		case east:
			return [][3]int{nextPos(row, col, north)}

		case south:
			return [][3]int{nextPos(row, col, west)}

		case west:
			return [][3]int{nextPos(row, col, south)}
		}

		panic("invalid direction")
	}

	if r == '\\' {
		switch direction {
		case north:
			return [][3]int{nextPos(row, col, west)}

		case east:
			return [][3]int{nextPos(row, col, south)}

		case south:
			return [][3]int{nextPos(row, col, east)}

		case west:
			return [][3]int{nextPos(row, col, north)}
		}

		panic("invalid direction")
	}

	if r == '-' {
		switch direction {
		case north, south:
			return [][3]int{
				nextPos(row, col, east),
				nextPos(row, col, west),
			}

		case east, west:
			return [][3]int{nextPos(row, col, direction)}
		}

		panic("invalid direction")
	}

	if r == '|' {
		switch direction {
		case north, south:
			return [][3]int{nextPos(row, col, direction)}

		case east, west:
			return [][3]int{
				nextPos(row, col, north),
				nextPos(row, col, south),
			}
		}

		panic("invalid direction")
	}

	panic("invalid character")
}

func nextPos(row, col, direction int) [3]int {
	switch direction {
	case north:
		return [3]int{row - 1, col, direction}

	case east:
		return [3]int{row, col + 1, direction}

	case south:
		return [3]int{row + 1, col, direction}

	case west:
		return [3]int{row, col - 1, direction}
	}

	panic("invalid direction")
}

func parse(input <-chan string) [][]rune {
	grid := make([][]rune, 0)
	for line := range input {
		grid = append(grid, []rune(line))
	}

	return grid
}
