package day12

import (
	"strconv"
	"strings"
	"sync"
)

type record struct {
	line   string
	groups []int
}

func SolvePart1(input <-chan string) int {
	var wg sync.WaitGroup
	ch := make(chan int, 200)

	for _, r := range parse(input) {
		wg.Add(1)
		go func(r record) {
			defer wg.Done()
			ch <- arrangements(r.line, r.groups)
		}(r)
	}

	go func() {
		wg.Wait()
		defer close(ch)
	}()

	sum := 0
	for v := range ch {
		sum += v
	}

	return sum
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func arrangements(line string, groups []int) int {
	missing := 0
	for _, v := range groups {
		missing += v
	}

	unknown := make([]int, 0)
	for i, r := range line {
		if r == '#' {
			missing--
		} else if r == '?' {
			unknown = append(unknown, i)
		}
	}

	return validVariations(line, groups, unknown, missing)
}

func validVariations(line string, groups []int, unknown []int, missing int) int {
	src := make([]rune, len(unknown))
	for i := 0; i < len(src); i++ {
		if i < missing {
			src[i] = '#'
		} else {
			src[i] = '.'
		}
	}

	variations := 0
	for perm := range uniqueCombinations(unknown, missing) {
		permutedLine := []rune(line)
		for _, idx := range perm {
			permutedLine[idx] = '#'
		}

		for _, idx := range unknown {
			if permutedLine[idx] == '?' {
				permutedLine[idx] = '.'
			}
		}

		l := string(permutedLine)
		if isValidVariation(l, groups) {
			variations++
		}
	}

	return variations
}

func uniqueCombinations[T any](s []T, n int) <-chan []T {
	ch := make(chan []T)
	go func() {
		defer close(ch)

		idx := make([]int, n)
		for i := 0; i < len(idx); i++ {
			idx[i] = i
		}

		for {
			p := make([]T, n)
			for i, v := range idx {
				p[i] = s[v]
			}

			ch <- p

			for i := len(idx) - 1; i >= 0; i-- {
				ref := len(s)
				if i+1 < len(idx) {
					ref = idx[i+1]
				}

				newV := idx[i] + 1
				if newV < ref {
					idx[i] = newV

					for j := i + 1; j < len(idx); j++ {
						newV++
						idx[j] = newV
					}
					break
				}

				if i == 0 {
					return
				}
			}
		}
	}()

	return ch
}

func incr(idx []int, pos int, ref int) {

}

func cartesianProduct[T any, R any](s []T, n int, fn func(p []T) R) <-chan R {
	ch := make(chan R, 100)
	go func() {
		defer close(ch)

		nextIndex := func(ix []int, l int) {
			for j := len(ix) - 1; j >= 0; j-- {
				ix[j]++
				if j == 0 || ix[j] < l {
					return
				}
				ix[j] = 0
			}
		}

		p := make([]T, n)
		for ix := make([]int, n); ix[0] < len(s); nextIndex(ix, len(s)) {
			for i, j := range ix {
				p[i] = s[j]
			}

			ch <- fn(p)
		}
	}()

	return ch
}

func isValidVariation(line string, groups []int) bool {
	groupIdx := 0
	springCount := 0

	for _, v := range line {
		if v == '#' {
			springCount++
		} else {
			if springCount > 0 {
				if springCount != groups[groupIdx] {
					return false
				}

				springCount = 0
				groupIdx++
			}
		}
	}

	if springCount == 0 {
		return groupIdx == len(groups)
	} else {
		return groupIdx+1 == len(groups) && springCount == groups[groupIdx]
	}
}

func parse(input <-chan string) []record {
	records := make([]record, 0)
	for line := range input {
		parts := strings.SplitN(line, " ", 2)
		r := record{
			line:   parts[0],
			groups: make([]int, 0),
		}

		for _, num := range strings.Split(parts[1], ",") {
			v, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			r.groups = append(r.groups, v)
		}

		records = append(records, r)
	}

	return records
}
