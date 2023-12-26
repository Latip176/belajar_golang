package main

import "fmt"

func main() {
	type latif bool
	type kalimat string
	type number int

	const (
		firstName kalimat = "Latif"
		lastName  kalimat = "Harkat"
	)
	var condition1 latif = firstName != lastName
	fmt.Println(condition1)

	var (
		x number = 10
		y number = 30
	)
	var condition2 latif = x >= y
	fmt.Println(condition2)

	var (
		nilai1 number = 97
		nilai2 number = 70
	)
	var (
		ujian1 number = 90
		ujian2 number = 88
	)
	var nilaiAkhir number = (nilai1 + nilai2) / 2
	var ujianAkhir number = (ujian1 + ujian2) / 2
	const syaratNilaiKelulusan number = 75
	const syaratNilaiUjian number = 80

	var cek_kelulusan latif = nilaiAkhir > syaratNilaiKelulusan
	var cek_ujian latif = ujianAkhir > syaratNilaiUjian
	var cek_akhir latif = cek_kelulusan && cek_ujian

	fmt.Println(cek_akhir)

}
