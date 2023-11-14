package iter

import "rseaon/option"

type Range[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	l, r, posL, posR T
}

func NewRange[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l, r T) *Range[T] {
	return &Range[T]{
		l:    l,
		r:    r,
		posL: l,
		posR: r,
	}
}

func (r *Range[T]) Next() *option.Option[T] {
	if r.posL == r.r+1 {
		r.posL = r.l
		return option.Nil[T]()
	}

	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *Range[T]) NextBack() *option.Option[T] {
	if r.posR == r.l-1 {
		r.posR = r.r
		return option.Nil[T]()
	}

	r.posR -= 1
	return option.NewWrap[T](r.posR + 1)
}

type RangeExclusiveR[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	l, r, posL, posR T
}

func NewRangeExclusiveR[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l, r T) *RangeExclusiveR[T] {
	return &RangeExclusiveR[T]{
		l:    l,
		r:    r,
		posL: l,
		posR: r,
	}
}

func (r *RangeExclusiveR[T]) Next() *option.Option[T] {
	if r.posL == r.r {
		r.posL = r.l
		return option.Nil[T]()
	}

	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *RangeExclusiveR[T]) NextBack() *option.Option[T] {
	if r.posR == r.l {
		r.posR = r.r
		return option.Nil[T]()
	}

	r.posR -= 1
	return option.NewWrap[T](r.posR - 1)
}

type RangeForever[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64] struct {
	posL T
	posR T
}

func NewRangeForever[T int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint | uint64](l T) *RangeForever[T] {
	return &RangeForever[T]{
		posL: l,
		posR: l,
	}
}

func (r *RangeForever[T]) Next() *option.Option[T] {
	r.posL += 1
	return option.NewWrap[T](r.posL - 1)
}

func (r *RangeForever[T]) NextBack() *option.Option[T] {
	r.posR -= 1
	return option.NewWrap[T](r.posR + 1)
}
