package main

import (
	"fmt"

	"gotptest/option"
)

func MakeURLWithPort(scheme, addr string, port int) string {
	return fmt.Sprintf("%s://%s:%d", scheme, addr, port)
}

func main() {

	resOpt := option.Applicative3(MakeURLWithPort).
		ApOption(option.Some("https")).
		Shift().
		Ap(8443).
		Ap("localhost")
	fmt.Println(resOpt)
}
