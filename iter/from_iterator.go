package iter

type FromIterator[T, S any] interface {
	fromIter(iter Iterator[T]) S
}
