package iter

import "github.com/Misonoi/rseaon/option"

type iRange[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	l, r, posL, posR T
}

func NewRange[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l, r T) *iRange[T] {
	return &iRange[T]{
		l:    l,
		r:    r,
		posL: l,
		posR: r,
	}
}

func (r *iRange[T]) Next() *option.Option[T] {
	if r.posL == r.r+1 {
		r.posL = r.l
		return option.Nil[T]()
	}

	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *iRange[T]) NextBack() *option.Option[T] {
	if r.posR == r.l-1 {
		r.posR = r.r
		return option.Nil[T]()
	}

	r.posR -= 1
	return option.NewWrap[T](r.posR + 1)
}

type rangeExclusiveR[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	l, r, posL, posR T
}

func NewRangeExclusiveR[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l, r T) *rangeExclusiveR[T] {
	return &rangeExclusiveR[T]{
		l:    l,
		r:    r,
		posL: l,
		posR: r,
	}
}

func (r *rangeExclusiveR[T]) Next() *option.Option[T] {
	if r.posL == r.r {
		r.posL = r.l
		return option.Nil[T]()
	}

	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *rangeExclusiveR[T]) NextBack() *option.Option[T] {
	if r.posR == r.l {
		r.posR = r.r
		return option.Nil[T]()
	}

	r.posR -= 1
	return option.NewWrap[T](r.posR - 1)
}

type rangeForever[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	posL T
	posR T
}

func NewRangeForever[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l T) *rangeForever[T] {
	return &rangeForever[T]{
		posL: l,
		posR: l,
	}
}

func (r *rangeForever[T]) Next() *option.Option[T] {
	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *rangeForever[T]) NextBack() *option.Option[T] {
	r.posR -= 1
	return option.NewWrap[T](r.posR + 1)
}
