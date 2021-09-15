package main

import "fmt"

func main() {
	ages1 := make(map[string]int)
	ages1["alice"] = 31
	ages1["charlie"] = 34

	ages2 := map[string]int{}
	ages2["alice"] = 31
	ages2["charlie"] = 34

	ages3 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	fmt.Println(ages3["alice"]) // 31
	delete(ages3, "alice")
	fmt.Println(ages3["alice"]) // 0

	ages3["bob"] = ages3["bob"] + 1
	fmt.Println(ages3["bob"]) // 1

	for name, age := range ages3 {
		fmt.Println(name, age)
	}
	// 	bob 1
	// charlie 34

	if age, ok := ages3["bob"]; !ok {
		fmt.Println("not exist")
	} else {
		fmt.Printf("found; age=%d\n", age)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
