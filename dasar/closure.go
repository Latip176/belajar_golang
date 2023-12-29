package main

import "fmt"

type kalimat string

func main() {
	nama := "Omen Craft"
	closure := func() {
		nama = "Latif Harkat" // ubah data dari closure
		fmt.Println("Ubah Nama")
	}
	closure()
	fmt.Println(nama)
}
