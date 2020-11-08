package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is a error type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt is a math method
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var previousZ float64
	z := 1.0
	// fmt.Printf("[%v] previousZ=%g z=%g \n", 0, previousZ, z)
	for i := 0; i < 10; i++ {	
		previousZ = z
		z -= (z*z - x) / (2*z)
		// fmt.Printf("[%v] previousZ=%g z=%g \n", i+1, previousZ, z)
		if math.Abs(z - previousZ) / previousZ < 0.001 {break}
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))	
}
