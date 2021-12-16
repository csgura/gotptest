package option

import (
	"gotptest/curried"
	"gotptest/fp"
)

type ApplicativeFunctor2[HT, A1, A2, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[HT, A1, A2, R]) Shift() ApplicativeFunctor2[HT, A2, A1, R] {

	nf := fp.Compose(curried.Revert2[A1, A2, R], fp.Compose(fp.Func2[A1, A2, R].Shift, fp.Func2[A2, A1, R].Curried))
	return ApplicativeFunctor2[HT, A2, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor2[HT, A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[A1, A2, R] {

	return ApplicativeFunctor1[A1, A2, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[HT, A1, A2, R]) Ap(a A1) ApplicativeFunctor1[A1, A2, R] {

	return r.ApOption(Some(a))

}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[fp.Unit, A1, A2, R] {
	return ApplicativeFunctor2[fp.Unit, A1, A2, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor3[HT, A1, A2, A3, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, R]]]]
}

func (r ApplicativeFunctor3[HT, A1, A2, A3, R]) Shift() ApplicativeFunctor3[HT, A2, A3, A1, R] {

	nf := fp.Compose(curried.Revert3[A1, A2, A3, R], fp.Compose(fp.Func3[A1, A2, A3, R].Shift, fp.Func3[A2, A3, A1, R].Curried))
	return ApplicativeFunctor3[HT, A2, A3, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor3[HT, A1, A2, A3, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor2[A1, A2, A3, R] {

	return ApplicativeFunctor2[A1, A2, A3, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor3[HT, A1, A2, A3, R]) Ap(a A1) ApplicativeFunctor2[A1, A2, A3, R] {

	return r.ApOption(Some(a))

}
func Applicative3[A1, A2, A3, R any](fn fp.Func3[A1, A2, A3, R]) ApplicativeFunctor3[fp.Unit, A1, A2, A3, R] {
	return ApplicativeFunctor3[fp.Unit, A1, A2, A3, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor4[HT, A1, A2, A3, A4, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, R]]]]]
}

func (r ApplicativeFunctor4[HT, A1, A2, A3, A4, R]) Shift() ApplicativeFunctor4[HT, A2, A3, A4, A1, R] {

	nf := fp.Compose(curried.Revert4[A1, A2, A3, A4, R], fp.Compose(fp.Func4[A1, A2, A3, A4, R].Shift, fp.Func4[A2, A3, A4, A1, R].Curried))
	return ApplicativeFunctor4[HT, A2, A3, A4, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor4[HT, A1, A2, A3, A4, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor3[A1, A2, A3, A4, R] {

	return ApplicativeFunctor3[A1, A2, A3, A4, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor4[HT, A1, A2, A3, A4, R]) Ap(a A1) ApplicativeFunctor3[A1, A2, A3, A4, R] {

	return r.ApOption(Some(a))

}
func Applicative4[A1, A2, A3, A4, R any](fn fp.Func4[A1, A2, A3, A4, R]) ApplicativeFunctor4[fp.Unit, A1, A2, A3, A4, R] {
	return ApplicativeFunctor4[fp.Unit, A1, A2, A3, A4, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor5[HT, A1, A2, A3, A4, A5, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, R]]]]]]
}

func (r ApplicativeFunctor5[HT, A1, A2, A3, A4, A5, R]) Shift() ApplicativeFunctor5[HT, A2, A3, A4, A5, A1, R] {

	nf := fp.Compose(curried.Revert5[A1, A2, A3, A4, A5, R], fp.Compose(fp.Func5[A1, A2, A3, A4, A5, R].Shift, fp.Func5[A2, A3, A4, A5, A1, R].Curried))
	return ApplicativeFunctor5[HT, A2, A3, A4, A5, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor5[HT, A1, A2, A3, A4, A5, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor4[A1, A2, A3, A4, A5, R] {

	return ApplicativeFunctor4[A1, A2, A3, A4, A5, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor5[HT, A1, A2, A3, A4, A5, R]) Ap(a A1) ApplicativeFunctor4[A1, A2, A3, A4, A5, R] {

	return r.ApOption(Some(a))

}
func Applicative5[A1, A2, A3, A4, A5, R any](fn fp.Func5[A1, A2, A3, A4, A5, R]) ApplicativeFunctor5[fp.Unit, A1, A2, A3, A4, A5, R] {
	return ApplicativeFunctor5[fp.Unit, A1, A2, A3, A4, A5, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor6[HT, A1, A2, A3, A4, A5, A6, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, R]]]]]]]
}

func (r ApplicativeFunctor6[HT, A1, A2, A3, A4, A5, A6, R]) Shift() ApplicativeFunctor6[HT, A2, A3, A4, A5, A6, A1, R] {

	nf := fp.Compose(curried.Revert6[A1, A2, A3, A4, A5, A6, R], fp.Compose(fp.Func6[A1, A2, A3, A4, A5, A6, R].Shift, fp.Func6[A2, A3, A4, A5, A6, A1, R].Curried))
	return ApplicativeFunctor6[HT, A2, A3, A4, A5, A6, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor6[HT, A1, A2, A3, A4, A5, A6, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor5[A1, A2, A3, A4, A5, A6, R] {

	return ApplicativeFunctor5[A1, A2, A3, A4, A5, A6, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor6[HT, A1, A2, A3, A4, A5, A6, R]) Ap(a A1) ApplicativeFunctor5[A1, A2, A3, A4, A5, A6, R] {

	return r.ApOption(Some(a))

}
func Applicative6[A1, A2, A3, A4, A5, A6, R any](fn fp.Func6[A1, A2, A3, A4, A5, A6, R]) ApplicativeFunctor6[fp.Unit, A1, A2, A3, A4, A5, A6, R] {
	return ApplicativeFunctor6[fp.Unit, A1, A2, A3, A4, A5, A6, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor7[HT, A1, A2, A3, A4, A5, A6, A7, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, R]]]]]]]]
}

