package main

import (
	"fmt"
	"math"
)

func del(first int, second int) (float64, bool) {
	if second == 0 {
		return 0, false
	}
	ch := first / second
	if ch == int(ch) {
		return float64(ch), true
	}
	return float64(math.Round(float64(ch*10^2)) / (100.0)), true
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(del(a, b))
}
