package main

import (
	"fmt"
	"math"
)

// Sqrt is a math.Sqrt function
func Sqrt(x float64) float64 {
	var previousZ float64
	z := 1.0
	fmt.Printf("[%v] previousZ=%g z=%g \n", 0, previousZ, z)
	for i := 0; i < 10; i++ {	
		previousZ = z
		z -= (z*z - x) / (2*z)
		fmt.Printf("[%v] previousZ=%g z=%g \n", i+1, previousZ, z)
		if math.Abs(z - previousZ) / previousZ < 0.001 {break}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
