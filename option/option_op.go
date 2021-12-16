//go:generate go run gotptest/internal/generator/option_gen
package option

import (
	"gotptest/fp"
)

func Some[T any](v T) fp.Option[T] {
	return fp.None[T]{}.Recover(func() T {
		return v
	})
}

func None[T any]() fp.Option[T] {
	return fp.None[T]{}
}

func Ap[T, U any](t fp.Option[fp.Func1[T, U]], a fp.Option[T]) fp.Option[U] {
	return FlatMap(t, func(f fp.Func1[T, U]) fp.Option[U] {
		return Map(a.(fp.Option[T]), func(a T) U {
			return f(a)
		})
	})
}

func Map[T, U any](opt fp.Option[T], f func(v T) U) fp.Option[U] {
	return FlatMap(opt, func(v T) fp.Option[U] {
		return Some(f(v))
	})
}

func FlatMap[T, U any](opt fp.Option[T], fn func(v T) fp.Option[U]) fp.Option[U] {
	if opt.IsDefined() {
		return fn(opt.Get())
	}
	return None[U]()
}

type ApplicativeFunctor1[HT,A, R any] struct {
	h  fp.Option[HT]
	fn fp.Option[fp.Func1[A, R]]
}


func (r ApplicativeFunctor1[HT,A, R]) ApOption(a fp.Option[A]) fp.Option[R] {
	return Ap(r.fn, a)
}

func (r ApplicativeFunctor1[HT,A, R]) Ap(a A) fp.Option[R] {
	return r.ApOption(Some(a))
}

func Applicative1[A, R any](fn fp.Func1[A, R]) ApplicativeFunctor1[fp.Unit, A, R] {
	return ApplicativeFunctor1[fp.Unit, A, R]{None[fp.Unit](), Some(fn)}
}
