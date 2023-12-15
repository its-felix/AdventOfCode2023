package day15

import (
	"slices"
	"strconv"
	"strings"
)

type step struct {
	label    string
	op       rune
	modifier int
}

type lens struct {
	label string
	focal int
}

func SolvePart1(input <-chan string) int {
	return sumHash(input)
}

func SolvePart2(input <-chan string) int {
	boxes := [256][]lens{}
	for _, s := range parse(input) {
		hash := genHash(s.label)
		if boxes[hash] == nil {
			boxes[hash] = make([]lens, 0)
		}

		idx := slices.IndexFunc(boxes[hash], func(l lens) bool {
			return l.label == s.label
		})

		if s.op == '=' {
			if idx != -1 {
				boxes[hash][idx].focal = s.modifier
			} else {
				boxes[hash] = append(boxes[hash], lens{
					label: s.label,
					focal: s.modifier,
				})
			}
		} else if s.op == '-' {
			if idx != -1 {
				boxes[hash] = slices.Delete(boxes[hash], idx, idx+1)
			}
		} else {
			panic("invalid op")
		}
	}

	sum := 0
	for boxNum, box := range boxes {
		if box == nil {
			continue
		}

		for lensPos, l := range box {
			power := boxNum + 1
			power *= lensPos + 1
			power *= l.focal
			sum += power
		}
	}

	return sum
}

func parse(input <-chan string) []step {
	steps := make([]step, 0)
	for line := range input {
		var s step
		var ok bool
		for {
			line, s, ok = parseStep(line)
			if !ok {
				break
			}

			steps = append(steps, s)
		}
	}

	return steps
}

func parseStep(line string) (string, step, bool) {
	part := ""
	for _, r := range line {
		if r == ',' {
			break
		}

		part += string(r)
	}

	if len(part) < 1 {
		return line, step{}, false
	}

	if part == line {
		line = ""
	} else {
		line = line[len(part)+1:]
	}

	s := step{}
	if part[len(part)-1] == '-' {
		s.label = part[:len(part)-1]
		s.op = '-'
		s.modifier = 0
	} else {
		idx := strings.LastIndex(part, "=")
		if idx == -1 {
			panic("invalid op")
		}

		var err error
		s.label = part[:idx]
		s.op = '='
		s.modifier, err = strconv.Atoi(part[idx+1:])

		if err != nil {
			panic(err)
		}
	}

	return line, s, true
}

func sumHash(input <-chan string) int {
	sum := 0
	for line := range input {
		var hash int
		var ok bool
		for {
			line, hash, ok = nextHash(line)
			if !ok {
				break
			}

			sum += hash
		}
	}

	return sum
}

func nextHash(s string) (string, int, bool) {
	hash := 0
	maxI := -1

	for i, r := range s {
		maxI = i
		if r == ',' {
			break
		}

		hash += int(r)
		hash *= 17
		hash %= 256
	}

	return s[maxI+1:], hash, maxI >= 0
}

func genHash(s string) int {
	hash := 0
	for _, r := range s {
		hash += int(r)
		hash *= 17
		hash %= 256
	}

	return hash
}
