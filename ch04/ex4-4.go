package main

import "fmt"

func reverse_ex4_4(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate_ex4_4(s []int, pos int) {
	reverse_ex4_4(s[:pos])
	reverse_ex4_4(s[pos:])
	reverse_ex4_4(s[:])
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	rotate_ex4_4(a, 2)
	fmt.Println(a) // [3 4 5 1 2]
}
