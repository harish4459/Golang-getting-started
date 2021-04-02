package main

import "fmt"

func main() {

	arr2 := [3]int{1, 2, 3}
	slice := arr2[:]

	fmt.Println(arr2)
	fmt.Println(slice)

	sunny := []int{1, 2, 3}
	fmt.Println(sunny)
	sunny = append(sunny, 45, 59, 444459)
	fmt.Println(sunny)

	s2 := sunny[1:]
	s3 := sunny[:2]
	s4 := sunny[1:2]

	fmt.Println(s2, s3, s4)

}
