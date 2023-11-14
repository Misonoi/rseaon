package iter

import "github.com/Misonoi/rseaon/option"

type IntoIter[T any] interface {
	intoIter() Iterator[T]
}

type SliceIter[T any] struct {
	slice []T
	idx   int
}

func NewSliceIter[T any](slice []T) *SliceIter[T] {
	return &SliceIter[T]{
		slice: slice,
		idx:   0,
	}
}

func (s *SliceIter[T]) Next() *option.Option[T] {
	if s.idx == len(s.slice) {
		s.idx = 0
		return option.Nil[T]()
	}

	s.idx += 1
	return option.NewOption(&s.slice[s.idx-1])
}

type SliceWrapper[T any] struct {
	slice []T
}

func (s *SliceWrapper[T]) Unwrap() []T {
	return s.slice
}

func NewSliceWrapper[T any]() *SliceWrapper[T] {
	return &SliceWrapper[T]{
		slice: make([]T, 0),
	}
}

func (s *SliceWrapper[T]) fromIter(iter Iterator[T]) *SliceWrapper[T] {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			break
		}

		s.slice = append(s.slice, iter.Next().Unwrap())
	}

	return s
}
