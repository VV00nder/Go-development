package main

import "fmt"

func main() {
	PI := 3.14
	var r float64
	fmt.Scan(&r)
	fmt.Printf("%.3f\n", PI*r*r)
}
