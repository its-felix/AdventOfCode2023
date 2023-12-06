package util

import (
	"strings"
	"unicode"
)

func ReadInts(s string) (string, []int) {
	res := make([]int, 0)
	for {
		var num int
		var anyMatched bool
		s, num, anyMatched = ReadInt(strings.TrimSpace(s))

		if anyMatched {
			res = append(res, num)
		} else {
			break
		}
	}

	return s, res
}

func ReadInt(s string) (string, int, bool) {
	maxI := -1
	num := 0
	anyMatched := false

	for i, r := range s {
		v := int(r) - '0'
		if v >= 0 && v < 10 {
			num *= 10
			num += v
			maxI = i
			anyMatched = true
		} else {
			break
		}
	}

	return s[maxI+1:], num, anyMatched
}

func TrimNonDigit(s string) string {
	return strings.TrimLeftFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
}
