package main

import "fmt"

func main() {
	type kalimat string // aliasi string menjadi 'kalimat'
	type number int     // alias integer menjadi 'number'

	const nama kalimat = "Latif Harkat"
	var umur number = 18

	fmt.Println("Nama saya:", nama)
	fmt.Println("Umur saya:", umur)

	const (
		firstName kalimat = "Latif"
		lastName  kalimat = "Harkat"
	)
	fmt.Println("Nama Depan saya:", firstName)
	fmt.Println("Nama Belakang saya:", lastName)
}
