package main

import "fmt"

/*
struct adalah template yg digunakan untuk menggabungkan 0 atau lebih, dari tipe data.
*/

type Information struct {
	Name, Address, NoTelp string
	Age, TB               int
}

func main() {
	// cara pertama
	var Data Information
	info := Data
	info.Name = "Latif Harkat"
	info.Address = "Sukabumi, Jawa Barat, Indonesia"
	info.NoTelp = "+6283172566909"
	info.Age = 18
	info.TB = 175

	fmt.Println(info)

	// cara kedua
	info2 := Information{
		Name:    "Latif Harkat",
		Age:     18,
		Address: "Sukabumi, Jawa Barat, Indonesia",
		NoTelp:  "+6283172566909",
		TB:      175,
	}
	fmt.Println(info2)
}
