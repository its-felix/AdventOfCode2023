package day18

import (
	"encoding/binary"
	"encoding/hex"
	"strconv"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

type instruction struct {
	direction int
	steps     int
	rgb       string
}

func SolvePart1(input <-chan string) int {
	return solve(parse(input))
}

func SolvePart2(input <-chan string) int {
	instructions := parse(input)
	fixInstructions(instructions)
	return solve(instructions)
}

func fixInstructions(instructions []instruction) {
	for i, ins := range instructions {
		stepsRaw, dirRaw := ins.rgb[:5], ins.rgb[5:]

		b, err := hex.DecodeString("0" + stepsRaw)
		if err != nil {
			panic(err)
		}

		ins.steps = int(binary.BigEndian.Uint64(append(make([]byte, 5), b...)))

		switch dirRaw {
		case "0":
			ins.direction = east
		case "1":
			ins.direction = south
		case "2":
			ins.direction = west
		case "3":
			ins.direction = north
		default:
			panic("invalid direction")
		}

		instructions[i] = ins
	}
}

func solve(instructions []instruction) int {
	acc, d := 0, 0
	count := 0

	for _, ins := range instructions {
		switch ins.direction {
		case north:
			d += ins.steps
		case east:
			acc += d * ins.steps
		case south:
			d -= ins.steps
		case west:
			acc -= d * ins.steps
		}

		count += ins.steps
	}

	if acc < 0 {
		acc = -acc
	}

	return acc - (count / 2) + 1 + count
}

func parse(input <-chan string) []instruction {
	instructions := make([]instruction, 0)
	for line := range input {
		ins := instruction{}

		parts := strings.SplitN(line, " ", 3)
		dirRaw, stepsRaw, rgbRaw := parts[0], parts[1], parts[2]

		switch dirRaw {
		case "U":
			ins.direction = north
		case "D":
			ins.direction = south
		case "L":
			ins.direction = west
		case "R":
			ins.direction = east
		default:
			panic("invalid direction")
		}

		var err error
		ins.steps, err = strconv.Atoi(stepsRaw)
		if err != nil {
			panic(err)
		}

		ins.rgb = rgbRaw[2 : len(rgbRaw)-1]

		instructions = append(instructions, ins)
	}

	return instructions
}
