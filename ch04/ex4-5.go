package main

import "fmt"

func dup(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	b := a[:1]
	for _, v := range a[1:] {
		if b[len(b)-1] != v {
			b = append(b, v)
		}
	}
	a = b
	return b
}

func dup2(a []string) []string {
	if len(a) == 0 {
		return []string{}
	}
	last := 0
	for _, v := range a[1:] {
		if a[last] != v {
			a[last+1] = v
			last++
		}
	}
	return a[:last+1]
}

func f1_ex4_5() {
	a := []string{"a", "a", "b", "c", "c", "c", "b", "a"}
	fmt.Println(a) // [a a b c c c b a]
	b := dup(a)
	fmt.Println(a) // [a b c b a c b a]
	fmt.Println(b) // [a b c b a]
}

func f2_ex4_5() {
	a := []string{"a", "a", "b", "c", "c", "c", "b", "a"}
	fmt.Println(a) // [a a b c c c b a]
	b := dup2(a)
	fmt.Println(a) // [a b c b a c b a]
	fmt.Println(b) // [a b c b a]
}

func main() {
	f1_ex4_5()
	f2_ex4_5()
}
