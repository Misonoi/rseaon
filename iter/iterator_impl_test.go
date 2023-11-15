package iter

import (
	"fmt"
	"github.com/Misonoi/rseaon"
	"github.com/stretchr/testify/assert"
	"testing"
)

func iter() *SliceIter[int] {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return NewSliceIter(slice)
}

func TestDefaultAll(t *testing.T) {
	slice := []int{2, 4, 6, 8, 10}
	iter := NewSliceIter(slice)

	assert.Equal(t, All[int](iter, func(t int) bool {
		return t%2 == 0
	}), true)

	assert.Equal(t, All[int](iter, func(t int) bool {
		return t%3 == 0
	}), false)
}

func TestDefaultAny(t *testing.T) {
	slice := []int{1, 3, 5, 7, 9}
	iter := NewSliceIter(slice)

	assert.Equal(t, Any[int](iter, func(i int) bool {
		return i%2 == 0
	}), false)

	iter.slice = append(iter.slice, 2)

	assert.Equal(t, Any[int](iter, func(i int) bool {
		return i%2 == 0
	}), true)
}

func TestDefaultFind(t *testing.T) {
	slice := []int{2, 4, 6, 8, 10}
	iter := NewSliceIter(slice)

	assert.Equal(t, Find[int](iter, func(i int) bool {
		return i%2 == 0
	}).Unwrap(), 2)

	assert.Equal(t, Find[int](iter, func(i int) bool {
		return i%2 == 0
	}).Unwrap(), 4)
}

func TestDefaultFilter(t *testing.T) {
	iter := iter()

	assert.Equal(t, Fold[int, int](Filter[int](iter, func(i *int) bool {
		return *i%2 == 0
	}), 0, func(i int, i2 *int) int {
		return i + *i2
	}), 30)
}

func TestDefaultEnumerate(t *testing.T) {
	iter := iter()

	enum := newEnumerate[int](iter)

	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
	fmt.Printf("%v\n", formatEnumerate(enum.Next().UnwrapPtr()))
}

func formatEnumerate[T any](e *rseaon.Tuple2[*T, int]) string {
	return fmt.Sprintf("{ %v %v }", *e.First, e.Second)
}

func TestDefaultCollect(t *testing.T) {
	iter := iter()

	mapped := Map[int, int](iter, func(i *int) int {
		return *i + 2
	})

	slice := NewSliceWrapper[int]().fromIter(mapped).Unwrap()

	for _, e := range slice {
		println(e)
	}
}

func TestMax(t *testing.T) {
	iter := iter()

	assert.Equal(t, Max[int](iter).Unwrap(), 10)
}

func TestMaxBy(t *testing.T) {
	iter := iter()

	assert.Equal(t, MaxBy[int](iter, func(a, b *int) bool {
		if *a < *b {
			return true
		}

		return false
	}).Unwrap(), 1)
}

func TestMap(t *testing.T) {
	iter := iter()

	Filter[int](iter, func(i *int) bool {
		return *i%2 == 0
	})

	slice := NewSliceWrapper[int]().fromIter(iter).Unwrap()

	for _, e := range slice {
		println(e)
	}
}
