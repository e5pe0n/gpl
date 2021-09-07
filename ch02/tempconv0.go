package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	absoluteZeroC Celsius = -273.15
	FreeaingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	fmt.Printf("%g\n", BoilingC-FreeaingC) // 100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreeaingC)) // 180
	// fmt.Printf("%g\n", boilingF-FreeaingC) // Compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // Compile error: type mismatch
	fmt.Println(c == Celsius(f)) // true

	c = FToC(212.0)
	fmt.Println(c.String()) // 100°C
	fmt.Printf("%v\n", c)   // 100°C
	fmt.Printf("%s\n", c)   // 100°C
	fmt.Println(c)          // 100°C
	fmt.Printf("%g\n", c)   // 100
	fmt.Println(float64(c)) // 100
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
