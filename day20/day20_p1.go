package day20

import "strings"

type signal struct {
	sender string
	pulse  int
}

type handler interface {
	handle(sig signal) pulseFn
}

type pulseFn func() []pulseFn

func (p pulseFn) invoke() {
	queue := p()
	for len(queue) > 0 {
		fn := queue[0]
		queue = queue[1:]
		queue = append(queue, fn()...)
	}
}

func makePulseFn(sig signal, downstream []handler) pulseFn {
	return func() []pulseFn {
		fns := make([]pulseFn, len(downstream))
		for i, downstreamH := range downstream {
			fns[i] = downstreamH.handle(sig)
		}

		return fns
	}
}

type broadcastHandler struct {
	name       string
	downstream []handler
}

func (h *broadcastHandler) handle(sig signal) pulseFn {
	return makePulseFn(signal{
		sender: h.name,
		pulse:  sig.pulse,
	}, h.downstream)
}

type flipFlopHandler struct {
	name       string
	off        bool
	downstream []handler
}

func (h *flipFlopHandler) handle(sig signal) pulseFn {
	if sig.pulse == high {
		return func() []pulseFn {
			return make([]pulseFn, 0)
		}
	}

	h.off = !h.off
	pulse := high
	if h.off {
		pulse = low
	}

	return makePulseFn(signal{
		sender: h.name,
		pulse:  pulse,
	}, h.downstream)
}

type conjunctionHandler struct {
	name       string
	state      map[string]int
	downstream []handler
}

func (h *conjunctionHandler) handle(sig signal) pulseFn {
	h.state[sig.sender] = sig.pulse

	pulse := low
	for _, v := range h.state {
		if v != high {
			pulse = high
			break
		}
	}

	return makePulseFn(signal{
		sender: h.name,
		pulse:  pulse,
	}, h.downstream)
}

type noopHandler string

func (h noopHandler) handle(sig signal) pulseFn {
	return func() []pulseFn {
		return make([]pulseFn, 0)
	}
}

type interceptHandler struct {
	counter *counter
	forward handler
}

func (h interceptHandler) handle(sig signal) pulseFn {
	if sig.pulse == high {
		h.counter.countHigh++
	} else {
		h.counter.countLow++
	}

	return h.forward.handle(sig)
}

type counter struct {
	countHigh uint64
	countLow  uint64
}

type lateInitHandler struct {
	h          handler
	downstream []*lateInitHandler
}

func SolvePart1(input <-chan string) uint64 {
	c := &counter{}
	broadcaster := parseP1(input, func(name string, h handler) handler {
		return interceptHandler{
			counter: c,
			forward: h,
		}
	})

	sig := signal{
		sender: "button",
		pulse:  low,
	}

	for i := 0; i < 1000; i++ {
		broadcaster.handle(sig).invoke()
	}

	return c.countHigh * c.countLow
}

func parseP1(input <-chan string, decorateFn func(name string, h handler) handler) handler {
	const flipFlop, conjunction, broadcast = 0, 1, 2

	handlerByName := make(map[string]*lateInitHandler)
	upstreamsByName := make(map[string][]string)

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

		downstreamNames := strings.Split(line[idx+4:], ",")
		downstream := make([]*lateInitHandler, len(downstreamNames))

		for i, receiver := range downstreamNames {
			receiver = strings.TrimSpace(receiver)
			downstreamNames[i] = receiver
			downstream[i] = getOrCreateP1(handlerByName, receiver, decorateFn)

			up, ok := upstreamsByName[receiver]
			if !ok {
				up = make([]string, 0)
			}

			upstreamsByName[receiver] = append(up, sender)
		}

		lih := getOrCreateP1(handlerByName, sender, decorateFn)
		lih.downstream = downstream

		switch kind {
		case flipFlop:
			lih.h = &flipFlopHandler{
				name: sender,
				off:  true,
			}

		case conjunction:
			lih.h = &conjunctionHandler{
				name: sender,
			}

		case broadcast:
			lih.h = &broadcastHandler{
				name: sender,
			}

		default:
			panic("invalid kind")
		}

		lih.h = decorateFn(sender, lih.h)
	}

	for _, lih := range handlerByName {
		downstream := make([]handler, len(lih.downstream))
		for i, h := range lih.downstream {
			downstream[i] = h.h
		}

		h := lih.h
		for v, ok := h.(interceptHandler); ok; v, ok = h.(interceptHandler) {
			h = v.forward
		}

		switch h := h.(type) {
		case *flipFlopHandler:
			h.downstream = downstream

		case *conjunctionHandler:
			initState := make(map[string]int)
			if up, ok := upstreamsByName[h.name]; ok {
				for _, upV := range up {
					initState[upV] = low
				}
			}

			h.downstream = downstream
			h.state = initState

		case *broadcastHandler:
			h.downstream = downstream

		case noopHandler:
			// no downstream

		default:
			panic("invalid handler type")
		}
	}

	return handlerByName["broadcaster"].h
}

func getOrCreateP1(handlerByName map[string]*lateInitHandler, name string, decorateFn func(name string, h handler) handler) *lateInitHandler {
	h, ok := handlerByName[name]
	if ok {
		return h
	}

	h = &lateInitHandler{
		h: decorateFn(name, noopHandler(name)),
	}
	handlerByName[name] = h
	return h
}