func (r ApplicativeFunctor7[HT, A1, A2, A3, A4, A5, A6, A7, R]) Shift() ApplicativeFunctor7[HT, A2, A3, A4, A5, A6, A7, A1, R] {

	nf := fp.Compose(curried.Revert7[A1, A2, A3, A4, A5, A6, A7, R], fp.Compose(fp.Func7[A1, A2, A3, A4, A5, A6, A7, R].Shift, fp.Func7[A2, A3, A4, A5, A6, A7, A1, R].Curried))
	return ApplicativeFunctor7[HT, A2, A3, A4, A5, A6, A7, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor7[HT, A1, A2, A3, A4, A5, A6, A7, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, A7, R] {

	return ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, A7, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor7[HT, A1, A2, A3, A4, A5, A6, A7, R]) Ap(a A1) ApplicativeFunctor6[A1, A2, A3, A4, A5, A6, A7, R] {

	return r.ApOption(Some(a))

}
func Applicative7[A1, A2, A3, A4, A5, A6, A7, R any](fn fp.Func7[A1, A2, A3, A4, A5, A6, A7, R]) ApplicativeFunctor7[fp.Unit, A1, A2, A3, A4, A5, A6, A7, R] {
	return ApplicativeFunctor7[fp.Unit, A1, A2, A3, A4, A5, A6, A7, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor8[HT, A1, A2, A3, A4, A5, A6, A7, A8, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, R]]]]]]]]]
}

func (r ApplicativeFunctor8[HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Shift() ApplicativeFunctor8[HT, A2, A3, A4, A5, A6, A7, A8, A1, R] {

	nf := fp.Compose(curried.Revert8[A1, A2, A3, A4, A5, A6, A7, A8, R], fp.Compose(fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R].Shift, fp.Func8[A2, A3, A4, A5, A6, A7, A8, A1, R].Curried))
	return ApplicativeFunctor8[HT, A2, A3, A4, A5, A6, A7, A8, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor8[HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, A8, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor8[HT, A1, A2, A3, A4, A5, A6, A7, A8, R]) Ap(a A1) ApplicativeFunctor7[A1, A2, A3, A4, A5, A6, A7, A8, R] {

	return r.ApOption(Some(a))

}
func Applicative8[A1, A2, A3, A4, A5, A6, A7, A8, R any](fn fp.Func8[A1, A2, A3, A4, A5, A6, A7, A8, R]) ApplicativeFunctor8[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, R] {
	return ApplicativeFunctor8[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor9[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, R]]]]]]]]]]
}

