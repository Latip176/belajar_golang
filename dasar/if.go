package main

import "fmt"

func main() {
	type k string
	type v int

	const nama = "Latif Harkat"

	data_nilai_ujian := map[k]v{
		"nilai1": 50,
		"nilai2": 34,
	}

	data_nilai_rapot := [...]v{
		30,
		20,
	}

	if data_nilai_ujian["nilai1"] >= data_nilai_rapot[0] {
		fmt.Println("Selamat Anda lulus!", nama)
	} else {
		fmt.Println("Maaf Anda tidak lulus!", nama)
	}

	// short statment
	sekarang := "panas"
	if cuaca := "hujan"; sekarang == cuaca {
		fmt.Println("Sekarang sedang Hujan!")
	} else if cuaca := "panas"; sekarang == cuaca {
		fmt.Println("Sekarang sedang Panas!")
	} else {
		fmt.Println("Cuaca Sekarang tidak diketahui.")
	}
}
