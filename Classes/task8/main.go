package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", a)
		c := a + b
		a = b
		b = c
	}
}
