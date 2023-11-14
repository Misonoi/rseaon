package iter

import (
	"cmp"
	"github.com/Misonoi/rseaon/option"
	"math/rand"
)

func InstantMap[T, F any](slice []T, fn func(T, int) F) []F {
	fs := make([]F, len(slice))

	for i, e := range slice {
		fs[i] = fn(e, i)
	}

	return fs
}

func InstantFilter[T any](slice []T, fn func(T, int) bool) []T {
	fs := make([]T, 0)

	for i, e := range slice {
		if fn(e, i) {
			fs = append(fs, e)
		}
	}

	return fs
}

func InstantFilterMap[T, F any](slice []T, fn func(T, int) (F, bool)) []F {
	fs := make([]F, 0)

	for i, e := range slice {
		if res, ok := fn(e, i); ok {
			fs = append(fs, res)
		}
	}

	return fs
}

func InstantFlatMap[T, F any](slice []T, fn func(T, int) []F) []F {
	fs := make([]F, 0, len(slice))

	for i, e := range slice {
		fs = append(fs, fn(e, i)...)
	}

	return fs
}

func InstantFold[T, I any](slice []T, init I, fn func(I, T, int) I) I {
	acc := init

	for i, e := range slice {
		acc = fn(acc, e, i)
	}

	return acc
}

func InstantRFold[T, I any](slice []T, init I, fn func(I, T, int) I) I {
	acc := init

	for i := len(slice) - 1; i >= 0; i-- {
		acc = fn(acc, slice[i], i)
	}

	return acc
}

func InstantForEach[T any](slice []T, fn func(T, int)) {
	for i, e := range slice {
		fn(e, i)
	}
}

func InstantCount[T comparable](slice []T, target T) int {
	cnt := 0

	for _, e := range slice {
		if e == target {
			cnt++
		}
	}

	return cnt
}

func InstantCountBy[T any](slice []T, fn func(T) bool) int {
	cnt := 0

	for _, e := range slice {
		if fn(e) {
			cnt++
		}
	}

	return cnt
}

func InstantFind[T comparable](slice []T, target T) int {
	for i, e := range slice {
		if e == target {
			return i
		}
	}

	return -1
}

func InstantAny[T any](slice []T, fn func(T, int) bool) bool {
	for i, e := range slice {
		if fn(e, i) {
			return true
		}
	}

	return false
}

func InstantAll[T any](slice []T, fn func(T, int) bool) bool {
	for i, e := range slice {
		if !fn(e, i) {
			return false
		}
	}

	return true
}

func InstantMax[T cmp.Ordered](slice []T) *option.Option[T] {
	if len(slice) == 0 {
		return option.Nil[T]()
	}

	maxElement := new(T)

	*maxElement = slice[0]

	for _, e := range slice {
		if *maxElement >= e {
			*maxElement = e
		}
	}

	return option.NewOption(maxElement)
}

func InstantMin[T cmp.Ordered](slice []T) *option.Option[T] {
	if len(slice) == 0 {
		return option.Nil[T]()
	}

	maxElement := new(T)

	*maxElement = slice[0]

	for _, e := range slice {
		if *maxElement <= e {
			*maxElement = e
		}
	}

	return option.NewOption(maxElement)
}

func InstantMaxBy[T any](slice []T, cmp func(a, b T) bool) *option.Option[T] {
	if len(slice) == 0 {
		return option.Nil[T]()
	}

	maxElement := new(T)

	*maxElement = slice[0]

	for _, e := range slice {
		if cmp(*maxElement, e) {
			*maxElement = e
		}
	}

	return option.NewOption(maxElement)
}

func InstantMinBy[T any](slice []T, cmp func(a, b T) bool) *option.Option[T] {
	if len(slice) == 0 {
		return option.Nil[T]()
	}

	maxElement := new(T)

	*maxElement = slice[0]

	for _, e := range slice {
		if !cmp(*maxElement, e) {
			*maxElement = e
		}
	}

	return option.NewOption(maxElement)
}

func InstantFlatten[T any](slice [][]T) []T {
	fs := make([]T, 0)

	for _, e := range slice {
		fs = append(fs, e...)
	}

	return fs
}

func InstantShuffle[T any](slice []T) []T {
	rand.Shuffle(len(slice), func(a, b int) {
		slice[a], slice[b] = slice[b], slice[a]
	})

	return slice
}

func InstantReverse[T any](slice []T) []T {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - 1 - i
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}

func InstantDropN[T any](slice []T, n int) []T {
	if len(slice) <= n {
		return make([]T, 0)
	}

	fs := make([]T, 0, len(slice)-n)

	return append(fs, slice[n:]...)
}

func InstantDropWhile[T any](slice []T, fn func(T) bool) []T {
	i := 0

	for ; i < len(slice); i++ {
		if !fn(slice[i]) {
			break
		}
	}

	fs := make([]T, 0, len(slice)-1)

	return append(fs, slice[i:]...)
}

func InstantLast[T any](slice []T) *option.Option[T] {
	if len(slice) == 0 {
		return option.Nil[T]()
	}

	return option.NewOption[T](&slice[len(slice)-1])
}
