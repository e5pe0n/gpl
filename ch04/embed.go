package main

import "fmt"

type Point struct{ X, Y int }

type Circle struct {
	Point  // embedding struct Point as anonymous field
	Radius int
}

type Wheel struct {
	Circle // embedding struct Circle as anonymous field
	Spokes int
}

func main() {
	var w Wheel

	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	fmt.Printf("%#v\n", w) // main.Wheel{Circle:main.Circle{Point:main.Point{X:8, Y:8}, Radius:5}, Spokes:20}
	w.X = 42
	fmt.Printf("%v\n", w) // {{{42 8} 5} 20}
}
