package main

import (
	"fmt"
)

const (
	test  = 1
	test2 = "harish"
)

const (
	first = iota + 6
	second
)

const (
	third = iota
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)

}
