package main

import "fmt"

func divmod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("division by zero")
	}
	return a / b, a % b, nil
}

func main() {
	var a, b int
	var c error
	fmt.Scan(&a)
	fmt.Scan(&b)
	a, b, c = divmod(a, b)
	fmt.Println(a, b, c)
}
