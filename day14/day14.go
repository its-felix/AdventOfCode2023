package day14

func SolvePart1(input <-chan string) int {
	return solve(tiltNorth(parse(input)))
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func solve(lines [][]rune) int {
	sum := 0

	for i, line := range lines {
		for _, r := range line {
			if r == 'O' {
				sum += len(lines) - i
			}
		}
	}

	return sum
}

func tiltNorth(lines [][]rune) [][]rune {
	for {
		anyMoved := false

		for row := 1; row < len(lines); row++ {
			for col := 0; col < len(lines[row]); col++ {
				if lines[row-1][col] == '.' && lines[row][col] == 'O' {
					lines[row-1][col] = lines[row][col]
					lines[row][col] = '.'
					anyMoved = true
				}
			}
		}

		if !anyMoved {
			break
		}
	}

	return lines
}

func parse(input <-chan string) [][]rune {
	lines := make([][]rune, 0)
	for line := range input {
		lines = append(lines, []rune(line))
	}

	return lines
}
