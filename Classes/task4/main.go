package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	if b == 0 {
		fmt.Println("Нельзя делить на 0")
	} else {
		fmt.Printf("%.4f\n", float64(a)/float64(b))
		fmt.Println(a % b)
	}

}