func (r ApplicativeFunctor9[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Shift() ApplicativeFunctor9[HT, A2, A3, A4, A5, A6, A7, A8, A9, A1, R] {

	nf := fp.Compose(curried.Revert9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R], fp.Compose(fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R].Shift, fp.Func9[A2, A3, A4, A5, A6, A7, A8, A9, A1, R].Curried))
	return ApplicativeFunctor9[HT, A2, A3, A4, A5, A6, A7, A8, A9, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor9[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor9[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) Ap(a A1) ApplicativeFunctor8[A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {

	return r.ApOption(Some(a))

}
func Applicative9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R any](fn fp.Func9[A1, A2, A3, A4, A5, A6, A7, A8, A9, R]) ApplicativeFunctor9[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, R] {
	return ApplicativeFunctor9[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor10[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, R]]]]]]]]]]]
}

func (r ApplicativeFunctor10[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Shift() ApplicativeFunctor10[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A1, R] {

	nf := fp.Compose(curried.Revert10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R], fp.Compose(fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R].Shift, fp.Func10[A2, A3, A4, A5, A6, A7, A8, A9, A10, A1, R].Curried))
	return ApplicativeFunctor10[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor10[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	return ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor10[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) Ap(a A1) ApplicativeFunctor9[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {

	return r.ApOption(Some(a))

}
func Applicative10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R any](fn fp.Func10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]) ApplicativeFunctor10[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R] {
	return ApplicativeFunctor10[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor11[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, R]]]]]]]]]]]]
}

func (r ApplicativeFunctor11[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Shift() ApplicativeFunctor11[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A1, R] {

	nf := fp.Compose(curried.Revert11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R], fp.Compose(fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R].Shift, fp.Func11[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A1, R].Curried))
	return ApplicativeFunctor11[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor11[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	return ApplicativeFunctor10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor11[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) Ap(a A1) ApplicativeFunctor10[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {

	return r.ApOption(Some(a))

}
func Applicative11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R any](fn fp.Func11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]) ApplicativeFunctor11[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R] {
	return ApplicativeFunctor11[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor12[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, R]]]]]]]]]]]]]
}

func (r ApplicativeFunctor12[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Shift() ApplicativeFunctor12[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A1, R] {

	nf := fp.Compose(curried.Revert12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R], fp.Compose(fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R].Shift, fp.Func12[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A1, R].Curried))
	return ApplicativeFunctor12[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor12[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	return ApplicativeFunctor11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor12[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) Ap(a A1) ApplicativeFunctor11[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {

	return r.ApOption(Some(a))

}
func Applicative12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R any](fn fp.Func12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]) ApplicativeFunctor12[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R] {
	return ApplicativeFunctor12[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor13[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, R]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor13[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Shift() ApplicativeFunctor13[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A1, R] {

	nf := fp.Compose(curried.Revert13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R], fp.Compose(fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R].Shift, fp.Func13[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A1, R].Curried))
	return ApplicativeFunctor13[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor13[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	return ApplicativeFunctor12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor13[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) Ap(a A1) ApplicativeFunctor12[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {

	return r.ApOption(Some(a))

}
func Applicative13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R any](fn fp.Func13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]) ApplicativeFunctor13[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R] {
	return ApplicativeFunctor13[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, R]{None[fp.Unit](), Some(fn.Curried())}
}

type ApplicativeFunctor14[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, fp.Func1[A3, fp.Func1[A4, fp.Func1[A5, fp.Func1[A6, fp.Func1[A7, fp.Func1[A8, fp.Func1[A9, fp.Func1[A10, fp.Func1[A11, fp.Func1[A12, fp.Func1[A13, fp.Func1[A14, R]]]]]]]]]]]]]]]
}

func (r ApplicativeFunctor14[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Shift() ApplicativeFunctor14[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A1, R] {

	nf := fp.Compose(curried.Revert14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R], fp.Compose(fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R].Shift, fp.Func14[A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A1, R].Curried))
	return ApplicativeFunctor14[HT, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor14[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	return ApplicativeFunctor13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]{a, Ap(r.fn, a)}
}
func (r ApplicativeFunctor14[HT, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) Ap(a A1) ApplicativeFunctor13[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {

	return r.ApOption(Some(a))

}
func Applicative14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R any](fn fp.Func14[A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]) ApplicativeFunctor14[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R] {
	return ApplicativeFunctor14[fp.Unit, A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12, A13, A14, R]{None[fp.Unit](), Some(fn.Curried())}
}
