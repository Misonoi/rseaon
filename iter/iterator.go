package iter

import (
	"cmp"
	"github.com/Misonoi/rseaon/option"
)

// Iterator is the core interface of package iter.
// It defines a method Next which returns the next element of iterator.
// The return value is option.Option of non-nil when the next element is non-nil otherwise nil.
// If you want to implement your own Iterator, noticed that consumers will consume the iterator.
// Reset the iterator after one iteration if you want to use the same iterator multiple times or
// create a new iterator.
type Iterator[T any] interface {
	Next() *option.Option[T]
}

// DoubleEndedIterator is the extension of Iterator.
// It contains NextBack additionally which returns the next element of reverse iteration.
// Note that forward iteration and reverse iteration are two separate process, which
// means that calling NextBack and then Iterator.Next does not cause the iterator to return to
// the previous element
type DoubleEndedIterator[T any] interface {
	Iterator[T]
	NextBack() *option.Option[T]
}

// Map receive an iterator and a function.
// Performs the function on each element of the iterator and returns a new iterator with all return values.
func Map[T any, F any](iter Iterator[T], fn func(*T) F) *iMap[T, F] {
	return newMap(iter, fn)
}

// Filter receive an iterator and a function.
// Performs the function on each element of the iterator and returns a new iterator with the elements which
// the function returns true.
func Filter[T any](iter Iterator[T], fn func(*T) bool) *filter[T] {
	return newFilter(iter, fn)
}

// Enumerate receive an iterator.
// It will make a new iterator that return a tuple (int, option.Option), just like for-range loop,
// where the first argument is the index of the element in the iterator and the second is
// the element itself.
func Enumerate[T any](iter Iterator[T]) *enumerate[T] {
	return newEnumerate(iter)
}

// Take receive an iterator.
// It returns an iterator contains the first n elements.
func Take[T any](iter Iterator[T], n int) *take[T] {
	return newTake(iter, n)
}

// Fold receive an iterator, an initialization value and a function.
// Generates a value from an iterator through a function.
func Fold[T any, I any](iter Iterator[T], init I, fn func(I, *T) I) I {
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

// Foreach receive an iterator and a function.
// Performs the function on each element.
func Foreach[T any](iter Iterator[T], fn func(*T)) {
	Fold(iter, 0, func(i int, t *T) int {
		fn(t)
		return i
	})
}

func Zip[T, S any](iter Iterator[T], iter2 Iterator[S]) *zip[T, S] {
	return newZip(iter, iter2)
}

func All[T any](iter Iterator[T], fn func(T) bool) bool {
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

func AllPtr[T any](iter Iterator[T], fn func(*T) bool) bool {
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

func Any[T any](iter Iterator[T], fn func(T) bool) bool {
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

func AnyPtr[T any](iter Iterator[T], fn func(*T) bool) bool {
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

func Find[T any](iter Iterator[T], fn func(T) bool) *option.Option[T] {
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

func FindPtr[T any](iter Iterator[T], fn func(*T) bool) *option.Option[T] {
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

func FindMap[T, I any](iter Iterator[T], fn func(*T) *option.Option[I]) *option.Option[I] {
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

func Position[T any](iter Iterator[T], fn func(T) bool) *option.Option[int] {
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

func Max[T cmp.Ordered](iter Iterator[T]) *option.Option[T] {
	maxElement := option.Nil[T]()

	for {
		peek := iter.Next()

		if peek.IsNil() {
			break
		}

		if maxElement.IsNil() {
			maxElement = option.NewOption[T](peek.UnwrapPtr())
		}

		maxElement = option.Map[T, T](maxElement, func(t *T) T {
			return max(*t, peek.Unwrap())
		})
	}

	return maxElement
}

func MaxBy[T any](iter Iterator[T], fn func(a, b *T) bool) *option.Option[T] {
	maxElement := option.Nil[T]()

	for {
		peek := iter.Next()

		if peek.IsNil() {
			break
		}

		if maxElement.IsNil() {
			maxElement = option.NewOption[T](peek.UnwrapPtr())
		}

		maxElement = option.Map[T, T](maxElement, func(t *T) T {
			if fn(peek.UnwrapPtr(), t) {
				return peek.Unwrap()
			} else {
				return *t
			}
		})
	}

	return maxElement
}

func Min[T cmp.Ordered](iter Iterator[T]) *option.Option[T] {
	maxElement := option.Nil[T]()

	for {
		peek := iter.Next()

		if peek.IsNil() {
			break
		}

		if maxElement.IsNil() {
			maxElement = option.NewOption[T](peek.UnwrapPtr())
		}

		maxElement = option.Map[T, T](maxElement, func(t *T) T {
			return min(*t, peek.Unwrap())
		})
	}

	return maxElement
}

func MinBy[T any](iter Iterator[T], fn func(a, b *T) bool) *option.Option[T] {
	return MaxBy(iter, fn)
}

func Collect[T, S any](iter Iterator[T], from FromIterator[T, S]) S {
	return from.fromIter(iter)
}
