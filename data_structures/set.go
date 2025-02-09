package data_structures

import "iter"

type Set[T comparable] struct {
	values map[T]bool
}

func CreateSet[T comparable]() Set[T] {
	return Set[T]{
		values: make(map[T]bool),
	}
}

func (s *Set[T]) Add(value T) bool {
	_, found := s.values[value]
	s.values[value] = true
	return !found
}

func (s *Set[T]) Remove(value T) bool {
	_, found := s.values[value]
	delete(s.values, value)
	return found
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.values[value]
	return found
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k, v := range s.values {
			if v {
				if !yield(k) {
					return
				}
			}
		}
	}
}
