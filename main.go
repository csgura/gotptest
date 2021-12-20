package main

import (
	"fmt"

	"gotptest/option"
)

type hello struct {
}

func (r hello) String() string {
	return "hello"
}

func main() {
	v := option.Some(hello{})
	_ = v

	res := option.Applicative2(func(a int, b int) int {
		fmt.Println(a, b)
		return 10
	}).
		Ap(1).
		Ap(20)
	fmt.Println(res)
}
