package main

import "fmt"

func reverse(s []int) {
	// reverse slice in place
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	reverse(a[:])
	fmt.Println(a) // [5 4 3 2 1]

	s := []int{0, 1, 2, 3, 4, 5}
	// rotate s left by two positions
	reverse(s[:2])
	reverse(s[2:])
	reverse(s[:])
	fmt.Println(s) // [2 3 4 5 0 1]
}
