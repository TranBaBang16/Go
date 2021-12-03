//Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	str := strings.Fields(s)
	for i := range str{
		_,ok := m[str[i]]
		if ok{
			m[str[i]]++
		}else{
			m[str[i]] = 1
		}
	}
	return m
}


func main() {
	wc.Test(WordCount)

}

