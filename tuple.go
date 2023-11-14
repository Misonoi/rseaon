package rseaon

type Tuple2[T, F any] struct {
	First  T
	Second F
}

func MakeTuple2[T, F any](first T, second F) *Tuple2[T, F] {
	return &Tuple2[T, F]{
		First:  first,
		Second: second,
	}
}

type Tuple3[T, F, K any] struct {
	First  T
	Second F
	Third  K
}

func MakeTuple3[T, F, K any](first T, second F, third K) *Tuple3[T, F, K] {
	return &Tuple3[T, F, K]{
		First:  first,
		Second: second,
		Third:  third,
	}
}

type Tuple4[T1, T2, T3, T4 any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
}

func MakeTuple4[T1, T2, T3, T4 any](
	first T1, second T2, third T3, fourth T4,
) *Tuple4[T1, T2, T3, T4] {
	return &Tuple4[T1, T2, T3, T4]{
		First:  first,
		Second: second,
		Third:  third,
		Fourth: fourth,
	}
}

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
	Fifth  T5
}

func MakeTuple5[T1, T2, T3, T4, T5 any](
	first T1, second T2, third T3, fourth T4, fifth T5,
) *Tuple5[T1, T2, T3, T4, T5] {
	return &Tuple5[T1, T2, T3, T4, T5]{
		First:  first,
		Second: second,
		Third:  third,
		Fourth: fourth,
		Fifth:  fifth,
	}
}

type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	First  T1
	Second T2
	Third  T3
	Fourth T4
	Fifth  T5
	Sixth  T6
}

func MakeTuple6[T1, T2, T3, T4, T5, T6 any](
	first T1, second T2, third T3, fourth T4, fifth T5, sixth T6,
) *Tuple6[T1, T2, T3, T4, T5, T6] {
	return &Tuple6[T1, T2, T3, T4, T5, T6]{
		First:  first,
		Second: second,
		Third:  third,
		Fourth: fourth,
		Fifth:  fifth,
		Sixth:  sixth,
	}
}

type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	First   T1
	Second  T2
	Third   T3
	Fourth  T4
	Fifth   T5
	Sixth   T6
	Seventh T7
}

func MakeTuple7[T1, T2, T3, T4, T5, T6, T7 any](
	first T1, second T2, third T3, fourth T4, fifth T5, sixth T6, seventh T7,
) *Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return &Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		First:   first,
		Second:  second,
		Third:   third,
		Fourth:  fourth,
		Fifth:   fifth,
		Sixth:   sixth,
		Seventh: seventh,
	}
}

type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	First   T1
	Second  T2
	Third   T3
	Fourth  T4
	Fifth   T5
	Sixth   T6
	Seventh T7
	Eighth  T8
}

func MakeTuple8[T1, T2, T3, T4, T5, T6, T7, T8 any](
	first T1, second T2, third T3, fourth T4, fifth T5, sixth T6, seventh T7, eighth T8,
) *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return &Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		First:   first,
		Second:  second,
		Third:   third,
		Fourth:  fourth,
		Fifth:   fifth,
		Sixth:   sixth,
		Seventh: seventh,
		Eighth:  eighth,
	}
}

type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	First   T1
	Second  T2
	Third   T3
	Fourth  T4
	Fifth   T5
	Sixth   T6
	Seventh T7
	Eighth  T8
	Ninth   T9
}

func MakeTuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](
	first T1, second T2, third T3, fourth T4, fifth T5, sixth T6, seventh T7, eighth T8, ninth T9,
) *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return &Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		First:   first,
		Second:  second,
		Third:   third,
		Fourth:  fourth,
		Fifth:   fifth,
		Sixth:   sixth,
		Seventh: seventh,
		Eighth:  eighth,
		Ninth:   ninth,
	}
}
