package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	x = x + 3
	fmt.Println(x)
	x += 2
	fmt.Println(x)
	x *= 4
	fmt.Println(x)
	x -= 10
	fmt.Println(x)
}
