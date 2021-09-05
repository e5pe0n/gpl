package main

import (
	"fmt"
	"math/rand"
)

func main() {
	heads, tails := 0, 0
	switch coinflip() {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on edge!")
	}
}

func coinflip() string {
	if rand.Intn(2) == 1 {
		return "heads"
	} else {
		return "tails"
	}
}
