package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, error := rot.r.Read(p)
	symb_a := 'a'
	symb_A := 'A'
	for i := range p {
		if p[i] >= 'a' && p[i] <= 'z' {
			if p[i] < byte(symb_a + 13) {
				p[i] += 13
			} else {
				p[i] -= 13
			}
		} else {
			if p[i] >= 'A' && p[i] <= 'Z' {
				if p[i] < byte(symb_A + 13) {
					p[i] += 13
				} else {
					p[i] -= 13
				}
			}
		}

	}
	return n, error
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
