package main

import "fmt"

func main() {
	fmt.Println("Hello, playground! ")
	port := 300
	startWebServer(port, 2)
}

func startWebServer(ports, retries int) {
	fmt.Println("Hello, starting server on port! ", ports)
	fmt.Println("num of retries", retries)
}
