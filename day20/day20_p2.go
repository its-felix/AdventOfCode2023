package day20

import (
	"math"
	"strings"
)

const (
	noop = iota
	broadcast
	flipFlop
	conjunction
)

type module struct {
	name     string
	kind     int
	upstream []*module
	ps       [2]int
}

type iterationState struct {
	m     *module
	pulse int
}

func SolvePart2(input <-chan string) int {
	m := parseP2(input, "rx")
	return presses(m, low)
}

func presses(m *module, pulse int) int {
	q := newQueue[iterationState]()
	q.push(iterationState{
		m:     m,
		pulse: pulse,
	})

	for {
		curr, ok := q.pop()
		if !ok {
			break
		}

		if m.ps[pulse] != 0 {
			return m.ps[pulse]
		}

		pressesInternal(q, curr.m, curr.pulse)
	}

	panic("invalid state")
}

func pressesInternal(q *queue[iterationState], m *module, pulse int) {
	switch m.kind {
	case noop:
		noopPresses(q, m, pulse)

	case broadcast:
		broadcastPresses(q, m, pulse)

	case flipFlop:
		flipFlopPresses(q, m, pulse)

	case conjunction:
		conjunctionPresses(q, m, pulse)
	}
}

func noopPresses(q *queue[iterationState], m *module, pulse int) {
	numP := 0
	missing := make([]*module, 0)

	for _, v := range m.upstream {
		if v.ps[pulse] == 0 {
			missing = append(missing, v)
		} else {
			if numP == 0 {
				numP = v.ps[pulse]
			} else {
				numP = min(numP, v.ps[pulse])
			}

			if numP == 1 {
				missing = missing[:0]
				break
			}
		}
	}

	if len(missing) < 1 {
		m.ps[pulse] = numP
		println("noop done")
	} else {
		for _, v := range missing {
			q.push(iterationState{
				m:     v,
				pulse: pulse,
			})
		}

		q.push(iterationState{
			m:     m,
			pulse: pulse,
		})
	}
}

func broadcastPresses(q *queue[iterationState], m *module, pulse int) {
	m.ps[high] = math.MaxInt
	m.ps[low] = 1
}

func flipFlopPresses(q *queue[iterationState], m *module, pulse int) {
	numP := 0
	missing := make([]*module, 0)

	for _, v := range m.upstream {
		if v.ps[low] == 0 {
			missing = append(missing, v)
		} else {
			if numP == 0 {
				numP = v.ps[low]
			} else {
				numP = min(numP, v.ps[low])
			}

			if numP == 1 {
				missing = missing[:0]
				break
			}
		}
	}

	if len(missing) < 1 {
		m.ps[high] = numP
		m.ps[low] = numP + 1
		println("flipflop done")
	} else {
		for _, v := range missing {
			q.push(iterationState{
				m:     v,
				pulse: low,
			})
		}

		q.push(iterationState{
			m:     m,
			pulse: pulse,
		})
	}
}

func conjunctionPresses(q *queue[iterationState], m *module, pulse int) {
	switch pulse {
	case high:
		conjunctionPressesHigh(q, m)

	case low:
		conjunctionPressesLow(q, m)
	}
}

func conjunctionPressesLow(q *queue[iterationState], m *module) {
	numP := 0
	allFound := true

	for _, v := range m.upstream {
		if v.ps[high] == 0 {
			q.push(iterationState{
				m:     v,
				pulse: high,
			})
			allFound = false
		} else {
			if numP == 0 {
				numP = v.ps[high]
			} else {
				numP = lcm(numP, v.ps[high])
			}
		}
	}

	if allFound {
		m.ps[low] = numP
		println("conjunction low done")
	} else {
		q.push(iterationState{
			m:     m,
			pulse: low,
		})
	}
}

func conjunctionPressesHigh(q *queue[iterationState], m *module) {
	if len(m.upstream) == 1 {
		if m.upstream[0].ps[low] == 0 {
			q.push(iterationState{
				m:     m.upstream[0],
				pulse: low,
			})
			q.push(iterationState{
				m:     m,
				pulse: high,
			})
		} else {
			m.ps[high] = m.upstream[0].ps[low]
			println("conjunction high (==1) done")
		}
	} else {
		m.ps[high] = 1
		println("conjunction high (>1) done")
	}
}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}

	return gcd(y, x%y)
}

func lcm(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	return x * y / gcd(x, y)
}

func parseP2(input <-chan string, name string) *module {
	moduleByName := make(map[string]*module)

	for line := range input {
		var kind int
		var offset int

		if strings.HasPrefix(line, "%") {
			kind = flipFlop
			offset = 1
		} else if strings.HasPrefix(line, "&") {
			kind = conjunction
			offset = 1
		} else {
			kind = broadcast
			offset = 0
		}

		line = line[offset:]
		idx := strings.Index(line, " -> ")
		sender := line[:idx]
		m := getOrCreateP2(moduleByName, sender)
		m.kind = kind

		if kind == broadcast {
			m.ps[high] = math.MaxInt
			m.ps[low] = 1
		}

		downstream := strings.Split(line[idx+4:], ",")
		for _, receiver := range downstream {
			down := getOrCreateP2(moduleByName, strings.TrimSpace(receiver))
			down.upstream = append(down.upstream, m)
		}
	}

	return moduleByName[name]
}

func getOrCreateP2(moduleByName map[string]*module, name string) *module {
	m, ok := moduleByName[name]
	if ok {
		return m
	}

	m = &module{
		name:     name,
		kind:     noop,
		upstream: make([]*module, 0),
	}
	moduleByName[name] = m
	return m
}

type queue[T comparable] struct {
	head *node[T]
	tail *node[T]
}

func newQueue[T comparable]() *queue[T] {
	return &queue[T]{}
}

type node[T comparable] struct {
	v    T
	tail *node[T]
}

func (q *queue[T]) push(v T) {
	n := &node[T]{v: v}

	if q.tail == nil {
		q.head = n
		q.tail = n
	} else {
		q.tail.tail = n
		q.tail = n
	}
}

func (q *queue[T]) pop() (T, bool) {
	if q.head == nil {
		var def T
		return def, false
	}

	n := q.head
	q.head = n.tail

	if q.tail == n {
		q.tail = q.head
	}

	return n.v, true
}
