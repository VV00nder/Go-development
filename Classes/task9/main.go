package main

import "fmt"

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}
func main() {
	var a, s, c int
	fmt.Println("Введите кол-во чисел:")
	fmt.Scan(&c)
	for i := 0; i < c; i++ {
		fmt.Scan(&a)
		s = sum(s, a)
	}
	fmt.Println(s)
}
