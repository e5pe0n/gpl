package main

import "fmt"

func main() {
	fmt.Println(gcd(12, 8))  // 4
	fmt.Println(gcd(8, 12))  // 4
	fmt.Println(gcd(12, 13)) // 1
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
