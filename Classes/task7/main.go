package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)
	if score > 100 || score < 0 {
		fmt.Println("Invalid")
	} else {
		switch score / 10 {
		case 10, 9:
			fmt.Println("A")
		case 8:
			fmt.Println("B")
		case 7:
			fmt.Println("C")
		case 6:
			fmt.Println("D")
		case 5, 4, 3, 2, 1, 0:
			fmt.Println("F")
		default:
			fmt.Println("Invalid")
		}
	}
}
