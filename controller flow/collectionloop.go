package main

func main() {
	knownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range knownPorts {
		println(k, v)
	}

}
