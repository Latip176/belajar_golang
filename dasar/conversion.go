package main

import "fmt"

func main() {
	var namaByte byte = "Latif Harkat"[0]
	var nama string = string(namaByte)

	fmt.Println(nama)
}
