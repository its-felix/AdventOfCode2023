package day14

func SolvePart1(input <-chan string) int {
	lines := parse(input)
	tiltNorth(lines)
	return solve(lines)
}

func SolvePart2(input <-chan string) int {
	lines := parse(input)
	runCycles(lines, 1000000000)
	return solve(lines)
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

func buildString(lines [][]rune) string {
	s := ""
	for _, line := range lines {
		s += string(line)
		s += "\n"
	}

	return s[:len(s)-1]
}

func runCycles(lines [][]rune, times int) {
	var cycleStart, cycleLen int
	cache := make(map[string]int)

	for i := 1; i <= times && cycleLen == 0; i++ {
		runCycle(lines)
		s := buildString(lines)
		if prev, ok := cache[s]; ok {
			cycleStart = prev
			cycleLen = i - prev
		} else {
			cache[s] = i
		}
	}

	if cycleLen > 0 {
		times = (times - cycleStart) % cycleLen
		for i := 0; i < times; i++ {
			runCycle(lines)
		}
	}
}

func runCycle(lines [][]rune) {
	tiltNorth(lines)
	tiltWest(lines)
	tiltSouth(lines)
	tiltEast(lines)
}

func tiltNorth(lines [][]rune) {
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
}

func tiltSouth(lines [][]rune) {
	for {
		anyMoved := false

		for row := len(lines) - 2; row >= 0; row-- {
			for col := 0; col < len(lines[row]); col++ {
				if lines[row+1][col] == '.' && lines[row][col] == 'O' {
					lines[row+1][col] = lines[row][col]
					lines[row][col] = '.'
					anyMoved = true
				}
			}
		}

		if !anyMoved {
			break
		}
	}
}

func tiltWest(lines [][]rune) {
	for {
		anyMoved := false

		for col := 1; col < len(lines[0]); col++ {
			for row := 0; row < len(lines); row++ {
				if lines[row][col-1] == '.' && lines[row][col] == 'O' {
					lines[row][col-1] = lines[row][col]
					lines[row][col] = '.'
					anyMoved = true
				}
			}
		}

		if !anyMoved {
			break
		}
	}
}

func tiltEast(lines [][]rune) {
	for {
		anyMoved := false

		for col := len(lines[0]) - 2; col >= 0; col-- {
			for row := 0; row < len(lines); row++ {
				if lines[row][col+1] == '.' && lines[row][col] == 'O' {
					lines[row][col+1] = lines[row][col]
					lines[row][col] = '.'
					anyMoved = true
				}
			}
		}

		if !anyMoved {
			break
		}
	}
}

func parse(input <-chan string) [][]rune {
	lines := make([][]rune, 0)
	for line := range input {
		lines = append(lines, []rune(line))
	}

	return lines
}
