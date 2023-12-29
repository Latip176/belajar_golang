package main

import "fmt"

type kalimat string

func setNama(nama kalimat) kalimat {
	return nama
}

func main() {
	namaSaya := setNama

	fmt.Printf("Nama Saya %s", namaSaya("Latif Harkat"))
}
