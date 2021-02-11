package main

import (
	"fmt"
	"math"
)

func main() {
	var diameters float64 = 64
	var luas, keliling = calculate(diameters)

	fmt.Printf("Luas lingkaran\t\t: %.2f \n", luas)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", keliling)
}

// func calculate(diameter float64) (float64, float64) {
// 	// hitung luas lingkaran
// 	var luas = math.Pi * math.Pow(diameter/2, 2)

// 	var keliling = math.Pi * diameter

// 	return luas, keliling
// }

// VERSI 2

func calculate(d float64) (luas float64, keliling float64) {
	luas = math.Pi * math.Pow(d/2, 2) // koma disini adalah pangkat <- fungsi dari Pow()

	keliling = math.Pi * d

	return
}
