package main

import "fmt"

func main() {
	m := map[string]int{"foo": 2}
	fmt.Println(m)

	fmt.Println(m["foo"])

	m["foo"] = 4459
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

}
