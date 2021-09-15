package main

import "fmt"

func main() {
	var s []int
	for i, v := range s {
		fmt.Println(i, v)
	}
}
