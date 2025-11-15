package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		if n > 0 {
			fmt.Println("positive even")
		} else {
			if n < 0 {
				fmt.Println("negative even")
			} else {
				fmt.Println("zero")
			}
		}
	} else {
		if n > 0 {
			fmt.Println("positive odd")
		} else {
			fmt.Println("negative odd")
		}
	}
}
