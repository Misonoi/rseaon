package option

type Option[T any] struct {
	value *T
}

func (t *Option[T]) IsNil() bool {
	return t.value == nil
}

func NewOption[T any](value *T) *Option[T] {
	return &Option[T]{
		value: value,
	}
}

func NewWrap[T any](value T) *Option[T] {
	v := new(T)
	*v = value

	return &Option[T]{
		value: v,
	}
}

func (t *Option[T]) UnwrapPtr() *T {
	return t.value
}

func (t *Option[T]) Unwrap() T {
	return *t.value
}

func (t *Option[T]) UnwrapOr(dft T) T {
	if t.IsNil() {
		return dft
	}

	return t.Unwrap()
}

func (t *Option[T]) UnwrapPtrOr(dft *T) *T {
	if t.IsNil() {
		return dft
	}

	return t.UnwrapPtr()
}

func (t *Option[T]) UnwrapOrElse(fn func() T) T {
	if t.IsNil() {
		return fn()
	}

	return t.Unwrap()
}

func (t *Option[T]) UnwrapPtrOrElse(fn func() *T) *T {
	if t.IsNil() {
		return fn()
	}

	return t.UnwrapPtr()
}

func andThen[T, F any](op *Option[T], fn func(T) *Option[F]) *Option[F] {
	if op.IsNil() {
		return NewOption[F](nil)
	}

	return fn(op.Unwrap())
}

func Map[T any, F any](t *Option[T], fn func(*T) F) *Option[F] {
	if t.IsNil() {
		return NewOption[F](nil)
	}

	return NewWrap(fn(t.value))
}

func Nil[T any]() *Option[T] {
	return &Option[T]{
		value: nil,
	}
}
