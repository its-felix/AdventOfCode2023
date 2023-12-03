package day3

import (
	"fmt"
	"unicode"
)

type Token struct {
	Num      int
	Symbol   rune
	IsNum    bool
	IsSymbol bool
}

func SolvePart1(input <-chan string) {
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

	println(sum)
}

func SolvePart2(input <-chan string) {
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

	println(sum)
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
	values := make([]*Token, 0)

	for _, tk := range FindAdjacent(matrix, line, col) {
		if tk.IsNum {
			if _, ok := seen[tk]; !ok {
				seen[tk] = true
				values = append(values, tk)
			}
		}
	}

	return values
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

	res := make([]*Token, 0)

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
	matrix := make([][]*Token, 0)

	for line := range input {
		cols := make([]*Token, 0)
		curr := &Token{}

		for _, r := range line {
			if unicode.IsDigit(r) {
				if int(r) < 48 || int(r) > 57 {
					panic(fmt.Sprintf("IsDigit but out of bounds: %v %d", r, int(r)))
				}

				v := int(r) - 48

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
