package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " ")) // /tmp/go-build1859771998/b001/exe/ex1-1 hello world
}
