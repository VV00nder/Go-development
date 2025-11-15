package main

import "fmt"

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c*9/5 + 32)
}
func (c *Celsius) Add(d Celsius) {
	*c += d
}
func main() {
	var a, b Celsius
	fmt.Println("Celsius_1: ")
	fmt.Scan(&a)
	fmt.Println("Celsius_2: ")
	fmt.Scan(&b)
	fmt.Printf("Fahrenheit_1: %.2f\n", a.ToFahrenheit())
	fmt.Printf("Fahrenheit_2: %.2f\n", b.ToFahrenheit())
	a.Add(b)
	fmt.Printf("Sum Celsius: %.2f\n", a)
	fmt.Printf("Sum Fahrenheit: %.2f\n", a.ToFahrenheit())
}
