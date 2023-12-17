package util

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddIfAbsent(v T) bool {
	if _, present := s[v]; !present {
		s[v] = struct{}{}
		return true
	}

	return false
}
