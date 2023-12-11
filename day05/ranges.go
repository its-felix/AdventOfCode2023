package day05

type ranges[T any] struct {
	entries []rangeEntry[T]
}

type rangeEntry[T any] struct {
	start  int
	end    int
	values []T
}

func (r *ranges[T]) add(start, end int, value T) {
	for i := 0; i < len(r.entries); i++ {
		if r.entries[i].start == start && r.entries[i].end == end {
			r.entries[i].values = append(r.entries[i].values, value)
			return
		}
	}

	r.entries = append(r.entries, rangeEntry[T]{
		start:  start,
		end:    end,
		values: []T{value},
	})
}

func (r *ranges[T]) find(idx int) []T {
	values := make([]T, 0)
	for _, e := range r.entries {
		if e.start <= idx && e.end >= idx {
			values = append(values, e.values...)
		}
	}

	return values
}

func (r *ranges[T]) subrange(start, end int, mark T) (int, int, bool) {
	for _, e := range r.entries {
		if e.start <= start && e.end >= start {
			start = e.end + 1
		}

		if e.start <= end && e.end >= end {
			end = e.start - 1
		}
	}

	if start > end {
		return 0, 0, false
	}

	r.add(start, end, mark)
	return start, end, true
}
