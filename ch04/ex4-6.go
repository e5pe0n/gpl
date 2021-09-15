package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squash(a []byte) []byte {
	b := []byte{}
	space := false
	for len(a) > 0 {
		r, size := utf8.DecodeRune(a)
		if unicode.IsSpace(r) {
			if !space {
				b = append(b, byte(' '))
			}
			space = true
		} else {
			buf := make([]byte, 3)
			n := utf8.EncodeRune(buf, r)
			b = append(b, buf[:n]...)
			space = false
		}
		a = a[size:]
	}
	a = b
	return b
}

func main() {
	ss := []string{
		"Hello,  世界",
		"Hello,	世界",
		"Hello, 　世界"}
	for _, s := range ss {
		t := s
		fmt.Printf("%s -> %s\n", t, string(squash([]byte(s))))
	}
}

// Hello,  世界 -> Hello, 世界
// Hello,  世界 -> Hello, 世界
// Hello, 　世界 -> Hello, 世界
