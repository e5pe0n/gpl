package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN

	fmt.Println(math.IsNaN(z / z)) // true

	nan := math.NaN()
	fmt.Println(nan == nan) // false
	fmt.Println(nan < nan)  // false
	fmt.Println(nan > nan)  // false
}
