package main

import (
	"fmt"
	"math"
)

func main() {
	// @see for Nextafter http: //golang.org/pkg/math/#Nextafter
	fmt.Printf("Now you have %g programs.",
		math.Nextafter(2, 3))
}
