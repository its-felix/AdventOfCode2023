package util

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) AddIfAbsent(v T) bool {
	_, present := s[v]
	s[v] = struct{}{}
	return !present
}
