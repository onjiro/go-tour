package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	z := 1.0 + 0i
	for i := 0; i < 100; i++ {
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(x, 2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
