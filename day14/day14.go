package day14

const (
	north = uint8(iota)
	east
	south
	west
)

func SolvePart1(input <-chan string) int {
	g := parse(input)
	g.tilt(north)
	return solve(g)
}

func SolvePart2(input <-chan string) int {
	g := parse(input)
	runCycles(g, 1000000000)
	return solve(g)
}

func solve(g grid) int {
	sum := 0

	for i, line := range g {
		for _, r := range line {
			if r == 'O' {
				sum += len(g) - i
			}
		}
	}

	return sum
}

func buildString(g grid) string {
	s := ""
	for _, line := range g {
		s += string(line)
		s += "\n"
	}

	return s[:len(s)-1]
}

func runCycles(g grid, times int) {
	var cycleStart, cycleLen int
	cache := make(map[string]int)

	for i := 1; i <= times && cycleLen == 0; i++ {
		runCycle(g)
		s := buildString(g)
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
			runCycle(g)
		}
	}
}

func runCycle(g grid) {
	g.tilt(north)
	g.tilt(west)
	g.tilt(south)
	g.tilt(east)
}

type grid [][]rune

func (g grid) tilt(tilt uint8) {
	var rowPrev, colPrev, row, col int
	var ok bool

	for {
		anyMoved := false
		iter := g.iter(tilt)
		for {
			rowPrev, colPrev, row, col, ok = iter()
			if !ok {
				break
			}

			if g[rowPrev][colPrev] == '.' && g[row][col] == 'O' {
				g[rowPrev][colPrev] = 'O'
				g[row][col] = '.'
				anyMoved = true
			}
		}

		if !anyMoved {
			break
		}
	}
}

func (g grid) iter(tilt uint8) func() (int, int, int, int, bool) {
	var rIdx, cIdx int
	var prevOffset int
	var rStart, rEnd, rIncr int
	var cStart, cEnd, cIncr int

	switch tilt {
	case north:
		rIdx, cIdx = 0, 1
		prevOffset = -1
		rStart, rEnd, rIncr = 1, len(g), 1
		cStart, cEnd, cIncr = 0, len(g[0]), 1
	case south:
		rIdx, cIdx = 0, 1
		prevOffset = 1
		rStart, rEnd, rIncr = len(g)-2, -1, -1
		cStart, cEnd, cIncr = 0, len(g[0]), 1
	case east:
		rIdx, cIdx = 1, 0
		prevOffset = 1
		rStart, rEnd, rIncr = len(g[0])-2, -1, -1
		cStart, cEnd, cIncr = 0, len(g), 1
	case west:
		rIdx, cIdx = 1, 0
		prevOffset = -1
		rStart, rEnd, rIncr = 1, len(g[0]), 1
		cStart, cEnd, cIncr = 0, len(g), 1
	default:
		panic("invalid tilt")
	}

	var idxs [2][2]int
	row, col := rStart, cStart
	exhausted := false

	return func() (int, int, int, int, bool) {
		if exhausted {
			return 0, 0, 0, 0, false
		}

		idxs[0][rIdx] = row + prevOffset
		idxs[0][cIdx] = col
		idxs[1][rIdx] = row
		idxs[1][cIdx] = col

		col += cIncr
		if col == cEnd {
			col = cStart
			row += rIncr
			if row == rEnd {
				exhausted = true
			}
		}

		return idxs[0][0], idxs[0][1], idxs[1][0], idxs[1][1], true
	}
}

func parse(input <-chan string) grid {
	lines := make([][]rune, 0)
	for line := range input {
		lines = append(lines, []rune(line))
	}

	return lines
}
