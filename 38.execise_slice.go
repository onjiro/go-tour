package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dy)
	for y := range ret {
		ret[y] = make([]uint8, dx)
		for x := range ret[y] {
			ret[y][x] = (uint8)(x * y)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
