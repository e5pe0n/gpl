package ifscope

import (
	"fmt"
)

var fname string

func main() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
}

func f() int {
	return 1
}

func g(x int) int {
	return x * 2
}
