package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])        // 0
	fmt.Println(a[len(a)-1]) //0

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Printf("%x\n", q) // [1 2 3]
	fmt.Println(r[2])     // 0

	q2 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q2)      // [3]int
	fmt.Printf("%t\n", q == q2) // true
}
