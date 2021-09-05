package main

func main() {}

func Signum(x int) int {
	// tagless switch
	switch {
	case x > 0:
		return +1
	default:
		return 0
	case x < 0:
		return -1
	}
}
