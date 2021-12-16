//go:generate go run gotptest/internal/generator/fp_gen
package fp

type Func0[R any] func() R

type Func1[A1, R any] func(a1 A1) R

func (r Func1[A1, R]) Curried() Func1[A1, R] {
	return r
}

func Compose[A, B, C any](f1 Func1[A, B], f2 Func1[B, C]) Func1[A, C] {
	return func(a A) C {
		return f2(f1(a))
	}
}


type Unit struct{}

type Option[T any] interface {
	IsDefined() bool
	Get() T
}

type Some[T any] struct {
	v T
}


func (r Some[T]) IsDefined() bool {
	return true
}


func (r Some[T]) Get() T {
	return r.v
}

type None[T any] struct{}

func (r None[T]) IsDefined() bool {
	return false
}

func (r None[T]) Get() T {
	panic("Option.empty")
}

func (r None[T]) Recover(f func() T) Option[T] {
	return Some[T]{f()}
}
