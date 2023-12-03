package day3

type Token struct {
	Num      int
	Symbol   rune
	IsNum    bool
	IsSymbol bool
}

func SolvePart1(input <-chan string) int {
	matrix := Parse(input)
	sum := 0

	for lineNum, cols := range matrix {
		var prevCounted *Token
		for colNum, col := range cols {
			if col != prevCounted && col.IsNum && IsAdjacentToSymbol(matrix, lineNum, colNum) {
				prevCounted = col
				sum += col.Num
			}
		}
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	matrix := Parse(input)
	sum := 0

	for lineNum, cols := range matrix {
		for colNum, col := range cols {
			if col.IsSymbol && col.Symbol == '*' {
				adjacentParts := FindAdjacentParts(matrix, lineNum, colNum)
				if len(adjacentParts) == 2 {
					sum += adjacentParts[0].Num * adjacentParts[1].Num
				}
			}
		}
	}

	return sum
}

func IsAdjacentToSymbol(matrix [][]*Token, line, col int) bool {
	for _, tk := range FindAdjacent(matrix, line, col) {
		if tk.IsSymbol {
			return true
		}
	}

	return false
}

func FindAdjacentParts(matrix [][]*Token, line, col int) []*Token {
	seen := make(map[*Token]bool)
	res := make([]*Token, 0, 8)

	for _, tk := range FindAdjacent(matrix, line, col) {
		if tk.IsNum {
			if _, ok := seen[tk]; !ok {
				seen[tk] = true
				res = append(res, tk)
			}
		}
	}

	return res
}

func FindAdjacent(matrix [][]*Token, line, col int) []*Token {
	positions := [][2]int{
		{line - 1, col - 1},
		{line - 1, col},
		{line - 1, col + 1},

		{line, col - 1},
		{line, col + 1},

		{line + 1, col - 1},
		{line + 1, col},
		{line + 1, col + 1},
	}

	res := make([]*Token, 0, len(positions))

	for _, check := range positions {
		l, c := check[0], check[1]
		if l >= 0 && l < len(matrix) {
			values := matrix[l]
			if c >= 0 && c < len(values) {
				res = append(res, values[c])
			}
		}
	}

	return res
}

func Parse(input <-chan string) [][]*Token {
	matrix := make([][]*Token, 0, 140)

	for line := range input {
		cols := make([]*Token, 0, len(line))
		curr := &Token{}

		for _, r := range line {
			v := int(r) - '0'
			if v >= 0 && v < 10 {
				if curr.IsNum {
					curr.Num *= 10
					curr.Num += v
				} else {
					curr = &Token{
						Num:   v,
						IsNum: true,
					}
				}
			} else if r == '.' {
				curr = &Token{}
			} else {
				curr = &Token{
					Symbol:   r,
					IsSymbol: true,
				}
			}

			cols = append(cols, curr)
		}

		matrix = append(matrix, cols)
	}

	return matrix
}
