package main

import (
	"fmt"
)

func main() {

	port := 300
	portsi, isstarted := startWebServer(port)
	fmt.Println(portsi, isstarted)

}

func startWebServer(ports int) (int, error) {
	fmt.Println("Hello, starting server on port! ", ports)
	return ports, nil
}
