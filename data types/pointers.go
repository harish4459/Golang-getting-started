package main

import (
	"fmt"
)

func main() {
	var firstName *string = new(string)

	*firstName = "Arthur"

	fmt.Println(*firstName)

	secondname := "sunny"
	fmt.Println(secondname)

	ptr := &secondname
	fmt.Println(ptr, &ptr, *ptr)

	secondname = "Harish"
	fmt.Println(ptr, *ptr)

}
