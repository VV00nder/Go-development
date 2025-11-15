package main

import "fmt"

func swap(a, b *int) (int, int) {
	*a, *b = *b, *a
	return *a, *b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	a, b = swap(&a, &b)
	fmt.Println(a, b)
}
