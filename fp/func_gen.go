package fp

type Func2[A1, A2, R any] func(a1 A1, a2 A2) R

type Func3[A1, A2, A3, R any] func(a1 A1, a2 A2, a3 A3) R

type Func4[A1, A2, A3, A4, R any] func(a1 A1, a2 A2, a3 A3, a4 A4) R

type Func5[A1, A2, A3, A4, A5, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R

type Func6[A1, A2, A3, A4, A5, A6, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R

type Func7[A1, A2, A3, A4, A5, A6, A7, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R

type Func8[A1, A2, A3, A4, A5, A6, A7, A8, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R

type Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R

type Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R

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

func (r Func3[A1, A2, A3, R]) Curried() Func1[A1, Func1[A2, Func1[A3, R]]] {
	return func(a1 A1) Func1[A2, Func1[A3, R]] {
		return Func2[A2, A3, R](func(a2 A2, a3 A3) R {
			return r(a1, a2, a3)
		}).Curried()
	}
}

func (r Func3[A1, A2, A3, R]) Shift() Func3[A2, A3, A1, R] {
	return func(a2 A2, a3 A3, a1 A1) R {
		return r(a1, a2, a3)
	}
}

func (r Func4[A1, A2, A3, A4, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, R]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, R]]] {
		return Func3[A2, A3, A4, R](func(a2 A2, a3 A3, a4 A4) R {
			return r(a1, a2, a3, a4)
		}).Curried()
	}
}

func (r Func4[A1, A2, A3, A4, R]) Shift() Func4[A2, A3, A4, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a1 A1) R {
		return r(a1, a2, a3, a4)
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, R]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, R]]]] {
		return Func4[A2, A3, A4, A5, R](func(a2 A2, a3 A3, a4 A4, a5 A5) R {
			return r(a1, a2, a3, a4, a5)
		}).Curried()
	}
}

func (r Func5[A1, A2, A3, A4, A5, R]) Shift() Func5[A2, A3, A4, A5, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a1 A1) R {
		return r(a1, a2, a3, a4, a5)
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, R]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, R]]]]] {
		return Func5[A2, A3, A4, A5, A6, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
			return r(a1, a2, a3, a4, a5, a6)
		}).Curried()
	}
}

func (r Func6[A1, A2, A3, A4, A5, A6, R]) Shift() Func6[A2, A3, A4, A5, A6, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6)
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, R]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, R]]]]]] {
		return Func6[A2, A3, A4, A5, A6, A7, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
			return r(a1, a2, a3, a4, a5, a6, a7)
		}).Curried()
	}
}

func (r Func7[A1, A2, A3, A4, A5, A6, A7, R]) Shift() Func7[A2, A3, A4, A5, A6, A7, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7)
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, R]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, R]]]]]]] {
		return Func7[A2, A3, A4, A5, A6, A7, A8, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8)
		}).Curried()
	}
}

func (r Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) Shift() Func8[A2, A3, A4, A5, A6, A7, A8, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8)
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, R]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, R]]]]]]]] {
		return Func8[A2, A3, A4, A5, A6, A7, A8, A9, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
		}).Curried()
	}
}

func (r Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Shift() Func9[A2, A3, A4, A5, A6, A7, A8, A9, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9)
	}
}

func (r Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, R]]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, R]]]]]]]]] {
		return Func9[A2, A3, A4, A5, A6, A7, A8, A9, A10, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
		}).Curried()
	}
}

func (r Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Shift() Func10[A2, A3, A4, A5, A6, A7, A8, A9, A10, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10)
	}
}
