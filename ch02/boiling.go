package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	fmt.Printf("boiling point = %f°F or %f°C\n", f, c)
}

// boiling point = 212°F or 100°C
// boiling point = 212.000000°F or 100.000000°C
