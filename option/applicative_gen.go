package option

import (
	"gotptest/curried"
	"gotptest/fp"
	"gotptest/hlist"
)

type ApplicativeFunctor2[H hlist.Header[HT], HT, A1, A2, R any] struct {
	h  fp.Option[H]
	fn fp.Option[fp.Func1[A1, fp.Func1[A2, R]]]
}

func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Shift() ApplicativeFunctor2[H, HT, A2, A1, R] {

	nf := fp.Compose(curried.Revert2[A1, A2, R], fp.Compose(fp.Func2[A1, A2, R].Shift, fp.Func2[A2, A1, R].Curried))
	return ApplicativeFunctor2[H, HT, A2, A1, R]{
		r.h,
		Map(r.fn, nf),
	}

}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) ApOption(a fp.Option[A1]) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	nh := FlatMap(r.h, func(hv H) fp.Option[hlist.Cons[A1, H]] {
		return Map(a, func(av A1) hlist.Cons[A1, H] {
			return hlist.Concat(av, hv)
		})
	})

	return ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R]{nh, Ap(r.fn, a)}
}
func (r ApplicativeFunctor2[H, HT, A1, A2, R]) Ap(a A1) ApplicativeFunctor1[hlist.Cons[A1, H], A1, A2, R] {

	return r.ApOption(Some(a))

}
func Applicative2[A1, A2, R any](fn fp.Func2[A1, A2, R]) ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R] {
	return ApplicativeFunctor2[hlist.Nil, hlist.Nil, A1, A2, R]{Some(hlist.Empty()), Some(fn.Curried())}
}
