package main

import "fmt"

func main() {
	/*
		constant variable tidak akan error walau tidak dipanggil. berbeda dengan var
		constant adalah variable yang tidak bisa di ubah datanya.
	*/

	const nama string = "Latif Harkat"
	const umur = 18

	fmt.Println("Nama saya:", nama)
	fmt.Println("Umur saya:", umur)

	// multiple const
	const (
		firstName = "Latif"
		lastName  = "Harkat"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
