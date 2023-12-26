package main

import "fmt"

func main() {
	var nama_mobil [3]string

	nama_mobil[0] = "Avanza"
	nama_mobil[1] = "Kijang"
	nama_mobil[2] = "Lamborghini"
	fmt.Println(nama_mobil)

	var nilai_ujian = [...]int{
		87,
		80,
		94,
	} // ... untuk jika belum menentukan panjang array
	fmt.Println(nilai_ujian)

	const jumlah_data = len(nilai_ujian) // menghitung panjang array bukan jumlah data
	fmt.Println(jumlah_data)

}
