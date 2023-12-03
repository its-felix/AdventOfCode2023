package day2

import (
	"strconv"
	"strings"
)

func SolvePart1(input <-chan string, maxRed, maxGreen, maxBlue int) {
	sumPossible := 0

	for line := range input {
		gameNum, values := Parse(line)
		if IsPossible(values, maxRed, maxGreen, maxBlue) {
			sumPossible += gameNum
		}
	}

	println(sumPossible)
}

func SolvePart2(input <-chan string) {
	sum := 0

	for line := range input {
		_, values := Parse(line)
		least := LeastPossible(values)
		sum += least[0] * least[1] * least[2]
	}

	println(sum)
}

func IsPossible(values [][3]int, maxRed, maxGreen, maxBlue int) bool {
	for _, set := range values {
		if set[0] > maxRed || set[1] > maxGreen || set[2] > maxBlue {
			return false
		}
	}

	return true
}

func LeastPossible(values [][3]int) [3]int {
	least := [3]int{0, 0, 0}
	for _, set := range values {
		least[0] = max(least[0], set[0])
		least[1] = max(least[1], set[1])
		least[2] = max(least[2], set[2])
	}

	return least
}

func Parse(line string) (int, [][3]int) {
	gameAndSets := strings.SplitN(line, ":", 2)
	if len(gameAndSets) != 2 {
		panic("couldnt split game and sets")
	}

	gameNum, err := strconv.Atoi(strings.TrimPrefix(gameAndSets[0], "Game "))
	if err != nil {
		panic(err)
	}

	sets := strings.Split(gameAndSets[1], ";")
	values := make([][3]int, 0, len(sets))

	for _, set := range sets {
		set := strings.Split(set, ",")
		rgb := [3]int{0, 0, 0}

		for _, v := range set {
			v = strings.TrimSpace(v)
			idx := 0

			if num := strings.TrimSuffix(v, " red"); num != v {
				idx = 0
				v = num
			} else if num := strings.TrimSuffix(v, " green"); num != v {
				idx = 1
				v = num
			} else if num := strings.TrimSuffix(v, " blue"); num != v {
				idx = 2
				v = num
			}

			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			rgb[idx] = num
		}

		values = append(values, rgb)
	}

	return gameNum, values
}
