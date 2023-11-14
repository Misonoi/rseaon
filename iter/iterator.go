package iter

import (
	"rseaon/option"
)

type Iterator[T any] interface {
	Next() *option.Option[T]
}

type DoubleEndedIterator[T any] interface {
	Iterator[T]
	NextBack() *option.Option[T]
}

func DefaultMap[T any, F any](iter Iterator[T], fn func(*T) F) *Map[T, F] {
	return NewMap(iter, fn)
}

func DefaultFilter[T any](iter Iterator[T], fn func(*T) bool) *Filter[T] {
	return NewFilter(iter, fn)
}

func DefaultEnumerate[T any](iter Iterator[T]) *Enumerate[T] {
	return NewEnumerate(iter)
}

func DefaultTake[T any](iter Iterator[T]) *Take[T] {
	return NewTake(iter)
}

func DefaultFold[T any, I any](iter Iterator[T], init I, fn func(I, *T) I) I {
	acc := init

	for {
		peek := iter.Next()

		if peek.IsNil() {
			break
		}

		acc = fn(acc, peek.UnwrapPtr())
	}

	return acc
}

func DefaultForeach[T any](iter Iterator[T]) {
	DefaultFold(iter, 0, func(i int, t *T) int {
		return i
	})
}

func DefaultAll[T any](iter Iterator[T], fn func(T) bool) bool {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return true
		}

		if !fn(peek.Unwrap()) {
			return false
		}
	}
}

func DefaultAllPtr[T any](iter Iterator[T], fn func(*T) bool) bool {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return true
		}

		if !fn(peek.UnwrapPtr()) {
			return false
		}
	}
}

func DefaultAny[T any](iter Iterator[T], fn func(T) bool) bool {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return false
		}

		if fn(peek.Unwrap()) {
			return true
		}
	}
}

func DefaultAnyPtr[T any](iter Iterator[T], fn func(*T) bool) bool {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return false
		}

		if !fn(peek.UnwrapPtr()) {
			return true
		}
	}
}

func DefaultFind[T any](iter Iterator[T], fn func(T) bool) *option.Option[T] {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return option.Nil[T]()
		}

		if fn(peek.Unwrap()) {
			return option.NewOption(peek.UnwrapPtr())
		}
	}
}

func DefaultFindPtr[T any](iter Iterator[T], fn func(*T) bool) *option.Option[T] {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return option.Nil[T]()
		}

		if fn(peek.UnwrapPtr()) {
			return option.NewOption(peek.UnwrapPtr())
		}
	}
}

func DefaultFindMap[T, I any](iter Iterator[T], fn func(*T) *option.Option[I]) *option.Option[I] {
	for {
		peek := iter.Next()

		if peek.IsNil() {
			return option.Nil[I]()
		}

		mapped := fn(peek.UnwrapPtr())

		if !mapped.IsNil() {
			return mapped
		}
	}
}

func DefaultPosition[T any](iter Iterator[T], fn func(T) bool) *option.Option[int] {
	cnt := 0

	for {
		peek := iter.Next()

		if peek.IsNil() {
			return option.Nil[int]()
		}

		if fn(peek.Unwrap()) {
			return option.NewWrap[int](cnt)
		}

		cnt++
	}
}

func DefaultCollect[T, S any](iter Iterator[T], from FromIterator[T, S]) S {
	return from.fromIter(iter)
}
