package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(strconv.FormatInt(int64(x), 2), s)

	i1, _ := strconv.Atoi("123")
	i2, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(i1, i2)
}
