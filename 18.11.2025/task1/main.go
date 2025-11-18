package main

import (
	"fmt"
)

func main() {
	week := []string{"ПН", "ВТ", "СР", "ЧТ", "ПТ", "СБ", "ВС"}
	week_last := week[3:]
	fmt.Println(week)
	fmt.Println(week[1])
	fmt.Println(week_last)
	week_dop := week[:]
	for i := 0; i < 7; i++ {
		week_dop = append(week_dop, "ПТ")
	}
	fmt.Println(week_dop)

}
