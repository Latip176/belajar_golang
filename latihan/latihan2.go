package main

import "fmt"

/**

Latihan Memanipulasi dan Mengakses array

*/

func main() {
	nama_buah := [...]interface{}{
		"Apel",
		"Pisang",
		"Semangka",
		"Rambutan",
		"Mangga",
		[...]string{
			"Kapal",
			"Mobil",
			"Pesawat",
		},
	} // array 2 dimensi

	fmt.Println(nama_buah)

	nama_buah[0] = "Manggis" // ubah data index ke 0

	fmt.Println(nama_buah)

	fmt.Println(nama_buah[5].([3]string)[1]) // akses Mobil

	fmt.Println(len(nama_buah)) // menghitung panjang data, bukan jumlah data
}
