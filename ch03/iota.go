package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
}

// Sunday 0
// Monday 1
// Tuesday 2
// Wednesday 3
// Thursday 4
// Friday 5
// Saturday 6
