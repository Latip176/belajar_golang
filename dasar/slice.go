package main

import "fmt"

func main() {
	/*
		Array dan Slice beda, walaupun sama.
	*/
	type kalimat string
	type number int

	nama_buah := [...]kalimat{
		"Apel",
		"Mangga",
		"Pisang",
		"Jeruk",
		"Sirsak",
	}
	slice1 := nama_buah[1:4]
	fmt.Printf("Slice: %s", slice1)
	slice1 = append(slice1, "Manggis") // menambahkan data baru ke array lewat slice.
	fmt.Printf("Slice: %s", slice1)

	nilai_ujian := [5]number{
		29,
		40,
		60,
		12,
		43,
	}
	slice2 := nilai_ujian[2:5]
	fmt.Printf("\nSlice: %d\n", slice2)

	const len = len(nama_buah)   // menghitung panjang slice
	const cap = cap(nilai_ujian) // menghitung capacity
	fmt.Println(len)
	fmt.Println(cap)
}
