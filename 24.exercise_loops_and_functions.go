package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// Newton 法を利用した近似解の導出
	// 差が小さくなるまでぶん回す
	v := 1.0
	for p := 0.0; math.Abs(v-p) > 0.001; {
		p = v
		v = v - (math.Pow(v, 2)-x)/2*v
	}
	return v
}
func main() {
	fmt.Println(Sqrt(2))
}
