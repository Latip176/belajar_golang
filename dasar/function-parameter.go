package main

import "fmt"

func info(name string, address string) {
	fmt.Printf("Name: %s\nAddress: %s\n", name, address)
}
func main() {
	name := "latif harkat"
	address := "sukabumi, jawa barat"
	info(name, address)
}
