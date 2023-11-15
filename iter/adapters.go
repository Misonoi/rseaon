package iter

import (
	"github.com/Misonoi/rseaon"
	"github.com/Misonoi/rseaon/option"
)

type iMap[T any, F any] struct {
	iter Iterator[T]
	fn   func(*T) F
}

func newMap[T any, F any](iter Iterator[T], fn func(*T) F) *iMap[T, F] {
	return &iMap[T, F]{
		iter: iter,
		fn:   fn,
	}
}

func (m *iMap[T, F]) Next() *option.Option[F] {
	return option.Map[T, F](m.iter.Next(), m.fn)
}

type enumerate[T any] struct {
	iter Iterator[T]
	cnt  int
}

func newEnumerate[T any](iter Iterator[T]) *enumerate[T] {
	return &enumerate[T]{
		iter: iter,
		cnt:  0,
	}
}

func (e *enumerate[T]) Next() *option.Option[rseaon.Tuple2[*T, int]] {
	peek := e.iter.Next()

	if peek.IsNil() {
		return option.Nil[rseaon.Tuple2[*T, int]]()
	}

	e.cnt++
	return option.NewOption[rseaon.Tuple2[*T, int]](rseaon.MakeTuple2(peek.UnwrapPtr(), e.cnt-1))
}

type filter[T any] struct {
	iter Iterator[T]
	fn   func(*T) bool
}

func newFilter[T any](iter Iterator[T], fn func(*T) bool) *filter[T] {
	return &filter[T]{
		iter: iter,
		fn:   fn,
	}
}

func (f *filter[T]) Next() *option.Option[T] {
	return FindPtr[T](f.iter, f.fn)
}

type take[T any] struct {
	iter Iterator[T]
	n    int
}

func newTake[T any](iter Iterator[T], n int) *take[T] {
	return &take[T]{
		iter: iter,
		n:    n,
	}
}

func (t *take[T]) Next() *option.Option[T] {
	if t.n > 0 {
		t.n--
		return t.iter.Next()
	}

	return option.Nil[T]()
}

type zip[T, S any] struct {
	iter1 Iterator[T]
	iter2 Iterator[S]
}

func newZip[T, S any](iter1 Iterator[T], iter2 Iterator[S]) *zip[T, S] {
	return &zip[T, S]{
		iter1: iter1,
		iter2: iter2,
	}
}

func (z *zip[T, S]) Next() *option.Option[rseaon.Tuple2[*T, *S]] {
	peek1, peek2 := z.iter1.Next(), z.iter2.Next()

	if peek1.IsNil() || peek2.IsNil() {
		return option.Nil[rseaon.Tuple2[*T, *S]]()
	}

	return option.NewOption(rseaon.MakeTuple2(peek1.UnwrapPtr(), peek2.UnwrapPtr()))
}
