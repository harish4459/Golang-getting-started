package main

import "fmt"

func main() {
	type sunny struct {
		Id         int
		firstname  string
		secondname string
	}

	var u sunny
	u.Id = 59
	u.firstname = "harish"
	u.secondname = "sunny"

	fmt.Println(u)

	u2 := sunny{
		Id:         59,
		firstname:  "sunny",
		secondname: "chowdary",
	}
	fmt.Println(u2)

}
