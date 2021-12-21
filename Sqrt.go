// given a number x, we want to find the number z for which z² is most nearly x.



package main

import (
	"fmt"
)
const m float64 = 0.0000000000000001
func Sqrt(x float64) float64 {
	z:=float64(1)
// so sanh voi m de vong lap dung lai khi ma z dat den gia tri ma sai so so voi sqrt(x) la chap nhan duoc
//trong bai nay em lay m=0.1*10^-16 => ket qua chinh xac den 8 chu so thap phan
	for (z*z-x)*(z*z-x) > m {
	z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(7))
}


