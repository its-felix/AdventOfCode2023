package day11

func SolvePart1(input <-chan string) int {
	galaxies, rowHasGalaxy, colHasGalaxy := parse(input)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j], rowHasGalaxy, colHasGalaxy, 2)
		}
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	galaxies, rowHasGalaxy, colHasGalaxy := parse(input)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j], rowHasGalaxy, colHasGalaxy, 1000000)
		}
	}

	return sum
}

func distance(a, b [2]int, rowHasGalaxy, colHasGalaxy []bool, expansion int) int {
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

	d := 0
	for row := startRow; row < endRow; row++ {
		if rowHasGalaxy[row] {
			d += 1
		} else {
			d += expansion
		}
	}

	for col := startCol; col < endCol; col++ {
		if colHasGalaxy[col] {
			d += 1
		} else {
			d += expansion
		}
	}

	return d
}

func parse(input <-chan string) ([][2]int, []bool, []bool) {
	rowHasGalaxy := make([]bool, 0)
	colHasGalaxy := make([]bool, 0)
	galaxies := make([][2]int, 0)

	row := 0
	for line := range input {
		rowHasGalaxy = expand(rowHasGalaxy, row+1)

		for col, r := range line {
			colHasGalaxy = expand(colHasGalaxy, col+1)

			if r == '#' {
				rowHasGalaxy[row] = true
				colHasGalaxy[col] = true
				galaxies = append(galaxies, [2]int{row, col})
			}
		}
		row++
	}

	return galaxies, rowHasGalaxy, colHasGalaxy
}

func expand[T any](s []T, size int) []T {
	missing := size - len(s) + 1
	if missing > 0 {
		s = append(s, make([]T, missing)...)
	}

	return s
}
