package iter

import (
	"rseaon"
	"rseaon/option"
)

type Map[T any, F any] struct {
	iter Iterator[T]
	fn   func(*T) F
}

func NewMap[T any, F any](iter Iterator[T], fn func(*T) F) *Map[T, F] {
	return &Map[T, F]{
		iter: iter,
		fn:   fn,
	}
}

func (m *Map[T, F]) Next() *option.Option[F] {
	return option.Map[T, F](m.iter.Next(), m.fn)
}

type Enumerate[T any] struct {
	iter Iterator[T]
	cnt  int
}

func NewEnumerate[T any](iter Iterator[T]) *Enumerate[T] {
	return &Enumerate[T]{
		iter: iter,
		cnt:  0,
	}
}

func (e *Enumerate[T]) Next() *option.Option[rseaon.Tuple2[*T, int]] {
	peek := e.iter.Next()

	if peek.IsNil() {
		return option.Nil[rseaon.Tuple2[*T, int]]()
	}

	e.cnt++
	return option.NewOption[rseaon.Tuple2[*T, int]](rseaon.MakeTuple2(peek.UnwrapPtr(), e.cnt-1))
}

type Filter[T any] struct {
	iter Iterator[T]
	fn   func(*T) bool
}

func NewFilter[T any](iter Iterator[T], fn func(*T) bool) *Filter[T] {
	return &Filter[T]{
		iter: iter,
		fn:   fn,
	}
}

func (f *Filter[T]) Next() *option.Option[T] {
	return DefaultFindPtr[T](f.iter, f.fn)
}

type Take[T any] struct {
	iter Iterator[T]
	n    int
}

func NewTake[T any](iter Iterator[T]) *Take[T] {
	return &Take[T]{
		iter: iter,
		n:    0,
	}
}

func (t *Take[T]) Next() *option.Option[T] {
	if t.n > 0 {
		t.n--
		return t.iter.Next()
	}

	return option.Nil[T]()
}
