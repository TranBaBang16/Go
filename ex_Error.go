package main

import (
	"fmt"
)

const m float64 = 0.0000000000000001

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return -1, &err
	}
	z := float64(1)
	for (z*z-x)*(z*z-x) > m {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
