package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g programs.",
		math.Nextafter(2, 3))
}
