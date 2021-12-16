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

type Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R

type Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R

type Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R

type Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any] func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R

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

func (r Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, R]]]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, R]]]]]]]]]] {
		return Func10[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
		}).Curried()
	}
}

func (r Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Shift() Func11[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11)
	}
}

func (r Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, R]]]]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, R]]]]]]]]]]] {
		return Func11[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
		}).Curried()
	}
}

func (r Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Shift() Func12[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12)
	}
}

func (r Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, Func1[A13, R]]]]]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, Func1[A13, R]]]]]]]]]]]] {
		return Func12[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
		}).Curried()
	}
}

func (r Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Shift() Func13[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13)
	}
}

func (r Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Curried() Func1[A1, Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, Func1[A13, Func1[A14, R]]]]]]]]]]]]]] {
	return func(a1 A1) Func1[A2, Func1[A3, Func1[A4, Func1[A5, Func1[A6, Func1[A7, Func1[A8, Func1[A9, Func1[A10, Func1[A11, Func1[A12, Func1[A13, Func1[A14, R]]]]]]]]]]]]] {
		return Func13[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R](func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
			return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
		}).Curried()
	}
}

func (r Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Shift() Func14[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A1, R] {
	return func(a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14, a1 A1) R {
		return r(a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14)
	}
}
