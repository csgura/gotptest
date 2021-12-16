package curried

import (
	"gotptest/fp"
)

func Revert2[A1, A2, R any](f fp.Func1[A1, fp.Func1[A2, R]]) fp.Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return f(a1)(a2)
	}
}
func Revert3[A1, A2, A3, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]) fp.Func3[A1, A2, A3, R] {
	return func(a1 A1, a2 A2, a3 A3) R {
		return f(a1)(a2)(a3)
	}
}
func Revert4[A1, A2, A3, A4, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]) fp.Func4[A1, A2, A3, A4, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4) R {
		return f(a1)(a2)(a3)(a4)
	}
}
func Revert5[A1, A2, A3, A4, A5, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]) fp.Func5[A1, A2, A3, A4, A5, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5) R {
		return f(a1)(a2)(a3)(a4)(a5)
	}
}
func Revert6[A1, A2, A3, A4, A5, A6, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]) fp.Func6[A1, A2, A3, A4, A5, A6, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)
	}
}
func Revert7[A1, A2, A3, A4, A5, A6, A7, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]) fp.Func7[A1, A2, A3, A4, A5, A6, A7, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)
	}
}
func Revert8[A1, A2, A3, A4, A5, A6, A7, A8, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]) fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)
	}
}
func Revert9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]) fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)
	}
}
func Revert10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, R]]]]]]]]]]) fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)(a10)
	}
}
func Revert11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, R]]]]]]]]]]]) fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)(a10)(a11)
	}
}
func Revert12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, R]]]]]]]]]]]]) fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)(a10)(a11)(a12)
	}
}
func Revert13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, R]]]]]]]]]]]]]) fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)(a10)(a11)(a12)(a13)
	}
}
func Revert14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](f fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, R]]]]]]]]]]]]]]) fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return func(a1 A1, a2 A2, a3 A3, a4 A4, a5 A5, a6 A6, a7 A7, a8 A8, a9 A9, a10 A10, a11 A11, a12 A12, a13 A13, a14 A14) R {
		return f(a1)(a2)(a3)(a4)(a5)(a6)(a7)(a8)(a9)(a10)(a11)(a12)(a13)(a14)
	}
}
