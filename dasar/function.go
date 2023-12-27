package main

import "fmt"

func info() {
	name := "latif harkat"
	address := "sukabumi, jawa barat."
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Address: %s\n", address)
}

func main() {
	info()
	for i := 1; i <= 10; i++ {
		info()
	}
}
