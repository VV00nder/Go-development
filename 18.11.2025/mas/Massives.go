package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := []byte{}
	for i := range b {
		c = append(c, b[i])
	}
	return c
}

func main() {
	digits := FindDigits("181125/mas/digits")
	fmt.Println(string(digits))
	digitsCopy := CopyDigits("181125/mas/digits")
	fmt.Println(string(digitsCopy))
}
