package main

import (
	"fmt"
	"strings"
)

func count(str string) int {
	count := 0
	mas := strings.Split(str, " ")
	fmt.Println(mas)
	for i := 0; i < len(mas); i++ {
		if len(mas[i]) > 1 {
			count++
			fmt.Println(mas[i])
		} else {
			if len(mas[i]) == 1 {
				for i, r := range "qwertyuiopasdgfhjklzxcvbnm" {
					if mas[i] == string(r) {
						count++
					}
				}

			}
		}
	}
	return count

}

func main() {
	var a string
	fmt.Scan(&a)
	if len(a) == 0 {
		fmt.Println("0")
	} else {
		fmt.Println(count(strings.ToLower(a)))
	}
}
