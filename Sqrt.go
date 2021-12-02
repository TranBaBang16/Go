// given a number x, we want to find the number z for which z² is most nearly x.



package main

import (
	"fmt"
)
const m float64 =0.0000000000000001
func Sqrt(x float64) float64 {
	z:=float64(1)
for (z*z-x)*(z*z-x)>m {
z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(7))
}
