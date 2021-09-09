package main

import (
	"bytes"
	"fmt"
	"strings"
)

func enhancedComma(s string) string {
	var buf bytes.Buffer
	start := 0
	if s[0] == '-' {
		buf.WriteByte('-')
		start++
	}
	if s[0] == '+' {
		buf.WriteByte('+')
		start++
	}
	n := len(s)
	dot := strings.LastIndex(s, ".")
	if dot < 0 {
		dot = n
	}
	for i, c := range s[start:dot] {
		if i != 0 && (dot-(i+start))%3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%c", c)
	}
	if dot != n {
		buf.WriteByte('.')
	}
	if dot+1 < n {
		for i, c := range s[dot+1:] {
			if i != 0 && i%3 == 0 {
				buf.WriteByte(',')
			}
			fmt.Fprintf(&buf, "%c", c)
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(enhancedComma("1234567.1234567"))
	fmt.Println(enhancedComma("+1234567.1234567"))
	fmt.Println(enhancedComma("-1234567.1234567"))
	fmt.Println(enhancedComma("12345678.12345678"))
	fmt.Println(enhancedComma("+12345678.12345678"))
	fmt.Println(enhancedComma("-12345678.12345678"))
	fmt.Println(enhancedComma("123456789.123456789"))
	fmt.Println(enhancedComma("+123456789.123456789"))
	fmt.Println(enhancedComma("-123456789.123456789"))
	fmt.Println()

	fmt.Println(enhancedComma("1234567"))
	fmt.Println(enhancedComma("+1234567"))
	fmt.Println(enhancedComma("-1234567"))
	fmt.Println(enhancedComma("12345678"))
	fmt.Println(enhancedComma("+12345678"))
	fmt.Println(enhancedComma("-12345678"))
	fmt.Println(enhancedComma("123456789"))
	fmt.Println(enhancedComma("+123456789"))
	fmt.Println(enhancedComma("-123456789"))
	fmt.Println()

	fmt.Println(enhancedComma("1234567."))
	fmt.Println(enhancedComma("+1234567."))
	fmt.Println(enhancedComma("-1234567."))
	fmt.Println(enhancedComma("12345678."))
	fmt.Println(enhancedComma("+12345678."))
	fmt.Println(enhancedComma("-12345678."))
	fmt.Println(enhancedComma("123456789."))
	fmt.Println(enhancedComma("+123456789."))
	fmt.Println(enhancedComma("-123456789."))
	fmt.Println()

	fmt.Println(enhancedComma(".1234567"))
	fmt.Println(enhancedComma("+.1234567"))
	fmt.Println(enhancedComma("-.1234567"))
	fmt.Println(enhancedComma(".12345678"))
	fmt.Println(enhancedComma("+.12345678"))
	fmt.Println(enhancedComma("-.12345678"))
	fmt.Println(enhancedComma(".123456789"))
	fmt.Println(enhancedComma("+.123456789"))
	fmt.Println(enhancedComma("-.123456789"))
	fmt.Println()
}

// 1,234,567.123,456,7
// +1,234,567.123,456,7
// -1,234,567.123,456,7
// 12,345,678.123,456,78
// +12,345,678.123,456,78
// -12,345,678.123,456,78
// 123,456,789.123,456,789
// +123,456,789.123,456,789
// -123,456,789.123,456,789

// 1,234,567
// +1,234,567
// -1,234,567
// 12,345,678
// +12,345,678
// -12,345,678
// 123,456,789
// +123,456,789
// -123,456,789

// 1,234,567.
// +1,234,567.
// -1,234,567.
// 12,345,678.
// +12,345,678.
// -12,345,678.
// 123,456,789.
// +123,456,789.
// -123,456,789.

// .123,456,7
// +.123,456,7
// -.123,456,7
// .123,456,78
// +.123,456,78
// -.123,456,78
// .123,456,789
// +.123,456,789
// -.123,456,789
