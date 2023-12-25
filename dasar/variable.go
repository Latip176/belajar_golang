package main

import "fmt"

func main() {
	/*
		varriable di golang harus dipanggil, kalo hanya di deklarasikan akan error.
	*/

	// deklarasi dengan tipe data
	var name string
	name = "Latif Harkat"
	fmt.Println(name)

	// tanpa deklarasi
	var age = 18
	fmt.Println(age)

	// tanpa inisialisasi var, harus pakai ':'
	addres := "Sukabumi, Jawa Barat, Indonesia."
	fmt.Println(addres)

	// multiple variable
	var (
		firstName = "Latif"
		lastName  = "Harkat"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
