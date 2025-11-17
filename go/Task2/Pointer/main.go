package main

import (
	"fmt"
)

func add(n *int) {
	*n += 10
}

func mult(s []int) {
	for i := range s {
		s[i] *= 2
	}
}

func main() {
	n := 0
	add(&n)
	fmt.Println("n=", n)

	s := []int{1, 2, 3}
	mult(s)
	fmt.Println("s=", s)
}
