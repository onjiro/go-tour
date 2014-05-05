package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Newton 法を利用した近似解の導出
	// 固定10回
	v := 1.0
	for i := 0; i < 10; i++ {
		v = v - (math.Pow(v, 2)-x)/2*v
	}
	return v
}
func main() {
	fmt.Println(Sqrt(2))
}
