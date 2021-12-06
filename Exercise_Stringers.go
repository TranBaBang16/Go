package main

import "fmt"

type IPAddr [4]byte

// Byte2String convert Byte type to String type
func Byte2String(x byte) string {
	var s string
	for x/10 != 0 {
		r := x % 10
		x = x / 10
		s = s + string(r+48)
	}
	s = s + string(x+48)
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}

// TODO: Add a "String() string" method to IPAddr.

func (ipaddr IPAddr) String() string {
	var result string = Byte2String(ipaddr[0])
	
	for i := 1 ; i < len(ipaddr) ; i++ {
		result += "." + Byte2String(ipaddr[i])
	}
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

