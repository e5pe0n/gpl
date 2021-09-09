package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buf bytes.Buffer
	for i, c := range s {
		if i != 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%c", c)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))    // 1,234,567
	fmt.Println(comma("12345678"))   // 12,345,678
	fmt.Println(comma("123456789"))  // 123,456,789
	fmt.Println(comma("1234567890")) // 1,234,567,890
}
