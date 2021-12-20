package curried

import (
	"gotptest/fp"
)

func Revert2[A1, A2, R any](f fp.Func1[A1, fp.Func1[A2, R]]) fp.Func2[A1, A2, R] {
	return func(a1 A1, a2 A2) R {
		return f(a1)(a2)
	}
}
