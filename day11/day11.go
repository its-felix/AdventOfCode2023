package day11

func SolvePart1(input <-chan string) int {
	return solve(input, 2)
}

func SolvePart2(input <-chan string) int {
	return solve(input, 1000000)
}

func solve(input <-chan string, expansion int) int {
	galaxies, expandedRows, expandedCols := parse(input)
	return sumDistance(galaxies, expandedRows, expandedCols, expansion)
}

func sumDistance(galaxies [][2]int, expandedRows, expandedCols []int, expansion int) int {
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j], expandedRows, expandedCols, expansion)
		}
	}

	return sum
}

func distance(a, b [2]int, expandedRows, expandedCols []int, expansion int) int {
	startRow, endRow, startCol, endCol := a[0], b[0], a[1], b[1]
	if startRow > endRow {
		temp := startRow
		startRow = endRow
		endRow = temp
	}
	if startCol > endCol {
		temp := startCol
		startCol = endCol
		endCol = temp
	}

	rawDistance := (endRow - startRow) + (endCol - startCol)
	expandedRowsBetween := expandedRows[endRow] - expandedRows[startRow]
	expandedColsBetween := expandedCols[endCol] - expandedCols[startCol]
	totalExpanded := expandedRowsBetween + expandedColsBetween

	return rawDistance - totalExpanded + (totalExpanded * expansion)
}

func parse(input <-chan string) ([][2]int, []int, []int) {
	rowHasGalaxy := make([]bool, 0)
	colHasGalaxy := make([]bool, 0)
	galaxies := make([][2]int, 0)

	row := 0
	for line := range input {
		rowHasGalaxy = grow(rowHasGalaxy, row+1)

		for col, r := range line {
			colHasGalaxy = grow(colHasGalaxy, col+1)

			if r == '#' {
				rowHasGalaxy[row] = true
				colHasGalaxy[col] = true
				galaxies = append(galaxies, [2]int{row, col})
			}
		}
		row++
	}

	expandedRows := make([]int, len(rowHasGalaxy))
	expandedCols := make([]int, len(colHasGalaxy))

	total := 0
	for i, v := range rowHasGalaxy {
		if !v {
			total++
		}

		expandedRows[i] = total
	}

	total = 0
	for i, v := range colHasGalaxy {
		if !v {
			total++
		}

		expandedCols[i] = total
	}

	return galaxies, expandedRows, expandedCols
}

func grow[T any](s []T, size int) []T {
	missing := size - len(s) + 1
	if missing > 0 {
		s = append(s, make([]T, missing)...)
	}

	return s
}
