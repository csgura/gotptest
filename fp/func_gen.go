package fp

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

func (r Func2[A1, A2, R]) Curried() Func1[A1, Func1[A2, R]] {
	return func(a1 A1) Func1[A2, R] {
		return Func1[A2, R](func(a2 A2) R {
			return r(a1, a2)
		}).Curried()
	}
}

func (r Func2[A1, A2, R]) Shift() Func2[A2, A1, R] {
	return func(a2 A2, a1 A1) R {
		return r(a1, a2)
	}
}
