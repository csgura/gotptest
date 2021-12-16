//go:generate go run gotptest/internal/generator/curried_gen
package curried

import (
	"gotptest/fp"
)

func Func1[A, R any](f func(A) R) fp.Func1[A, R] {
	return fp.Func1[A, R](f)
}
