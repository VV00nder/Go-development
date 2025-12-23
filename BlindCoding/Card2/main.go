package main

import (
	"fmt"
	"math"
)

func del(first int, second int) (float64, bool) {
	if second == 0 {
		return 0, false
	}
	ch := float64(first) / float64(second)
	fmt.Println(ch)
	return math.Round(ch*(100)) / 100, true
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(del(a, b))
}
