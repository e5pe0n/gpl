package main

import "fmt"

const size = 5

func reverse_arr(s *[size]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func main() {
	a := [size]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]
	reverse_arr(&a)
	fmt.Println(a) // [5 4 3 2 1]
}
