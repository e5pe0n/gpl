package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1)
	fmt.Println(x) // [1]

	x = append(x, 2, 3)
	fmt.Println(x) // [1 2 3]

	x = append(x, 4, 5, 6)
	fmt.Println(x) // [1 2 3 4 5 6]

	x = append(x, x...)
	fmt.Println(x) // [1 2 3 4 5 6 1 2 3 4 5 6]
}
