package main

import (
	"fmt"
	"math"
)

// An error, when given negative number to Sqrt
type ErrNegativeSqrt float64

// Implement implicitly `error.Error()`
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative numeber: %g", e)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
