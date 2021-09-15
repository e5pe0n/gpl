package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string { return fmt.Sprintf("%q", list) }

func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	s1 := []string{"a", "b", "c"}
	s2 := []string{"a", "b"}
	s3 := []string{"a", "b", "c"}
	ss := [][]string{s1, s2, s3}
	for _, s := range ss {
		Add(s)
	}
	for _, s := range ss {
		fmt.Println(s, Count(s))
	}
}

// [a b c] 2
// [a b] 1
// [a b c] 2
