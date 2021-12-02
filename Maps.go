
//Implement WordCount. It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.


package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m:=make(map[string]int)
	var word string
	for i:=0;i<len(s);i++{
		if string(s[i])==""{
			continue
		}
		
		for j:=i+1;j<len(s);j++{
			if j==len(s)-1{
				word=s[i:len(s)]
				i=len(s)
			}else{
			if s[j]==' '{
				word=s[i:j]
				i=j;
			}else{
				continue
			}
				
			}
			
			_,ok:=m[word]
			if ok {
				m[word]++
			} else{
				m[word]=1
			}
			break
		}
		
	}
	
	return m
}


func main() {
	wc.Test(WordCount)

}