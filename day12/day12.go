package day12

import (
	"strconv"
	"strings"
	"sync"
)

type record struct {
	uGroups        []*group // unique groups
	groups         []*group // groups by index (any r.groups[g.start..g.end] == g)
	expectedGroups []int    // expected spring groups
}

type group struct {
	r     rune
	start int
	end   int
}

func (g *group) len() int {
	return g.end - g.start + 1
}

func SolvePart1(input <-chan string) int {
	return solve(parse(input))
}

func SolvePart2(input <-chan string) int {
	return solve(parse(repeat(input, 5)))
}

func solve(records []record) int {
	var wg sync.WaitGroup
	ch := make(chan int, 200)

	for _, r := range records {
		wg.Add(1)
		go func(r record) {
			defer wg.Done()
			ch <- arrangements(r)
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

func arrangements(r record) int {
	missing := 0
	for _, v := range r.expectedGroups {
		missing += v
	}

	gCount := 0
	for _, g := range r.uGroups {
		if g.r == '#' {
			missing -= g.len()
			gCount++
		}
	}

	return missing
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
			uGroups:        make([]*group, 0),
			groups:         make([]*group, 0),
			expectedGroups: make([]int, 0),
		}

		var g *group
		for i, c := range parts[0] {
			if g == nil || g.r != c {
				if g != nil {
					r.uGroups = append(r.uGroups, g)
				}

				g = &group{
					r:     c,
					start: i,
				}
			}

			g.end = i
			r.groups = append(r.groups, g)
		}

		r.uGroups = append(r.uGroups, g)

		for _, num := range strings.Split(parts[1], ",") {
			v, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			r.expectedGroups = append(r.expectedGroups, v)
		}

		records = append(records, r)
	}

	return records
}

func repeat(input <-chan string, times int) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		for line := range input {
			parts := strings.SplitN(line, " ", 2)
			seps := []string{"?", ","}

			for i, part := range parts {
				copies := make([]string, times)
				for j := 0; j < times; j++ {
					copies[j] = part
				}

				parts[i] = strings.Join(copies, seps[i])
			}

			ch <- strings.Join(parts, " ")
		}
	}()

	return ch
}
