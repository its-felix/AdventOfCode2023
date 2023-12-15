package dayxx

func SolvePart1(input <-chan string) int {
	return parse(input)
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func parse(input <-chan string) int {
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
