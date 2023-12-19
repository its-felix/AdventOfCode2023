package day19

import "strings"

var (
	gt = operator(func(lhs int, rhs int) bool {
		return lhs > rhs
	})
	lt = operator(func(lhs, rhs int) bool {
		return lhs < rhs
	})
	accept = new(workflow)
	reject = new(workflow)
)

type part map[rune]int

type workflow struct {
	name      string
	rules     []rule
	otherwise *workflow
}

type rule struct {
	category rune
	op       operator
	value    int
	next     *workflow
}

type operator func(lhs, rhs int) bool

func SolvePart1(input <-chan string) int {
	return solve(parse(input))
}

func SolvePart2(input <-chan string) int {
	parse(input)
	return 0
}

func solve(wf *workflow, parts []part) int {
	sum := 0
	for _, p := range parts {
		if eval(wf, p) {
			sum += sumPart(p)
		}
	}

	return sum
}

func eval(wf *workflow, p part) bool {
	if wf == accept {
		return true
	}

	if wf == reject {
		return false
	}

	for _, r := range wf.rules {
		v := p[r.category]
		if r.op(v, r.value) {
			return eval(r.next, p)
		}
	}

	return eval(wf.otherwise, p)
}

func sumPart(p part) int {
	sum := 0
	for _, v := range p {
		sum += v
	}

	return sum
}

func parse(input <-chan string) (*workflow, []part) {
	workflowByName := make(map[string]*workflow)
	workflowByName["A"] = accept
	workflowByName["R"] = reject

	for line := range input {
		if line == "" {
			break
		}

		parseWorkflow(workflowByName, line)
	}

	parts := make([]part, 0)
	for line := range input {
		parts = append(parts, parsePart(line))
	}

	return workflowByName["in"], parts
}

func parseWorkflow(workflowByName map[string]*workflow, line string) *workflow {
	idx1, idx2 := strings.IndexRune(line, '{'), strings.IndexRune(line, '}')
	if idx1 == -1 || idx2 == -1 {
		panic("invalid workflow")
	}

	wf := getOrCreateWorkflow(workflowByName, line[:idx1])

	parts := strings.Split(line[idx1+1:idx2], ",")
	for i := 0; i < len(parts)-1; i++ {
		wf.rules = append(wf.rules, parseRule(workflowByName, parts[i]))
	}

	wf.otherwise = getOrCreateWorkflow(workflowByName, parts[len(parts)-1])

	return wf
}

func parseRule(workflowByName map[string]*workflow, line string) rule {
	r := rule{}

	state := 0
	nextWorkflowName := ""
	for _, c := range line {
		if state == 0 {
			r.category = c
			state = 1
		} else if state == 1 {
			switch c {
			case '>':
				r.op = gt
			case '<':
				r.op = lt
			default:
				panic("invalid operator")
			}
			state = 2
		} else if state == 2 {
			if c == ':' {
				state = 3
				continue
			}

			v := int(c) - '0'
			if v < 0 || v > 9 {
				panic("invalid value")
			}

			r.value *= 10
			r.value += v
		} else if state == 3 {
			nextWorkflowName += string(c)
		}
	}

	if state != 3 {
		panic("invalid state")
	}

	r.next = getOrCreateWorkflow(workflowByName, nextWorkflowName)

	return r
}

func parsePart(line string) part {
	line = line[1 : len(line)-1]
	p := make(part)

	var category rune
	var value int

	state := 0
	for _, c := range line {
		if state == 0 {
			category = c
			state = 1
		} else if state == 1 {
			if c == '=' {
				state = 2
			} else {
				panic("expected '='")
			}
		} else if state == 2 {
			if c == ',' {
				p[category] = value

				value = 0
				state = 0
				continue
			}

			v := int(c) - '0'
			if v < 0 || v > 9 {
				panic("invalid value")
			}

			value *= 10
			value += v
		}
	}

	if state != 2 {
		panic("invalid state")
	}

	p[category] = value

	return p
}

func getOrCreateWorkflow(workflowByName map[string]*workflow, name string) *workflow {
	if wf, ok := workflowByName[name]; ok {
		return wf
	}

	wf := &workflow{
		name:  name,
		rules: make([]rule, 0),
	}
	workflowByName[name] = wf

	return wf
}
