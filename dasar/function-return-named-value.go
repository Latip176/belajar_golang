package main

import "fmt"

type kalimat string
type number int

func myData() (nama kalimat, umur number, alamat kalimat) {
	nama = "Latif Harkat"
	umur = 18
	alamat = "Sukabumi, Jawa Barat, Indonesia."
	return nama, umur, alamat
}

func main() {
	n, u, _ := myData()

	fmt.Println(n, u)
}
