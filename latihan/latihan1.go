package main

import "fmt"

type kalimat string
type number int

func main() {
	var (
		name kalimat = "Latif Harkat"
		age  number  = 18
	)

	fmt.Printf("Nama Saya: %s\nUmur Saya: %d\n", name, age)
}
