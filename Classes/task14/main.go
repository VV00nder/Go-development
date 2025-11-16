package main

import (
	"fmt"
)

func bubbleSort(a *[5]int) int {
	flag := 1
	cnt := 0
	for flag != 0 {
		flag = 0
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				cnt++
				flag += 1
			}
		}
	}
	return cnt
}

func bubbleSortdown(a *[5]int) int {
	flag := 1
	cnt := 0
	for flag != 0 {
		flag = 0
		for i := 1; i < len(a); i++ {
			if a[i-1] < a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				cnt++
				flag += 1
			}
		}
	}
	return cnt
}

func main() {
	var arr [5]int
	for i := 0; i < 5; i++ {
		if _, err := fmt.Scan(&arr[i]); err != nil {
			return
		}
	}
	arr2 := arr
	n := bubbleSort(&arr)
	for i, v := range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Printf("Кол-во перестановок: %v", n)
	fmt.Println()
	n = bubbleSortdown(&arr2)
	for i, v := range arr2 {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Printf("Кол-во перестановок: %v", n)
}
