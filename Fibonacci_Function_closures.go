//Implement a fibonacci function that returns a function (a closure) 
//that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func( int) int {
	preNumber1 := 0
	preNumber2 := 1
	return func(x int) int{
		if x == 0 { 
			return 0 
		}
		if x == 1 {
			return 1 
		}
		res := preNumber1 + preNumber2
		preNumber1 = preNumber2
		preNumber2 = res
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 24; i++ {
		fmt.Println(f(i))
	}
}

