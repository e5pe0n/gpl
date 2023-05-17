package main

import "fmt"

const builingF = 212.0

func main() {
	var f = builingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

// boiling point = 212°F or 100°C
