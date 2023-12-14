package day13

type pattern struct {
	rows    []string
	columns []string
}

func SolvePart1(input <-chan string) int {
	return solve(parse(input), false)
}

func SolvePart2(input <-chan string) int {
	return solve(parse(input), true)
}

func solve(patterns []pattern, smudges bool) int {
	sum := 0
	for _, p := range patterns {
		rowReflections := maxReflectionsLeft(p.rows, smudges)
		colReflections := maxReflectionsLeft(p.columns, smudges)

		if rowReflections > colReflections {
			sum += rowReflections * 100
		} else {
			sum += colReflections
		}
	}

	return sum
}

func maxReflectionsLeft(s []string, smudges bool) int {
	a, b := 0, 0

	// fold right onto left
	for i := len(s) - 2; i >= 0; i-- {
		cmp := s[i:] // cut off left side
		l := len(cmp)

		if l%2 == 0 {
			if isReflected(cmp, smudges) {
				a = len(s) - (l / 2)
				break
			}
		} else {
			// if the portion to compare is not even and its the first iteration (=nothing was cut off left yet)
			// we can ignore the last element of the right side for the comparison
			if i == 0 {
				cmp = cmp[:l-1]
				if isReflected(cmp, smudges) {
					a = len(s) - ((l + 1) / 2)
					break
				}
			}
		}
	}

	// fold left onto right
	for i := len(s); i >= 2; i-- {
		cmp := s[:i] // cut off right side
		l := len(cmp)

		if l%2 == 0 {
			if isReflected(cmp, smudges) {
				b = l / 2
				break
			}
		} else {
			// if the portion to compare is not even and its the first iteration (=nothing was cut off right yet)
			// we can ignore the last element of the left side for the comparison
			if i == len(s) {
				cmp = cmp[1:]
				if isReflected(cmp, smudges) {
					b = (l + 1) / 2
					break
				}
			}
		}
	}

	return max(a, b)
}

func isReflected(s []string, smudges bool) bool {
	ch := make(chan []string, 1)
	if smudges {
		go func(s []string) {
			defer close(ch)

			rowI, colI := 0, 0
			for {
				dst := append(make([]string, 0, len(s)), s...)
				row := []rune(dst[rowI])
				if row[colI] == '.' {
					row[colI] = '#'
				} else {
					row[colI] = '.'
				}

				dst[rowI] = string(row)
				ch <- dst

				colI++
				if colI >= len(row) {
					colI = 0
					rowI++

					if rowI >= len(dst) {
						break
					}
				}
			}
		}(s)
	} else {
		ch <- s
		close(ch)
	}

	for s := range ch {
		match := true

		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-i-1] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

func parse(input <-chan string) []pattern {
	patterns := make([]pattern, 0)

	p := pattern{
		rows:    make([]string, 0),
		columns: make([]string, 0),
	}

	for line := range input {
		if line == "" {
			if len(p.rows) > 0 {
				patterns = append(patterns, p)
			}

			p = pattern{
				rows:    make([]string, 0),
				columns: make([]string, 0),
			}
			continue
		}

		p.rows = append(p.rows, line)

		for col, r := range line {
			p.columns = grow(p.columns, col+1)
			p.columns[col] = p.columns[col] + string(r)
		}
	}

	if len(p.rows) > 0 {
		patterns = append(patterns, p)
	}

	return patterns
}

func grow[T any](s []T, size int) []T {
	missing := size - len(s)
	if missing > 0 {
		s = append(s, make([]T, missing)...)
	}

	return s
}
